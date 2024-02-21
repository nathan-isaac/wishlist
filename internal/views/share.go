package views

type Share struct {
	Id       string
	Code     string
	Wishlist Wishlist
	Items    []Item
}

type Wishlist struct {
	ID          string
	Name        string
	Owner       string
	Description string
}

type Item struct {
	Id          string
	Link        string
	ImageUrl    string
	Description string
	Quantity    int
}
