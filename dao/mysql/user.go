package mysql

import (
	"BlueBell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

const secret = "woshi2guotou"

var (
	ErrorUserExist       = errors.New("user Exists")
	ErrorUserNotExist    = errors.New("user does not exist")
	ErrorInvalidPassword = errors.New("username or password error")
)

func IfUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username =?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	//encrypt
	password := encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return
}

func encryptPassword(originalPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(originalPassword)))
}

func Login(user *models.User) (err error) {
	originPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(originPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserByID(idStr string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, idStr)
	if err != nil {
		zap.L().Error("GetUserByID() failed", zap.String("sql", sqlStr), zap.Error(err))
		return
	}
	return
}
