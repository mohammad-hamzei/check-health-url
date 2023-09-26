package controllers

import (
	"git.tashilcar.com/core/tashilcar/app/di"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get all items in the entities.Tashilcar list
// @ID get-all-entities.Tashilcar
// @Produce json
// @Success 200 {object} []entities.Tashilcar
// @Router /urls [get]
func GetAPIs(c *gin.Context) {
	/*var db *gorm.DB
	var err error
	db, err = gorm.Open(
		postgres.Open(
			"host="+configs.Env("DB_HOST")+
				" user="+configs.Env("DB_USERNAME")+
				" password="+configs.Env("DB_PASSWORD")+
				" dbname="+configs.Env("DB_DATABASE")+
				" port="+configs.Env("DB_PORT")+
				" sslmode=disable",
		), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		},
	)
	if err != nil {
		c.String(400, "errrrr")
	}*/
	//cbuc := usecases.NewGetAPIs(repositories.NewTashilcarRepository(databases.NewDBDataSource(db)))
	cbuc := di.GetAPIsInstance()
	urls, error := cbuc.Exec()
	if error != nil {
		c.JSON(http.StatusOK, gin.H{"data": error})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": urls})
	}
}
