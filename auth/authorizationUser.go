package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"task-api/database"
	"time"
)

func AuthorizationUser(c *gin.Context) {
	//Receive request body with user information
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cannot read body request"})
		log.Fatal(err)
		return
	}
	//Check in database if email and password is correct
	row := database.DB.QueryRow(`SELECT email,password_hash,secret FROM users WHERE email = $1`, body.Email)
	var emailFromDB string
	var hashedPass []byte
	var secret string
	if row.Scan(&emailFromDB, &hashedPass, &secret) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Incorrect email"})
		log.Fatal(errors.New("incorrect email"))
		return
	}
	if bcrypt.CompareHashAndPassword(hashedPass, []byte(body.Password)) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		log.Fatal(errors.New("incorrect password"))
		return
	}
	//Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": body.Email,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cannot create JWT"})
		log.Fatal(err)
		return
	}
	//Give JWT in cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "localhost", false, true)
}
