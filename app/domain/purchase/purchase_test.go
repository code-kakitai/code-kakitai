package purchase

import (
	"testing"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"
)

func TestNewPurchaseProduct(t *testing.T) {
	productID := ulid.NewULID()
	type args struct {
		productID string
		count     int
	}
	tests := []struct {
		name    string
		args    args
		want    *PurchaseProduct
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				productID: productID,
				count:     1,
			},
			want: &PurchaseProduct{
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
			got, err := NewPurchaseProduct(tt.args.productID, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPurchaseProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(PurchaseProduct{}),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
