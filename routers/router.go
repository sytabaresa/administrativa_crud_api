// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/administrativa_crud_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),

		beego.NSNamespace("/vigencia_contrato",
			beego.NSInclude(
				&controllers.VigenciaContratoController{},
			),
		),

		beego.NSNamespace("/contrato_disponibilidad",
			beego.NSInclude(
				&controllers.ContratoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/detalle_servicio_necesidad",
			beego.NSInclude(
				&controllers.DetalleServicioNecesidadController{},
			),
		),

		beego.NSNamespace("/tipo_resolucion",
			beego.NSInclude(
				&controllers.TipoResolucionController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/fuente_financiacion_rubro_necesidad",
			beego.NSInclude(
				&controllers.FuenteFinanciacionRubroNecesidadController{},
			),
		),

		beego.NSNamespace("/marco_legal_necesidad",
			beego.NSInclude(
				&controllers.MarcoLegalNecesidadController{},
			),
		),

		beego.NSNamespace("/marco_legal",
			beego.NSInclude(
				&controllers.MarcoLegalController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion_docente",
			beego.NSInclude(
				&controllers.ResolucionVinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/catalogo_elemento",
			beego.NSInclude(
				&controllers.CatalogoElementoController{},
			),
		),

		beego.NSNamespace("/catalogo_elemento_grupo",
			beego.NSInclude(
				&controllers.CatalogoElementoGrupoController{},
			),
		),

		beego.NSNamespace("/componente_resolucion",
			beego.NSInclude(
				&controllers.ComponenteResolucionController{},
			),
		),

		beego.NSNamespace("/disponibilidad_apropiacion_solicitud_rp",
			beego.NSInclude(
				&controllers.DisponibilidadApropiacionSolicitudRpController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/relacion_parametro",
			beego.NSInclude(
				&controllers.RelacionParametroController{},
			),
		),

		beego.NSNamespace("/tipo_financiacion_necesidad",
			beego.NSInclude(
				&controllers.TipoFinanciacionNecesidadController{},
			),
		),

		beego.NSNamespace("/actividad_economica_necesidad",
			beego.NSInclude(
				&controllers.ActividadEconomicaNecesidadController{},
			),
		),

		beego.NSNamespace("/especificacion_tecnica",
			beego.NSInclude(
				&controllers.EspecificacionTecnicaController{},
			),
		),

		beego.NSNamespace("/tipo_contrato_necesidad",
			beego.NSInclude(
				&controllers.TipoContratoNecesidadController{},
			),
		),

		beego.NSNamespace("/actividad_especifica",
			beego.NSInclude(
				&controllers.ActividadEspecificaController{},
			),
		),

		beego.NSNamespace("/estado_necesidad",
			beego.NSInclude(
				&controllers.EstadoNecesidadController{},
			),
		),

		beego.NSNamespace("/modalidad_seleccion",
			beego.NSInclude(
				&controllers.ModalidadSeleccionController{},
			),
		),

		beego.NSNamespace("/requisito_minimo",
			beego.NSInclude(
				&controllers.RequisitoMinimoController{},
			),
		),

		beego.NSNamespace("/lugar_ejecucion",
			beego.NSInclude(
				&controllers.LugarEjecucionController{},
			),
		),

		beego.NSNamespace("/supervisor_contrato",
			beego.NSInclude(
				&controllers.SupervisorContratoController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/acta_inicio",
			beego.NSInclude(
				&controllers.ActaInicioController{},
			),
		),

		beego.NSNamespace("/solicitud_disponibilidad",
			beego.NSInclude(
				&controllers.SolicitudDisponibilidadController{},
			),
		),

		beego.NSNamespace("/necesidad",
			beego.NSInclude(
				&controllers.NecesidadController{},
			),
		),

		beego.NSNamespace("/dependencia_necesidad",
			beego.NSInclude(
				&controllers.DependenciaNecesidadController{},
			),
		),

		beego.NSNamespace("/estado_contrato",
			beego.NSInclude(
				&controllers.EstadoContratoController{},
			),
		),

		beego.NSNamespace("/contrato_estado",
			beego.NSInclude(
				&controllers.ContratoEstadoController{},
			),
		),

		beego.NSNamespace("/parametros",
			beego.NSInclude(
				&controllers.ParametrosController{},
			),
		),

		beego.NSNamespace("/poliza",
			beego.NSInclude(
				&controllers.PolizaController{},
			),
		),

		beego.NSNamespace("/tr_necesidad",
			beego.NSInclude(
				&controllers.TrNecesidadController{},
			),
		),

		beego.NSNamespace("/persona_escalafon",
			beego.NSInclude(
				&controllers.PersonaEscalafonController{},
			),
		),

		beego.NSNamespace("/estado_resolucion",
			beego.NSInclude(
				&controllers.EstadoResolucionController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/precontratado",
			beego.NSInclude(
				&controllers.PrecontratadoController{},
			),
		),

		beego.NSNamespace("/solicitud_rp",
			beego.NSInclude(
				&controllers.SolicitudRpController{},
			),
		),

		beego.NSNamespace("/resolucion_estado",
			beego.NSInclude(
				&controllers.ResolucionEstadoController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion",
			beego.NSInclude(
				&controllers.ResolucionVinculacionController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/tipo_necesidad",
			beego.NSInclude(
				&controllers.TipoNecesidadController{},
			),
		),

		beego.NSNamespace("/necesidad_proceso_externo",
			beego.NSInclude(
				&controllers.NecesidadProcesoExternoController{},
			),
		),

		beego.NSNamespace("/contenido_resolucion",
			beego.NSInclude(
				&controllers.ResolucionCompletaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
