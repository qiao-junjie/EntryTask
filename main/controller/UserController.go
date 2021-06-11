package controller

import (
	"entryTask/main/cgo"
	"entryTask/main/constant"
	"entryTask/main/entity"
	"entryTask/main/service"
	"entryTask/main/utils"
	"log"
	"net/http"
)

type UserConterller struct {
}

var userService = new(service.UserService)

func (p *UserConterller) Router(router *cgo.RouterHandler) {
	router.Router("/register", p.register)
	router.Router("/login", p.login)
	router.Router("/findAll", p.findAll)
	router.Router("/user_login", p.userLogin)
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserConterller) register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		cgo.ResultFail(w, "username or password can not be empty")
		return
	}
	id := userService.Insert(username, password)
	if id <= 0 {
		cgo.ResultFail(w, "register fail")
		return
	}
	cgo.ResultOk(w, "register success")
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserConterller) login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		cgo.ResultFail(w, "username or password can not be empty")
		return
	}
	users := userService.SelectUserByName(username)
	if len(users) == 0 {
		cgo.ResultFail(w, "user does not exist")
		return
	}
	if users[0].Password != password {
		cgo.ResultFail(w, "password error")
		return
	}

	//session
	session := cgo.GlobalSession().SessionStart(w, r)
	session.Set(constant.KEY_USER, &users[0])
	cgo.ResultOk(w, "login success")
}

// GET/POST
func (p *UserConterller) findAll(w http.ResponseWriter, r *http.Request) {
	coikie, err := r.Cookie("GSESSION")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(coikie.Value)
	}
	users := userService.SelectAllUser()
	cgo.ResultJsonOk(w, users)
}

func (p *UserConterller) userLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		cgo.ResultFail(w, "username or password can not be empty")
		return
	}
	users := userService.VerifyUserByName(username, password)
	if len(users) == 0 {
		cgo.ResultFail(w, "user does not exist")
		return
	}

	userResp := entity.UserResp{200, " ", users}

	cgo.ResultJsonOk(w, userResp)
}
