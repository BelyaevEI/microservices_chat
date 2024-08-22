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
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		chatName = gofakeit.Animal()
		ids      = []int64{1, 2, 3, 4, 5}
		names    = []string{gofakeit.Animal(), gofakeit.Animal()}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Usernames: names,
			Id:        ids,
			Chatname:  chatName,
		}

		res = &desc.CreateResponse{
			Id: id,
		}

		chatCreate = model.ChatCreate{
			Name:   chatName,
			UserID: ids,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateChatMock.Expect(ctx, &chatCreate).Return(id, nil)
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
				mock.CreateChatMock.Expect(ctx, &chatCreate).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			createServiceMock := test.chatServiceMock(mc)
			api := chat.NewImplementation(createServiceMock)

			newID, err := api.CreateChat(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, newID)
		})
	}

}
