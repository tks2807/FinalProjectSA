package repositories

import (
	"Catalog_User/core"
	"Catalog_User/core/interfaces"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)
type CatalogRepository struct {
	pool pgxpool.Pool
}

func NewCatalogRepository(conn *pgxpool.Pool) interfaces.ICatalogRepository {
	return &CatalogRepository{*conn}
}

func (r CatalogRepository) GetAll() []*core.Product{
	stmt := "SELECT * FROM catalog"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		return nil
	}
	defer rows.Close()
	products := []*core.Product{}
	for rows.Next() {
		p := &core.Product{}
		err = rows.Scan(&p.Id, &p.Name, &p.Price, &p.Quantity)
		if err != nil {
			return nil
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return products
}
