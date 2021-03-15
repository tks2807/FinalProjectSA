package repositories

import (
	"Catalog_Admin/core"
	"Catalog_Admin/core/interfaces"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)
type CatalogRepository struct {
	pool pgxpool.Pool
}

func NewCatalogRepository(conn *pgxpool.Pool) interfaces.ICatalogRepository {
	return &CatalogRepository{*conn}
}

func (c *CatalogRepository) CreateProduct(product core.Product) bool {
	sql := "INSERT INTO catalod(id, name, price, quantity) " +
		"VALUES($1, $2, $3, $4)";
	row := c.pool.QueryRow(context.Background(),
		sql, product.Id, product.Name, product.Price, product.Quantity)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}

