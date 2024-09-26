package types

import "time"

type Proposta struct {
	ID              int64
	OrcamentoItemID int64 `db:"orcamento_item_id"`
	EmpresaID       int64 `db:"empresa_id"`
	Detalhes        string
	Preco           float32
	Aceita          bool
	DataCriacao     time.Time `db:"data_criacao"`
	DataAtualizacao time.Time `db:"data_atualizacao"`
}

type PropostaModel struct {
	Proposta
	EmpresaNome string `db:"empresa_nome"`
}

type PropostaOrcamentoItensViewModel struct {
	Propostas []PropostaModel
	Errors    map[string]string
}

type AutoCategoriaTipo string

const (
	AutoCategoriaTipoProduto AutoCategoriaTipo = "PRODUTO"
	AutoCategoriaTipoServico AutoCategoriaTipo = "SERVICO"
)

type AutoCategoria struct {
	ID        int64
	Descricao string
	Tipo      AutoCategoriaTipo
}
