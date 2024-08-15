package app

import (
	"context"
	"log"

	"github.com/BelyaevEI/microservices_chat/internal/api/chat"
	"github.com/BelyaevEI/microservices_chat/internal/config"
	"github.com/BelyaevEI/microservices_chat/internal/repository"
	chatRepository "github.com/BelyaevEI/microservices_chat/internal/repository/chat"
	"github.com/BelyaevEI/microservices_chat/internal/service"
	chatService "github.com/BelyaevEI/microservices_chat/internal/service/chat"
	"github.com/BelyaevEI/platform_common/pkg/closer"
	"github.com/BelyaevEI/platform_common/pkg/db"
	"github.com/BelyaevEI/platform_common/pkg/db/pg"
	"github.com/BelyaevEI/platform_common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	pgClient  db.Client
	txManager db.TxManager

	chatImpl       *chat.Implementation
	chatRepository repository.ChatRepository
	chatService    service.ChatService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {

	if s.grpcConfig == nil {

		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) PostgresClient(ctx context.Context) db.Client {
	if s.pgClient == nil {
		client, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(client.Close)

		s.pgClient = client
	}

	return s.pgClient
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.PostgresClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.PostgresClient(ctx).DB())
	}

	return s.txManager
}
