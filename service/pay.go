package service

import (
	"payments/domain"
	"payments/dto"
	"payments/errs"
)

type PayService interface {
	Pay(dto.RequestPay) (string, *errs.AppError)
	Payments() (*[]dto.ResponsePay, *errs.AppError)
}

type DefaultPayService struct {
	repo domain.PayStorage
}

func NewPayService(repo domain.PayStorage) DefaultPayService {
	return DefaultPayService{
		repo,
	}
}

func (s DefaultPayService) Pay(req dto.RequestPay) (msg string, err *errs.AppError) {
	err = req.Validate()
	if err != nil {
		return msg, err
	}

	r := domain.NewPay(req.DocumentoIdentificacion, req.CodigoInmueble, req.ValorPagado, req.FechaPago)

	paychecks := s.repo.PayChecks(r)
	msg = paychecks.PayMessage(r.ValorPagado)
	if !paychecks.PayComplete() {
		err := s.repo.InsertPay(r)
		if err != nil {
			return msg, err
		}
	} else {
		msg = "no procede pago, su mensualidad ya esta pagada comenpletamente"
	}

	return msg, err
}

func (s DefaultPayService) Payments() (*[]dto.ResponsePay, *errs.AppError) {
	res := make([]dto.ResponsePay, 0)
	payments, err := s.repo.SelectPayments()
	if err != nil {
		return &res, err
	}

	for _, p := range *payments {
		res = append(res, p.ToDto())
	}

	return &res, err
}
