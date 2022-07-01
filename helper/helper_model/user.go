package helper_model

import (
	"github.com/faridlan/web-basketball-extra/model/domain"
	"github.com/faridlan/web-basketball-extra/model/web/web_user"
)

func UserResponse(user domain.User) web_user.UserResponse {
	return web_user.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		RoleId:   user.RoleId,
		StatusId: user.StatusId,
	}
}

func UserResponses(users []domain.User) []web_user.UserResponse {
	userResponse := []web_user.UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, UserResponse(user))
	}

	return userResponse
}
