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
	"wishlist/internal/views/item"
	"wishlist/internal/views/list"
)

func (s *Server) ListsShowHandler(c echo.Context) error {
	id := c.Param("list_id")

	response, err := s.domain.FindList(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, list.WishlistShowView(response.List, response.Items))
}

func (s *Server) ListsEditHandler(c echo.Context) error {
	id := c.Param("list_id")

	response, err := s.domain.FindList(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, list.WishlistEditView(response.List))
}

func (s *Server) ListsUpdateHandler(c echo.Context) error {
	id := c.Param("list_id")

	response, err := s.domain.UpdateList(domain.UpdateListParams{
		ListId:      id,
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error updating wishlist: %s", err))
	}

	return HxRedirect(c, fmt.Sprintf("/admin/lists/%s", response.List.ListId))
}

func (s *Server) ListsDeleteHandler(c echo.Context) error {
	id := c.Param("list_id")

	err := s.domain.DeleteList(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error deleting wishlist: %s", err))
	}

	return HxRedirect(c, "/admin/lists")
}

func (s *Server) ListsIndexHandler(c echo.Context) error {
	lists, err := s.queries.FilterLists(s.ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting list: %s", err))
	}

	return Render(c, list.WishlistIndexView(
		domain.ListIndex{
			NewWishlistURL: "/admin/lists/new",
			Lists:          utils.Map(lists, domain.ToList),
		},
	))
}

func (s *Server) ListsNewHandler(c echo.Context) error {
	return Render(c, list.WishlistNewView())
}

func (s *Server) ListsCreateHandler(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	id, err := domain.GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating id: %s", err))
	}

	shareId, err := domain.GenerateShareId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating share id: %s", err))
	}

	err = s.queries.CreateList(s.ctx, gateway.CreateListParams{
		ListID:      id,
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

	return c.Redirect(http.StatusFound, "/admin/lists")
}

func (s *Server) ItemsNewHandler(c echo.Context) error {
	id := c.Param("list_id")

	wishlist, err := s.queries.FindList(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, item.ItemNewView(domain.ToList(wishlist)))
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
	wishlistId := c.FormValue("list_id")

	wishlist, err := s.queries.FindList(s.ctx, wishlistId)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	itemInput, err := NewItemFormInput(c)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error parsing input: %s", err))
	}

	id, err := domain.GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating id: %s", err))
	}

	err = s.queries.CreateListItem(s.ctx, gateway.CreateListItemParams{
		ListItemID:  id,
		ListID:      wishlistId,
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

	return c.Redirect(http.StatusFound, fmt.Sprintf("/admin/lists/%s", wishlist.ListID))
}

func (s *Server) ItemsEditHandler(c echo.Context) error {
	id := c.Param("list_item_id")

	log.Info("id: ", id)

	listItem, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting item: %s", err))
	}

	wishlist, err := s.queries.FindList(s.ctx, listItem.ListID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return Render(c, item.ItemEditView(domain.ToList(wishlist), domain.ToItem(listItem)))
}

func (s *Server) ItemsUpdateHandler(c echo.Context) error {
	id := c.Param("list_item_id")

	listItem, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	itemInput, err := NewItemFormInput(c)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error parsing input: %s", err))
	}

	err = s.queries.UpdateItem(s.ctx, gateway.UpdateItemParams{
		ListItemID:  listItem.ListItemID,
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

	return HxRedirect(c, fmt.Sprintf("/admin/lists/%s", listItem.ListID))
}

func (s *Server) ItemsDeleteHandler(c echo.Context) error {
	id := c.Param("list_item_id")

	listItem, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	err = s.queries.DeleteItem(s.ctx, listItem.ListItemID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error updating wishlist: %s", err))
	}

	return HxRedirect(c, fmt.Sprintf("/admin/lists/%s", listItem.ListID))
}
