package controllers

import (
	"goshop/models"
	"goshop/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type IndexController struct {
	beego.Controller
}

type IndexReturnJSON struct {
	Banners  []models.NideshopAd      `json:"banner"`
	Channels []models.NideshopChannel `json:"channel"`
	Newgoods []orm.Params             `json:"newGoodsList"`
	Hotgoods []orm.Params             `json:"hotGoodsList"`
}

func updateJsonKeysIndex(vals []orm.Params) {
	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicUrl":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			case "GoodsBrief":
				delete(val, k)
				val["goods_brief"] = v
			}
		}
	}
}

// @Title Get
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *IndexController) Get() {
	var banners []models.NideshopAd
	var channels []models.NideshopChannel
	var newgoods []orm.Params
	var hotgoods []orm.Params

	db := orm.NewOrm()

	ad := new(models.NideshopAd)
	db.QueryTable(ad).Filter("ad_position_id", 1).All(&banners)

	ch := new(models.NideshopChannel)
	db.QueryTable(ch).OrderBy("sort_order").All(&channels)

	ng := new(models.NideshopGoods)
	db.QueryTable(ng).Filter("is_new", 1).Limit(8).Values(&newgoods, "id", "name", "list_pic_url", "retail_price")

	db.QueryTable(ng).Filter("is_hot", 1).Limit(8).Values(&hotgoods, "id", "name", "list_pic_url", "retail_price", "goods_brief")
	updateJsonKeysIndex(hotgoods)

	utils.ReturnSuccess(&o.Controller, IndexReturnJSON{banners, channels, newgoods, hotgoods})
	o.ServeJSON()
}
