package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"task-api/database"
)

func GetUserId(c *gin.Context) (userId int, _ error) {
	//get token from cookie
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		return -1, errors.New("unauthorized person")
	}
	//divide token on parts
	parts := strings.Split(tokenStr, ".")
	//decode claims on json
	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return -1, errors.New("invalid token")
	}
	var body struct {
		Exp   int    `json:"exp"`
		Email string `json:"user"`
	}
	//get from json claims email
	if json.Unmarshal(claimsJSON, &body) != nil {
		return -1, errors.New("invalid token")
	}
	//search user by email
	row := database.DB.QueryRow(`SELECT id FROM task.public.users WHERE email = $1`, body.Email)
	//get their id
	if row.Scan(&userId) != nil {
		return -1, errors.New("can not find user")
	}
	//return it
	return userId, nil
}
