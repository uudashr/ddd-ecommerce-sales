package product

import "errors"

type Product struct {
	id   string
	name string
}

func New(id, name string) (*Product, error) {
	if id == "" {
		return nil, errors.New("empty product id")
	}

	if name == "" {
		return nil, errors.New("empty product name")
	}

	return &Product{
		id:   id,
		name: name,
	}, nil
}

func (p Product) ID() string {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p *Product) ChangeName(s string) error {
	if s == "" {
		return errors.New("empty product name")
	}

	p.name = s
	return nil
}

type Repository interface {
	Store(*Product) error
	ProductByID(id string) (*Product, error)
	Update(*Product) error
}
