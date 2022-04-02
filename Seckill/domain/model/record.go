package model

import "gorm.io/plugin/soft_delete"
type Product_Record struct{
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	PID int64 `json:"pid"` //对应产品ID
	UID int64 `json:"uid"` //对应用户ID
	Status int64 `json:"status"` //是否成功，1->成功，2->失败
	FailReason int64 `json:"fail_reason"` 
	// -1->成功，1->用户不符合管理员配置规则,2->库存不足,其它->其它原因

	

	Created  int64 `gorm:"autoCreateTime"`//创建时间
	Updated  int64 `gorm:"autoUpdateTime"`//修改时间
	DeletedAt soft_delete.DeletedAt//删除时间，数据库存档一段时间，可以做个定时脚本定时删除
}
