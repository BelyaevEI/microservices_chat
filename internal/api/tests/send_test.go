package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/BelyaevEI/microservices_chat/internal/api/chat"
	"github.com/BelyaevEI/microservices_chat/internal/model"
	"github.com/BelyaevEI/microservices_chat/internal/service"
	"github.com/BelyaevEI/microservices_chat/internal/service/mocks"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id         = gofakeit.UUID()
		fromUserID = gofakeit.Int64()
		toChatID   = gofakeit.Int64()
		text       = gofakeit.BeerName()

		serviceErr = fmt.Errorf("service error")

		req = &desc.SendMessageRequest{
			FromUserId: fromUserID,
			ToChatId:   toChatID,
			Text:       text,
		}

		createMessage = model.MessageCreate{
			Info: model.MessageInfo{
				ChatID: toChatID,
				UserID: fromUserID,
				Text:   text,
			},
		}

		res = &desc.SendMessageResponse{
			Id:     id,
			ChatId: toChatID,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.SendMessageResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, &createMessage).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, &createMessage).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			sendMessageServiceMock := test.chatServiceMock(mc)
			api := chat.NewImplementation(sendMessageServiceMock)

			newID, err := api.SendMessage(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, newID)
		})
	}

}
