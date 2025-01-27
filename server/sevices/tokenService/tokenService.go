package tokenService

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"os"
	"server/models"
	"server/utils/logger"
	"server/utils/response"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

type UserToken struct {
	models.Users
}

type JwtClaims struct {
	UserId       uint   `json:"user_id"`
	Mail         string `json:"mail"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	TwoFaEnabled bool   `json:"2fa_enabled"`
	jwt.RegisteredClaims
}

type TempToken struct {
	UserId uint
	exp    time.Time
}

func (user UserToken) IssueJWTToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		UserId:       user.Id,
		Mail:         user.Mail,
		Type:         user.Type,
		Status:       user.Status,
		TwoFaEnabled: user.TwoFaEnabled,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(90 * 24 * time.Hour)),
		},
	})

	return token.SignedString([]byte(os.Getenv("HMAC_SECRET")))
}

func VerifyJWTToken(tokenString string) (bool, *JwtClaims) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Replace this with your actual JWT secret key
		return []byte(os.Getenv("HMAC_SECRET")), nil
	})

	if err != nil {
		logger.Error(err, "Error validating token")
		return false, nil
	}

	if !token.Valid || time.Now().Unix() > claims.ExpiresAt.Unix() {
		logger.Error(nil, "Token expired")
		return false, nil
	}

	return true, claims
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

func (u UserToken) GenerateTopt() (string, string, *response.ErrorObject) {
	logger.Info("Generating totp")

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NutriCamAi",
		AccountName: u.Mail,
	})

	if err != nil {
		logger.Error(err, "Error generating totp")
		errObj := response.NewErrorResponse(500, nil, "Something went wrong")
		return "", "", &errObj
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		logger.Error(err, "Error generating totp")
		errObj := response.NewErrorResponse(500, nil, "Something went wrong")
		return "", "", &errObj
	}
	png.Encode(&buf, img)

	return key.Secret(), base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func VerifyTotp(code string, secret string) bool {
	valid := totp.Validate(code, secret)
	if !valid {
		return false
	}
	return true
}
