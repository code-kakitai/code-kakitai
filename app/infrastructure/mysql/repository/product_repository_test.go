package repository

import (
	"context"
	"testing"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"

	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

func Test_productRepository_Save_And_Find(t *testing.T) {
	ownerID := "01HCNYK3F7RJTWJ7GAQHPZDVE3"
	testStr := "test"
	productID1 := ulid.NewULID()
	productID2 := ulid.NewULID()
	product1, _ := productDomain.Reconstruct(
		productID1,
		ownerID,
		testStr,
		testStr,
		100,
		100,
	)
	product2, _ := productDomain.Reconstruct(
		productID2,
		ownerID,
		testStr,
		testStr,
		100,
		100,
	)
	tests := []struct {
		name    string
		product []*productDomain.Product
		wantErr bool
	}{
		{
			name:    "正常系",
			product: []*productDomain.Product{product1, product2},
			wantErr: false,
		},
	}
	repo := NewProductRepository()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// productを全て保存
			for _, p := range tt.product {
				if err := repo.Save(ctx, p); (err != nil) != tt.wantErr {
					t.Errorf("productRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			// IDで検索
			p, err := repo.FindByID(ctx, tt.product[0].ID())
			if err != nil {
				t.Errorf("productRepository.FindByID() error = %v", err)
			}
			if diff := cmp.Diff(
				tt.product[0],
				p,
				cmp.AllowUnexported(productDomain.Product{}),
			); diff != "" {
				t.Errorf("productRepository.FindByID() = %v, want %v. error is %s", p, tt.product, err)
			}

			// IDsで検索
			ps, err := repo.FindByIDs(ctx, []string{productID1, productID2})
			if err != nil {
				t.Errorf("productRepository.FindByIDs() error = %v", err)
			}
			if diff := cmp.Diff(
				tt.product,
				ps,
				cmp.AllowUnexported(productDomain.Product{}),
			); diff != "" {
				t.Errorf("productRepository.FindByIDs() = %v, want %v. error is %s", ps, tt.product, err)
			}
		})
	}
}
