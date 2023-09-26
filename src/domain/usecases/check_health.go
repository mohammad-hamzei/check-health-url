package usecases

import (
	"fmt"
	"git.tashilcar.com/core/tashilcar/core/configs"
	"git.tashilcar.com/core/tashilcar/data/datasources/databases"
	"git.tashilcar.com/core/tashilcar/data/repositories"
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"sync"
	"time"
)

func CheckHealth() {
	urls, error := getUrls()
	if error != nil {
		fmt.Println(error)
		return
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		healthCheck := url.HealthCheck
		if healthCheck {
			updatedAt := url.UpdatedAt
			interval := url.HealthCheckTimeInterval
			totalTime := updatedAt + (interval * 60)
			now := time.Now().Unix()
			if totalTime < uint64(now) {
				go getWebsite(url.RequestURL, url.RequestHTTPMethod, &wg)
				wg.Add(1)
			}
		}

	}
	wg.Wait()
}

func getWebsite(website string, method string, wg *sync.WaitGroup) {
	if method == "GET" {
		if res, err := http.Get(website); err != nil {

			fmt.Println(website, "is down!")
		} else {
			fmt.Printf("[%d] %s is up \n", res.StatusCode, website)
		}
	} else {
		fmt.Println("method is not allowed!")
	}
	wg.Done()

}

func getUrls() ([]*entities.Tashilcar, error) {
	var db *gorm.DB
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
		fmt.Println(err)
	}
	cbuc := NewGetAPIs(repositories.NewTashilcarRepository(databases.NewDBDataSource(db)))
	//cbuc := di.GetAPIsInstance()
	urls, err := cbuc.Exec()
	if err != nil {
		return nil, err
	} else {
		return urls, nil
	}
	return nil, nil
}
