package server

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wishlist/internal/domain"
	"wishlist/internal/gateway"
	"wishlist/internal/utils"
	"wishlist/internal/views"
)

func (s *Server) WishlistsShowHandler(c echo.Context) error {
	id := c.Param("id")

	response, err := s.domain.FindWishlist(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, views.WishlistShowView(response.Wishlist, response.Items))
}

func (s *Server) WishlistsEditHandler(c echo.Context) error {
	id := c.Param("id")

	response, err := s.domain.FindWishlist(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, views.WishlistEditView(response.Wishlist))
}

func (s *Server) WishlistsUpdateHandler(c echo.Context) error {
	id := c.Param("id")

	response, err := s.domain.UpdateWishlist(domain.UpdateWishlistParams{
		ID:          id,
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error updating wishlist: %s", err))
	}

	return HxRedirect(c, fmt.Sprintf("/admin/wishlists/%s", response.Wishlist.ID))
}

func (s *Server) WishlistsDeleteHandler(c echo.Context) error {
	id := c.Param("id")

	err := s.domain.DeleteWishlist(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error deleting wishlist: %s", err))
	}

	return HxRedirect(c, "/admin/wishlists")
}

func (s *Server) WishlistsIndexHandler(c echo.Context) error {
	wishlists, err := s.queries.ListWishlists(s.ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlists: %s", err))
	}

	wishlistsView := utils.Map(wishlists, func(wishlist gateway.Wishlist) domain.Wishlist {
		return domain.ToWishlist(wishlist)
	})

	return Render(c, views.WishlistIndexView(
		domain.WishlistIndex{
			NewWishlistURL: "/admin/wishlists/new",
			Wishlists:      wishlistsView,
		},
	))
}

func (s *Server) WishlistsNewHandler(c echo.Context) error {
	return Render(c, views.WishlistNewView())
}

func (s *Server) WishlistsCreateHandler(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	id, err := GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating id: %s", err))
	}

	shareId, err := GenerateShareId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating share id: %s", err))
	}

	err = s.queries.CreateWishlist(s.ctx, gateway.CreateWishlistParams{
		ID:          id,
		Name:        name,
		Description: description,
		ShareCode:   shareId,
		Public:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating wishlist: %s", err))
	}

	return c.Redirect(http.StatusFound, "/admin/wishlists")
}

func (s *Server) ItemsNewHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, views.ItemNewView(domain.ToWishlist(wishlist)))
}

type ItemFormInput struct {
	Name        string
	Link        string
	Description string
	ImageURL    string
	Quantity    int64
	Price       int64
}

func NewItemFormInput(c echo.Context) (ItemFormInput, error) {
	quantity, err := strconv.ParseInt(c.FormValue("quantity"), 10, 64)

	if err != nil {
		return ItemFormInput{}, err
	}

	price, err := parsePriceInput(c.FormValue("price"))

	if err != nil {
		return ItemFormInput{}, err
	}

	return ItemFormInput{
		Name:        strings.Trim(c.FormValue("name"), ""),
		Link:        strings.Trim(c.FormValue("link"), ""),
		Description: strings.Trim(c.FormValue("description"), ""),
		ImageURL:    strings.Trim(c.FormValue("image_url"), ""),
		Quantity:    quantity,
		Price:       price,
	}, nil
}

func parsePriceInput(price string) (int64, error) {
	priceReplacer := strings.NewReplacer("$", "", ",", "")
	price = priceReplacer.Replace(strings.Trim(price, ""))
	floatPrice, err := strconv.ParseFloat(price, 64)

	if err != nil {
		return 0, err
	}

	moneyPrice := money.NewFromFloat(floatPrice, money.USD)
	return moneyPrice.Amount(), nil
}

func (s *Server) ItemsCreateHandler(c echo.Context) error {
	wishlistId := c.FormValue("wishlist_id")

	wishlist, err := s.queries.FindWishlist(s.ctx, wishlistId)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	itemInput, err := NewItemFormInput(c)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error parsing input: %s", err))
	}

	id, err := GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating id: %s", err))
	}

	err = s.queries.CreateWishlistItem(s.ctx, gateway.CreateWishlistItemParams{
		ID:          id,
		WishlistID:  wishlistId,
		Name:        itemInput.Name,
		Description: itemInput.Description,
		Link:        itemInput.Link,
		ImageUrl:    itemInput.ImageURL,
		Quantity:    itemInput.Quantity,
		Price:       itemInput.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating wishlist: %s", err))
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/admin/wishlists/%s", wishlist.ID))
}

func (s *Server) ItemsEditHandler(c echo.Context) error {
	id := c.Param("id")

	log.Info("id: ", id)

	item, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting item: %s", err))
	}

	wishlist, err := s.queries.FindWishlist(s.ctx, item.WishlistID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, views.ItemEditView(domain.ToWishlist(wishlist), domain.ToItem(item)))
}

func (s *Server) ItemsUpdateHandler(c echo.Context) error {
	id := c.Param("id")

	item, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	itemInput, err := NewItemFormInput(c)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error parsing input: %s", err))
	}

	err = s.queries.UpdateItem(s.ctx, gateway.UpdateItemParams{
		ID:          item.ID,
		Name:        itemInput.Name,
		Description: itemInput.Description,
		Link:        itemInput.Link,
		ImageUrl:    itemInput.ImageURL,
		Quantity:    itemInput.Quantity,
		Price:       itemInput.Price,
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error updating wishlist: %s", err))
	}

	return HxRedirect(c, fmt.Sprintf("/admin/wishlists/%s", item.WishlistID))
}

func (s *Server) ItemsDeleteHandler(c echo.Context) error {
	id := c.Param("id")

	item, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	err = s.queries.DeleteItem(s.ctx, item.ID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error updating wishlist: %s", err))
	}

	return HxRedirect(c, fmt.Sprintf("/admin/wishlists/%s", item.WishlistID))
}
