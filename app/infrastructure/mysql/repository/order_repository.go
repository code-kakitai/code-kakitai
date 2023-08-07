package repository

import (
	"context"

	"github.com/code-kakitai/go-pkg/ulid"

	"github/code-kakitai/code-kakitai/domain/purchase"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type orderRepository struct {
}

func NewPurchaseRepository() purchase.PurchaseHistoryRepository {
	return &orderRepository{}
}

func (r *orderRepository) Save(ctx context.Context, order *purchase.PurchaseHistory) error {
	query := db.GetQuery(ctx)
	if err := query.InsertPurchaseHistory(ctx, dbgen.InsertPurchaseHistoryParams{
		ID:          order.ID(),
		UserID:      order.UserID(),
		TotalAmount: order.TotalAmount(),
		OrderedAt:   order.PurchasedAt(),
	}); err != nil {
		return err
	}
	pp := order.Products()
	for _, p := range pp {
		id := ulid.NewULID()
		op := dbgen.InsertOrderProductParams{
			ID:        id,
			OrderID:   order.ID(),
			ProductID: p.ProductID(),
			Price:     100, // todo domainロジック修正したらここも修正
			Quantity:  int32(p.Count()),
		}
		if err := query.InsertOrderProduct(ctx, op); err != nil {
			return err
		}
	}
	return nil
}
