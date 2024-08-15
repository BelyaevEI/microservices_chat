package chat

import (
	"context"

	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteChat Delete deletes a chat
func (i *Implementation) DeleteChat(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	err := i.chatService.DeleteChat(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
