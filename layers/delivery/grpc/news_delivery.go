package grpc

import (
	"context"
	"microservice/app"
	"microservice/app/core"
	"microservice/layers/domain"
	pb "microservice/pkg/pb/api"
)

type NewsDeliveryService struct {
	pb.NewsServiceServer
	log       core.Logger
	newsUcase domain.NewsUseCase
}

func NewNewsService(log core.Logger, newsUCase domain.NewsUseCase) *NewsDeliveryService {
	return &NewsDeliveryService{
		log:       log,
		newsUcase: newsUCase,
	}
}

func (d *NewsDeliveryService) Init() error {
	app.InitGRPCService(pb.RegisterNewsServiceServer, pb.NewsServiceServer(d))
	return nil
}

func (d *NewsDeliveryService) GetNews(ctx context.Context, r *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	uCaseRes, err := d.newsUcase.GetNews(1)
	if err != nil {
		print(err.Error())
	}

	return &pb.GetNewsResponse{
		Data: []*pb.NewsCard{
			&pb.NewsCard{
				Id:    int32(uCaseRes[0].Id),
				Title: uCaseRes[0].Title,
			},
		},
	}, nil
}
