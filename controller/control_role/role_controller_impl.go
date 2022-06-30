package control_role

import (
	"html/template"
	"net/http"

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
	panic("not implemented") // TODO: Implement
}

func (controller *RoleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *RoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleResponses := controller.RoleService.FindAll(request.Context())
	tmpl.ExecuteTemplate(writer, "View", roleResponses)
}
