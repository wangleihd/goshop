package utils

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type HttpData struct {
	ErrNo  int         `json: "error"`
	ErrMsg string      `json: "errmsg`
	Data   interface{} `json: data`
}

func ReturnSuccess(this *beego.Controller, val interface{}) {
	retdata := HttpData{
		ErrNo:  0,
		ErrMsg: "Success",
		Data:   val,
	}

	data, err := json.Marshal(retdata)

	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
}

func ReturnError(errno int, errmsg string) interface{} {
	retdata := HttpData{
		ErrNo:  errno,
		ErrMsg: errmsg,
		Data:   nil,
	}
	data, _ := json.Marshal(retdata)
	return json.RawMessage(string(data))
}
