package repository

import (
	"context"
	"testing"
	"time"

	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

func TestOrderRepository_Save(t *testing.T) {
	orderProduct, err := orderDomain.NewOrderProduct(
		"01HCNYK4MQNC6G6X3F3DGXZ2J8",
		1000,
		1,
	)
	now := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	order, err := orderDomain.NewOrder(
		"01HCNYK0PKYZWB0ZT1KR0EPWGP",
		1000,
		[]orderDomain.OrderProduct{*orderProduct},
		now)
	if err != nil {
		t.Error(err)
	}
	orderRepository := NewOrderRepository()

	tests := []struct {
		name  string
		input *orderDomain.Order
	}{
		{
			name:  "保存ができること",
			input: order,
		},
	}
	resetTestData(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := orderRepository.Save(context.Background(), tt.input); err != nil {
				t.Errorf("Save() error = %v", err)
			}
		})
	}
}
