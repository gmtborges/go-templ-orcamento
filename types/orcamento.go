package types

import (
	"time"
)

type OrcamentoStatus string

const (
	OrcamentoStatusAguardandoProposta OrcamentoStatus = "AGUARDANDO_PROPOSTA"
	OrcamentoStatusPendente           OrcamentoStatus = "PENDENTE"
	OrcamentoStatusCancelado          OrcamentoStatus = "CANCELADO"
	OrcamentoStatusFinalizado         OrcamentoStatus = "FINALIZADO"
)

type OrcamentoItemStatus string

const (
	OrcamentoItemStatusAberto           OrcamentoItemStatus = "ABERTO"
	OrcamentoItemStatusPropostaRecebida OrcamentoItemStatus = "PROPOSTA_RECEBIDA"
	OrcamentoItemStatusPropostaAceita   OrcamentoItemStatus = "PROPOSTA_ACEITA"
	OrcamentoItemStatusCancelado        OrcamentoItemStatus = "CANCELADO"
)

type Orcamento struct {
	ID              int64
	EmpresaID       int64  `db:"empresa_id"`
	UsuarioID       int64  `db:"usuario_id"`
	AssociadoNome   string `db:"associado_nome" form:"associadoNome"`
	VeiculoMarca    string `db:"veiculo_marca" form:"veiculoMarca"`
	VeiculoNome     string `db:"veiculo_nome" form:"veiculoNome"`
	VeiculoAno      int    `db:"veiculo_ano" form:"veiculoAno"`
	VeiculoCor      string `db:"veiculo_cor" form:"veiculoCor"`
	Observacao      string `form:"observacao"`
	Status          OrcamentoStatus
	DataCriacao     time.Time `db:"data_criacao"`
	DataAtualizacao time.Time `db:"data_atualizacao"`
}

type OrcamentoItem struct {
	ID              int64
	OrcamentoID     int64  `db:"orcamento_id"`
	Observacao      string `json:"observacao"`
	Status          OrcamentoItemStatus
	AutoCategoriaID int64     `db:"auto_categoria_id" json:"autoCategoriaID"`
	DataCriacao     time.Time `db:"data_criacao"`
	DataAtualizacao time.Time `db:"data_atualizacao"`
}

type OrcamentoItemModel struct {
	OrcamentoItem
	AutoCategoriaDescricao string `db:"auto_categoria_descricao"`
	AutoCategoriaTipo      string `db:"auto_categoria_tipo"`
}

type OrcamentoModel struct {
	Orcamento
	Itens []OrcamentoItemModel
}

type OrcamentoResultSet struct {
	Count int
	Data  []OrcamentoModel
}

type OrcamentoIndexViewModel struct {
	Count       int
	CurrentPage int
	TotalPages  int
	SeqNumber   int
	Orcamentos  []OrcamentoModel
	Errors      []string
}

type OrcamentoCreateViewModel struct {
	OrcamentoModel
	AutoCategorias map[string][]AutoCategoria
	Errors         map[string]string
}

type OrcamentoShowViewModel struct {
	OrcamentoModel
	Errors map[string]string
}

type OrcamentoFiltros struct {
	Limit      int
	Offset     int
	OrderBy    string
	Order      string
	FilterBy   string
	SearchTerm string
}
