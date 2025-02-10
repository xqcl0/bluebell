package mysql

import (
	"bluebell/models"
	"bluebell/settings"

	"testing"
)

func init() {
	dbCfg := &settings.MySQLConfig{
		Host:         "117.72.13.155",
		User:         "root",
		Password:     "?",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 100,
		MaxIdleConns: 10,
	}
	err := Init(dbCfg)
	if err != nil {
		panic(err)
	}
	return
}
func TestInsertUser(t *testing.T) {

	user := &models.User{
		ID:       1100,
		UserID:   11111,
		Username: "testUser",
		Password: "??????",
	}
	err := InsertUser(user)
	if err != nil {
		t.Fatalf("insert user test err: %v\n", err)
	}
	t.Logf("insert user test success")
}
