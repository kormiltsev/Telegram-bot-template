package product

import (
	"log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Sku {
	return catalog.Products
}

func (s *Service) Get(idx int) (*Sku, bool) {
	if idx < 0 || idx >= len(catalog.Products) {
		log.Println("there are no element in Products with number ", idx)
		return nil, false
	}
	return &catalog.Products[idx], true
}

func (s *Service) Addprod(newsku *Sku) int {

	catalog.Products = append(catalog.Products, *newsku)
	log.Printf("New Sku %s added with id = #%v.", newsku.Title, len(catalog.Products))
	return len(catalog.Products) - 1
}

func (s *Service) RewriteStorage() error {
	return RewriteStorage()
}

func (s *Service) Editprod(newsku *Sku, idx int) bool {

	catalog.Products[idx] = *newsku
	return true
}

func (s *Service) Delprod(idx int) bool {
	if idx < 0 || idx >= len(catalog.Products) {
		log.Println("there are no element in Products with number ", idx)
		return false
	}
	if idx == len(catalog.Products)-1 {
		catalog.Products = catalog.Products[:len(catalog.Products)-1]
	} else {
		for idx < len(catalog.Products)-1 {
			catalog.Products[idx] = catalog.Products[idx+1]
			idx++
		}
		catalog.Products = catalog.Products[:len(catalog.Products)-1]
	}
	RewriteStorage()
	return true
}
