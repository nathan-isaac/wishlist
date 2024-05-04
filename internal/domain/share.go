package domain

type Share struct {
	Id             string
	Code           string
	List           List
	Items          []Item
	PurchasedCount int
	CheckoutUrl    string
	CheckoutId     string
}
