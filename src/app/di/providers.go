package di

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"git.tashilcar.com/core/tashilcar/core/configs"
	errs "git.tashilcar.com/core/tashilcar/core/yerrors"
)

func provideDB() *gorm.DB {
	dbConnection := configs.Env("DB_CONNECTION")
	var db *gorm.DB
	var err error
	switch dbConnection {
	case "mysql":
		db, err = gorm.Open(
			mysql.Open(
				configs.Env("DB_USERNAME")+":"+configs.Env("DB_PASSWORD")+
					"@tcp("+configs.Env("DB_HOST")+":"+configs.Env("DB_PORT")+")"+
					"/"+configs.Env("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local",
			), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
				SkipDefaultTransaction: true,
				PrepareStmt:            true,
			},
		)
		if errs.IsNotNil(err) {
			panic("failed to connect database")
		}
		break
	case "postgres":
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
		if errs.IsNotNil(err) {
			panic("failed to connect database")
		}
		break
	default:
		panic("please choose valid db name")
	}
	return db
}
