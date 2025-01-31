package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var p = new(models.ParamSignUp)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("Sign up with error param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err = logic.SignUp(p)

	if err != nil {
		zap.L().Error("Sign up with error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExit) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	zap.L().Info("user sign up ", zap.Any("user", *p))
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var p = new(models.ParamSignIn)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("login in with error param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err = logic.Login(p)
	if err != nil {
		zap.L().Error("login in with error", zap.Error(err))
		if errors.Is(err, mysql.ErrorPasswordWrong) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		if errors.Is(err, mysql.ErrorUserNotExit) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	zap.L().Info("user login in ", zap.Any("user", *p))
	ResponseSuccess(c, nil)
}
