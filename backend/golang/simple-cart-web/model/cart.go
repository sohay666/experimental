package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	con "shop-cart/config"
	"strconv"
)

type Cart struct {
	CartId string `json:"cartId"`
	SkuNo  string `json:"skuNo"`
	Qty    int    `json:"qty"`
}

func Add2Cart(cartId, skuNo, qtyStr string) (err error) {
	log.Printf("cartId (%s)", cartId)
	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		log.Printf("(%s)", err.Error())
		return
	}

	if qty == 0 {
		err = fmt.Errorf("Qty can't be zero")
		return
	}

	//insert data to db
	db, err := con.Db()
	defer db.Close()
	if err != nil {
		log.Printf("GetProducts (%s)", err.Error())
		return
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = fmt.Errorf("fail to create transaction")
		return
	}

	defer tx.Rollback()

	product := Product{}
	err = tx.QueryRow("SELECT skuNo, name, price, inventoryQty FROM products WHERE skuNo = ? AND deletedAt is null", skuNo).
		Scan(&product.SkuNo, &product.Name, &product.Price, &product.InventoryQty)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("That product not exist!")
		}
		log.Printf("Check Qty is available or not (%s)", err.Error())
		return
	}

	// do add to cart
	cart := Cart{}
	err = tx.QueryRow("SELECT cartId, skuNo, qty FROM cart WHERE skuNo = ? AND cartId = ? FOR UPDATE", skuNo, cartId).
		Scan(&cart.CartId, &cart.SkuNo, &cart.Qty)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Check Cart (%s)", err.Error())
		return
	}

	totalQty := cart.Qty + qty
	if product.InventoryQty < qty || totalQty > product.InventoryQty {
		err = fmt.Errorf("Supply item for %s not available, left %d", product.Name, product.InventoryQty)
		return
	}

	if len(cart.SkuNo) == 0 && len(cart.CartId) == 0 {
		_, err = tx.ExecContext(ctx, "INSERT INTO cart SET cartId = ?, skuNo= ?, qty = ?", cartId, skuNo, qty)
		if err != nil {
			log.Printf("Insert cart (%s)", err.Error())
			return
		}
	} else {
		_, err = tx.ExecContext(ctx, "UPDATE cart SET qty = ? WHERE cartId = ? AND skuNo= ?", totalQty, cartId, skuNo)
		if err != nil {
			log.Printf("Update cart (%s)", err.Error())
			return
		}
	}

	// update inventoryQty at table product
	inventoryQty := product.InventoryQty - qty
	_, err = tx.ExecContext(ctx, "UPDATE products SET inventoryQty = ? WHERE skuNo= ? AND deletedAt is null", inventoryQty, skuNo)
	if err != nil {
		log.Printf("Update product inventoryQty (%s)", err.Error())
		return
	}

	if err = tx.Commit(); err != nil {
		err = fmt.Errorf("fail to commit trx")
		return
	}

	return
}

type CartProduct struct {
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
	Price float64 `json:"price"`
}

type ListCart struct {
	CartProducts []CartProduct `json:"cartProducts"`
	Total        float64       `json:"total"`
}

func GetListCart(cartId string) (listCart ListCart, err error) {
	log.Printf("cartId (%s)", cartId)

	db, err := con.Db()
	defer db.Close()
	if err != nil {
		log.Printf("ConnectDb (%s)", err.Error())
		return
	}

	//SELECT name, qty, price FROM product
	rows, err := db.Query("SELECT p.name, c.qty, p.price FROM cart c JOIN products p ON p.skuNo = c.skuNo WHERE cartId = ?", cartId)
	if err != nil {
		log.Printf("QueryListCart (%s)", err.Error())
		return
	}
	defer rows.Close()

	var totalPrice float64
	cartProducts := []CartProduct{}
	for rows.Next() {
		var cart = CartProduct{}
		err = rows.Scan(&cart.Name, &cart.Qty, &cart.Price)
		if err != nil {
			log.Printf("ScanListCart (%s)", err.Error())
			return
		}
		cartProducts = append(cartProducts, cart)
		totalPrice += float64(cart.Qty) * cart.Price
	}

	listCart.Total = totalPrice
	listCart.CartProducts = cartProducts
	return
}
