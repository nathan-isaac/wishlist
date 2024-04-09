package server

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
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

	checkoutResponse, err := s.queries.FindCheckoutResponse(s.ctx, id)

	if err != nil {
		checkoutResponse = gateway.CheckoutResponse{}
	}

	return Render(c, checkout.CheckoutShowView(checkout.CheckoutShowParams{
		Checkout: domain.Checkout{
			CheckoutId: checkoutRecord.Checkout.CheckoutID,
			CreatedAt:  checkoutRecord.Checkout.CreatedAt,
			UpdatedAt:  checkoutRecord.Checkout.UpdatedAt,
			List:       domain.ToList(checkoutRecord.List),
			CheckoutItems: utils.Map(itemRecords, func(t gateway.FilterCheckoutItemsRow) domain.CheckoutItem {
				return domain.CheckoutItem{
					ID:         t.CheckoutItem.CheckoutItemID,
					CheckoutID: t.CheckoutItem.CheckoutID,
					Quantity:   t.CheckoutItem.Quantity,
					CreatedAt:  t.CheckoutItem.CreatedAt,
					UpdatedAt:  t.CheckoutItem.UpdatedAt,
					Item:       domain.ToItem(t.ListItem),
				}
			}),
			Response: domain.CheckoutResponse{
				ID:             checkoutResponse.CheckoutResponseID,
				CheckoutID:     checkoutResponse.CheckoutID,
				Name:           checkoutResponse.Name,
				AddressLineOne: checkoutResponse.AddressLineOne,
				AddressLineTwo: checkoutResponse.AddressLineTwo,
				City:           checkoutResponse.City,
				State:          checkoutResponse.State,
				Zip:            checkoutResponse.Zip,
				Message:        checkoutResponse.Message,
			},
		},
	}))
}

type CheckoutCreateRequest struct {
	ListId     string `form:"list_id"`
	ItemId     string `form:"item_id"`
	CheckoutId string `form:"checkout_id"`
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

	if !s.domain.ListContainsItem(list.Items, req.ItemId) {
		return c.String(http.StatusBadRequest, "item not found")
	}

	checkoutResponse, err := s.queries.FindCheckout(s.ctx, req.CheckoutId)
	checkoutId := ""

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return c.String(http.StatusInternalServerError, "error fetching checkout id")
		}

		newId, err := domain.GenerateId()
		checkoutId = newId

		if err != nil {
			return c.String(http.StatusInternalServerError, "error generating checkout id")
		}

		err = s.queries.CreateCheckout(s.ctx, gateway.CreateCheckoutParams{
			CheckoutID: checkoutId,
			ListID:     req.ListId,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, "error creating checkout")
		}
	} else {
		checkoutId = checkoutResponse.Checkout.CheckoutID
		err = s.queries.UpdateCheckout(s.ctx, gateway.UpdateCheckoutParams{
			UpdatedAt:  time.Now(),
			CheckoutID: checkoutId,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, "error updating checkout")
		}
	}

	checkoutItem, err := s.queries.FindCheckoutItemByItemId(s.ctx, gateway.FindCheckoutItemByItemIdParams{
		CheckoutID: checkoutId,
		ListItemID: req.ItemId,
	})
	itemId := ""

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		newId, err := domain.GenerateId()
		itemId = newId

		if err != nil {
			return c.String(http.StatusInternalServerError, "error generating checkout item id")
		}

		err = s.queries.CreateCheckoutItem(s.ctx, gateway.CreateCheckoutItemParams{
			CheckoutItemID: itemId,
			CheckoutID:     checkoutId,
			ListItemID:     req.ItemId,
			Quantity:       1,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, "error creating checkout item")
		}
	} else {
		itemId = checkoutItem.CheckoutItem.CheckoutItemID

		err = s.queries.UpdateCheckoutItem(s.ctx, gateway.UpdateCheckoutItemParams{
			Quantity:       checkoutItem.CheckoutItem.Quantity + 1,
			UpdatedAt:      time.Now(),
			CheckoutItemID: itemId,
		})

		if err != nil {
			return err
		}
	}

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
			responseId, err := domain.GenerateId()

			if err != nil {
				return c.String(http.StatusInternalServerError, "error generating checkout id")
			}

			err = s.queries.CreateCheckoutResponse(s.ctx, gateway.CreateCheckoutResponseParams{
				CheckoutResponseID: responseId,
				CheckoutID:         checkoutRecord.Checkout.CheckoutID,
				Name:               req.Name,
				AddressLineOne:     req.AddressLineOne,
				AddressLineTwo:     req.AddressLineTwo,
				City:               req.City,
				State:              req.Region,
				Zip:                req.PostalCode,
				Message:            req.Message,
				CreatedAt:          time.Now(),
				UpdatedAt:          time.Now(),
			})

			if err != nil {
				return c.String(http.StatusInternalServerError, "error fetching checkout response")
			}
		} else {
			slog.Info("checkout response not found", "record", checkoutResponseRecord, "err", err)
			return c.String(http.StatusInternalServerError, "error fetching checkout response")
		}
	}

	err = s.queries.UpdateCheckoutResponse(s.ctx, gateway.UpdateCheckoutResponseParams{
		Name:               req.Name,
		AddressLineOne:     req.AddressLineOne,
		AddressLineTwo:     req.AddressLineTwo,
		City:               req.City,
		State:              req.Region,
		Zip:                req.PostalCode,
		Message:            req.Message,
		UpdatedAt:          time.Now(),
		CheckoutResponseID: checkoutResponseRecord.CheckoutResponseID,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "error updating checkout response")
	}

	redirectParams := url.Values{
		"checkoutId": {checkoutRecord.Checkout.CheckoutID},
	}

	return HxRedirect(c, fmt.Sprintf("/share/%s?%s", checkoutRecord.List.ShareCode, redirectParams.Encode()))
}

type CheckoutItemUpdateRequest struct {
	Id       string `param:"id"`
	Quantity string `form:"quantity"`
}

func (s *Server) CheckoutItemUpdateHandler(c echo.Context) error {
	var req CheckoutItemUpdateRequest
	err := c.Bind(&req)

	if err != nil {
		return err
	}

	checkoutItem, err := s.queries.FindCheckoutItem(s.ctx, req.Id)

	if err != nil {
		return err
	}

	quantity, err := strconv.ParseInt(req.Quantity, 10, 64)

	if err != nil {
		return err
	}

	if quantity <= 0 {
		err = s.queries.DeleteCheckoutItem(s.ctx, req.Id)

		if err != nil {
			return err
		}

		return HxRedirect(c, fmt.Sprintf("/checkout/%s", checkoutItem.CheckoutItem.CheckoutID))
	}

	err = s.queries.UpdateCheckoutItem(s.ctx, gateway.UpdateCheckoutItemParams{
		Quantity:       quantity,
		UpdatedAt:      time.Now(),
		CheckoutItemID: checkoutItem.CheckoutItem.CheckoutItemID,
	})

	if err != nil {
		return err
	}

	return HxRedirect(c, fmt.Sprintf("/checkout/%s", checkoutItem.CheckoutItem.CheckoutID))
}

// how does a user cancel a checkout?
// maybe we need a confirmation page?
