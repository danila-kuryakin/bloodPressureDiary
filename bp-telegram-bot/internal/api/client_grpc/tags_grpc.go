package client_grpc

import (
	"bp-telegram-bot/internal/model"
	proto2 "bp-telegram-bot/internal/pb/grpc-tag/proto"
	"context"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagGRPC struct {
	TagClient proto2.TagServiceClient
}

func NewTagGRPC(conn *grpc.ClientConn) *TagGRPC {
	tagClient := proto2.NewTagServiceClient(conn)

	return &TagGRPC{
		TagClient: tagClient,
	}
}

func (c *TagGRPC) CreateTag(name model.UserTag, userId string) error {
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	msg, err := c.TagClient.CreateTag(ctx, &proto2.UserTagRequest{
		ID:        name.ID,
		UserID:    userId,
		Name:      name.Name,
		IsActive:  name.IsActive,
		CreatedAt: timestamppb.New(name.CreatedAt),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("CreateTag err: %v", err))
	}
	if msg.GetStatus() == "1" {
		return errors.New(fmt.Sprintf("CreateTag err: %v", msg))
	}
	return nil
}

func (c *TagGRPC) ListTags(userId string) ([]*model.Tag, error) {
	fmt.Println("ListTags")
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)

	tags, err := c.TagClient.ListTag(ctx, &proto2.ListTagRequest{
		UserID: userId,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ListTags err: %v", err))
	}
	if len(tags.GetUTR()) == 0 {
		return nil, errors.New(fmt.Sprintf("ListTags err: %v", tags.GetUTR()))
	}

	var lTags []*model.Tag
	for _, tag := range tags.GetUTR() {
		urt := &model.Tag{
			Name: tag.Name,
		}
		lTags = append(lTags, urt)
	}
	return lTags, err
}

func (c *TagGRPC) UpdateTag(id int64, name string, userId string) error {
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	msg, err := c.TagClient.UpdateTag(ctx, &proto2.UserTagRequest{
		ID:     id,
		UserID: userId,
		Name:   name,
	})
	if err != nil {
		return err
	}
	if msg.GetStatus() == "1" {
		return errors.New(fmt.Sprintf("UpdateTag err: %v", msg))
	}
	return nil
}

func (c *TagGRPC) DeleteTag(name string, userId string) error {
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"token_api", os.Getenv("SERVICE_TOKEN_API"),
	)
	msg, err := c.TagClient.DeleteTag(ctx, &proto2.DeleteTagRequest{
		UserID: userId,
		Name:   name,
	})
	if err != nil {
		return err
	}
	if msg.GetStatus() == "1" {
		return errors.New(fmt.Sprintf("DeleteTag err: %v", msg))
	}
	return nil
}
