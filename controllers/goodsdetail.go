package controllers

import (
	"fmt"
	"goshop/models"
	"goshop/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type DetailController struct {
	beego.Controller
}

type DetailReturnJSON struct {
	Banners  []models.NideshopAd      `json:"banner"`
	Channels []models.NideshopChannel `json:"channel"`
	Newgoods []orm.Params             `json:"newGoodsList"`
	Hotgoods []orm.Params             `json:"hotGoodsList"`
}

// @Title Get
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *DetailController) Get() {
	var id int
	o.Ctx.Input.Bind(&id, "id") //id ==123

	goodsId := o.GetString("id")

	var banners []models.NideshopAd
	var channels []models.NideshopChannel
	var newgoods []orm.Params
	var hotgoods []orm.Params

	db := orm.NewOrm()
	fmt.Println("111===", goodsId, id)

	ad := new(models.NideshopAd)
	db.QueryTable(ad).Filter("ad_position_id", 1).All(&banners)

	ch := new(models.NideshopChannel)
	db.QueryTable(ch).OrderBy("sort_order").All(&channels)

	ng := new(models.NideshopGoods)
	db.QueryTable(ng).Filter("is_new", 1).Limit(8).Values(&newgoods, "id", "name", "list_pic_url", "retail_price")

	db.QueryTable(ng).Filter("is_hot", 1).Limit(8).Values(&hotgoods, "id", "name", "list_pic_url", "retail_price", "goods_brief")

	utils.ReturnSuccess(&o.Controller, IndexReturnJSON{banners, channels, newgoods, hotgoods})
	o.ServeJSON()
}
