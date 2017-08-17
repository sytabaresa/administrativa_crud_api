package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "InsertarContratos",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion/:idPersona`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAllContratado",
			Router: `/Contratado/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "ResolucionTemplate",
			Router: `/ResolucionTemplate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GenerarResolucion",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "CancelarResolucion",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "RestaurarResolucion",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VigenciaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VigenciaContratoController"],
		beego.ControllerComments{
			Method: "VigenciaContrato",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "InsertarVinculaciones",
			Router: `/InsertarVinculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
