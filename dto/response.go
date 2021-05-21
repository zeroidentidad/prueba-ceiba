package dto

type ResponsePay struct {
	DocumentoIdentificacion string `json:"documentoIdentificacionArrendatario"`
	CodigoInmueble          string `json:"codigoInmueble"`
	ValorPagado             string `json:"valorPagado"`
	FechaPago               string `json:"fechaPago"`
}
