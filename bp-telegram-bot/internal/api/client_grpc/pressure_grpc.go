package client_grpc

import (
	"bp-telegram-bot/internal/model"
	"bp-telegram-bot/internal/pb/grpc-pressure/proto"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PressureGRPC struct {
	PressureClient proto.BloodPressureServiceClient
}

func NewPressureGRPC(conn *grpc.ClientConn) *PressureGRPC {
	pressureClient := proto.NewBloodPressureServiceClient(conn)

	return &PressureGRPC{
		PressureClient: pressureClient,
	}
}

func (c *PressureGRPC) CreatePressure(p model.BloodPressure) error {
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	msg, err := c.PressureClient.CreatePressure(ctx, &proto.BloodPressureRequest{
		ID:        p.ID,
		UserID:    p.UserID,
		Systolic:  p.Systolic,
		Diastolic: p.Diastolic,
		Pulse:     p.Pulse,
		TagNames:  p.TagNames,
		CreatedAt: timestamppb.New(p.CreatedAt),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("CreatePressure: %v", err))
	}
	if msg.GetStatus() != "1" {
		return errors.New(fmt.Sprintf("The CreatePressure operation ended with the GetStatus %v code.", msg.GetStatus()))
	}
	return nil
}

func (c *PressureGRPC) ListPressure(userId string) ([]*model.BloodPressure, error) {
	log.Println("ListPressure Out")
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	list, err := c.PressureClient.ListPressure(ctx, &proto.ListPressureRequest{
		UserID: userId,
	})
	log.Println(list, err)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ListPressure: %v", err))
	}
	if len(list.GetBPR()) == 0 {
		return nil, errors.New(fmt.Sprintln("ListPressure. The length of the array is zero."))
	}

	var press []*model.BloodPressure
	for _, pressure := range list.GetBPR() {
		bpr := &model.BloodPressure{
			ID:        pressure.ID,
			UserID:    pressure.UserID,
			Systolic:  pressure.Systolic,
			Diastolic: pressure.Diastolic,
			Pulse:     pressure.Pulse,
			TagNames:  pressure.TagNames,
			CreatedAt: pressure.CreatedAt.AsTime(),
		}

		press = append(press, bpr)
	}
	return press, err
}

func (c *PressureGRPC) UpdatePressure(p model.BloodPressure) error {
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	msg, err := c.PressureClient.UpdatePressure(ctx, &proto.BloodPressureRequest{
		ID:        p.ID,
		UserID:    p.UserID,
		Systolic:  p.Systolic,
		Diastolic: p.Diastolic,
		Pulse:     p.Pulse,
		TagNames:  p.TagNames,
		CreatedAt: timestamppb.New(p.CreatedAt),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("UpdatePressure: %v", err))
	}
	if msg.GetStatus() != "1" {
		return errors.New(fmt.Sprintf("The UpdatePressure operation ended with the GetStatus %v code.", msg.GetStatus()))
	}
	return nil
}

func (c *PressureGRPC) DeletePressure(id int64, userId string) error {
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	msg, err := c.PressureClient.DeletePressure(ctx, &proto.DeletePressureRequest{
		ID:     id,
		UserID: userId,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("DeletePressure: %v", err))
	}
	if msg.GetStatus() != "1" {
		return errors.New(fmt.Sprintf("The DeletePressure operation ended with the GetStatus %v code.", msg.GetStatus()))
	}
	return nil
}
