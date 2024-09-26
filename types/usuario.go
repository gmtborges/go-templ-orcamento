package types

import (
	"database/sql"
	"time"
)

type Usuario struct {
	ID              int64
	Email           string
	Senha           string
	Nome            string
	EmpresaID       sql.NullInt64 `db:"empresa_id"`
	DataCriacao     time.Time     `db:"data_criacao"`
	DataAtualizacao time.Time     `db:"data_atualizacao"`
}

type Funcao struct {
	ID   int64
	Nome string
}

type UsuarioAutenticacao struct {
	ID        int64
	EmpresaID int64 `db:"empresa_id"`
	Nome      string
	Senha     string
	Funcoes   []string
}
