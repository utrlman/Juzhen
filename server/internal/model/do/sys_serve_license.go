// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLicense is the golang structure of table hg_sys_serve_license for DAO operations like Where/Data.
type SysServeLicense struct {
	g.Meta       `orm:"table:hg_sys_serve_license, do:true"`
	Id           interface{} // 许可ID
	Group        interface{} // 分组
	Name         interface{} // 许可名称
	Appid        interface{} // 应用ID
	SecretKey    interface{} // 应用秘钥
	Desc         interface{} // 授权说明
	RemoteAddr   interface{} // 最后连接地址
	Online       interface{} // 在线数量
	OnlineLimit  interface{} // 在线数量限制，默认1
	LoginTimes   interface{} // 登录次数
	LastLoginAt  *gtime.Time // 最后登录时间
	LastActiveAt *gtime.Time // 最后活跃时间
	Routes       *gjson.Json // 路由表，空使用默认分组路由
	AllowedIps   interface{} // 白名单，*代表所有，只有允许的IP才能连接到tcp服务
	EndAt        *gtime.Time // 授权结束时间
	Remark       interface{} // 备注
	Status       interface{} // 状态
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 修改时间
}