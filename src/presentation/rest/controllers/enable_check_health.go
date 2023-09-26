package controllers

import (
	"git.tashilcar.com/core/tashilcar/app/di"
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary	enable cheack health a entities.Tashilcar item by ID
// @ID			enable-check-health-entities.Tashilcar-by-id
// @Produce	json
// @Param		id	path		string	true	"entities.Tashilcar.ID ID"
// @Success	200	{object}	error
// @Failure	404	{object}	error
// @Router		/url/:id [put]
func EnableCheckHealth(c *gin.Context) {
	tashilcar := entities.Tashilcar{}

	if err := c.BindJSON(&tashilcar); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var getId string = c.Param("id")
	id, er := strconv.ParseInt(getId, 10, 64)
	if er != nil {
		panic(er)
	}
	enable := tashilcar.HealthCheck

	cbuc := di.EnableCheckHealthInstance()
	error := cbuc.Exec(uint64(id), enable)
	if error != nil {
		c.String(200, "change status `check health` is failed!")
	} else {
		c.String(200, "change status `check health` is successful!")
	}
}
