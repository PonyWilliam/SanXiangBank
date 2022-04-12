package model

import (
	"github.com/PonyWilliam/tarsUser/domain/model"
	"gorm.io/plugin/soft_delete"
)
type Order struct{
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Product Product
	ProductID int64 `json:"pid"` //对应产品ID
	User model.User
	UserID int64 `json:"uid"`
	Status int64 `json:"status"` //是否成功，1->成功，2->失败
	Reason int64 `json:"reason"` 
	// 1->成功，2->用户不符合管理员配置规则,3->库存不足,其它->其它原因


	Created  int64 `gorm:"autoCreateTime"`//创建时间
	Updated  int64 `gorm:"autoUpdateTime"`//修改时间
	DeletedAt soft_delete.DeletedAt//删除时间，数据库存档一段时间，可以做个定时脚本定时删除
}
