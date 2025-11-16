package repository

import (
	"arukabe/gen/sqlc"
	"context"
)

func (cr *ChatRepository) GetAllChats(ctx context.Context) ([]sqlc.GetAllChatsRow, error) {
	dbData, err := cr.db.GetAllChats(ctx)
	if err != nil {
		return nil, err
	}
	return dbData, nil
}
