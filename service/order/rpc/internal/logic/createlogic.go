package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozeroDemo/service/order/model"
	"gozeroDemo/service/product/rpc/product"
	"gozeroDemo/service/user/rpc/user"

	"gozeroDemo/service/order/rpc/internal/svc"
	"gozeroDemo/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}
	// 查询产品是否存在
	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})
	if err != nil {
		return nil, err
	}
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "产品库存不足")
	}
	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}
	// 创建订单
	res, err := l.svcCtx.OrderModel.Insert(&newOrder)
	if err != nil {
		return nil, err
	}
	// 获取最新更新的ID
	newOrder.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 更新产品详情
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     productRes.Id,
		Name:   productRes.Name,
		Desc:   productRes.Desc,
		Stock:  productRes.Stock - 1,
		Amount: productRes.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{
		Id: newOrder.Id,
	}, nil

}
