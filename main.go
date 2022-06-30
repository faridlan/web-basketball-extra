package main

import (
	"net/http"

	"github.com/faridlan/web-basketball-extra/app"
	"github.com/faridlan/web-basketball-extra/controller/control_role"
	"github.com/faridlan/web-basketball-extra/repository/repo_role"
	"github.com/faridlan/web-basketball-extra/service/service_role"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDatabase()
	router := httprouter.New()

	roleRepository := repo_role.NewRoleRepository()
	roleService := service_role.NewRoleService(roleRepository, db)
	roleController := control_role.NewRoleController(roleService)

	router.GET("/role", roleController.FindAll)
	router.GET("/role/create", roleController.CreateView)
	router.POST("/role/create", roleController.Create)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
