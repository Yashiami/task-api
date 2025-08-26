package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"task-api/database"
)

func CreateUser(c *gin.Context) {
	//Receive body request
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cannot read body request"})
		log.Fatal(err)
		return
	}
	//Hash password
	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 15)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cannot encode password"})
		log.Fatal(err)
		return
	}
	cryptedSecret, err := bcrypt.GenerateFromPassword([]byte(body.Username+body.Email), 10)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cannot encode secret"})
		log.Fatal(err)
		return
	}
	//Save it in db
	if _, err = database.DB.Exec(`INSERT INTO users (username,email,password_hash,secret) values ($1,$2,$3,$4)`, body.Username, body.Email, cryptedPass, cryptedSecret); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cannot save data"})
		log.Fatal(err)
		return
	}
	//Send response
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created"})
}
