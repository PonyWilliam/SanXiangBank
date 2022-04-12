package repository

import (
	"github.com/PonyWilliam/tarsSeckill/domain/model"
	"gorm.io/gorm"
)

type IOrder interface{
	Init()error
	Create(uid int64,pid int64,status int64,Reason int64)(error)
	DelByID(id int64)error
	FindByID(id int64)model.Order
	FindByUserID(uid int64)[]model.Order
	FindByProductID(pid int64)[]model.Order
	FindAll()[]model.Order
	Update(model.Order)error
}
type OrderRepo struct{
	Db *gorm.DB
	PdtRepo IProduct
}

func InitOrderRepo(db *gorm.DB,PdtRepo IProduct)IOrder{
	return &OrderRepo{
		Db:db,
		PdtRepo: PdtRepo,
	}
}

func(r OrderRepo)Init()error{
	if(r.Db.Migrator().HasTable(&model.Order{})){
		return nil
	}
	r.Db.Migrator().CreateTable(&model.Order{})
	return nil
}
func(r OrderRepo)Create(uid int64,pid int64,status int64,Reason int64)(error){
	order := &model.Order{}
	order.UserID = uid
	order.ProductID = pid
	order.Status = status
	order.Reason = Reason //全部在网关层面进行判断，调用product->cfg,user->user两个微服务
	res := r.Db.Create(order)
	return res.Error
}
func(r OrderRepo)DelByID(id int64)error{
	return r.Db.Delete(&model.Order{},id).Error
}
func(r OrderRepo)FindByID(id int64)model.Order{
	res := &model.Order{}
	r.Db.Where("id = ?",id).First(res)
	return *res
}
func(r OrderRepo)FindByUserID(uid int64)[]model.Order{
	var res []model.Order
	r.Db.Where("uid = ?",uid).Find(res)
	return res
}
func(r OrderRepo)FindByProductID(pid int64)[]model.Order{
	var res []model.Order
	r.Db.Where("pid = ?",pid).Find(res)
	return res
}
func(r OrderRepo)FindAll()[]model.Order{
	var res []model.Order
	r.Db.Model(&model.Order{}).Find(res)
	return res
}
func(r OrderRepo)Update(order model.Order)error{
	res := r.Db.Model(&model.Order{}).Where("id = ?",order.ID).Updates(order)
	return res.Error
}