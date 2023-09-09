package repository

import (
	"context"

	"github.com/code-kakitai/go-pkg/ulid"

	"github/code-kakitai/code-kakitai/domain/order"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type orderRepository struct {
}

func NewOrderRepository() order.OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) Save(ctx context.Context, order *order.Order) error {
	query := db.GetQuery(ctx)
	if err := query.InsertOrder(ctx, dbgen.InsertOrderParams{
		ID:          order.ID(),
		UserID:      order.UserID(),
		TotalAmount: order.TotalAmount(),
		OrderedAt:   order.OrderedAt(),
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
			Price:     p.Price(),
			Quantity:  int32(p.Quantity()),
		}
		if err := query.InsertOrderProduct(ctx, op); err != nil {
			return err
		}
	}
	return nil
}
