package tokenService

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"server/models"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserToken struct {
	models.Users
}

type TempToken struct {
	UserId uint
	exp    time.Time
}

func (user UserToken) IssueJWTToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.Id,
		"mail":        user.Mail,
		"type":        user.Type,
		"status":      user.Status,
		"exp":         time.Now().Unix() + (60 * 60 * 24 * 90),
		"2fa_enabled": user.TwoFaEnabled,
	})

	return token.SignedString([]byte(os.Getenv("HMAC_SECRET")))
}

func (user UserToken) GenerateTempToken() (string, error) {
	token := TempToken{
		UserId: user.Id,
		exp:    time.Now().Add(5 * time.Minute),
	}
	tokenJson, err := json.Marshal(token)

	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, []byte(os.Getenv("HMAC_SECRET")))
	mac.Write(tokenJson)
	signature := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(tokenJson) + "." + string(signature), nil
}

func (user UserToken) VerifyTempToken(signedToken string) (uint, error) {
	parts := strings.Split(signedToken, ".")

	if len(parts) != 2 {
		return 0, errors.New("Invalid token")
	}

	tokenJSON, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return 0, errors.New("invalid token content")
	}

	// Verify the signature
	mac := hmac.New(sha256.New, []byte(os.Getenv("HMAC_SECRET")))
	mac.Write(tokenJSON)
	expectedSignature := mac.Sum(nil)
	actualSignature := parts[1]
	if !hmac.Equal(expectedSignature, []byte(actualSignature)) {
		return 0, errors.New("invalid token signature")
	}

	// Deserialize token
	var token TempToken

	if err := json.Unmarshal(tokenJSON, &token); err != nil {
		return 0, errors.New("invalid token content")
	}

	// Check expiration
	if time.Now().After(token.exp) {
		return 0, errors.New("token expired")
	}

	return token.UserId, nil
}
