package grpc

import (
	"context"
	"microservice/app"
	"microservice/app/core"
	"microservice/layers/domain"
	pb "microservice/pkg/pb/api"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	uCaseRes, err := d.newsUcase.GetNews(ctx, r.Page)

	if err != nil {
		return &pb.GetNewsResponse{
			Status: &pb.Status{
				Code:    domain.ServerError,
				Message: err.Error(),
			},
			Data: nil,
		}, nil
	}

	if r == nil {
		return &pb.GetNewsResponse{
			Status: &pb.Status{
				Code:    domain.ValidationError,
				Message: domain.FieldRequired,
			},
			Data: nil,
		}, nil
	}

	response := &pb.GetNewsResponse{
		Status: &pb.Status{
			Code:    uCaseRes.Status.Code,
			Message: uCaseRes.Status.Message,
		},
		Data: nil,
	}

	for i := range uCaseRes.News {

		r := &pb.NewsCard{
			Id:        uCaseRes.News[i].Id,
			Title:     uCaseRes.News[i].Title,
			Image:     uCaseRes.News[i].Image,
			Type:      uCaseRes.News[i].Type,
			CreatedAt: timestamppb.New(uCaseRes.News[i].CreatedAt),
		}
		response.Data = append(response.Data, r)

	}
	// For Debug
	// println(response.Data[0].Id)
	// println(response.Data[1].Id)
	// println(response.Data[2].Id)
	// println(response.Data[3].Id)
	// println(response.Data[4].Id)
	// println(response.Status.Code)
	// println(response.Status.Message)

	return response, nil
}

func (d *NewsDeliveryService) AddNewsCard(ctx context.Context, r *pb.CreateNewsCardRequest) (*pb.CreateNewsCardResponse, error) {
	_, err := d.newsUcase.AddNewsCard(ctx, domain.NewsCard{})
	if err != nil {
		return &pb.CreateNewsCardResponse{}, errors.Wrap(err, "")
	}
	return &pb.CreateNewsCardResponse{}, nil
}
