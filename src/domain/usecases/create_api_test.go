package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"testing"
)

func (m *MockRepository) Insert(tashilcar *entities.Tashilcar) error {
	// در اینجا می‌توانید عملیات موک‌آپ ریپازیتوری را انجام دهید.
	// در این مثال، هیچ خطایی را برنمی‌گردانیم.
	return nil
}

func TestExecCreateAPI(t *testing.T) {
	// ایجاد یک نمونه از ریپازیتوری موک‌آپ.
	mockRepo := &MockRepository{}

	// ایجاد نمونه از usecase و تزریق موک‌آپ ریپازیتوری به آن.
	usecase := NewCreateAPI(mockRepo)

	// ایجاد یک نمونه از Tashilcar برای تست.
	tashilcar := entities.Tashilcar{
		ID:                      0,
		HealthCheck:             false,
		HealthCheckTimeInterval: 10,
		RequestURL:              "TestRequestURL",
		RequestHTTPMethod:       "RequestHTTPMethodTest",
		RequestHeaders:          "RequestHeadersTest",
		RequestBody:             "RequestBodyTest",
		ResponseStatus:          "ResponseStatusTest",
		CreatedAt:               0,
		UpdatedAt:               0,
	}

	// فراخوانی تابع Exec برای تست.
	err := usecase.Exec(tashilcar)

	// بررسی خطا (در اینجا ما از یک موک‌آپ ریپازیتوری خالی استفاده کرده‌ایم که هیچ خطایی ندارد).
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// می‌توانید تست‌های بیشتری برای بررسی رفتار در صورت خطا ایجاد کنید.
}

/*func TestExecCreateAPIWithError(t *testing.T) {
	mockRepo := &MockRepository{}

	mockRepo.Err = errors.New("some error")

	usecase := NewCreateAPI(mockRepo)

	tashilcar := entities.Tashilcar{
		ID:                      0,
		HealthCheck:             false,
		HealthCheckTimeInterval: 10,
		RequestURL:              "TestRequestURL",
		RequestHTTPMethod:       "RequestHTTPMethodTest",
		RequestHeaders:          "RequestHeadersTest",
		RequestBody:             "RequestBodyTest",
		ResponseStatus:          "ResponseStatusTest",
		CreatedAt:               0,
		UpdatedAt:               0,
	}

	err := usecase.Exec(tashilcar)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
*/
