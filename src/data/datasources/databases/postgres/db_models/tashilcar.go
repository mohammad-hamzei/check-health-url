package db_models

type Tashilcar struct {
	ID                      uint64 `gorm:"<-:create;primarykey"`
	HealthCheck             bool   `gorm:"<-:create"`
	HealthCheckTimeInterval uint64 `gorm:"<-:create"`
	RequestURL              string `gorm:"<-:create"`
	RequestHTTPMethod       string `gorm:"<-:create"`
	RequestHeaders          string `gorm:"<-:create"`
	RequestBody             string `gorm:"<-:create"`
	ResponseStatus          string `gorm:"<-:create"`
	CreatedAt               uint64 `gorm:"<-:create"`
	UpdatedAt               uint64 `gorm:"<-:create"`
}

/* `gorm:"<-:create"` it will alter column to allow only read and create */
/* `gorm:"index:idx_userID"` will index */
