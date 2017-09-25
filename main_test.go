package main

import "testing"

func TestItCanGetACoupon(t *testing.T) {

	var coupons CouponList
	coupons.generateCouponsList(10)
	cp := coupons.get("Tesco", 5)

	if len(cp) != 1 {
		t.Errorf("Expected exactly one result but got %d results, with %+v instead.", len(cp), cp)
	}

	if cp[0].Brand != "Tesco" {
		t.Errorf("Expected brand %s, but it was %s instead.", "Tesco", cp[0].Brand)
	}

	if cp[0].Value != 5 {
		t.Errorf("Expected value %d, but it was %f instead.", 5, cp[0].Value)
	}
}

func TestItGetsACouponOnlyOnce(t *testing.T) {

	var coupons CouponList
	coupons.generateCouponsList(10)

	cp := coupons.get("Tesco", 5)
	cp = coupons.get("Tesco", 5)

	if len(cp) > 0 {
		t.Errorf("Expected empty slice but got %+v", cp)
	}
}

func TestItCannotGetInexistentCoupon(t *testing.T) {

	var coupons CouponList
	coupons.generateCouponsList(10)

	cp := coupons.get("Tesco", 11)

	if len(cp) > 0 {
		t.Errorf("Expected empty slice but got %+v", cp)
	}
}
