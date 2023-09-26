package postgres

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"git.tashilcar.com/core/tashilcar/core/configs"
	errs "git.tashilcar.com/core/tashilcar/core/yerrors"
	"git.tashilcar.com/core/tashilcar/data/datasources/databases/postgres/db_models"
	"git.tashilcar.com/core/tashilcar/data/models"
)

type Postgres struct {
	DB *gorm.DB
}

func (p *Postgres) AllModels() []interface{} {
	return []interface{}{
		(*db_models.Tashilcar)(nil),
	}
}
func (p *Postgres) CreateTables() error {
	const op errs.Op = "data_sources.createTables"
	for _, model := range p.AllModels() {
		productionMode, _ := strconv.ParseBool(configs.Env("PRODUCTION_MODE"))
		if !productionMode {
			if p.DB.Migrator().HasTable(model) {
				dropErr := p.DB.Debug().Migrator().DropTable(model)
				if dropErr != nil {
					return errs.E(op, errs.KindInternal, dropErr)
				}
			}
		}
		migrationErr := p.DB.Migrator().AutoMigrate(model)
		if errs.IsNotNil(migrationErr) {
			return errs.E(op, errs.KindInternal, migrationErr)
		}
	}
	return nil
}

func (p *Postgres) Insert(tashilcarModel models.TashilcarModel) error {
	const op errs.Op = "postgres.Insert"
	/* create instances from tashilcar db models */
	dbTC := &db_models.Tashilcar{
		HealthCheck:             tashilcarModel.HealthCheck,
		HealthCheckTimeInterval: tashilcarModel.HealthCheckTimeInterval,
		RequestURL:              tashilcarModel.RequestURL,
		RequestHTTPMethod:       tashilcarModel.RequestHTTPMethod,
		RequestHeaders:          tashilcarModel.RequestHeaders,
		RequestBody:             tashilcarModel.RequestBody,
	}
	/* save TC in db */
	dbErr := p.DB.WithContext(context.Background()).Create(dbTC).Error
	if errs.IsNotNil(dbErr) {
		return errs.E(op, errors.New("error while setting TC record"), errs.LevelError, errs.KindInternal)
	}
	return nil
}

func (p *Postgres) Get() ([]*models.TashilcarModel, error) {
	const op errs.Op = "postgres.Get"

	var urls []*models.TashilcarModel
	dbErr := p.DB.WithContext(context.Background()).Table("tashilcar").Find(&urls).Error
	if errs.IsNotNil(dbErr) {
		return nil, errs.E(op, errors.New("error while setting TC record"), errs.LevelError, errs.KindInternal)
	}
	tashilcarModel := make([]*models.TashilcarModel, len(urls))
	for i, url := range urls {
		u := &models.TashilcarModel{
			ID:                      url.ID,
			HealthCheck:             url.HealthCheck,
			HealthCheckTimeInterval: url.HealthCheckTimeInterval,
			RequestURL:              url.RequestURL,
			RequestHTTPMethod:       url.RequestHTTPMethod,
			RequestHeaders:          url.RequestHeaders,
			RequestBody:             url.RequestBody,
			ResponseStatus:          url.ResponseStatus,
			CreatedAt:               url.CreatedAt,
			UpdatedAt:               url.UpdatedAt,
		}
		tashilcarModel[i] = u
	}
	return tashilcarModel, nil
}

func (p *Postgres) EnableCheckHealth(id uint64, enable bool) error {
	const op errs.Op = "postgres.EnableCheckHealth"
	err := p.DB.WithContext(context.Background()).Table("tashilcar").Where("id = ?", id).Updates(map[string]interface{}{"health_check": enable}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (p *Postgres) DeleteAPI(id uint64) error {
	const op errs.Op = "postgres.DeleteAPI"
	err := p.DB.WithContext(context.Background()).Table("tashilcar").Where("id = ?", id).Delete(&models.TashilcarModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UpdateResponseStatus(id uint64, status uint16) error {
	const op errs.Op = "postgres.EnableCheckHealth"
	err := p.DB.WithContext(context.Background()).Table("tashilcar").Where("id = ?", id).Updates(
		&models.TashilcarModel{
			ResponseStatus: status,
		}).Error
	if err != nil {
		return err
	}
	return nil
}
