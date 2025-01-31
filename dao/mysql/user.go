package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
)

func QueryUserByUsername(username string) (*models.User, error) {
	sqlStr := "select id, user_id, username, password from user where username = ?"
	var user = new(models.User)
	err := db.Get(user, sqlStr, username)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrorUserNotExit
	}
	return user, err
}

// CheckUserExist check if user already exist in db
func CheckUserExist(username string) error {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return ErrorUserExit
	}
	return err
}

// InsertUser insert user data into db
func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user(user_id, username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//func CheckUserLogin(p *models.ParamSignIn) (err error) {
//	user, err := QueryUserByUsername(p.Username)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return errors.New("user not exist")
//		}
//		return err
//	}
//	h := md5.New()
//	h.Write([]byte(secret))
//	pwd := hex.EncodeToString(h.Sum([]byte(p.Password)))
//	if pwd != user.Password {
//		return errors.New("password wrong")
//	}
//	return
//}
