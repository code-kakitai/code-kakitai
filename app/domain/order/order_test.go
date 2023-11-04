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
	price := int64(100)
	type args struct {
		productID string
		price     int64
		quantity  int
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
				price:     price,
				quantity:  1,
			},
			want: &OrderProduct{
				productID: productID,
				price:     price,
				quantity:  1,
			},
			wantErr: false,
		},
		{
			name: "異常系: 商品IDが不正",
			args: args{
				productID: "test",
				price:     price,
				quantity:  1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 購入数が不正",
			args: args{
				productID: productID,
				price:     price,
				quantity:  0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrderProduct(tt.args.productID, price, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrderProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(OrderProduct{}),
			)
			if diff != "" {
				t.Errorf("NewOrder() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	productID1 := ulid.NewULID()
	productID2 := ulid.NewULID()
	userID := ulid.NewULID()
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
						quantity:  1,
					},
					{
						productID: productID2,
						quantity:  2,
					},
				},
			},
			want: &Order{
				totalAmount: 100,
				userID:      userID,
				products: []OrderProduct{
					{
						productID: productID1,
						quantity:  1,
					},
					{
						productID: productID2,
						quantity:  2,
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
						quantity:  1,
					},
					{
						productID: productID2,
						quantity:  2,
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
				cmpopts.IgnoreFields(Order{}, "id", "orderedAt"),
			)
			if diff != "" {
				t.Errorf("NewOrder() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}

func TestOrderProducts_TotalAmount(t *testing.T) {
	tests := []struct {
		name string
		p    OrderProducts
		want int64
	}{
		{
			name: "正常系",
			p: OrderProducts{
				{
					price:    100,
					quantity: 1,
				},
				{
					price:    200,
					quantity: 2,
				},
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.TotalAmount(); got != tt.want {
				t.Errorf("OrderProducts.TotalAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
