# Coupon-Box Golang service

### Description

This service accepts a GET request from a client or service and returns zero or one coupon corresponding to the querystring.

### Scope

Part of the microservices suite for the Coupon Box project

### Dependencies

None

### Installation

Clone this repository.

##### Requirements

Requires <= Go 1.9

### API

Accepts a RESTful GET request :

`/get-coupon?brand={brand}&value={value}`

returns zero or one results in a json array :

`[{"brand":"Tesco","value":4}]`

##### Exceptions

Following limitations are implemented :

- *http.StatusBadRequest* : when `value` is not numeric
- *http.StatusInternalServerError* : when an error occurs during the json marshalling. 

### Deployment

`go build`

`./coupons-box-go`

### Tests

Includes a test suite :

`go test`

