package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozeroDemo/service/product/model"

	"gozeroDemo/service/product/rpc/internal/svc"
	"gozeroDemo/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	err = l.svcCtx.ProductModel.Update(&model.Product{
		Id:     in.Id,
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.UpdateResponse{}, nil
}
