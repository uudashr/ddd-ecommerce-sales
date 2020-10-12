package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/uudashr/ddd-ecommerce-sales/internal/shop"
)

type CartRepsitory struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) (*CartRepsitory, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	return &CartRepsitory{db}, nil
}

func (r *CartRepsitory) Store(cart *shop.Cart) (retErr error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if retErr != nil {
			err = tx.Rollback()
			if err != nil {
				log.Println("rollback fail:", err)
			}

			retErr = tx.Commit()
		}
	}()

	_, err = tx.Exec("INSERT INTO cart (id) VALUES (?)", cart.ID()) // assume we have lot of columns writen to this table
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO cart_items (cart_id, line_items, quantity) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	for _, cartItem := range cart.Items() {
		_, err = stmt.Exec(cart.ID(), cartItem.ItemID, cartItem.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CartRepsitory) Update(cart *shop.Cart) error {
	// TODO: need to implement this
	panic("not implemented")
}

func (r *CartRepsitory) CartByID(id string) (retCart *shop.Cart, retErr error) {
	// setting transaction is important to maintain consistency when constructing the `shop.Cart`
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSnapshot, ReadOnly: true})
	if err != nil {
		return nil, err
	}

	defer func() {
		if retErr != nil {
			err = tx.Rollback()
			if err != nil {
				// TODO: logging should use `logger` passed through somehow, instead of using package level log
				log.Println("rollback fail:", err)
			}

			retErr = tx.Commit()
		}
	}()

	var cartID string
	err = tx.QueryRow("SELECT id FROM carts WHERE id = ?", id).Scan(&cartID)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT item_id, quantity FROM cart_items WHERE cart_id = ?")
	if err != nil {
		return nil, err
	}

	var items []shop.CartItem
	for rows.Next() {
		var (
			itemID   string
			quantity int
		)
		err = rows.Scan(&itemID, &quantity)
		if err != nil {
			return nil, err
		}

		items = append(items, shop.CartItem{ItemID: itemID, Quantity: quantity})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return shop.NewCart(cartID, items)
}
