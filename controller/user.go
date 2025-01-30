package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var p = new(models.ParamSignUp)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("Sign up with error param", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "wrong params",
			"err": err.Error(),
		})
		return
	}
	err = logic.SignUp(p)
	if err != nil {
		zap.L().Error("Sign up with error", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "sign up fail",
			"err": err.Error(),
		})
		return
	}
	zap.L().Info("user sign up ", zap.Any("user", *p))
	c.JSON(http.StatusOK, "ok")
}
