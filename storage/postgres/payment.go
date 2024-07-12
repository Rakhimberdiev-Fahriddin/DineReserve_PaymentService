package postgres

import (
	"database/sql"
	pb "payment-service/generated/payment_service"
)

type PaymentRepo struct {
	DB *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{DB: db}
}

func (p *PaymentRepo) CreatePayment(payment *pb.CreatePaymentRequest)(*pb.CreatePaymentResponse,error){
	paym := pb.Payment{}
	err := p.DB.QueryRow(`
		INSERT INTO Payments(
			reservation_id,
			amount,
			payment_method,
			payment_status
		)
		VALUES(
			$1,
			$2,
			$3,
			$4
		)
		returning
			id,
			reservation_id,
			amount,
			payment_method,
			payment_status`,
		payment.ReservationId,payment.Amount,payment.PaymentMethod,payment.PaymentStatus).Scan(
			&paym.Id, &paym.ReservationId, &paym.Amount, &paym.PaymentMethod, &paym.PaymentStatus,
		)		
	if err != nil{
		return nil,err
	}
	return &pb.CreatePaymentResponse{Payment: &paym},err
}

func (p *PaymentRepo) GetPayment(payment *pb.GetPaymentRequest)(*pb.GetPaymentResponse,error){
	resPayment := pb.Payment{}
	err := p.DB.QueryRow(`
		SELECT
			id,
			reservation_id,
			amount,
			payment_method,
			payment_status
		FROM
			Payments
		WHERE
			deleted_at = 0 and id = $1
		`,
		payment.Id).Scan(
			&resPayment.Id, &resPayment.ReservationId, &resPayment.Amount, &resPayment.PaymentMethod, &resPayment.PaymentStatus,
		)
	if err != nil{
		return nil,err
	}
	return &pb.GetPaymentResponse{
		Payment: &resPayment,
	},nil
}

func (p *PaymentRepo) UpdatePayment(updatePayment *pb.UpdatePaymentRequest)(*pb.UpdatePaymentResponse,error){
	resPayment := pb.Payment{}
	err := p.DB.QueryRow(`UPDATE
		Payments
		SET
			reservation_id = $1,
			amount = $2,
			payment_method = $3,
			payment_status = $4,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $5
		returning
			id,
			reservation_id,
			amount,
			payment_method,
			payment_status
		`,
		updatePayment.ReservationId,updatePayment.Amount,updatePayment.PaymentMethod,updatePayment.PaymentStatus,updatePayment.Id,
	).Scan(
		&resPayment.Id, &resPayment.ReservationId, &resPayment.Amount, &resPayment.PaymentMethod, &resPayment.PaymentStatus,
	)
	if err != nil{
		return nil,err
	}
	return &pb.UpdatePaymentResponse{
		Payment: &resPayment,
	},err
}