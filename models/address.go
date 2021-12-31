package models

type GeoLocation struct {
	Lat  string `json:"lat"`
	Long string `json:"lng"`
}

type Address struct {
	Street  string      `json:"street"`
	Suite   string      `json:"suite"`
	City    string      `json:"city"`
	Zipcode string      `json:"zipcode"`
	Geo     GeoLocation `json:"geo"`
}
