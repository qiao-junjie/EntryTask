package service

import (
	"entryTask/main/dao"
	"entryTask/main/entity"
)

type UserService struct {
}

var userDao = new(dao.UserDao)

func (p *UserService) Insert(username, password string) int64 {
	//return userDao.Insert(&entity.User{Username:username,Password:password,CreateTime:time.Now()})
	return userDao.Insert(&entity.User{Username: username, Password: password, CreateTime: 0})
}

func (p *UserService) SelectUserByName(username string) []entity.User {
	return userDao.SelectUserByName(username)
}

func (p *UserService) VerifyUserByName(username, password string) entity.Data {
	return userDao.VerifyUserByName(username, password)
}

func (p *UserService) SelectAllUser() []entity.User {
	return userDao.SelectAllUser()
}
