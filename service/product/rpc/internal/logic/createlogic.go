package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozeroDemo/service/product/model"

	"gozeroDemo/service/product/rpc/internal/svc"
	"gozeroDemo/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

/**
 * Author: coder279
 * Desc: 商品创建逻辑
 * Date: 2022-07-27
 */

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	newproduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}
	res, err := l.svcCtx.ProductModel.Insert(&newproduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newproduct.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &product.CreateResponse{
		Id: newproduct.Id,
	}, nil
}
