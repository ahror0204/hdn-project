package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/hdn-project/review-service/storage"
	l "github.com/hdn-project/review-service/pkg/logger"
	cl "github.com/hdn-project/review-service/service/grpc_client"
)

// ReviewService ...
type ReviewService struct {
	storage storage.IStorage
	logger  l.Logger
	client cl.GrpcClientI
}

/*
// PostService ...
type PostService struct {
	storage storage.IStorage
	client  grpc_client.IServiceManager
	logger  l.Logger
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpc_client.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}
*/
// NewReviewService ...
func NewReviewService(db *sqlx.DB, log l.Logger,client cl.GrpcClientI) *ReviewService {
	return &ReviewService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client: client,
	}
}
