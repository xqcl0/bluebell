package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if p.Password != p.RePassword {
		zap.L().Error("re_password does not match password")
		return errors.New("re_password does not match password")
	}
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		//查询出错
		return err
	}

	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.InsertUser(user)
	return
}
