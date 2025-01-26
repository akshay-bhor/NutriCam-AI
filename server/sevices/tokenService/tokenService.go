package tokenService

import (
	"os"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserToken struct {
	models.Users
}

func (user UserToken) IssueToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.Id,
		"mail":   user.Mail,
		"type":   user.Type,
		"status": user.Status,
		"exp":    time.Now().Unix() + (60 * 60 * 24 * 90),
	})

	return token.SignedString([]byte(os.Getenv("HMAC_SECRET")))
}
