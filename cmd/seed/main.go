package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/auth"
	"github.com/gmtborges/orcamento-auto/db"
	"github.com/gmtborges/orcamento-auto/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error on loading .env: %v")
	}
	connStr := os.Getenv("DB_URL")
	db := db.Conn(connStr)
	cleanUp(db)

	org := types.Empresa{
		Nome:   "Org 1",
		Tipo:   types.EmpresaTipoOrg,
		Estado: "GO",
		Cidade: "Goiania",
	}

	tx := db.MustBegin()
	orgID := int64(0)
	err = tx.QueryRow(`INSERT INTO empresas 
  (nome, tipo, estado, cidade, data_criacao, data_atualizacao) 
  VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`, org.Nome, org.Tipo, org.Estado, org.Cidade).Scan(&orgID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert into empresas")
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	auto := types.Empresa{
		Nome:   "Auto 1",
		Tipo:   types.EmpresaTipoAuto,
		Estado: "GO",
		Cidade: "Goiania",
	}

	tx = db.MustBegin()
	autoID := int64(0)
	err = tx.QueryRow(`INSERT INTO empresas 
  (nome, tipo, estado, cidade, data_criacao, data_atualizacao) 
  VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`, auto.Nome, auto.Tipo, auto.Estado, auto.Cidade).Scan(&autoID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert into empresas")
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	hash, err := auth.GeneratePasswordHash("123")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate password hash")
	}
	usuarios := []types.Usuario{
		{
			Nome:      "User Org",
			Email:     "org@test.com",
			Senha:     hash,
			EmpresaID: sql.NullInt64{Int64: orgID, Valid: true},
		},
		{
			Nome:      "User Auto",
			Email:     "auto@test.com",
			Senha:     hash,
			EmpresaID: sql.NullInt64{Int64: autoID, Valid: true},
		},
		{
			Nome:      "User Standalone",
			Email:     "stand@test.com",
			EmpresaID: sql.NullInt64{Valid: false},
			Senha:     hash,
		},
	}

	tx = db.MustBegin()
	for i, u := range usuarios {
		var uID int64
		err = tx.QueryRow(`INSERT INTO usuarios 
	      (nome, email, senha, empresa_id, data_criacao, data_atualizacao) 
	      VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`, u.Nome, u.Email, u.Senha, u.EmpresaID).Scan(&uID)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to insert into usuarios")
		}
		usuarios[i].ID = uID
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	setUsuarioFuncoes(db, usuarios)
	autoCategoriaIDs := seedAutoCategorias(db)
	seedOrcamentos(db, orgID, autoCategoriaIDs)
}

func setUsuarioFuncoes(db *sqlx.DB, usuarios []types.Usuario) {
	var roleID int64
	err := db.Get(&roleID, `SELECT id FROM funcoes WHERE nome = 'ADMIN'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get funcao ID")
	}

	tx := db.MustBegin()
	for _, u := range usuarios {
		_, err = tx.Exec(`INSERT INTO usuarios_funcoes 
    (usuario_id, funcao_id, data_criacao, data_atualizacao) 
    VALUES ($1, $2, now(), now())`,
			u.ID, roleID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert into usuarios_funcoes")
		}
	}
	if err = tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}
}

func seedOrcamentos(db *sqlx.DB, orgID int64, categoriaIDs []int64) {
	var uID int64
	err := db.Get(&uID, `SELECT id FROM usuarios WHERE empresa_id = $1`, orgID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get usuario")
	}
	for i := 1; i <= 34; i++ {

		b := types.Orcamento{
			EmpresaID:     orgID,
			UsuarioID:     uID,
			AssociadoNome: fmt.Sprintf("Associado %d", i),
			VeiculoMarca:  "Ford",
			VeiculoNome:   fmt.Sprintf("veiculo %d", i),
			VeiculoAno:    1990 + i,
			VeiculoCor:    "Preto",
			Observacao:    "veiculo todo fudido",
			Status:        getRandomStatus(),
			DataCriacao:   time.Now().AddDate(0, 0, -i),
		}

		oi := []types.OrcamentoItem{
			{
				AutoCategoriaID: categoriaIDs[0],
				Observacao:      "",
				Status:          types.OrcamentoItemStatusAberto,
			},
			{
				AutoCategoriaID: categoriaIDs[1],
				Observacao:      "Uma observacao bem grande que passa de 30 caracteres.",
				Status:          types.OrcamentoItemStatusPropostaAceita,
			},
			{
				AutoCategoriaID: categoriaIDs[2],
				Observacao:      "Uma observacao bem grande que passa de 30 caracteres.",
				Status:          types.OrcamentoItemStatusPropostaRecebida,
			},
			{
				AutoCategoriaID: categoriaIDs[3],
				Observacao:      "",
				Status:          types.OrcamentoItemStatusAberto,
			},
			{
				AutoCategoriaID: categoriaIDs[4],
				Observacao:      "",
				Status:          types.OrcamentoItemStatusCancelado,
			},
			{
				AutoCategoriaID: categoriaIDs[5],
				Observacao:      "",
				Status:          types.OrcamentoItemStatusAberto,
			},
		}

		tx := db.MustBegin()
		var orcamentoID int64
		err := tx.QueryRow(`
		INSERT INTO orcamentos (empresa_id, usuario_id, associado_nome, veiculo_marca, veiculo_nome, 
    veiculo_ano, veiculo_cor, observacao, status, data_criacao, data_atualizacao)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now())
		RETURNING id`, b.EmpresaID, b.UsuarioID, b.AssociadoNome, b.VeiculoMarca, b.VeiculoNome, b.VeiculoAno,
			b.VeiculoCor, b.Observacao, b.Status, b.DataCriacao).Scan(&orcamentoID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert into orcamentos")
		}

		for _, item := range oi {
			item.OrcamentoID = orcamentoID
			_, err := tx.Exec(`
			INSERT INTO orcamento_itens (orcamento_id, auto_categoria_id, observacao, status)
			VALUES ($1, $2, $3, $4)`, item.OrcamentoID, item.AutoCategoriaID, item.Observacao, item.Status)
			if err != nil {
				tx.Rollback()
				log.Fatal().Err(err).Msg("Failed to insert into orcamento_itens")
			}
		}

		if err := tx.Commit(); err != nil {
			log.Fatal().Err(err).Msg("Failed to commit transaction")
		}
	}
}

func getRandomStatus() types.OrcamentoStatus {
	status := []types.OrcamentoStatus{
		types.OrcamentoStatusAguardandoProposta,
		types.OrcamentoStatusPendente,
		types.OrcamentoStatusFinalizado,
		types.OrcamentoStatusCancelado,
	}
	return status[rand.Intn(len(status))]
}

func seedAutoCategorias(db *sqlx.DB) []int64 {
	autoCategorias := []types.AutoCategoria{
		{Descricao: "Lanternagem", Tipo: types.AutoCategoriaTipoServico},
		{Descricao: "Pintura", Tipo: types.AutoCategoriaTipoServico},
		{Descricao: "Mecânica", Tipo: types.AutoCategoriaTipoServico},
		{Descricao: "Parachoque", Tipo: types.AutoCategoriaTipoProduto},
		{Descricao: "Retrovisor", Tipo: types.AutoCategoriaTipoProduto},
		{Descricao: "Porta", Tipo: types.AutoCategoriaTipoProduto},
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start transaction")
	}
	var autoCategoryIDs []int64

	for _, ac := range autoCategorias {
		var autoCategoryID int64
		err = tx.QueryRow(`
			INSERT INTO auto_categorias (descricao, tipo)
			VALUES ($1, $2)
			RETURNING id`, ac.Descricao, ac.Tipo).Scan(&autoCategoryID)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert into auto_categorias")
		}

		autoCategoryIDs = append(autoCategoryIDs, autoCategoryID)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	return autoCategoryIDs
}

func cleanUp(db *sqlx.DB) {
	_, err := db.Exec(`
		DELETE FROM usuarios_funcoes;
		DELETE FROM usuarios;
		DELETE FROM empresas;
		DELETE FROM auto_categorias;
		DELETE FROM orcamentos;
	`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to clean up database")
	}
}
