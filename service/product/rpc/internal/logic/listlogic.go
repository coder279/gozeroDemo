package logic

import (
	"context"
	"gozeroDemo/service/product/rpc/internal/svc"
	"gozeroDemo/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *product.ListRequest) (*product.ListResponse, error) {
	var (
		list   []*product.DetailResponse
		signle *product.DetailResponse
	)
	products, err := l.svcCtx.ProductModel.FindAll(in.GetPage(), in.GetSize(), in.GetName())
	if err != nil {
		return nil, err
	}
	for _, v := range products {
		signle = &product.DetailResponse{
			Id:     v.Id,
			Name:   v.Name,
			Desc:   v.Desc,
			Stock:  v.Stock,
			Status: v.Status,
			Amount: v.Amount,
		}
		list = append(list, signle)
	}

	return &product.ListResponse{
		List: list,
	}, nil
}
