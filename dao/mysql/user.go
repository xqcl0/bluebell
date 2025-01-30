package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const (
	secret = "wlllz"
)

func queryUserById() {

}

// CheckUserExist check if user already exist in db
func CheckUserExist(username string) error {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return errors.New("user exist")
	}
	return err
}

// InsertUser insert user data into db
func InsertUser(user *models.User) (err error) {
	//加密
	pwd := encryptPassword(user.Password)
	sqlStr := "insert into user(user_id, username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, pwd)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
