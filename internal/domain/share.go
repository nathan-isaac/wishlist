package domain

type Share struct {
	Id       string
	Code     string
	Wishlist Wishlist
	Items    []Item
}
