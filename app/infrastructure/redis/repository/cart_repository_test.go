package repository

import (
	"context"
	"reflect"
	"testing"

	domainCart "github/code-kakitai/code-kakitai/domain/cart"
)

func Test_cartRepository_FindByUserID_Save(t *testing.T) {
	cart, err := domainCart.NewCart("01HCNYK0PKYZWB0ZT1KR0EPWGP")
	if err != nil {
		t.Error(err)
	}
	count := 3
	if err := cart.AddProduct("01HCNYK0PKYZWB0ZT1KR0EPWGP", count); err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		want    *domainCart.Cart
		wantErr bool
	}{
		{
			name:    "正常系: 保存したカートを取得できる",
			want:    cart,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &cartRepository{}
			err := r.Save(context.Background(), tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := r.FindByUserID(context.Background(), tt.want.UserID())
			if (err != nil) != tt.wantErr {
				t.Errorf("cartRepository.FindByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cartRepository.FindByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
