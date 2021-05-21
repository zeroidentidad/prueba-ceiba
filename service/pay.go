package service

import (
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

func (s DefaultRoleService) Pay(req dto.RequestPay) (msg string, err *errs.AppError) {
	r := domain.NewPay(req.ID, req.Name)

	message, err := s.repo.InsertPay(r)
	if err != nil {
		return msg, err
	}

	return message, err
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
