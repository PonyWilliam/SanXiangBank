package main

import (
	"context"

	"github.com/PonyWilliam/tarsSeckill/domain/model"
	"github.com/PonyWilliam/tarsSeckill/tars-protocol/SanXiangBank"
)

// PublishImp servant implementation
type PublishImp struct {
}

// Init servant init
func (imp *PublishImp) Init() error {
	return nil
}

// Destroy servant destory
func (imp *PublishImp) Destroy() {
}

func (imp *PublishImp) Publish(ctx context.Context,pdt *SanXiangBank.Product) (bool,error) {
	err := rp.Publish(ImpToModel(*pdt))
	return err == nil,err
}
func (imp *PublishImp) Update(ctx context.Context,pdt *SanXiangBank.Product) (bool,error) {
	err := rp.Update(ImpToModel(*pdt))
	return err == nil,err
}
func (imp *PublishImp) DelByID(ctx context.Context,id int32) (bool,error) {
	err := rp.DelProductByID(id)
	return err == nil,err
}
func (imp *PublishImp) FindAllProduct(ctx context.Context,pdts *[]SanXiangBank.Product) (bool,error) {
	res := rp.FindAllProduct()
	var u []SanXiangBank.Product
	for _,v := range res{
		r := ModelToImp(v)
		u = append(u, r)
	}
	pdts = &u
	return true,nil
}
func (imp *PublishImp) FindProductByID(ctx context.Context,id int32,pdt *SanXiangBank.Product) (bool,error) {
	res := rp.FindProductByID(id)
	res2 := ModelToImp(res)
	pdt = &res2
	return true,nil
}

func ImpToModel(Iproduct SanXiangBank.Product)model.Product{
	mproduct := model.Product{}
	mproduct.ConfigID = int64(Iproduct.Cid)
	mproduct.Name = Iproduct.Name
	mproduct.Nums = int64(Iproduct.Nums)
	mproduct.Price = int64(Iproduct.Nums)
	mproduct.Stock = int64(Iproduct.Stock)
	mproduct.URL = Iproduct.Url
	mproduct.StartTime = int64(Iproduct.Start)
	mproduct.EndTime = int64(Iproduct.End)
	return mproduct
}
func ModelToImp(mproduct model.Product)SanXiangBank.Product{
	iproduct := SanXiangBank.Product{}
	iproduct.Id = int32(mproduct.ID)
	iproduct.Cfg.Id = int32(mproduct.Config.ID)
	iproduct.Cfg.Name = mproduct.Config.Name
	iproduct.Cfg.Rage = int32(mproduct.Config.RequireAge)
	iproduct.Cfg.Rlevel = int32(mproduct.Config.RequireLevel)
	iproduct.Cfg.Rpromise = mproduct.Config.RequirePromise
	iproduct.Cfg.Rwork = int32(mproduct.Config.RequireWork)
	iproduct.Cid = int32(mproduct.ConfigID)
	iproduct.Name = mproduct.Name
	iproduct.Nums = int32(mproduct.Nums)
	iproduct.Price = int32(mproduct.Price)
	iproduct.Stock = int32(mproduct.Stock)
	iproduct.Url = mproduct.URL
	iproduct.Start = int32(mproduct.StartTime)
	iproduct.End = iproduct.End
	return iproduct
}
// bool Publish(Product pdt);
// bool Update(Product pdt);
// bool DelByID(int id);
// bool FindAllProduct(out vector<Product> pdts);
// bool FindProductByID(int id,out Product pdt);