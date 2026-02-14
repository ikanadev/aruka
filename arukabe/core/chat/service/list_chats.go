package service

import (
	"arukabe/core/common/mappers"
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	modelsv1 "arukabe/gen/connect/aruka/models/v1"
	"arukabe/gen/sqlc"
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (cs *ChatService) ListChats(ctx context.Context, req *chatv1.ListChatsRequest) (*chatv1.ListChatsResponse, error) {
	// Set defaults
	page := uint32(1)
	limit := uint32(20)

	// Override with request values if provided
	if req.Pagination != nil {
		if req.Pagination.GetPage() > 0 {
			page = req.Pagination.GetPage()
		}
		if req.Pagination.GetLimit() > 0 {
			limit = req.Pagination.GetLimit()
		}
	}

	// Calculate offset
	offset := int32((page - 1) * limit)

	// Get total count
	totalCount, err := cs.db.CountChats(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Get paginated chats
	dbRows, err := cs.db.GetAllChats(ctx, sqlc.GetAllChatsParams{
		Limit: int32(limit),
		Offset: offset,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Convert to protobuf
	chats := make([]*modelsv1.Chat, 0, len(dbRows))
	for _, row := range dbRows {
		chats = append(chats, chatToPB(row.Chat, row.Model, row.Provider))
	}

	// Calculate pagination metadata
	totalPages := uint32(totalCount) / limit
	if uint32(totalCount)%limit > 0 {
		totalPages++
	}

	paginationData := &modelsv1.PaginationData{
		Page:       page,
		Limit:      limit,
		TotalItems: uint32(totalCount),
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}

	return &chatv1.ListChatsResponse{
		Chats:      chats,
		Pagination: paginationData,
	}, nil
}

func chatToPB(chat sqlc.Chat, model sqlc.Model, provider sqlc.Provider) *modelsv1.Chat {
	return &modelsv1.Chat{
		Id:     chat.ID.String(),
		Title:  chat.Title,
		Prompt: chat.Prompt,
		Pinned: chat.Pinned,
		Model:  mappers.ModelToPB(model),
		Provider: &modelsv1.ProviderBase{
			Id:   provider.ID.String(),
			Name: provider.Name,
		},
		TimeData: &modelsv1.TimeData{
			CreatedAt:  timestamppb.New(chat.CreatedAt),
			UpdatedAt:  timestamppb.New(chat.UpdatedAt),
			ArchivedAt: mappers.TimestampFromTime(chat.ArchivedAt),
			DeletedAt:  mappers.TimestampFromTime(chat.DeletedAt),
		},
	}
}
