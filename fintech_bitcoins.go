package learngo

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Retirar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
