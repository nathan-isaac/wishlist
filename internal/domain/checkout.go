package domain

import (
	"fmt"
	"time"
)

type Checkout struct {
	CheckoutId    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	List          List
	CheckoutItems []CheckoutItem
	Response      CheckoutResponse
}

func (c Checkout) UpdateUrl() string {
	return fmt.Sprintf("/checkouts/%s", c.CheckoutId)
}

type CheckoutItem struct {
	ID         string
	CheckoutID string
	Quantity   int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Item       Item
}

func (it *CheckoutItem) QuantityOptions() []int64 {
	quantity := it.Item.Quantity

	if it.Quantity > quantity {
		quantity = it.Quantity
	}

	var options []int64

	for i := int64(1); i <= quantity; i++ {
		options = append(options, i)
	}

	return options
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
