package chat

import (
	"context"

	"github.com/BelyaevEI/microservices_chat/internal/converter"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
)

// SendMessage sends a new message
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {

	id, err := i.chatService.SendMessage(ctx, converter.ToMessageCreateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.SendMessageResponse{
		Id:     id,
		ChatId: req.GetToChatId(),
	}, nil
}
