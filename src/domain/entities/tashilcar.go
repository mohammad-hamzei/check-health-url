package entities

type Tashilcar struct {
	ID                      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	HealthCheck             bool   `gorm:"default:true" json:"health_check"`
	HealthCheckTimeInterval uint64 `json:"health_check_time_interval"`
	RequestURL              string `gorm:"size:255;not null;unique" json:"request_url"`
	RequestHTTPMethod       string `gorm:"size:255;not null" json:"request_http_method"`
	RequestHeaders          string `gorm:"size:255;not null" json:"request_headers"`
	RequestBody             string `gorm:"size:255;not null" json:"request_body"`
	ResponseStatus          uint16 `json:"response_status"`
	CreatedAt               uint64 `json:"created_at"`
	UpdatedAt               uint64 `json:"updated_at"`
}
