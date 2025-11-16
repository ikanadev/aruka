package repository

import (
	"arukabe/gen/sqlc"
	"context"
)

func (cr *ChatRepository) GetAllChats(ctx context.Context, limit, offset int32) ([]sqlc.GetAllChatsRow, int64, error) {
	// Get total count
	totalCount, err := cr.db.CountChats(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated chats
	dbData, err := cr.db.GetAllChats(ctx, sqlc.GetAllChatsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, 0, err
	}

	return dbData, totalCount, nil
}
