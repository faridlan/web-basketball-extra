package control_achievment

import (
	"net/http"

	"github.com/faridlan/web-basketball-extra/service/service_achievment"
	"github.com/julienschmidt/httprouter"
)

type AchievmentControllerImpl struct {
	ServiceAchiev service_achievment.AchievmentService
}

func (controller *AchievmentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *AchievmentControllerImpl) CreateView(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *AchievmentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *AchievmentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *AchievmentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *AchievmentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
