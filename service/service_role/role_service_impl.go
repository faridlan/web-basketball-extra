package service_role

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/helper"
	"github.com/faridlan/web-basketball-extra/helper/helper_model"
	"github.com/faridlan/web-basketball-extra/model/domain"
	"github.com/faridlan/web-basketball-extra/model/web/web_role"
	"github.com/faridlan/web-basketball-extra/repository/repo_role"
)

type RoleServiceImpl struct {
	RepoRole repo_role.RoleRepository
	DB       *sql.DB
}

func NewRoleService(RepoRole repo_role.RoleRepository, DB *sql.DB) RoleService {
	return RoleServiceImpl{
		RepoRole: RepoRole,
		DB:       DB,
	}
}

func (service RoleServiceImpl) Create(ctx context.Context, request web_role.RoleCreateReq) int64 {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.Role{
		Name: request.Name,
	}

	i := service.RepoRole.Save(ctx, tx, role)
	return i
}

func (service RoleServiceImpl) Update(ctx context.Context, request web_role.RoleUpdateReq) {
	panic("not implemented") // TODO: Implement
}

func (service RoleServiceImpl) Delete(ctx context.Context, roleId int) {
	panic("not implemented") // TODO: Implement
}

func (service RoleServiceImpl) FindAll(ctx context.Context) []web_role.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roleResponse := service.RepoRole.FindAll(ctx, tx)
	return helper_model.RoleResponses(roleResponse)
}
