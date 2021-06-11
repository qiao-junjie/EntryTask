package entity

type User struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	CreateTime int    `json:"create_time"`
}

type CommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User1 struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}

type Data []User1

type UserResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data1   Data   `json:"data"`
}
