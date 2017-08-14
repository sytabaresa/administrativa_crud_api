package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNeceisdadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNeceisdadController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNeceisdadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNeceisdadController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNeceisdadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNeceisdadController"],
			beego.ControllerComments{
				Method:           "Put",
				Router:           `/:id`,
				AllowHTTPMethods: []string{"put"},
				Params:           nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/:idResolucion`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
			beego.ControllerComments{
				Method: "GetAllContratado",
				Router: `Contratado/:idResolucion`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:idResolucion/:idPersona`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PersonaEscalafonController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
			beego.ControllerComments{
				Method: "InsertarVinculaciones",
				Router: `InsertarVinculaciones`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:idResolucion`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
			beego.ControllerComments{
				Method: "Put",
				Router: `/:idResolucion`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
			beego.ControllerComments{
				Method: "ResolucionTemplate",
				Router: `ResolucionTemplate`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
			beego.ControllerComments{
				Method: "CancelarResolucion",
				Router: `CancelarResolucion/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
			beego.ControllerComments{
				Method: "InsertarContratos",
				Router: `InsertarContratos`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
			beego.ControllerComments{
				Method: "GenerarResolucion",
				Router: `GenerarResolucion`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
			beego.ControllerComments{
				Method: "RestaurarResolucion",
				Router: `RestaurarResolucion/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		}
