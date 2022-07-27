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
 * Author: Coder279
 * Desc: 商品详情接口
 * Date: 2022-07-27 10:57
 */
type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	res,err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil{
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	return &product.DetailResponse{
		Id:     res.Id,
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  res.Stock,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
