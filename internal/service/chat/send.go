package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	var id string

	id, err := s.chatRepository.SendMessage(ctx, createMessage)
	if err != nil {
		return "", err
	}

	return id, nil
}
