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

func (d PayStorageDb) InsertPay(pay Pay) *errs.AppError {
	sqlInsert := "INSERT INTO pagos(documentoIdentificacionArrendatario, codigoInmueble, valorPagado, fechaPago) values(?,?,?,?)"

	_, err := d.client.Exec(sqlInsert, pay.DocumentoIdentificacion, pay.CodigoInmueble, pay.ValorPagado, pay.FechaPago)
	if err != nil {
		logs.Error("Error inserting pago: " + err.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}

func (d PayStorageDb) PayChecks(pay Pay) (paycheck *PayCheck) {
	sqlAbonos := `SELECT COUNT(valorPagado) AS abonos FROM pagos 
	WHERE (SELECT MONTH(fechaPago)) = ?
	AND (SELECT YEAR(fechaPago)) = ?
	AND documentoIdentificacionArrendatario = ?
	AND codigoInmueble = ?`

	sqlPagado := `SELECT t.pagado, (1000000-t.pagado) AS restante 
	FROM (SELECT documentoIdentificacionArrendatario, codigoInmueble, SUM(valorPagado) AS pagado
	FROM pagos WHERE (SELECT MONTH(fechaPago)) = ?
	AND (SELECT YEAR(fechaPago)) = ?
	AND documentoIdentificacionArrendatario = ?
	AND codigoInmueble = ?
	GROUP BY documentoIdentificacionArrendatario, codigoInmueble) t`

	var abonos PayCheck
	err := d.client.Get(&abonos, sqlAbonos, pay.FechaPago.Month(), pay.FechaPago.Year(), pay.DocumentoIdentificacion, pay.CodigoInmueble)
	if err != nil {
		abonos.Pagado = 0
		abonos.Restante = 0
		abonos.Abonos = 0
	}

	var pagado PayCheck
	err = d.client.Get(&pagado, sqlPagado, pay.FechaPago.Month(), pay.FechaPago.Year(), pay.DocumentoIdentificacion, pay.CodigoInmueble)
	if err != nil {
		pagado.Pagado = 0
		pagado.Restante = 1000000
		pagado.Abonos = 0
	}

	paychecks := PayCheck{
		Pagado:   pagado.Pagado,
		Restante: pagado.Restante,
		Abonos:   abonos.Abonos,
	}

	return &paychecks
}

func (d PayStorageDb) SelectPayments() (*[]Pay, *errs.AppError) {
	sqlSelect := `SELECT * FROM pagos ORDER BY documentoIdentificacionArrendatario, codigoInmueble, fechaPago ASC`
	payments := make([]Pay, 0)

	err := d.client.Select(&payments, sqlSelect)
	if err != nil {
		if err == sql.ErrNoRows {
			return &payments, errs.NewUnexpectedError("Payments not found")
		}
		logs.Error("Error querying table: " + err.Error())
		return &payments, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &payments, nil
}
