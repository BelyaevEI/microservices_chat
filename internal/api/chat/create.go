package chat

import (
	"context"
	"log"

	"github.com/BelyaevEI/microservices_chat/internal/converter"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
)

// CreateChat creates a new chat
func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := i.chatService.CreateChat(ctx, converter.ToChatCreateFromDesc(req))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
