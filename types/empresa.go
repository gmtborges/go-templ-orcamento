package types

import "time"

type EmpresaTipo string

const (
	EmpresaTipoOrg  EmpresaTipo = "ORG"
	EmpresaTipoAuto EmpresaTipo = "AUTO"
)

type Empresa struct {
	ID              int64
	Nome            string
	Tipo            EmpresaTipo
	Estado          string
	Cidade          string
	Endereco        string
	NumeroTelefone  string    `db:"numero_telefone"`
	DataCriacao     time.Time `db:"data_criacao"`
	DataAtualizacao time.Time `db:"data_atualizacao"`
}
