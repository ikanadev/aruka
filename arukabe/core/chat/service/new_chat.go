package service

import (
	"arukabe/core/common/mappers"
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	modelsv1 "arukabe/gen/connect/aruka/models/v1"
	"arukabe/gen/sqlc"
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (cs *ChatService) NewChat(
	ctx context.Context,
	req *chatv1.NewChatRequest,
) (*chatv1.NewChatResponse, error) {
	modelUuid, err := uuid.Parse(req.ModelId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	dbRow, err := cs.db.CreateChat(ctx, sqlc.CreateChatParams{
		ID:       id,
		Title:    "",
		Prompt:   strings.TrimSpace(req.Prompt),
		ModelID:  modelUuid,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp := &chatv1.NewChatResponse{
		Chat: &modelsv1.Chat{
			Id:       dbRow.ID.String(),
			Title:    dbRow.Title,
			Prompt:   dbRow.Prompt,
			Pinned:   dbRow.Pinned,
			Model:    mappers.ModelToPB(dbRow.Model),
			Provider: &modelsv1.ProviderBase{
				Id:   dbRow.Provider.ID.String(),
				Name: dbRow.Provider.Name,
			},
			TimeData: &modelsv1.TimeData{
				CreatedAt:  timestamppb.New(dbRow.CreatedAt),
				UpdatedAt:  timestamppb.New(dbRow.UpdatedAt),
				ArchivedAt: mappers.TimestampFromTime(dbRow.ArchivedAt),
				DeletedAt:  mappers.TimestampFromTime(dbRow.DeletedAt),
			},
		},
	}
	return resp, nil
}
