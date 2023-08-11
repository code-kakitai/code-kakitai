package order

import (
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewOrderProduct(t *testing.T) {
	productID := ulid.NewULID()
	type args struct {
		productID string
		count     int
	}
	tests := []struct {
		name    string
		args    args
		want    *OrderProduct
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				productID: productID,
				count:     1,
			},
			want: &OrderProduct{
				productID: productID,
				count:     1,
			},
			wantErr: false,
		},
		{
			name: "異常系: 商品IDが不正",
			args: args{
				productID: "test",
				count:     1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 購入数が不正",
			args: args{
				productID: productID,
				count:     0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrderProduct(tt.args.productID, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrderProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(OrderProduct{}),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	productID1 := ulid.NewULID()
	productID2 := ulid.NewULID()
	type args struct {
		totalAmount int64
		products    []OrderProduct
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				totalAmount: 100,
				products: []OrderProduct{
					{
						productID: productID1,
						count:     1,
					},
					{
						productID: productID2,
						count:     2,
					},
				},
			},
			want: &Order{
				totalAmount: 100,
				products: []OrderProduct{
					{
						productID: productID1,
						count:     1,
					},
					{
						productID: productID2,
						count:     2,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "異常系: 合計金額が不正",
			args: args{
				totalAmount: -100,
				products: []OrderProduct{
					{
						productID: productID1,
						count:     1,
					},
					{
						productID: productID2,
						count:     2,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 購入商品がない",
			args: args{
				totalAmount: 100,
				products:    []OrderProduct{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	userID := ulid.NewULID()
	OrderedAt := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrder(userID, tt.args.totalAmount, tt.args.products, OrderedAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Order{}, OrderProduct{}),
				cmpopts.IgnoreFields(Order{}, "id", "OrderedAt"),
			)
			if diff != "" {
				t.Errorf("NewOrder() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
