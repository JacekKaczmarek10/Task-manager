package product

import (
	"fmt"
)

type ProductRepo struct {
	products []Product
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{make([]Product, 0)}
}

func (p *ProductRepo) Create(partial Product) Product {
	newItem := partial
	newItem.ID = uint(len(p.products)) + 1
	p.products = append(p.products, newItem)
	return newItem
}

func (p *ProductRepo) GetList() []Product {
	return p.products
}

func (p *ProductRepo) GetOne(id uint) (Product, error) {
	for _, it := range p.products {
		if it.ID == id {
			return it, nil
		}
	}
	return Product{}, fmt.Errorf("key '%d' not found", id)
}

func (p *ProductRepo) Update(id uint, amended Product) (Product, error) {
	for i, it := range p.products {
		if it.ID == id {
			amended.ID = id
			p.products = append(p.products[:i], p.products[i+1:]...)
			p.products = append(p.products, amended)
			return amended, nil
		}
	}
	return Product{}, fmt.Errorf("key '%d' not found", amended.ID)
}

func (p *ProductRepo) DeleteOne(id uint) (bool, error) {
	for i, it := range p.products {
		if it.ID == id {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}
