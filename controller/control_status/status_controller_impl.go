package control_status

import (
	"net/http"

	"github.com/faridlan/web-basketball-extra/service/service_status"
	"github.com/julienschmidt/httprouter"
)

type StatusControllerImpl struct {
	ServiceStatus service_status.StatusService
}

func (controller *StatusControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *StatusControllerImpl) CreateView(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *StatusControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *StatusControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *StatusControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *StatusControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
