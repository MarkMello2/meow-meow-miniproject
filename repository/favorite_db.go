package repository

import (
	"gorm.io/gorm"
)

type favoriteRepositoryDb struct {
	gorm *gorm.DB
}

func NewFavoriteRepositoryDb(gorm *gorm.DB) FavoriteRepository {
	return favoriteRepositoryDb{gorm: gorm}
}

func (f favoriteRepositoryDb) GetById(id int) ([]FavoriteGet, error) {
	sql := "select f.id, p.price, f.quantity, f.created_at, p.code, p.name, p.description , p.rating ,p.image, f.user_id from favorites as f inner join products p on f.product_id = p.id where f.user_id = ? and f.deleted_at IS NULL"
	condition := f.gorm.Raw(sql, id)
	result := []FavoriteGet{}
	err := condition.Find(&result).Error
	return result, err
}

func (f favoriteRepositoryDb) Save(favData Favorite) error {
	favSave := Favorite{
		Id:        favData.Id,
		Price:     favData.Price,
		Quantity:  favData.Quantity,
		ProductId: favData.ProductId,
		UserId:    favData.UserId,
	}

	tx := f.gorm.Save(&favSave)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (f favoriteRepositoryDb) DeleteById(id int) error {
	err := f.gorm.Delete(&Favorite{Id: id}).Error
	return err
}
