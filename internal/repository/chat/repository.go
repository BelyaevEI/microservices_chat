package chat

import (
	"github.com/BelyaevEI/microservices_chat/internal/repository"
	"github.com/BelyaevEI/platform_common/pkg/db"
)

const (
	tableName  = "chat"
	idColumn   = "id"
	nameColumn = "name"
	userIDs    = "user_ids"

	tableNameMessage = "message"
	chatIDColumn     = "chat_id"
	userIDColumn     = "user_id"
	textColumn       = "text"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}
