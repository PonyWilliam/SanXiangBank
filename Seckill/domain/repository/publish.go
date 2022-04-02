package repository

import (
	"github.com/PonyWilliam/tarsSeckill/domain/model"
	"gorm.io/gorm"
)
var db *gorm.DB

func init(){}

type IProduct interface{
	Init() error
	Publish(model.Product) (error)
	Update(product model.Product) (error)
	DelProductByID(id int32)(error)
	FindAllProduct()(res []model.Product)
	FindProductByID(id int32)(res model.Product)
}

type ProductRepo struct{
	Db *gorm.DB
}

func InitProductRepo(db *gorm.DB)IProduct{
	return &ProductRepo{Db: db}
}

func(r *ProductRepo)Init() (error){
	if(!r.Db.Migrator().HasTable(&model.Product{})){
		r.Db.Migrator().CreateTable(&model.Product{})
	}
	return nil
}

func(r *ProductRepo)Publish(product model.Product) (error){
	res := r.Db.Create(product)
	return res.Error
}
func(r *ProductRepo)Update(product model.Product) (error){
	res := r.Db.Model(&model.Product{}).Where("id = ?",product.ID).Updates(product)
	return res.Error
}
func(r *ProductRepo)DelProductByID(id int32)(error){
	return r.Db.Delete(&model.Product{},id).Error
}
func(r *ProductRepo)FindAllProduct()(res []model.Product){
	r.Db.Find(&res)
	return res
}
func(r *ProductRepo)FindProductByID(id int32)(res model.Product){
	r.Db.Where("id = ?",id).First(&res)
	return res
}
