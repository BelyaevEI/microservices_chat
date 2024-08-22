package chat

import (
	"github.com/BelyaevEI/microservices_chat/internal/repository"
	"github.com/BelyaevEI/microservices_chat/internal/service"
	"github.com/BelyaevEI/platform_common/pkg/db"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

// NewService creates a new chat service.
func NewService(chatRepository repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
