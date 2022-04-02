package model

import (
	"github.com/PonyWilliam/tarsConfig/domain/model"
	"gorm.io/plugin/soft_delete"
)
type Product struct{
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name string `json:"name"` //产品名称
	Price int64 `json:"price"` //产品单价
	Nums int64 `json:"nums"` //产品数量
	Stock int64 `json:"stock"` //剩余数量（要上锁）
	StartTime int64 `json:"start_time"` //秒杀开始时间
	EndTime int64 `json:"end_time"` //秒杀结束时间
	ConfigID int64 `json:"configid"` //秒杀默认配置模板
	Config model.RequireConfig
	URL string `json:"url"` //系统随机生成的秒杀链接，未开始前对用户不可见
	Created  int64 `gorm:"autoCreateTime"`//创建时间
	Updated  int64 `gorm:"autoUpdateTime"`//修改时间
	DeletedAt soft_delete.DeletedAt//删除时间，数据库存档一段时间，可以做个定时脚本定时删除
}
