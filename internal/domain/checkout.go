package domain

import "time"

type Checkout struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	List          List
	CheckoutItems []CheckoutItem
	Response      CheckoutResponse
}

type CheckoutItem struct {
	ID         string
	CheckoutID string
	Quantity   int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Item       Item
}

type CheckoutResponse struct {
	ID             string
	CheckoutID     string
	Name           string
	AddressLineOne string
	AddressLineTwo string
	City           string
	State          string
	Zip            string
	Message        string
}
