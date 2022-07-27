package logic

import (
	"context"
	"gozeroDemo/service/product/rpc/product"

	"gozeroDemo/service/product/api/internal/svc"
	"gozeroDemo/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListRequest) (resp *types.ListResponse, err error) {
	var productList = make([]types.DetailResponse, 0)
	res, err := l.svcCtx.ProductRpc.List(l.ctx, &product.ListRequest{
		Name: req.Name,
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	for _, value := range res.GetList() {
		productList = append(productList, types.DetailResponse{
			Id:     value.Id,
			Name:   value.Name,
			Desc:   value.Desc,
			Stock:  value.Stock,
			Amount: value.Amount,
			Status: value.Status,
		})
	}

	return &types.ListResponse{List: productList}, nil
}
