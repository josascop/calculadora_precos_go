package modelos

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

const (
	g  = "g"
	ml = "ml"
	un = "un"

	ingrediente = "ingrediente"
	operacional = "operacional"
)

type Insumo struct {
	ID                   uint64  `json:"id,omitempty"`
	Nome                 string  `json:"nome,omitempty"`
	PrecoUnitarioFloat   float64 `json:"preco_unitario,omitempty"`
	PrecoUnitarioDecimal decimal.Decimal
	Tipo                 string `json:"tipo,omitempty"`
	Unidade              string `json:"unidade,omitempty"`
}

func (i *Insumo) formatar() {
	i.Nome = strings.TrimSpace(strings.ToUpper(i.Nome))
	i.PrecoUnitarioDecimal = decimal.NewFromFloat(i.PrecoUnitarioFloat)
	i.Tipo = strings.TrimSpace(strings.ToLower((i.Tipo)))
	i.Unidade = strings.TrimSpace(strings.ToLower((i.Unidade)))
}

func (i *Insumo) Validar() error {
	i.formatar()

	if i.Nome == "" {
		return errors.New("informe um nome válido")
	}
	if len(i.Nome) < 3 {
		return errors.New("o nome deve ter ao menos três caracteres")
	}
	if !i.PrecoUnitarioDecimal.IsPositive() {
		return errors.New("informe um preço válido")
	}
	if !(i.Tipo == ingrediente || i.Tipo == operacional) {
		return errors.New("tipo inválido de produto")
	}
	if !(i.Unidade == g || i.Unidade == ml || i.Unidade == un) {
		return errors.New("unidade inválida")
	}

	return nil
}
