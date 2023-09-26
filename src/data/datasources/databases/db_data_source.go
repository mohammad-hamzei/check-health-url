package databases

import (
	"gorm.io/gorm"

	"git.tashilcar.com/core/tashilcar/data/datasources/databases/postgres"
	"git.tashilcar.com/core/tashilcar/data/models"
)

type DBDataSource interface {
	AllModels() []interface{}
	CreateTables() error
	Insert(tashilcarModel models.TashilcarModel) error
	Get() ([]*models.TashilcarModel, error)
	EnableCheckHealth(id uint64, enable bool) error
	DeleteAPI(id uint64) error
	UpdateResponseStatus(id uint64, status string) error
}

func NewDBDataSource(db *gorm.DB) DBDataSource {
	if db == nil {
		panic("db instance should not be nil")
	}
	return &postgres.Postgres{
		DB: db,
	}
}
