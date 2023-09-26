package usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.tashilcar.com/core/tashilcar/core/configs"
	"git.tashilcar.com/core/tashilcar/data/datasources/databases"
	"git.tashilcar.com/core/tashilcar/data/repositories"
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"os"
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
		if healthCheck == true {
			updatedAt := url.UpdatedAt
			interval := url.HealthCheckTimeInterval
			totalTime := updatedAt + (interval * 60)
			now := time.Now().Unix()
			if totalTime > uint64(now) {
				go getWebsite(url.ID, url.RequestURL, url.ResponseStatus, &wg)
				wg.Add(1)
			}
		}

	}
	wg.Wait()
}

func getWebsite(id uint64, url string, responseStatus uint16, wg *sync.WaitGroup) {
	statusCode, err := checkURLStatus(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("URL: %s - status code: %d\n", url, statusCode)

		err := setStatus(id, uint16(statusCode))
		if err != nil {
			return
		}
		if int(responseStatus) != statusCode {
			sendWebhook(url, statusCode)
		}
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

func setStatus(id uint64, statusCode uint16) error {
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
	cbuc := NewUpdateResponseStatus(repositories.NewTashilcarRepository(databases.NewDBDataSource(db)))
	//cbuc := di.UpdateResponseStatusInstance()
	er := cbuc.Exec(id, statusCode)
	if er != nil {
		return er
	}
	return nil
}

func checkURLStatus(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// وضعیت بالا
		return response.StatusCode, nil
	} else if response.StatusCode >= 400 && response.StatusCode < 500 {
		// وضعیت پایین (معمولا خطاهای کاربر)
		return response.StatusCode, nil
	} else {
		// وضعیت نامشخص یا خطای سرور
		return response.StatusCode, fmt.Errorf("HTTP Status Code: %d", response.StatusCode)
	}
}

func sendWebhook(url string, statusCode int) {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		fmt.Println("please set `WEBHOOK_URL`")
		return
	}

	data := map[string]interface{}{
		"url":        url,
		"statusCode": statusCode,
	}
	webhookMessage := os.Getenv("WEBHOOK_MESSAGE")
	if webhookMessage != "" {
		data = map[string]interface{}{
			"url":        url,
			"statusCode": statusCode,
			"message":    webhookMessage,
		}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error []byte:", err)
		return
	}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

}
