package server_grpc

import (
	"bp-service/internal/model"
	"bp-service/internal/pb/grpc-pressure/proto"
	"bp-service/internal/service"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PressureGRPC struct {
	proto.UnimplementedBloodPressureServiceServer
	MU     sync.Mutex
	svc    *service.Service
	buffer map[string]map[string]any
}

func NewPressureGRPC(svc *service.Service) *PressureGRPC {
	return &PressureGRPC{
		svc: svc,
	}
}

func (s *PressureGRPC) CreatePressure(ctx context.Context, req *proto.BloodPressureRequest) (*proto.PressureResponse, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	p := &model.BloodPressure{
		ID:        req.ID,
		UserID:    req.UserID,
		Systolic:  req.Systolic,
		Diastolic: req.Diastolic,
		Pulse:     req.Pulse,
		TagNames:  req.TagNames,
		CreatedAt: req.CreatedAt.AsTime(),
	}

	err := s.svc.Pressure.Create(ctx, p)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("pressure create error: %v", err))
	}

	msg := &proto.PressureResponse{
		Status: "1",
	}

	fmt.Println("Создано сообщение:", msg)
	fmt.Println("\n")

	return msg, nil
}

func (s *PressureGRPC) UpdatePressure(ctx context.Context, req *proto.BloodPressureRequest) (*proto.PressureResponse, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	p := &model.BloodPressure{
		ID:        req.ID,
		UserID:    req.UserID,
		Systolic:  req.Systolic,
		Diastolic: req.Diastolic,
		Pulse:     req.Pulse,
		TagNames:  req.TagNames,
		CreatedAt: req.CreatedAt.AsTime(),
	}

	fmt.Println("UpdateMessage", p)

	msg := &proto.PressureResponse{
		Status: "1",
	}

	fmt.Println("Обновлено сообщение:", msg)
	fmt.Println("\n")

	return msg, nil
}

func (s *PressureGRPC) DeletePressure(ctx context.Context, req *proto.DeletePressureRequest) (*proto.PressureResponse, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	pressureIdInt := strconv.FormatInt(req.ID, 10)
	fmt.Println("MethodDelete", req.UserID, pressureIdInt)

	keyStr := fmt.Sprintf("LastPressureList%v", req)

	pressure, ok := s.buffer[req.UserID][keyStr].(*model.BloodPressure)
	if !ok {
		return nil, fmt.Errorf("DeletePressure: user buffer not found")
	}
	fmt.Println("pressure", pressure)

	if s.svc.Pressure.Delete(ctx, pressure.ID, req.UserID) != nil {
		return &proto.PressureResponse{Status: "0"}, errors.New("pressure delete error")
	}

	msg := &proto.PressureResponse{
		Status: "1",
	}

	fmt.Println("Удалено сообщение:", req, msg)
	fmt.Println("\n")

	return msg, nil
}

func (s *PressureGRPC) ListPressure(ctx context.Context, req *proto.ListPressureRequest) (*proto.ListBloodPressureRequest, error) {

	log.Println("Start ListPressure")
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	pressures, err := s.svc.Pressure.List(ctx, req.UserID)
	if err != nil {
		log.Println("List. Pressure list error", err)
		return nil, errors.New(fmt.Sprintf("List. Pressure list error: %v", err))
	}

	if s.buffer == nil {
		s.buffer = make(map[string]map[string]any)
	}

	if s.buffer[req.UserID] == nil {
		s.buffer[req.UserID] = map[string]any{}
	}

	var arrBPR []*proto.BloodPressureRequest
	for id, pressure := range pressures {
		keyStr := fmt.Sprintf("LastPressureList%v", id)
		s.buffer[req.UserID][keyStr] = pressure

		bpr := &proto.BloodPressureRequest{
			ID:        pressure.ID,
			UserID:    pressure.UserID,
			Systolic:  pressure.Systolic,
			Diastolic: pressure.Diastolic,
			Pulse:     pressure.Pulse,
			TagNames:  pressure.TagNames,
			CreatedAt: timestamppb.New(pressure.CreatedAt),
		}

		arrBPR = append(arrBPR, bpr)
	}

	lBPR := &proto.ListBloodPressureRequest{
		BPR: arrBPR,
	}
	fmt.Println("Отправлен список: ", lBPR)
	fmt.Println("\n")

	return lBPR, nil
}
