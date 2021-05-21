package domain

import (
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

type PayStorage interface {
	InsertPay(Pay) (string, *errs.AppError)
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
