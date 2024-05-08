package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/jwt"
	"BlueBell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// If user exists
	if err := mysql.IfUserExist(p.Username); err != nil {
		return err
	}
	userID, _ := snowflake.GetID()
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (atoken, rtoken string, err error) {
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", "", err
	}
	return jwt.GenToken(user.UserID)
}
