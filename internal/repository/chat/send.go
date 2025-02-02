package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
	"github.com/BelyaevEI/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	builderInsert := sq.Insert(tableNameMessage).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn, textColumn).
		Values(createMessage.Info.ChatID, createMessage.Info.UserID, createMessage.Info.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "message_repository.SendMessage",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
