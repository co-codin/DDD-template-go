package aggregate

import (
	"ddd-template/entity"
	"errors"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	
}