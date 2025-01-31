package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

const (
	secret = "wlllz"
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
		Password: encryptPassword(p.Password),
	}
	err = mysql.InsertUser(user)
	return
}

func Login(p *models.ParamSignIn) (err error) {
	user, err := mysql.QueryUserByUsername(p.Username)
	if err != nil {
		return err
	}
	h := md5.New()
	h.Write([]byte(secret))
	pwd := hex.EncodeToString(h.Sum([]byte(p.Password)))
	if pwd != user.Password {
		return mysql.ErrorPasswordWrong
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
