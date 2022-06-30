package service_role

import (
	"context"
	"database/sql"
	"log"

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

func (service RoleServiceImpl) Update(ctx context.Context, request web_role.RoleUpdateReq) int64 {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.Role{
		Id:   request.Id,
		Name: request.Name,
	}

	i := service.RepoRole.Update(ctx, tx, role)
	return i
}

func (service RoleServiceImpl) Delete(ctx context.Context, roleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.RepoRole.Delete(ctx, tx, roleId)
}

func (service RoleServiceImpl) FindById(ctx context.Context, roleId int) web_role.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roleResponse := service.RepoRole.FindId(ctx, tx, roleId)
	log.Print(roleResponse)
	return helper_model.RoleResponse(roleResponse)
}

func (service RoleServiceImpl) FindAll(ctx context.Context) []web_role.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roleResponse := service.RepoRole.FindAll(ctx, tx)
	return helper_model.RoleResponses(roleResponse)
}
