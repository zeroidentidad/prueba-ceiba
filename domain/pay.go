package domain

import (
	"fmt"
	"payments/dto"
	"payments/errs"
	"strconv"
	"time"
)

type Pay struct {
	DocumentoIdentificacion int       `db:"documentoIdentificacionArrendatario"`
	CodigoInmueble          string    `db:"codigoInmueble"`
	ValorPagado             int       `db:"valorPagado"`
	FechaPago               time.Time `db:"fechaPago"`
}

type PayCheck struct {
	Pagado   int `db:"pagado"`
	Restante int `db:"restante"`
	Abonos   int `db:"abonos"`
}

type PayStorage interface {
	InsertPay(Pay) *errs.AppError
	PayChecks(Pay) *PayCheck
	SelectPayments() (*[]Pay, *errs.AppError)
}

func NewPay(doc, cod, valor, fecha string) Pay {
	docId, _ := strconv.Atoi(doc)
	valorPagado, _ := strconv.Atoi(valor)
	fechaPago, _ := time.Parse("02/01/2006", fecha)
	return Pay{
		DocumentoIdentificacion: docId,
		CodigoInmueble:          cod,
		ValorPagado:             valorPagado,
		FechaPago:               fechaPago,
	}
}

func (p Pay) ToDto() dto.ResponsePay {
	docId := strconv.Itoa(p.DocumentoIdentificacion)
	valorPagado := strconv.Itoa(p.ValorPagado)
	fechaPago := p.FechaPago.Format("02/01/2006")
	return dto.ResponsePay{
		DocumentoIdentificacion: docId,
		CodigoInmueble:          p.CodigoInmueble,
		ValorPagado:             valorPagado,
		FechaPago:               fechaPago,
	}
}

func (p PayCheck) PayMessage(pago int) string {
	if p.Restante == 1000000 && p.Abonos == 0 && pago == 1000000 {
		return "gracias por pagar todo tu arriendo"
	}
	if p.Restante > 0 && p.Abonos >= 0 && pago+p.Pagado < 1000000 {
		return fmt.Sprintf("gracias por tu abono, sin embargo recuerda que te hace falta pagar $%d", p.Restante-pago)
	}
	if p.Restante > 0 && p.Abonos >= 0 && pago+p.Pagado == 1000000 {
		return "gracias por pagar todo tu arriendo"
	}

	return ""
}

func (p PayCheck) PayComplete() bool {
	if p.Restante == 0 && p.Abonos > 0 {
		return true
	}
	return false
}
