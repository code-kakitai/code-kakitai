package repository

import (
	"context"

	"github/code-kakitai/code-kakitai/domain/product"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type productRepository struct {
	query *dbgen.Queries
}

func NewProductRepository() product.ProductRepository {
	return &productRepository{query: db.GetQuery()}
}

func (r *productRepository) Save(ctx context.Context, product *product.Product) error {
	if err := r.query.UpsertProduct(ctx, dbgen.UpsertProductParams{
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
	p, err := r.query.ProductFindById(ctx, id)
	if err != nil {
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
	ps, err := r.query.ProductFindByIds(ctx, ids)
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
