package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozeroDemo/service/product/model"
	"gozeroDemo/service/product/rpc/internal/config"
)

/**
 * Author:Coder279
 * Desc: 将Model、配置文件写入上下文中
 * Date: 2022-07-22
 */
type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
	}
}
