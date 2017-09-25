package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

type requestHandler struct{}

type Coupon struct {
	Brand string `json:"brand,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type CouponList []Coupon

var Coupons CouponList

func handleGetCoupon(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	brand := r.URL.Query().Get("brand")
	value := r.URL.Query().Get("value")

	if brand != "" && value != "" {

		val, err := strconv.ParseFloat(value, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cp := Coupons.get(brand, val)

		response, err := json.Marshal(cp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}

func (coupons CouponList) get(brand string, value float64) (CouponList) {

	var responseCoupons CouponList

	for i, cp := range coupons {
		if cp.Brand == brand && cp.Value == value {
			coupons.remove(i)
			responseCoupons = append(responseCoupons, cp)
			return responseCoupons
		}
	}
	return CouponList{}
}

func (coupons CouponList) remove(i int) {

	coupons[i] = Coupon{}
}

func (coupons *CouponList) generateCouponsList(number int) {
	for i := 0; i < number; i++ {
		*coupons = append(*coupons, Coupon{"Tesco", float64(i)})
	}
}

func main() {
	Coupons.generateCouponsList(20)

	http.HandleFunc("/get-coupon", handleGetCoupon)

	err := http.ListenAndServe(":9999", nil)

	log.Fatal(err)
}
