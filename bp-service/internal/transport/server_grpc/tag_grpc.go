package server_grpc

import (
	"bp-service/internal/model"
	"bp-service/internal/pb/grpc-tag/proto"
	"bp-service/internal/service"
	"context"
	"errors"
	"fmt"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagGRPC struct {
	proto.UnimplementedTagServiceServer
	MU  sync.Mutex
	svc *service.Service
}

func NewTagGRPC(svc *service.Service) *TagGRPC {
	return &TagGRPC{
		svc: svc,
	}
}

func (s *TagGRPC) CreateTag(ctx context.Context, in *proto.UserTagRequest) (*proto.TagResponse, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	tag := &model.UserTag{
		ID:        in.GetID(),
		UserID:    in.GetUserID(),
		Name:      in.GetName(),
		IsActive:  in.GetIsActive(),
		CreatedAt: in.GetCreatedAt().AsTime(),
	}

	err := s.svc.Tag.Create(ctx, tag)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating tag: %v", err))
	}

	msg := &proto.TagResponse{Status: "1"}

	fmt.Printf("Создано сообщение: %s\n", msg)

	return msg, nil
}

func (s *TagGRPC) UpdateTag(ctx context.Context, in *proto.UserTagRequest) (*proto.TagResponse, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	// ерунда
	err := s.svc.Tag.Update(ctx, in.GetName(), in.UserID, in.GetName())
	if err != nil {
		return nil, err
	}

	msg := &proto.TagResponse{Status: "1"}

	fmt.Printf("Создано сообщение: %s\n", msg)

	return msg, nil

}
func (s *TagGRPC) DeleteTag(ctx context.Context, in *proto.DeleteTagRequest) (*proto.TagResponse, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	err := s.svc.Tag.Delete(ctx, in.GetName(), in.UserID)
	if err != nil {
		return nil, err
	}

	msg := &proto.TagResponse{Status: "1"}

	return msg, nil
}
func (s *TagGRPC) ListTag(ctx context.Context, in *proto.ListTagRequest) (*proto.ListUserTagRequest, error) {
	s.MU.Lock()
	defer s.MU.Unlock()

	if tokenValid(ctx) != 1 {
		return nil, errors.New("invalid token")
	}

	tags, err := s.svc.Tag.List(ctx, in.UserID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error listing tags: %v", err))
	}

	var arrUTR []*proto.UserTagRequest
	for _, tag := range tags {
		utr := &proto.UserTagRequest{
			ID:        tag.ID,
			UserID:    tag.UserID,
			Name:      tag.Name,
			IsActive:  tag.IsActive,
			CreatedAt: timestamppb.New(tag.CreatedAt),
		}
		arrUTR = append(arrUTR, utr)
	}

	lUTR := &proto.ListUserTagRequest{
		UTR: arrUTR,
	}
	fmt.Printf("Отправлен список Tag: %s\n", lUTR)
	return lUTR, nil
}
