package dal

import (
	"github.com/li-jin-gou/kitex-server-demo/biz/dal/mysql"
	"github.com/li-jin-gou/kitex-server-demo/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
