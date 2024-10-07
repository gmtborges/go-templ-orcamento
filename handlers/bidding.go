package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/services"
	"github.com/gmtborges/orcamento-auto/types"
	views "github.com/gmtborges/orcamento-auto/views/bidding"
	"github.com/gmtborges/orcamento-auto/views/bidding/partials"
)

type BiddingHandler struct {
	biddingSvc *services.BiddingService
}

func NewBiddingHandler(biddingSvc *services.BiddingService) *BiddingHandler {
	return &BiddingHandler{biddingSvc: biddingSvc}
}

func (h *BiddingHandler) Index(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("p"))
	searchTerm := c.QueryParam("q")
	if err != nil {
		page = 1
	}
	offset := (page - 1) * 10
	filters := types.BiddingFilters{
		Limit:      10,
		Offset:     offset,
		OrderBy:    "created_at",
		Order:      "asc",
		FilterBy:   "",
		SearchTerm: searchTerm,
	}
	companyID := c.Get("companyID").(int64)
	isHTMX := c.Request().Header.Get("HX-Request") == "true"
	result, err := h.biddingSvc.GetAllBiddingsByCompanyID(c.Request().Context(), companyID, filters)
	if err != nil {
		vm := types.BiddingIndexViewModel{
			Errors: []string{"Erro ao buscar os orçamentos. Tente novamente mais tarde."},
		}
		return views.BiddingIndex(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.BiddingIndexViewModel{
		Count:       result.Count,
		CurrentPage: page,
		TotalPages:  (result.Count + filters.Limit - 1) / filters.Limit,
		SeqNumber:   (page-1)*filters.Limit + 1,
		Biddings:    result.Data,
		Errors:      []string{},
	}
	if isHTMX {
		return partials.BiddingList(vm).Render(c.Request().Context(), c.Response())
	}
	return views.BiddingIndex(vm).Render(c.Request().Context(), c.Response())
}

func (h *BiddingHandler) New(c echo.Context) error {
	// TODO: Search for auto categories client-side in the categoryType onchange event.
	ac, err := h.biddingSvc.GetAutoCategories(c.Request().Context())
	vm := types.BiddingCreateViewModel{
		BiddingModel: types.BiddingModel{
			Bidding: types.Bidding{VehicleYear: time.Now().Year()},
		},
		AutoCategories: ac,
	}
	if err != nil {
		vm.Errors = map[string]string{"ac": "Não foi possível buscar os itens do orçamento. Tente novamente mais tarde."}
	}
	return views.BiddingNew(vm).Render(c.Request().Context(), c.Response())
}

func (h *BiddingHandler) Create(c echo.Context) error {
	bidding := new(types.Bidding)
	err := c.Bind(bidding)
	if err != nil {
		log.Error().Err(err).Msg("error on unmarshal bidding")
		vm := types.BiddingCreateViewModel{
			BiddingModel: types.BiddingModel{
				Bidding: types.Bidding{VehicleYear: time.Now().Year()},
			},
			Errors: map[string]string{"json": "Não foi possível salvar o orçamento. Tente novamente mais tarde."},
		}
		c.Response().WriteHeader(http.StatusBadRequest)
		c.Response().Header().Set("HX-Push-Url", "/orcamentos/novo")
		return views.BiddingNew(vm).Render(c.Request().Context(), c.Response())
	}

	formItems := c.Request().FormValue("items")
	var items []struct {
		types.BiddingItem
	}
	err = json.Unmarshal([]byte(formItems), &items)
	if err != nil {
		log.Error().Err(err).Msg("error on Unmarshal itens")
	}
	if len(items) == 0 {
		vm := types.BiddingCreateViewModel{
			BiddingModel: types.BiddingModel{
				Bidding: types.Bidding{VehicleYear: time.Now().Year()},
			},
			Errors: map[string]string{"itens": "Pelo menos um item de orçamento deve ser cadastrado."},
		}
		c.Response().WriteHeader(http.StatusBadRequest)
		c.Response().Header().Set("HX-Push-Url", "/orcamentos/novo")
		return views.BiddingNew(vm).Render(c.Request().Context(), c.Response())
	}

	userID := c.Get("userID").(int64)
	companyID := c.Get("companyID").(int64)
	err = h.biddingSvc.CreateBidding(c.Request().Context(), userID, companyID, *bidding, items)
	if err != nil {
		vm := types.BiddingCreateViewModel{
			BiddingModel: types.BiddingModel{
				Bidding: types.Bidding{VehicleYear: time.Now().Year()},
			},
			Errors: map[string]string{"db": "Erro ao salvar orçamento. Tente novamente mais tarde."},
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Header().Set("HX-Push-Url", "/orcamentos/novo")
		return views.BiddingNew(vm).Render(c.Request().Context(), c.Response())
	}

	return c.Redirect(http.StatusSeeOther, "/orcamentos")
}

func (h *BiddingHandler) Show(c echo.Context) error {
	biddingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		vm := types.BiddingShowViewModel{
			Errors: map[string]string{"404": "404"},
		}
		c.Response().WriteHeader(http.StatusNotFound)
		return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
	}
	bidding, err := h.biddingSvc.GetBidding(c.Request().Context(), int64(biddingID))
	if err == sql.ErrNoRows {
		vm := types.BiddingShowViewModel{
			Errors: map[string]string{"404": "404"},
		}
		c.Response().WriteHeader(http.StatusNotFound)
		return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
	}
	if err != nil {
		vm := types.BiddingShowViewModel{
			Errors: map[string]string{"db": "Erro ao buscar o orçamento. Tente novamente mais tarde."},
		}
		return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.BiddingShowViewModel{
		BiddingModel: *bidding,
	}
	return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
}

func (h *BiddingHandler) Edit(c echo.Context) error {
	biddingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		vm := types.BiddingShowViewModel{
			Errors: map[string]string{"404": "404"},
		}
		return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
	}
	bidding, err := h.biddingSvc.GetBidding(c.Request().Context(), int64(biddingID))
	if err != nil {
		vm := types.BiddingShowViewModel{
			Errors: map[string]string{"db": "Erro ao buscar o orçamento. Tente novamente mais tarde."},
		}
		return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.BiddingShowViewModel{
		BiddingModel: *bidding,
	}
	return views.BiddingShow(vm).Render(c.Request().Context(), c.Response())
}
