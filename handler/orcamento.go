package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/types"
	views "github.com/gmtborges/orcamento-auto/views/orcamento"
	"github.com/gmtborges/orcamento-auto/views/orcamento/partials"
)

type OrcamentoHandler struct {
	orcamentoSvc *svc.OrcamentoService
}

func NewOrcamentoHandler(orcamentoSvc *svc.OrcamentoService) *OrcamentoHandler {
	return &OrcamentoHandler{orcamentoSvc: orcamentoSvc}
}

func (h *OrcamentoHandler) Index(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("p"))
	searchTerm := c.QueryParam("q")
	if err != nil {
		page = 1
	}
	offset := (page - 1) * 10
	filtros := types.OrcamentoFiltros{
		Limit:      10,
		Offset:     offset,
		OrderBy:    "data_criacao",
		Order:      "asc",
		FilterBy:   "",
		SearchTerm: searchTerm,
	}
	empID := c.Get("empID").(int64)
	isHTMX := c.Request().Header.Get("HX-Request") == "true"
	result, err := h.orcamentoSvc.GetAllOrcamentos(c.Request().Context(), empID, filtros)
	if err != nil {
		vm := types.OrcamentoIndexViewModel{
			Errors: []string{"Erro ao buscar os orçamentos. Tente novamente mais tarde."},
		}
		return views.OrcamentoIndex(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.OrcamentoIndexViewModel{
		Count:       result.Count,
		CurrentPage: page,
		TotalPages:  (result.Count + filtros.Limit - 1) / filtros.Limit,
		SeqNumber:   (page-1)*filtros.Limit + 1,
		Orcamentos:  result.Data,
		Errors:      []string{},
	}
	if isHTMX {
		return partials.OrcamentoList(vm).Render(c.Request().Context(), c.Response())
	}
	return views.OrcamentoIndex(vm).Render(c.Request().Context(), c.Response())
}

func (h *OrcamentoHandler) Create(c echo.Context) error {
	// TODO: buscar as auto categorias via client no event de change da categoriaTipo
	ac, err := h.orcamentoSvc.GetAutoCategorias(c.Request().Context())
	vm := types.OrcamentoCreateViewModel{
		OrcamentoModel: types.OrcamentoModel{
			Orcamento: types.Orcamento{VeiculoAno: time.Now().Year()},
		},
		AutoCategorias: ac,
	}
	if err != nil {
		vm.Errors = map[string]string{"ac": "Não foi possível buscar os itens do orçamento. Tente novamente mais tarde."}
	}
	return views.OrcamentoCreate(vm).Render(c.Request().Context(), c.Response())
}

func (h *OrcamentoHandler) Save(c echo.Context) error {
	orcamento := new(types.Orcamento)
	err := c.Bind(orcamento)
	if err != nil {
		log.Error().Err(err).Msg("error on unmarshal orcamento")
		vm := types.OrcamentoCreateViewModel{
			OrcamentoModel: types.OrcamentoModel{
				Orcamento: types.Orcamento{VeiculoAno: time.Now().Year()},
			},
			Errors: map[string]string{"json": "Não foi possível salvar o orçamento. Tente novamente mais tarde."},
		}
		c.Response().WriteHeader(http.StatusBadRequest)
		c.Response().Header().Set("HX-Push-Url", "/orcamentos/novo")
		return views.OrcamentoCreate(vm).Render(c.Request().Context(), c.Response())
	}

	formItens := c.Request().FormValue("itens")
	var itens []struct {
		types.OrcamentoItem
	}
	err = json.Unmarshal([]byte(formItens), &itens)
	if err != nil {
		log.Error().Err(err).Msg("error on Unmarshal itens")
	}
	if len(itens) == 0 {
		vm := types.OrcamentoCreateViewModel{
			OrcamentoModel: types.OrcamentoModel{
				Orcamento: types.Orcamento{VeiculoAno: time.Now().Year()},
			},
			Errors: map[string]string{"itens": "Pelo menos um item de orçamento deve ser cadastrado."},
		}
		c.Response().WriteHeader(http.StatusBadRequest)
		c.Response().Header().Set("HX-Push-Url", "/orcamentos/novo")
		return views.OrcamentoCreate(vm).Render(c.Request().Context(), c.Response())
	}

	uID := c.Get("uID").(int64)
	empID := c.Get("empID").(int64)
	err = h.orcamentoSvc.CreateOrcamento(c.Request().Context(), uID, empID, *orcamento, itens)
	if err != nil {
		vm := types.OrcamentoCreateViewModel{
			OrcamentoModel: types.OrcamentoModel{
				Orcamento: types.Orcamento{VeiculoAno: time.Now().Year()},
			},
			Errors: map[string]string{"db": "Erro ao salvar orçamento. Tente novamente mais tarde."},
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Header().Set("HX-Push-Url", "/orcamentos/novo")
		return views.OrcamentoCreate(vm).Render(c.Request().Context(), c.Response())
	}

	return c.Redirect(http.StatusSeeOther, "/orcamentos")
}

func (h *OrcamentoHandler) Show(c echo.Context) error {
	orcamentoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		vm := types.OrcamentoShowViewModel{
			Errors: map[string]string{"404": "404"},
		}
		c.Response().WriteHeader(http.StatusNotFound)
		return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
	}
	orcamento, err := h.orcamentoSvc.GetOrcamento(c.Request().Context(), int64(orcamentoID))
	if err == sql.ErrNoRows {
		vm := types.OrcamentoShowViewModel{
			Errors: map[string]string{"404": "404"},
		}
		c.Response().WriteHeader(http.StatusNotFound)
		return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
	}
	if err != nil {
		vm := types.OrcamentoShowViewModel{
			Errors: map[string]string{"db": "Erro ao buscar o orçamento. Tente novamente mais tarde."},
		}
		return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.OrcamentoShowViewModel{
		OrcamentoModel: *orcamento,
	}
	return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
}

func (h *OrcamentoHandler) Edit(c echo.Context) error {
	orcamentoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		vm := types.OrcamentoShowViewModel{
			Errors: map[string]string{"404": "404"},
		}
		return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
	}
	orcamento, err := h.orcamentoSvc.GetOrcamento(c.Request().Context(), int64(orcamentoID))
	if err != nil {
		vm := types.OrcamentoShowViewModel{
			Errors: map[string]string{"db": "Erro ao buscar o orçamento. Tente novamente mais tarde."},
		}
		return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.OrcamentoShowViewModel{
		OrcamentoModel: *orcamento,
	}
	return views.OrcamentoShow(vm).Render(c.Request().Context(), c.Response())
}
