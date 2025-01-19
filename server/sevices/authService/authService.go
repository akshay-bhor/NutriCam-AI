package authService

import (
	"net/http"
	"server/middlewares/validation"
	"server/utils/logger"
	"server/utils/response"

	"github.com/gin-gonic/gin"
)

type googleRegistration struct {
	Mail string `json:"mail"`
	Gid  string `json:"gid"`
}

type passRegistration struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type registration interface {
	CreateUser()
}

func RegistrationFactory(c *gin.Context) registration {
	body, _ := c.Get("body")
	reqBody, ok := body.(validation.RegistrationRequest)

	if !ok {
		response.ErrorResponse(c, http.StatusBadRequest, "", "Failed")
	}

	if reqBody.Gid != nil {
		return googleRegistration{
			Mail: reqBody.Mail,
			Gid:  *reqBody.Gid,
		}
	} else {
		return passRegistration{
			Mail: reqBody.Mail,
			Pass: *reqBody.Pass,
		}
	}
}

func (u googleRegistration) CreateUser() {
	logger.Info("google")
}

func (u passRegistration) CreateUser() {
	logger.Info("pass")
}
