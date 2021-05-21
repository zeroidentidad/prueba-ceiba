package domain

import (
	"database/sql"
	"payments/errs"
	"payments/logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type PayStorageDb struct {
	client *sqlx.DB
}

func NewPayStorageDb(dbClient *sqlx.DB) PayStorageDb {
	return PayStorageDb{dbClient}
}

func (d PayStorageDb) InsertPay(pay Pay) (string, *errs.AppError) {
	i := "INSERT INTO pagos(documentoIdentificacionArrendatario, codigoInmueble, valorPagado, fechaPago) values(?,?,?,?)"

	_, err := d.client.Exec(i, pay.DocumentoIdentificacion, pay.CodigoInmueble, pay.ValorPagado, pay.FechaPago)
	if err != nil {
		logs.Error("Error inserting pago: " + err.Error())
		return "", errs.NewUnexpectedError("Unexpected error from database")
	}

	msg := "pago insertado"

	return msg, nil
}

func (d PayStorageDb) SelectPayments() (*[]Pay, *errs.AppError) {
	s := `SELECT documentoIdentificacionArrendatario, codigoInmueble, valorPagado, fechaPago FROM pagos`
	payments := make([]Pay, 0)

	err := d.client.Select(&payments, s)
	if err != nil {
		if err == sql.ErrNoRows {
			return &payments, errs.NewUnexpectedError("Payments not found")
		}
		logs.Error("Error querying table: " + err.Error())
		return &payments, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &payments, nil
}
