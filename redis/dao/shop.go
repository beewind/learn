package dao

import "redis-learn/entity"

const TbShop string = "tb_shop"
const TbShopType string = "tb_shop_type"

func SelectShopById(id string) (entity.Shop, error) {
	var shop entity.Shop
	err := db.Table(TbShop).Where("id = ?", id).First(&shop).Error
	return shop, err
}
func SelectShopList() ([]entity.ShopType, error) {
	shopTypeList := []entity.ShopType{}
	err := db.Table(TbShopType).Order("sort Asc").Find(&shopTypeList).Error
	return shopTypeList, err
}
func UpdateShopById(shop entity.Shop) error {
	return db.Table(TbShop).Updates(shop).Error

}
func SaveShop(shop entity.Shop) error {
	return db.Table(TbShop).Omit("distance", "id").Create(shop).Error
}
