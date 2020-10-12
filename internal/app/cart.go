package app

import (
	"github.com/uudashr/ddd-ecommerce-sales/internal/product"
	"github.com/uudashr/ddd-ecommerce-sales/internal/shop"
)

type CartService struct {
	cartRepo    shop.CartRepository
	productRepo product.Repository
}

func (cs CartService) AddItemToCart(cmd AddItemToCartCommand) error {
	// TODO: we need to define transaction here, to indicate all the block should be in single transaction
	cart, err := cs.cartRepo.CartByID(cmd.CartID)
	if err != nil {
		return err
	}

	if cart == nil {
		return NotFoundError("cart not found")
	}

	prod, err := cs.productRepo.ProductByID(cmd.ItemID)
	if err != nil {
		return err
	}

	if prod == nil {
		return NotFoundError("product not found")
	}

	err = cart.AddItem(prod, cmd.Quantity)
	if err != nil {
		return err
	}

	return cs.cartRepo.Update(cart)
}

type AddItemToCartCommand struct {
	CartID   string
	ItemID   string
	Quantity int
}
