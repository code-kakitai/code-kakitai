package purchase

import (
	"context"
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

func Test_purchaseDomainService_PurchaseProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPurchaseHistoryRepo := NewMockPurchaseHistoryRepository(ctrl)
	mockProductRepo := productDomain.NewMockProductRepository(ctrl)
	ds := NewPurchaseDomainService(
		mockPurchaseHistoryRepo,
		mockProductRepo,
	)

	product1, _ := productDomain.NewProduct(
		ulid.NewULID(),
		"product1",
		"description",
		100,
		2,
	)
	product2, _ := productDomain.NewProduct(
		ulid.NewULID(),
		"product1",
		"description",
		200,
		1,
	)
	productIDs := []string{product1.ID(), product2.ID()}
	products := []*productDomain.Product{
		product1,
		product2,
	}
	userID := ulid.NewULID()
	tests := []struct {
		name             string
		purchaseProducts []PurchaseProduct
		mockFunc         func()
		wantErr          bool
	}{
		{
			name: "正常系",
			purchaseProducts: []PurchaseProduct{
				{
					productID: productIDs[0],
					count:     1,
				},
				{
					productID: productIDs[1],
					count:     1,
				},
			},
			mockFunc: func() {
				gomock.InOrder(
					mockProductRepo.EXPECT().FindByIDs(gomock.Any(), productIDs).Return(products, nil),
					mockProductRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Do(
						func(ctx context.Context, p *productDomain.Product) {
							pp := product1
							pp.Consume(1)
							diff := cmp.Diff(
								p,
								pp,
								cmpopts.IgnoreFields(productDomain.Product{}, "id"),
								cmp.AllowUnexported(productDomain.Product{}),
							)
							if diff != "" {
								t.Errorf("purchaseDomainService.PurchaseProducts() diff = %v", diff)
							}
						}).Return(nil),
					mockProductRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Do(
						func(ctx context.Context, p *productDomain.Product) {
							pp := product2
							pp.Consume(0)
							diff := cmp.Diff(
								p,
								pp,
								cmpopts.IgnoreFields(productDomain.Product{}, "id"),
								cmp.AllowUnexported(productDomain.Product{}),
							)
							if diff != "" {
								t.Errorf("purchaseDomainService.PurchaseProducts() diff = %v", diff)
							}
						}).Return(nil),
					mockPurchaseHistoryRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Do(
						func(ctx context.Context, ph *PurchaseHistory) {
							diff := cmp.Diff(
								ph,
								&PurchaseHistory{
									id:          ulid.NewULID(),
									totalAmount: 300,
									products: []PurchaseProduct{
										{
											productID: productIDs[0],
											count:     1,
										},
										{
											productID: productIDs[1],
											count:     1,
										},
									},
								},
								cmpopts.IgnoreFields(PurchaseHistory{}, "purchasedAt", "id"),
								cmp.AllowUnexported(PurchaseHistory{}, PurchaseProduct{}),
							)
							if diff != "" {
								t.Errorf("purchaseDomainService.PurchaseProducts() diff = %v", diff)
							}
						},
					).Return(nil),
				)
			},
			wantErr: false,
		},
		{
			name: "購入した商品の商品詳細が見つからない場合は購入できない",
			purchaseProducts: []PurchaseProduct{
				{
					productID: productIDs[0],
					count:     1,
				},
				{
					productID: productIDs[1],
					count:     10,
				},
			},
			mockFunc: func() {
				gomock.InOrder(
					mockProductRepo.EXPECT().FindByIDs(gomock.Any(), productIDs).Return([]*productDomain.Product{product1}, nil),
				)
			},
			wantErr: true,
		},
		{
			name: "在庫が不足している場合は購入できない",
			purchaseProducts: []PurchaseProduct{
				{
					productID: productIDs[0],
					count:     1,
				},
				{
					productID: productIDs[1],
					count:     10,
				},
			},
			mockFunc: func() {
				gomock.InOrder(
					mockProductRepo.EXPECT().FindByIDs(gomock.Any(), productIDs).Return(products, nil),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if err := ds.PurchaseProducts(context.Background(), userID, tt.purchaseProducts, time.Now()); (err != nil) != tt.wantErr {
				t.Errorf("purchaseDomainService.PurchaseProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
