package purchase

import (
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewPurchaseHistory(t *testing.T) {
	userID := ulid.NewULID()
	productID1 := ulid.NewULID()
	productID2 := ulid.NewULID()
	type args struct {
		totalAmount int64
		products    []PurchaseProduct
	}
	tests := []struct {
		name    string
		args    args
		want    *PurchaseHistory
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				totalAmount: 100,
				products: []PurchaseProduct{
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
			want: &PurchaseHistory{
				totalAmount: 100,
				products: []PurchaseProduct{
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
				products: []PurchaseProduct{
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
				products:    []PurchaseProduct{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	purchasedAt := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPurchaseHistory(tt.args.totalAmount, tt.args.products, purchasedAt, userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPurchaseHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(PurchaseHistory{}, PurchaseProduct{}),
				cmpopts.IgnoreFields(PurchaseHistory{}, "id", "purchasedAt"),
			)
			if diff != "" {
				t.Errorf("NewPurchaseHistory() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
