package service

import (
	"arukabe/core/chat/repository"
	"arukabe/core/common/mappers"
	chatv1 "arukabe/gen/connect/chat/v1"
	modelsv1 "arukabe/gen/connect/models/v1"
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (cs *ChatService) NewChat(
	ctx context.Context,
	req *chatv1.NewChatRequest,
) (*chatv1.NewChatResponse, error) {
	uuid, err := uuid.Parse(req.ModelId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	dbChat, dbModel, dbProvider, err := cs.repo.NewChat(ctx, repository.NewChatData{
		Prompt:  req.Prompt,
		ModelID: uuid,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp := &chatv1.NewChatResponse{
		Chat: &modelsv1.Chat{
			Id:       dbChat.ID.String(),
			Title:    dbChat.Title,
			Prompt:   dbChat.Prompt,
			Pinned:   dbChat.Pinned,
			Model:    mappers.ModelToPB(dbModel),
			Provider: &modelsv1.ProviderBase{
				Id:   dbProvider.ID.String(),
				Name: dbProvider.Name,
			},
			TimeData: &modelsv1.TimeData{
				CreatedAt:  timestamppb.New(dbChat.CreatedAt),
				UpdatedAt:  timestamppb.New(dbChat.UpdatedAt),
				ArchivedAt: mappers.TimestampFromTime(dbChat.ArchivedAt),
				DeletedAt:  mappers.TimestampFromTime(dbChat.DeletedAt),
			},
		},
	}
	return resp, nil
}
