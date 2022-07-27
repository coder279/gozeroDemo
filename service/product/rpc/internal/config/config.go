package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)
/**
 * Author:Coder279
 * Desc: RPC配置文件
 * Date: 2022-07-22 15:00
 */
type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
