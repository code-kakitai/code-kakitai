package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github/code-kakitai/code-kakitai/domain/error"
	"github/code-kakitai/code-kakitai/domain/product"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type productRepository struct {
}

func NewProductRepository() product.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Save(ctx context.Context, product *product.Product) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertProduct(ctx, dbgen.UpsertProductParams{
		ID:          product.ID(),
		OwnerID:     product.OwnerID(),
		Name:        product.Name(),
		Description: product.Description(),
		Price:       product.Price(),
		Stock:       int32(product.Stock()),
	}); err != nil {
		return err
	}
	return nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*product.Product, error) {
	query := db.GetQuery(ctx)
	p, err := query.ProductFindById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NotFoundErr
		}
		return nil, err
	}
	pd, err := product.Reconstruct(
		p.ID,
		p.OwnerID,
		p.Name,
		p.Description,
		p.Price,
		int(p.Stock),
	)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (r *productRepository) FindByIDs(ctx context.Context, ids []string) ([]*product.Product, error) {
	query := db.GetQuery(ctx)
	ps, err := query.ProductFindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	var products []*product.Product
	for _, p := range ps {
		pd, err := product.Reconstruct(
			p.ID,
			p.OwnerID,
			p.Name,
			p.Description,
			p.Price,
			int(p.Stock),
		)
		if err != nil {
			return nil, err
		}
		products = append(products, pd)
	}
	return products, nil
}
