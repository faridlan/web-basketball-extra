package service_role

import (
	"context"

	"github.com/faridlan/web-basketball-extra/model/web/web_role"
)

type RoleService interface {
	Create(ctx context.Context, request web_role.RoleCreateReq) int64
	Update(ctx context.Context, request web_role.RoleUpdateReq)
	Delete(ctx context.Context, roleId int)
	FindAll(ctx context.Context) []web_role.RoleResponse
}
