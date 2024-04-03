package server

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"time"
	"wishlist/internal/domain"
	"wishlist/internal/gateway"
	"wishlist/internal/utils"
	"wishlist/internal/views/checkout"
)

func (s *Server) CheckoutShowHandler(c echo.Context) error {
	id := c.Param("id")

	slog.Info("checkout show handler", slog.String("id", id))

	checkoutRecord, err := s.queries.FindCheckout(s.ctx, id)

	if err != nil {
		return c.String(http.StatusNotFound, "checkout not found")
	}

	itemRecords, err := s.queries.FilterCheckoutItems(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "error fetching checkout items")
	}

	return Render(c, checkout.CheckoutShowView(checkout.CheckoutShowParams{
		Checkout: domain.Checkout{
			ID:        checkoutRecord.Checkout.ID,
			CreatedAt: checkoutRecord.Checkout.CreatedAt,
			UpdatedAt: checkoutRecord.Checkout.UpdatedAt,
			List:      domain.ToList(checkoutRecord.List),
			CheckoutItems: utils.Map(itemRecords, func(t gateway.FilterCheckoutItemsRow) domain.CheckoutItem {
				return domain.CheckoutItem{
					ID:         t.CheckoutItem.ID,
					CheckoutID: t.CheckoutItem.CheckoutID,
					Quantity:   t.CheckoutItem.Quantity,
					CreatedAt:  t.CheckoutItem.CreatedAt,
					UpdatedAt:  t.CheckoutItem.UpdatedAt,
					Item:       domain.ToItem(t.ListItem),
				}
			}),
			Response: domain.CheckoutResponse{},
		},
	}))
}

type CheckoutCreateRequest struct {
	ListId string `form:"list_id"`
	ItemId string `form:"item_id"`
}

func (s *Server) CheckoutCreateHandler(c echo.Context) error {
	var req CheckoutCreateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	slog.Info("checkout show handler", "req", req)

	list, err := s.domain.FindList(req.ListId)

	if err != nil {
		return c.String(http.StatusBadRequest, "list not found")
	}

	for _, item := range list.Items {
		if item.Id != req.ItemId {
			return c.String(http.StatusBadRequest, "item not found")
		}
	}

	checkoutId, err := domain.GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, "error generating checkout id")
	}

	err = s.queries.CreateCheckout(s.ctx, gateway.CreateCheckoutParams{
		ID:        checkoutId,
		ListID:    req.ListId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "error creating checkout")
	}

	itemId, err := domain.GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, "error generating checkout item id")
	}

	err = s.queries.CreateCheckoutItem(s.ctx, gateway.CreateCheckoutItemParams{
		ID:         itemId,
		CheckoutID: checkoutId,
		ListItemID: req.ItemId,
		Quantity:   1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	return HxRedirect(c, "/checkout/"+checkoutId)
}

type CheckoutUpdateRequest struct {
	Id             string `param:"id"`
	Name           string `form:"name"`
	AddressLineOne string `form:"address_line_one"`
	AddressLineTwo string `form:"address_line_two"`
	City           string `form:"city"`
	Region         string `form:"region"`
	PostalCode     string `form:"postal_code"`
	Anonymous      bool   `form:"anonymous"`
	Message        string `form:"message"`
}

func (s *Server) CheckoutUpdateHandler(c echo.Context) error {
	var req CheckoutUpdateRequest
	err := c.Bind(&req)
	if err != nil {
		slog.Info("checkoutRecord update handler", "err", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	slog.Info("checkoutRecord show handler", "req", req, "id", req.Id)

	checkoutRecord, err := s.queries.FindCheckout(s.ctx, req.Id)

	if err != nil {
		slog.Info("checkoutRecord not found", "record", checkoutRecord, "err", err)
		return c.String(http.StatusBadRequest, "checkoutRecord not found")
	}

	checkoutResponseRecord, err := s.queries.FindCheckoutResponse(s.ctx, req.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// create a new checkout response
		} else {
			slog.Info("checkout response not found", "record", checkoutResponseRecord, "err", err)
			return c.String(http.StatusInternalServerError, "error fetching checkout response")
		}
	}

	// update the checkout response

	return HxRedirect(c, fmt.Sprintf("/share/%s", checkoutRecord.List.ShareCode))
}

// how does a user cancel a checkout?
// maybe we need a confirmation page?
