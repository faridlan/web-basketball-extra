package helper_model

import (
	"github.com/faridlan/web-basketball-extra/model/domain"
	"github.com/faridlan/web-basketball-extra/model/web/web_role"
)

func RoleResponse(role domain.Role) web_role.RoleResponse {
	return web_role.RoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}
}

func RoleResponses(roles []domain.Role) []web_role.RoleResponse {
	roleResponses := []web_role.RoleResponse{}
	for _, role := range roles {
		roleResponses = append(roleResponses, RoleResponse(role))
	}

	return roleResponses
}
