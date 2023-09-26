package repositories

import (
	"git.tashilcar.com/core/tashilcar/core/yerrors"
	"git.tashilcar.com/core/tashilcar/data/datasources/databases"
	"git.tashilcar.com/core/tashilcar/data/models"
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"git.tashilcar.com/core/tashilcar/domain/repositories"
)

//	func NewTashilcarRepository(dbDS databases.DBDataSource) repositories.TashilcarRepository {
//		return &TashilcarRepository{
//			databaseDS: dbDS,
//		}
//	}
//
//	type TashilcarRepository struct {
//		databaseDS databases.DBDataSource
//	}
//
//	func (g *TashilcarRepository) Insert(tashilcar entities.Tashilcar) error {
//		const op errs.Op = "tashilcar_repository.Insert"
//
//		/* call insert function from database datasource */
//		err := g.databaseDS.Insert(models.TashilcarModel{
//			HealthCheck: false,
//			HealthCheckTimeInterval: 3,
//			RequestURL: "google.com",
//			RequestHTTPMethod: "POST",
//			RequestHeaders: "heasres",
//			RequestBody: "test",
//		})
//		/* if has error return errs.E() */
//		if errs.IsNotNil(err) {
//			return errs.E(op, err)
//		}
//		return nil
//	}
func NewTashilcarRepository(dbDS databases.DBDataSource) repositories.TashilcarRepository {
	return &tashilcarRepository{
		databaseDS: dbDS,
	}
}

type tashilcarRepository struct {
	databaseDS databases.DBDataSource
}

func (t *tashilcarRepository) Insert(tashilcar *entities.Tashilcar) error {
	const op yerrors.Op = "tashilcar_repository.Insert"
	err := t.databaseDS.Insert(models.TashilcarModel{
		HealthCheck:             tashilcar.HealthCheck,
		HealthCheckTimeInterval: tashilcar.HealthCheckTimeInterval,
		RequestURL:              tashilcar.RequestURL,
		RequestHTTPMethod:       tashilcar.RequestHTTPMethod,
		RequestHeaders:          tashilcar.RequestHeaders,
		RequestBody:             tashilcar.RequestBody,
		ResponseStatus:          tashilcar.ResponseStatus,
		CreatedAt:               tashilcar.CreatedAt,
		UpdatedAt:               tashilcar.UpdatedAt,
	})
	if yerrors.IsNotNil(err) {
		return yerrors.E(op, err)
	}
	return nil
}

func (t *tashilcarRepository) Get() ([]*entities.Tashilcar, error) {
	const op yerrors.Op = "tashilcar.data.repository.Get"
	urls, err := t.databaseDS.Get()
	if yerrors.IsNotNil(err) {
		return nil, yerrors.E(op, err)
	}
	tashilcarModel := make([]*entities.Tashilcar, len(urls))
	for i, url := range urls {
		ur := &entities.Tashilcar{
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
		tashilcarModel[i] = ur
	}
	return tashilcarModel, nil
}

func (t *tashilcarRepository) EnableCheckHealth(id uint64, enable bool) error {
	const op yerrors.Op = "tashilcar_repository.EnableCheckHealth"
	err := t.databaseDS.EnableCheckHealth(id, enable)
	if yerrors.IsNotNil(err) {
		return yerrors.E(op, err)
	}
	return nil
}

func (t *tashilcarRepository) DeleteAPI(id uint64) error {
	const op yerrors.Op = "tashilcar_repository.DeleteAPI"
	err := t.databaseDS.DeleteAPI(id)
	if yerrors.IsNotNil(err) {
		return yerrors.E(op, err)
	}
	return nil
}

func (t *tashilcarRepository) UpdateResponseStatus(id uint64, status string) error {
	const op yerrors.Op = "tashilcar_repository.UpdateResponseStatus"
	err := t.databaseDS.UpdateResponseStatus(id, status)
	if yerrors.IsNotNil(err) {
		return yerrors.E(op, err)
	}
	return nil
}
