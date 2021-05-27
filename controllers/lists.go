package controllers

import (
	"goshop/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ListsController struct {
	beego.Controller
}

// @Title Get
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ListsController) Get() {
	obs := models.GetIndex()
	o.Data["json"] = obs
	o.ServeJSON()
}
