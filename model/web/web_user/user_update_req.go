package web_user

type UserUpdateReq struct {
	Id       int
	Username string
	Email    string
	Password string
	RoleId   int
	StatusId int
}
