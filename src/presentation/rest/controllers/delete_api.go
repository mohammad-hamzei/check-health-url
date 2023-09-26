package controllers

import (
	"git.tashilcar.com/core/tashilcar/app/di"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary	delete a entities.Tashilcar item by ID
// @ID			delete-entities.Tashilcar-by-id
// @Produce	json
// @Param		id	path		string	true	"entities.Tashilcar.ID ID"
// @Success	200	{object}	error
// @Failure	404	{object}	error
// @Router		/url/:id [delete]
func DeleteAPI(c *gin.Context) {
	var getId string = c.Param("id")
	id, er := strconv.ParseInt(getId, 10, 64)
	if er != nil {
		panic(er)
	}
	cbuc := di.DeleteAPIInstance()
	error := cbuc.Exec(uint64(id))
	if error != nil {
		c.String(200, "Deleted is failed!")
	} else {
		c.String(200, "Deleted is successful!")
	}
}
