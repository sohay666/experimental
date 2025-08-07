package controller

import (
	"shop-cart/model"

	"encoding/json"
	"net/http"
)

type RespProduct struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []model.Product `json:"data"`
}

func Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		resp := RespProduct{
			Status:  http.StatusOK,
			Message: "success",
		}
		products, err := model.GetProducts()
		if err != nil {
			resp.Status = http.StatusInternalServerError
			resp.Message = "error"
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp.Data = products
		json.NewEncoder(w).Encode(resp)
		return
	}
}
