package authService

import (
	"net/http"
	"server/db"
	"server/middlewares/validation"
	"server/models"
	"server/utils/logger"
	"server/utils/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type googleRegistration struct {
	Mail string  `json:"mail"`
	Gid  *string `json:"gid"`
}

type passRegistration struct {
	Mail string  `json:"mail"`
	Pass *string `json:"pass"`
}

type registration interface {
	CreateUser(*gin.Context) (models.Users, *response.ErrorObject)
}

func RegistrationFactory(c *gin.Context) (registration, *response.ErrorObject) {
	body, _ := c.Get("body")
	reqBody, ok := body.(validation.RegistrationRequest)

	if !ok {
		errObj := response.NewErrorResponse(http.StatusBadRequest, nil, "Failed")
		return nil, &errObj
	}

	if reqBody.Gid != nil {
		return googleRegistration{
			Mail: reqBody.Mail,
			Gid:  reqBody.Gid,
		}, nil
	} else if reqBody.Pass != nil {
		return passRegistration{
			Mail: reqBody.Mail,
			Pass: reqBody.Pass,
		}, nil
	} else {
		errObj := response.NewErrorResponse(http.StatusBadRequest, nil, "Bad request")
		return nil, &errObj
	}
}

func (u googleRegistration) CreateUser(c *gin.Context) (models.Users, *response.ErrorObject) {
	logger.Info("Initiating google reggistration flow")
	mail := u.Mail
	gid := u.Gid

	newUser := models.Users{
		Mail: mail,
		Gid:  gid,
	}

	results := db.DB.Create(&newUser)

	if results.Error != nil {
		logger.Error(results.Error, "Error inserting user in db")

		errObj := response.NewErrorResponse(400, nil, "Something went wrong")

		return models.Users{}, &errObj
	}
	logger.Info("User registration complete")

	return newUser, nil
}

func (u passRegistration) CreateUser(c *gin.Context) (models.Users, *response.ErrorObject) {
	logger.Info("Initiating internal reggistration flow")
	mail := u.Mail
	pass := u.Pass

	bytes, err := bcrypt.GenerateFromPassword([]byte(*pass), 14)

	if err != nil {
		logger.Error(err, "Error hashing password")
		response.NewErrorResponse(400, nil, "Something went wrong")
	}

	hashedPass := string(bytes)

	newUser := models.Users{
		Mail: mail,
		Pass: &hashedPass,
	}

	results := db.DB.Create(&newUser)

	if results.Error != nil {
		logger.Error(results.Error, "Error inserting user in db")

		errObj := response.NewErrorResponse(400, nil, "Something went wrong")
		return models.Users{}, &errObj
	}
	logger.Info("User registration complete")

	return newUser, nil
}
