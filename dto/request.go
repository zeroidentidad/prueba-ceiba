package dto

import (
	"payments/errs"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RequestPay struct {
	DocumentoIdentificacion string `json:"documentoIdentificacionArrendatario"`
	CodigoInmueble          string `json:"codigoInmueble"`
	ValorPagado             string `json:"valorPagado"`
	FechaPago               string `json:"fechaPago"`
}

func (r RequestPay) Validate() *errs.AppError {
	var validaciones strings.Builder
	alfanumeric, _ := regexp.Match(`^[a-zA-Z0-9]*$`, []byte(r.CodigoInmueble))
	numeric, _ := regexp.Match(`^[0-9]*$`, []byte(r.DocumentoIdentificacion))
	date, date_err := time.Parse("02/01/2006", r.FechaPago)
	// date, _ := regexp.Match(`\d{1,2}/\d{1,2}/\d{4}`, []byte(r.FechaPago))
	valorPagado, _ := strconv.Atoi(r.ValorPagado)

	if !alfanumeric {
		validaciones.WriteString(`“documentoIdentificacionArrendatario” debe ser solo númerico `)
	}
	if !numeric {
		validaciones.WriteString(`“codigoInmueble” debe ser alfanúmerico `)
	}
	if valorPagado < 1 || valorPagado > 1000000 {
		validaciones.WriteString(`“valorPagado” deberá estar entre 1 y 1000000 `)
	}
	if date_err != nil {
		validaciones.WriteString(`“fechaPago” formato de fecha incorrecto `)
	}
	if date.Year() != time.Now().Year() {
		validaciones.WriteString(`“fechaPago” debe ser fecha válida existente en el calendario `)
	}
	if date.Day()%2 == 0 {
		validaciones.WriteString(`“fechaPago” lo siento pero no se puede recibir el pago este día por decreto de administración `)
	}

	if validaciones.String() != "" {
		return errs.NewBadRequestError(validaciones.String())
	}

	return nil
}
