package views

type Share struct {
	Id       string
	Code     string
	Wishlist Wishlist
	Items    []Item
}
