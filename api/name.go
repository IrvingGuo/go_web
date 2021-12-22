package api

import (
	"go_web/model"
	"go_web/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Savename(c *gin.Context) {

}

func GetAllNames(c *gin.Context) {

}

func GetNameById(c *gin.Context) {
	var nameId int
	var err error
	if nameId, err = strconv.Atoi(c.Param("id")); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var name model.Album
	if name, err = model.FindNameById(uint(nameId)); err != nil {
		AbortFindFailed(c, err)
		return
	}

	OkWithData(c, name)

}

func DeleteNameById(c *gin.Context) {
	var nameId int
	var err error
	if nameId, err = strconv.Atoi(c.Param("name_id")); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = service.DeleteName(uint(nameId)); err != nil {
		FailWithErr(c, err)
		return
	}
	Ok(c)
}
