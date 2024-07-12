package service

import (
	"context"
	"log/slog"
	pb "payment-service/generated/payment_service"
	"payment-service/storage/postgres"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
	Logger  *slog.Logger
	Payment postgres.PaymentRepo
}

func NewPaymentService(payment postgres.PaymentRepo, logger *slog.Logger) *PaymentService {
	return &PaymentService{Payment: payment, Logger: logger}
}

func (s *PaymentService) CreatePayment(ctx context.Context, payment *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	s.Logger.Info("Request received to CreatePayment")
	res, err := s.Payment.CreatePayment(payment)
	if err != nil {
		s.Logger.Error("Failed to create payment", "error", err.Error())
		return nil, err
	}
	s.Logger.Info("Create payment successfully")
	return res, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, payment *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	s.Logger.Info("Request received to GetPayment")
	res, err := s.Payment.GetPayment(payment)
	if err != nil {
		s.Logger.Error("Failed to get payment", "error", err.Error())
		return nil, err
	}
	s.Logger.Info("Get payment successfully")
	return res, nil
}

func (s *PaymentService) UpdatePayment(ctx context.Context, updatePayment *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error) {
	s.Logger.Info("Request received to UpdatePayment")
	res, err := s.Payment.UpdatePayment(updatePayment)
	if err != nil {
		s.Logger.Error("Failed to update payment", "error", err.Error())
		return nil, err
	}
	s.Logger.Info("Update payment successfully")
	return res, nil
}
