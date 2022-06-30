package control_role

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/faridlan/web-basketball-extra/helper"
	"github.com/faridlan/web-basketball-extra/model/web/web_role"
	"github.com/faridlan/web-basketball-extra/service/service_role"
	"github.com/julienschmidt/httprouter"
)

var tmpl = template.Must(template.Must(template.ParseGlob("templates/*.html")).ParseGlob("templates/roles/*.html"))

type RoleControllerImpl struct {
	RoleService service_role.RoleService
}

func NewRoleController(service service_role.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: service,
	}
}

func (controller *RoleControllerImpl) CreateView(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tmpl.ExecuteTemplate(writer, "Create", nil)
}

func (controller *RoleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleCreate := &web_role.RoleCreateReq{}

	roleCreate.Name = request.FormValue("name")

	i := controller.RoleService.Create(request.Context(), *roleCreate)

	if i == 1 {
		// http.Redirect(writer, request, "/role", http.StatusMovedPermanently)
		tmpl.ExecuteTemplate(writer, "Create", map[string]interface{}{
			"Alert": template.HTML(`<script>
			alert("BERHASIL");
			window.location.href='/role';
			</script>`),
		})
	} else {
		tmpl.ExecuteTemplate(writer, "Create", map[string]interface{}{
			"Alert": template.HTML(`<script>
			alert("GAGAL");
			</script>`),
		})
	}

}

func (controller *RoleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleUpdate := &web_role.RoleUpdateReq{}

	roleUpdate.Name = request.FormValue("name")
	Id, err := strconv.ParseInt(request.FormValue("id")[0:], 10, 64)
	helper.PanicIfError(err)
	roleUpdate.Id = Id
	i := controller.RoleService.Update(request.Context(), *roleUpdate)

	if i == 1 {
		// http.Redirect(writer, request, "/role", http.StatusMovedPermanently)
		tmpl.ExecuteTemplate(writer, "Update", map[string]interface{}{
			"Alert": template.HTML(`<script>
			alert("BERHASIL");
			window.location.href='/role';
			</script>`),
		})
	} else {
		tmpl.ExecuteTemplate(writer, "Update", map[string]interface{}{
			"Alert": template.HTML(`<script>
			alert("GAGAL");
			</script>`),
		})
	}
}

func (controller *RoleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := request.URL.Query().Get("id")
	Id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	controller.RoleService.Delete(request.Context(), Id)
	http.Redirect(writer, request, "/role", http.StatusMovedPermanently)
}

func (controller *RoleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := request.URL.Query().Get("id")
	Id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)
	roleResponse := controller.RoleService.FindById(request.Context(), Id)

	x := map[string]interface{}{
		"Name": roleResponse.Name,
		"Id":   roleResponse.Id,
	}
	tmpl.ExecuteTemplate(writer, "Update", x)
	log.Print(roleResponse)
}

func (controller *RoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleResponses := controller.RoleService.FindAll(request.Context())
	tmpl.ExecuteTemplate(writer, "View", roleResponses)
}
