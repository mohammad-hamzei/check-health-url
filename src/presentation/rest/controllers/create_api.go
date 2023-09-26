package controllers

import (
	"git.tashilcar.com/core/tashilcar/app/di"
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateAPI @Summary	add a new item to the entities.Tashilcar list
// @ID			create-entities.Tashilcar
// @Produce	json
// @Param		data	body		entities.Tashilcar	true	"entities.Tashilcar data"
// @Success	200		{object}	entities.Tashilcar
// @Failure	400		{object}	error
// @Router		/url [post]
func CreateAPI(c *gin.Context) {

	tashilcar := entities.Tashilcar{}

	if err := c.BindJSON(&tashilcar); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	url := entities.Tashilcar{
		HealthCheck:             tashilcar.HealthCheck,
		HealthCheckTimeInterval: tashilcar.HealthCheckTimeInterval,
		RequestURL:              tashilcar.RequestURL,
		RequestHTTPMethod:       tashilcar.RequestHTTPMethod,
		RequestHeaders:          tashilcar.RequestHeaders,
		RequestBody:             tashilcar.RequestBody,
		ResponseStatus:          tashilcar.ResponseStatus,
	}

	//cbuc := usecases.NewCreateAPI(repositories.NewTashilcarRepository(databases.NewDBDataSource(db)))
	cbuc := di.CreateAPIInstance()
	error := cbuc.Exec(url)
	if error != nil {
		c.String(200, "create is failed!")
	} else {
		c.String(200, "create is successful!")
	}

}
