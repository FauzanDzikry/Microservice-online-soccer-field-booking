package response

import (
	"net/http"
	"user-service/constants"
	errConstant "user-service/constants/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
	Token   *string     `json:"token,omitempty"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
	Token   *string
}

func HttpResponse(param ParamHTTPResp) {
	if param.Err == nil {
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Success,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
			Token:   param.Token,
		})
	} else {
		message := errConstant.ErrInternalServer.Error()
		if param.Message != nil {
			message = *param.Message
		} else if param.Err != nil {
			if errConstant.ErrMapping(param.Err) {
				message = param.Err.Error()
			}
		}
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Error,
			Message: message,
			Data:    param.Data,
		})
	}
}
