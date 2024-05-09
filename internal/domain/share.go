package domain

type Share struct {
	Id             string
	Code           string
	List           List
	Items          []Item
	PurchasedItems []Item
	PurchasedCount int
	CheckoutUrl    string
	CheckoutId     string
}
