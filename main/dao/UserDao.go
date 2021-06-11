package dao

import (
	"entryTask/main/cgo"
	"entryTask/main/entity"
	"log"
)

type UserDao struct {
}

func (p *UserDao) Insert(user *entity.User) int64 {
	result, err := cgo.DB.Exec("INSERT INTO user(`username`,`password`,`create_time`) value(?,?,?)", user.Username, user.Password, user.CreateTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (p *UserDao) SelectUserByName(username string) []entity.User {
	rows, err := cgo.DB.Query("SELECT id, username, password, ctime  FROM user_tab WHERE username = ?", username)
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	rows.Close()
	return users
}

func (p *UserDao) SelectAllUser() []entity.User {
	rows, err := cgo.DB.Query("SELECT * FROM user")
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreateTime)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	rows.Close()
	return users
}

func (p *UserDao) VerifyUserByName(username, password string) entity.Data {
	rows, err := cgo.DB.Query("SELECT id, email, profile FROM user_tab WHERE username = ? and password = ?", username, password)
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User1
	for rows.Next() {
		var user entity.User1
		err := rows.Scan(&user.ID, &user.Email, &user.Profile)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	rows.Close()
	return users
}
