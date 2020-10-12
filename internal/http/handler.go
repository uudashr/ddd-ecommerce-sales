package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uudashr/ddd-ecommerce-sales/internal/app"
)

type CartApp interface {
	AddItemToCart(app.AddItemToCartCommand) error
}

type delegate struct {
	cartApp CartApp // app.CartService
}

func (d *delegate) addItemToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["id"]
	var p addItemToCartPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = d.cartApp.AddItemToCart(app.AddItemToCartCommand{
		CartID:   cartID,
		ItemID:   p.ItemID,
		Quantity: p.Quantity,
	})
	if err != nil {
		switch err.(type) {
		case app.NotFoundError:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(cartApp CartApp) http.Handler {
	d := &delegate{cartApp}
	r := mux.NewRouter()
	r.HandleFunc("/carts/{id}/items", d.addItemToCart).Methods(http.MethodPost)
	return r
}

type addItemToCartPayload struct {
	ItemID   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}
