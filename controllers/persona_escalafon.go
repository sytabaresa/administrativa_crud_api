package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_crud_api/models"
)

type PersonaEscalafonController struct {
	beego.Controller
}

func (c *PersonaEscalafonController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}

func (c *PersonaEscalafonController) GetAll() {
	listaPersonas := models.GetAllPersonaEscalafon()
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = listaPersonas
	c.ServeJSON()
}
