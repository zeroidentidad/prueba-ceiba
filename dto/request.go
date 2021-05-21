package dto

import (
	"payments/errs"
	"regexp"
	"strconv"
	"time"
)

type RequestPay struct {
	DocumentoIdentificacion string `json:"documentoIdentificacionArrendatario"`
	CodigoInmueble          string `json:"codigoInmueble"`
	ValorPagado             string `json:"valorPagado"`
	FechaPago               string `json:"fechaPago"`
}

func (r RequestPay) Validate() *errs.AppError {
	var validaciones string
	alfanumeric, _ := regexp.Match(`^[a-zA-Z0-9]*$`, []byte(r.CodigoInmueble))
	numeric, _ := regexp.Match(`^[0-9]*$`, []byte(r.DocumentoIdentificacion))
	date, date_err := time.Parse("02/01/2006", r.FechaPago)
	// date_format, _ := regexp.Match(`\d{1,2}/\d{1,2}/\d{4}`, []byte(r.FechaPago))
	valorPagado, _ := strconv.Atoi(r.ValorPagado)

	if !alfanumeric {
		validaciones += "\n'documentoIdentificacionArrendatario' debe ser solo númerico"
	} else if !numeric {
		validaciones += "\n'codigoInmueble' debe ser alfanúmerico"
	} else if valorPagado < 1 || valorPagado > 1000000 {
		validaciones += "\n'valorPagado' deberá estar entre 1 y 1000000"
	} else if date_err != nil {
		validaciones += "\n'fechaPago' formato de fecha incorrecto"
	} else if date.Year() != time.Now().Year() {
		validaciones += "\n'fechaPago' debe ser fecha válida existente en el calendario"
	} else if date.Day()%2 == 0 {
		validaciones += "\n'fechaPago' lo siento pero no se puede recibir el pago este día por decreto de administración"
	}

	if validaciones != "" {
		return errs.NewBadRequestError(validaciones)
	}

	return nil
}
