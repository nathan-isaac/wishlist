package domain

type Share struct {
	Id          string
	Code        string
	List        List
	Items       []Item
	CheckoutUrl string
	CheckoutId  string
}
