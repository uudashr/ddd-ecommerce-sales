package order

import "errors"

type Order struct {
	id               string
	itemQuantities   map[string]int
	shippingAddresss *PostalAddress
}

func New(id string, items []LineItem, shippingAddresss *PostalAddress) (*Order, error) {
	if id == "" {
		return nil, errors.New("empty order id")
	}

	if len(items) == 0 {
		return nil, errors.New("no line items")
	}

	if shippingAddresss == nil {
		return nil, errors.New("nil shipping address")
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
			return nil, errors.New("found duplicate item on line items")
		}

		quantities[item.ItemID] = item.Quantity
	}

	return &Order{
		id:               id,
		itemQuantities:   quantities,
		shippingAddresss: shippingAddresss,
	}, nil
}

func (o Order) ID() string {
	return o.id
}

func (o Order) LineItems() []LineItem {
	var items []LineItem
	for itemID, qty := range o.itemQuantities {
		items = append(items, LineItem{itemID, qty})
	}
	return items
}

type LineItem struct {
	ItemID   string
	Quantity int
}
