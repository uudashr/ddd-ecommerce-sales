package shop_test

import (
	"testing"

	"github.com/uudashr/ddd-ecommerce-sales/internal/product"
	"github.com/uudashr/ddd-ecommerce-sales/internal/shop"
)

func TestCart(t *testing.T) {
	cart, err := shop.EmptyCart("a-cart")
	if err != nil {
		t.Fatal(err)
	}

	prod, err := product.New("btl-blue-170", "Blue Bottle 170ml")
	if err != nil {
		t.Fatal(err)
	}

	err = cart.AddItem(prod, 10)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(cart.Items()), 1; got != want {
		t.Errorf("cart items got: %d, want: %d", got, want)
	}

	items := cart.Items()
	items[0].Quantity = -1

	items = cart.Items()
	if got, want := items[0].Quantity, 10; got != want {
		t.Errorf("cart items[0] quantity got: %d, want: %d", got, want)
	}
}
