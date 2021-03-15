package interfaces

import "Catalog_Admin/core"

type ICatalogRepository interface {
	CreateProduct(product core.Product) bool
}
