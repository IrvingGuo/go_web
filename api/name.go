package api

import (
	"github.com/gin-gonic/gin"
	"go_web/model"
)

func Savename(c *gin.Context) {
	var err error
	var name model.Name
	if err = c.BindJSON(&name); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = name.Save(); err != nil {
		FailWithErr(c, err)
		return
	}
	OkWithData(c, name)
}

func GetAllNames(c *gin.Context) {
	if names, err := model.FindAllNames(); err != nil {
		AbortFindFailed(c, err)
	} else {
		OkWithData(c, names)
	}
}

func GetNameById(c *gin.Context) {

}

func DeleteNameById(c *gin.Context) {

}
