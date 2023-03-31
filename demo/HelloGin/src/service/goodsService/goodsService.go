package goodsService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/pojo"
)

// var goods=pojo.Goods{}
var goodsServiceImpl = pojo.GoodsServiceImpl()

func GoodsAdd(add reqDto.GoodsAdd) error {
	goods := pojo.Goods{
		Name:        add.Name,
		Price:       add.Price,
		Address:     add.Address,
		Fms:         add.Fms,
		Brand:       add.Brand,
		Producer:    add.Producer,
		PriceRange:  add.PriceRange,
		Material:    add.Material,
		ArticleNo:   add.ArticleNo,
		Appraise:    add.Appraise,
		Question:    add.Question,
		ShowPic:     add.ShowPic,
		Rotation:    add.Rotation,
		Description: add.Description,
		Num:         add.Num,
		SellNum:     add.SellNum,
		Status:      add.Status}
	return goodsServiceImpl.GoodsAdd(goods)
}
func GoodsFindById(id int) (interface{}, error) {
	//var info=resDto.GoodsInfo{}
	info, err := goodsServiceImpl.GoodsById(uint(id))
	if err != nil {
		return nil, err
	}
	return info, nil
}

func GoodsFindByName(name string) (interface{}, error) {
	info, err := goodsServiceImpl.GoodsByName(name)
	if err != nil {
		return nil, err
	}
	return info, nil
}
func GoodsUpdate(update reqDto.GoodsUpdate) error {
	goods := pojo.Goods{
		Name:        update.Name,
		Price:       update.Price,
		Address:     update.Address,
		Fms:         update.Fms,
		Brand:       update.Brand,
		Producer:    update.Producer,
		PriceRange:  update.PriceRange,
		Material:    update.Material,
		ArticleNo:   update.ArticleNo,
		Appraise:    update.Appraise,
		Question:    update.Question,
		ShowPic:     update.ShowPic,
		Rotation:    update.Rotation,
		Description: update.Description,
		Num:         update.Num,
		SellNum:     update.SellNum,
		Status:      update.Status}
	goods.ID = update.Id
	return goodsServiceImpl.GoodsUpdate(goods)
}
