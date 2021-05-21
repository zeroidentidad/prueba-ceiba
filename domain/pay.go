package domain

import (
	"payments/dto"
	"payments/errs"
	"strconv"
)

//const dbTSLayout = "2006-01-02 15:04:05"

type Pay struct {
	DocumentoIdentificacion int    `db:"documentoIdentificacionArrendatario"`
	CodigoInmueble          string `db:"codigoInmueble"`
	ValorPagado             int    `db:"valorPagado"`
	FechaPago               string `db:"fechaPago"`
}

type PayStorage interface {
	InsertPay(Pay) (string, *errs.AppError)
	SelectPayments() (*[]Pay, *errs.AppError)
}

func NewPay(doc, cod, valor, fecha string) Pay {
	docId, _ := strconv.Atoi(doc)
	valorPagado, _ := strconv.Atoi(valor)
	return Pay{
		DocumentoIdentificacion: docId,
		CodigoInmueble:          cod,
		ValorPagado:             valorPagado,
		FechaPago:               fecha,
	}
}

func (p Pay) ToDto() dto.ResponsePay {
	docId := strconv.Itoa(p.DocumentoIdentificacion)
	valorPagado := strconv.Itoa(p.ValorPagado)
	return dto.ResponsePay{
		DocumentoIdentificacion: docId,
		CodigoInmueble:          p.CodigoInmueble,
		ValorPagado:             valorPagado,
		FechaPago:               p.FechaPago,
	}
}
