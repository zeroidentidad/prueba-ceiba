package dto

import "payments/errs"

type RequestPay struct {
	DocumentoIdentificacion string `json:"documentoIdentificacionArrendatario"`
	CodigoInmueble          string `json:"codigoInmueble"`
	ValorPagado             string `json:"valorPagado"`
	FechaPago               string `json:"fechaPago"`
}

func (r RequestPay) EmptyPassword() bool {
	return r.CodigoInmueble == ""
}

func (r RequestPay) Validate() *errs.AppError {
	if r.EmptyPassword() {
		return errs.NewBadRequestError("The password must not be empty")
	}
	if len(r.CodigoInmueble) < 6 {
		return errs.NewBadRequestError("The password is too short")
	}

	return nil
}
