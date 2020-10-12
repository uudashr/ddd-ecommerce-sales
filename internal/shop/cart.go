package shop

import (
	"errors"

	"github.com/uudashr/ddd-ecommerce-sales/internal/order"
	"github.com/uudashr/ddd-ecommerce-sales/internal/product"
)

type Cart struct {
	id             string
	itemQuantities map[string]int
}

func NewCart(id string, items []CartItem) (*Cart, error) {
	if id == "" {
		return nil, errors.New("empty cart id")
	}

	quantities := make(map[string]int)
	for _, item := range items {
		if item.ItemID == "" {
			return nil, errors.New("empty item id")
		}

		if item.Quantity <= 0 {
			return nil, errors.New("quantity can only larger than 0")
		}

		if _, found := quantities[item.ItemID]; found {
			return nil, errors.New("found duplicate item on cart items")
		}

		quantities[item.ItemID] = item.Quantity
	}

	return &Cart{
		id:             id,
		itemQuantities: quantities,
	}, nil
}

func EmptyCart(id string) (*Cart, error) {
	return NewCart(id, nil)
}

func (c Cart) ID() string {
	return c.id
}

func (c *Cart) AddItem(item *product.Product, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity can only larger than 0")
	}

	c.itemQuantities[item.ID()] = quantity
	return nil
}

func (c Cart) Items() []CartItem {
	var items []CartItem
	for itemID, qty := range c.itemQuantities {
		items = append(items, CartItem{itemID, qty})
	}
	return items
}

type CartItem struct {
	ItemID   string
	Quantity int
}

type CartRepository interface {
	Store(*Cart) error
	CartByID(id string) (*Cart, error)
	Update(*Cart) error
}

func (c *Cart) PlaceOder(orderID string, shippingAddress *order.PostalAddress) (*order.Order, error) {
	var items []order.LineItem
	for itemID, qty := range c.itemQuantities {
		items = append(items, order.LineItem{ItemID: itemID, Quantity: qty})
	}

	return order.New(orderID, items, shippingAddress)
}
