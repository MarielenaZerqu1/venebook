package controllers

import(
	"github.com/gominima/minima"
	"github.com/hexacry/venebook/utils"
)

func Home(res *minima.Response, req *minima.Request) {
	res.OK().JSON(&utils.Json{
		"msg": "done",
		"code": 200,
	})
}