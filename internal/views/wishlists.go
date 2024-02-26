package views

type WishlistIndex struct {
	NewWishlistURL string
	Wishlists      []Wishlist
}

type Wishlist struct {
	ID          string
	Name        string
	Owner       string
	Description string
	EditURL     string
	ShareCode   string
}

type Item struct {
	Id                string
	Link              string
	ImageUrl          string
	Description       string
	Name              string
	Price             string
	PurchasedQuantity string
	NeededQuantity    string
}
