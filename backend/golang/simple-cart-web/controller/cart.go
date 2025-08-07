package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"shop-cart/model"
)

type RespAddToCart struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RespListCart struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    model.ListCart `json:"data"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userAgent := r.UserAgent()
	date := time.Now().Format("2006-01-02 15")
	cartId := fmt.Sprintf("%x", md5.Sum([]byte(userAgent+date)))

	if r.Method == "POST" {
		r.ParseForm()

		qty := r.Form["qty"][0]
		skuNo := r.Form["skuNo"][0]

		resp := RespAddToCart{}
		if err := model.Add2Cart(cartId, skuNo, qty); err != nil {
			resp.Status = http.StatusBadRequest
			resp.Message = err.Error()
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp.Message = "success"
		resp.Status = http.StatusOK
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func ListCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userAgent := r.UserAgent()
	date := time.Now().Format("2006-01-02 15")
	cartId := fmt.Sprintf("%x", md5.Sum([]byte(userAgent+date)))

	resp := RespListCart{}
	data, err := model.GetListCart(cartId)
	if err != nil {
		resp.Status = http.StatusBadRequest
		resp.Message = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp.Message = "success"
	resp.Status = http.StatusOK
	resp.Data = data
	json.NewEncoder(w).Encode(resp)
	return
}
