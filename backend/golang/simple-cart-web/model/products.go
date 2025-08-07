package model

import (
	"log"
	con "shop-cart/config"
)

type Product struct {
	SkuNo        string  `json:"skuNo"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	InventoryQty int     `json:"inventoryQty"`
}

func GetProducts() (products []Product, err error) {
	db, err := con.Db()
	defer db.Close()
	if err != nil {
		log.Printf("GetProducts (%s)", err.Error())
		return
	}
	rows, err := db.Query("SELECT skuNo, name, price, inventoryQty FROM products WHERE deletedAt is null")
	if err != nil {
		log.Printf("GetProducts (%s)", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product = Product{}
		err = rows.Scan(&product.SkuNo, &product.Name, &product.Price, &product.InventoryQty)
		if err != nil {
			log.Printf("GetProducts (%s)", err.Error())
			return
		}
		products = append(products, product)
	}
	return
}
