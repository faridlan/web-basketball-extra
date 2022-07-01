package web_user

type UserCreateReq struct {
	Username string
	Email    string
	Password string
	RoleId   int
	StatusId int
}
