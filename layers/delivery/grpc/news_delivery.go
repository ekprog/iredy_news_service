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

func (d *NewsDeliveryService) GetNewsDetails(ctx context.Context, r *pb.GetNewsDetailsRequest) (*pb.GetNewsDetailsResponse, error) {
	uCaseRes, err := d.newsUcase.GetNewsDetails(ctx, r.Page, r.NewsId)

	if err != nil {
		return &pb.GetNewsDetailsResponse{
			Status: &pb.Status{
				Code:    domain.ServerError,
				Message: err.Error(),
			},
			Data: nil,
		}, nil
	}

	if r == nil {
		return &pb.GetNewsDetailsResponse{
			Status: &pb.Status{
				Code:    domain.ValidationError,
				Message: domain.FieldRequired,
			},
			Data: nil,
		}, nil
	}

	response := &pb.GetNewsDetailsResponse{
		Status: &pb.Status{
			Code:    uCaseRes.Status.Code,
			Message: uCaseRes.Status.Message,
		},
		Data: nil,
	}

	for i := range uCaseRes.NewsDetails {

		r := &pb.NewsDetails{
			Id:    uCaseRes.NewsDetails[i].Id,
			Title: uCaseRes.NewsDetails[i].Title,
			Image: uCaseRes.NewsDetails[i].Image,
			Type:  uCaseRes.NewsDetails[i].Type,
		}
		response.Data = append(response.Data, r)

	}

	return response, nil
}

func (d *NewsDeliveryService) AddNewsCard(ctx context.Context, r *pb.CreateNewsCardRequest) (*pb.CreateNewsCardResponse, error) {
	card := domain.NewsCard{
		Title: r.Title,
		Image: r.Image,
		Type:  r.Type,
	}

	res, err := d.newsUcase.AddNewsCard(ctx, card)
	if err != nil {
		return &pb.CreateNewsCardResponse{}, errors.Wrap(err, "Error at AddNewsCard UseCase Call")
	}
	return &pb.CreateNewsCardResponse{
		Status: &pb.Status{
			Code:    res.Status.Code,
			Message: res.Status.Message,
		},
		Id: res.Id,
	}, nil
}

func (d *NewsDeliveryService) AddNewsDetails(ctx context.Context, r *pb.CreateNewsDetailsRequest) (*pb.CreateNewsDetailsResponse, error) {

	news_details := []*domain.NewsDetails{}

	for i := range r.Data {

		detail := domain.NewsDetails{
			Title:      r.Data[i].Title,
			Image:      r.Data[i].Image,
			Type:       r.Data[i].Type,
			NewsID:     r.NewsId,
			SwipeDelay: r.Data[i].SwipeDelay,
		}

		news_details = append(news_details, &detail)

	}

	var uCaseRes domain.CreateNewsDetailesResponse

	uCaseRes, err := d.newsUcase.AddNewsDetails(ctx, news_details, r.NewsId)
	if err != nil {
		return &pb.CreateNewsDetailsResponse{
			Status: &pb.Status{
				Code:    uCaseRes.Status.Code,
				Message: uCaseRes.Status.Message,
			},
		}, errors.Wrap(err, "Error at AddNewsDetails UseCase Call")
	}

	return &pb.CreateNewsDetailsResponse{
		Status: &pb.Status{
			Code:    uCaseRes.Status.Code,
			Message: uCaseRes.Status.Message,
		},
	}, nil

}

func (d *NewsDeliveryService) DeleteNewsCard(ctx context.Context, r *pb.DeleteNewsCardRequest) (*pb.Status, error) {
	uCaseRes, err := d.newsUcase.DeleteNewsCard(ctx, r.Id)
	if err != nil {
		return &pb.Status{
			Code:    uCaseRes.Code,
			Message: uCaseRes.Message,
		}, err
	}

	return &pb.Status{
		Code:    uCaseRes.Code,
		Message: uCaseRes.Message,
	}, nil
}

func (d *NewsDeliveryService) DeleteNewsDetails(ctx context.Context, r *pb.DeleteNewsDetailsRequest) (*pb.Status, error) {
	uCaseRes, err := d.newsUcase.DeleteNewsDetails(ctx, r.Id)
	if err != nil {
		return &pb.Status{
			Code:    uCaseRes.Code,
			Message: uCaseRes.Message,
		}, err
	}

	return &pb.Status{
		Code:    uCaseRes.Code,
		Message: uCaseRes.Message,
	}, nil
}
