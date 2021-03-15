package interfaces

import "Catalog_User/core"

type ICatalogRepository interface {
	GetAll() []*core.Product
}
