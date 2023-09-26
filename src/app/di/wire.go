//go:build wireinject
// +build wireinject

package di

import (
	"git.tashilcar.com/core/tashilcar/data/datasources/databases"
	"git.tashilcar.com/core/tashilcar/data/repositories"
	"git.tashilcar.com/core/tashilcar/domain/usecases"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	//	Third-Party
	db = wire.NewSet(provideDB)

	//  DataSources
	dbDS = wire.NewSet(databases.NewDBDataSource, db)

	// Repositories
	repository = wire.NewSet(
		repositories.NewTashilcarRepository,
		dbDS,
	)
)

func DB() *gorm.DB {
	wire.Build(db)
	return nil
}

func DBDS() databases.DBDataSource {
	wire.Build(dbDS)
	return nil
}

func GetAPIsInstance() usecases.GetAPIs {
	wire.Build(
		usecases.NewGetAPIs,
		repositories.NewTashilcarRepository,
		databases.NewDBDataSource,
		db,
	)
	return nil
}

func CreateAPIInstance() usecases.CreateAPI {
	wire.Build(
		usecases.NewCreateAPI,
		repositories.NewTashilcarRepository,
		databases.NewDBDataSource,
		db,
	)
	return nil
}

func DeleteAPIInstance() usecases.DeleteAPI {
	wire.Build(
		usecases.NewDeleteAPI,
		repositories.NewTashilcarRepository,
		databases.NewDBDataSource,
		db,
	)
	return nil
}

func EnableCheckHealthInstance() usecases.EnableCheckHealth {
	wire.Build(
		usecases.NewEnableCheckHealth,
		repositories.NewTashilcarRepository,
		databases.NewDBDataSource,
		db,
	)
	return nil
}

func UpdateResponseStatusInstance() usecases.UpdateResponseStatus {
	wire.Build(
		usecases.NewUpdateResponseStatus,
		repositories.NewTashilcarRepository,
		databases.NewDBDataSource,
		db,
	)
	return nil
}
