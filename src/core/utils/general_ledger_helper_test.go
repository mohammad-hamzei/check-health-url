package utils_test

import (
	"testing"

	"git.tashilcar.com/core/tashilcar/core/utils"
)

func TestGetDepositServiceName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "deposit_service_name",
			want: "deposit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetDepositServiceName(); got != tt.want {
				t.Errorf("GetDepositServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDepositTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "deposit_tefsil_ID",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetDepositTafsilID(); got != tt.want {
				t.Errorf("GetDepositTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrderServiceName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "order_service_name",
			want: "order",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetOrderServiceName(); got != tt.want {
				t.Errorf("GetOrderServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCancelOrderTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "cancel_order_tafsil_ID",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetCancelOrderTafsilID(); got != tt.want {
				t.Errorf("GetCancelOrderTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatchOrderTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "match_order_tafsil_ID",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetMatchOrderTafsilID(); got != tt.want {
				t.Errorf("GetMatchOrderTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChangeCurrencyTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "change_currency_tafsil_ID",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetChangeCurrencyTafsilID(); got != tt.want {
				t.Errorf("GetChangeCurrencyTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWithdrawServiceName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "withdraw_service_name",
			want: "withdraw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetWithdrawServiceName(); got != tt.want {
				t.Errorf("GetWithdrawServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWithdrawRequestTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "withdraw_request_tafsil_ID",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetWithdrawRequestTafsilID(); got != tt.want {
				t.Errorf("GetWithdrawRequestTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCancelWithdrawTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "cancel_withdraw_tafsil_ID",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetCancelWithdrawTafsilID(); got != tt.want {
				t.Errorf("GetCancelWithdrawTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfirmWithdrawTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "confirm_withdraw_tafsil_ID",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetConfirmWithdrawTafsilID(); got != tt.want {
				t.Errorf("GetConfirmWithdrawTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBankPortWageTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "bank_port_wage_tafsil_ID",
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetBankPortWageTafsilID(); got != tt.want {
				t.Errorf("GetBankPortWageTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWithdrawFeeCustomerTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "withdraw_fee_customer_tafsil_ID",
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetWithdrawFeeCustomerTafsilID(); got != tt.want {
				t.Errorf("GetWithdrawFeeCustomerTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWithdrawFeeBusinessTafsilID(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "withdraw_fee_business_tafsil_ID",
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetWithdrawFeeBusinessTafsilID(); got != tt.want {
				t.Errorf("GetWithdrawFeeBusinessTafsilID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNull(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "null_ID",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetNull(); got != tt.want {
				t.Errorf("GetNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCash(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "cash_ID",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetCash(); got != tt.want {
				t.Errorf("GetCash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEquity(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "equity_ID",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetEquity(); got != tt.want {
				t.Errorf("GetEquity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCashReceivable(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "cash_receivable_ID",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetCashReceivable(); got != tt.want {
				t.Errorf("GetCashReceivable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCashPayable(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "cash_payable_ID",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetCashPayable(); got != tt.want {
				t.Errorf("GetCashPayable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetExpense(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "expense_ID",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetExpense(); got != tt.want {
				t.Errorf("GetExpense() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRevenue(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		{
			name: "revenue_ID",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetRevenue(); got != tt.want {
				t.Errorf("GetRevenue() = %v, want %v", got, tt.want)
			}
		})
	}
}
