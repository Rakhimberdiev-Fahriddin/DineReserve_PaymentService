package postgres

import (
	"log"
	pb "payment-service/generated/payment_service"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePayment(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	payment := NewPaymentRepo(db)

	resp, err := payment.CreatePayment(&pb.CreatePaymentRequest{
		ReservationId: "e38c1bb5-98b4-4af7-b895-d2a2aa3da99c",
		Amount:        30.5,
		PaymentMethod: "Credit Card",
		PaymentStatus: "Successful",
	})

	if err != nil {
		log.Fatal(err)
	}

	expectedResponse := &pb.CreatePaymentResponse{
		Payment: &pb.Payment{
			Id:            "generated-uuid",
			ReservationId: "e38c1bb5-98b4-4af7-b895-d2a2aa3da99c",
			Amount:        30.5,
			PaymentMethod: "Credit Card",
			PaymentStatus: "Successful",
		},
	}

	assert.NoError(t, err)

	// Compare the fields individually
	assert.Equal(t, expectedResponse.Payment.ReservationId, resp.Payment.ReservationId)
	assert.Equal(t, expectedResponse.Payment.Amount, resp.Payment.Amount)
	assert.Equal(t, expectedResponse.Payment.PaymentMethod, resp.Payment.PaymentMethod)
	assert.Equal(t, expectedResponse.Payment.PaymentStatus, resp.Payment.PaymentStatus)

	// Ensure that the ID is set (UUID format check can be done here)
	assert.NotEmpty(t, resp.Payment.Id)

}

func TestGetPayment(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	payment := NewPaymentRepo(db)

	resp, err := payment.GetPayment(&pb.GetPaymentRequest{
		Id: "01f553d0-db6b-4204-a109-e0d605df82ea",
	})

	if err != nil {
		log.Fatal(err)
	}

	expectedResponse := pb.GetPaymentResponse{
		Payment: &pb.Payment{
			Id:            "01f553d0-db6b-4204-a109-e0d605df82ea",
			ReservationId: "e38c1bb5-98b4-4af7-b895-d2a2aa3da99c",
			Amount:        30.5,
			PaymentMethod: "Credit Card",
			PaymentStatus: "Successful",
		},
	}

	if !reflect.DeepEqual(resp, &expectedResponse) {
		t.Errorf("have %v , wont %v", resp, &expectedResponse)
	}
}

func TestUpdatePayment(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	payment := NewPaymentRepo(db)

	resp, err := payment.UpdatePayment(&pb.UpdatePaymentRequest{
		Id: "0d9b32c8-2a80-4445-afa4-3d82175e8b13",
		ReservationId: "0963662b-9789-4f1d-8d80-cea06d4b4d28",
		Amount: 150.0,
		PaymentMethod: "Credit Card",
		PaymentStatus: "Successful",
	})

	if err != nil {
		log.Fatal(err)
	}

	expectedResponse := pb.UpdatePaymentResponse{
		Payment: &pb.Payment{
			Id:            "0d9b32c8-2a80-4445-afa4-3d82175e8b13",
			ReservationId: "0963662b-9789-4f1d-8d80-cea06d4b4d28",
			Amount:        150.0,
			PaymentMethod: "Credit Card",
			PaymentStatus: "Successful",
		},
	}

	if !reflect.DeepEqual(resp, &expectedResponse) {
		t.Errorf("have %v , wont %v", resp, &expectedResponse)
	}
}
