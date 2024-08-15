package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/BelyaevEI/microservices_chat/internal/api/chat"
	"github.com/BelyaevEI/microservices_chat/internal/service"
	"github.com/BelyaevEI/microservices_chat/internal/service/mocks"
	desc "github.com/BelyaevEI/microservices_chat/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id  = gofakeit.Int64()
		req = &desc.DeleteRequest{
			Id: id,
		}

		serviceErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, id).Return(nil)
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
				mock.DeleteChatMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			createServiceMock := test.chatServiceMock(mc)
			api := chat.NewImplementation(createServiceMock)

			newID, err := api.DeleteChat(test.args.ctx, test.args.req)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, newID)
		})
	}

}
