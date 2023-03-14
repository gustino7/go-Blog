package repository //BERINTERAKSI DENGAN DATABASE

import (
	"context"
	"fmt"
	"tugas4/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	FindUser(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
	FindUserByID(ctx context.Context, userID uint64) (entity.User, error)
	Update(ctx context.Context, tx *gorm.DB, nama string, userid uint64) (entity.User, error)
	Delete(ctx context.Context, tx *gorm.DB, userID uint64) (bool, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	var err error
	if tx == nil {
		tx = r.db.Create(&user)
		err = tx.Error
	} else {
		err = r.db.Create(&user).Error
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, tx *gorm.DB, nama string, userid uint64) (entity.User, error) {
	var updUser entity.User

	if tx == nil {
		fmt.Println(nama, userid)
		tx = r.db.WithContext(ctx).Debug().Model(entity.User{}).Where("id = ?", userid).Update("nama", nama)
		if tx.Error != nil {
			fmt.Println("error update")
			return entity.User{}, tx.Error
		}

		//take data user update
		tx = r.db.WithContext(ctx).Debug().Model(entity.User{}).Where("id = ?", userid).Take(&updUser)
		if tx.Error != nil {
			fmt.Println("gagal ambil user")
			return entity.User{}, tx.Error
		}
	} else {
		return entity.User{}, tx.Error
	}

	return updUser, nil
}

func (r *userRepository) Delete(ctx context.Context, tx *gorm.DB, userID uint64) (bool, error) {
	var err error

	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Model(entity.User{}).Where("id = ?", userID).Delete(&entity.User{})
		err = tx.Error

	} else {
		err = r.db.WithContext(ctx).Debug().Model(entity.User{}).Where("id = ?", userID).Delete(&entity.User{}).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// JANGAN RUBAH FIND USER UDAH BENER
func (r *userRepository) FindUser(ctx context.Context, tx *gorm.DB, email string) (entity.User, error) {
	var err error
	var user entity.User
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Where(("email = ?"), email).Take(&user)
		err = tx.Error

	} else {
		err = r.db.WithContext(ctx).Debug().Where(("email = ?"), email).Take(&user).Error
	}

	if err == nil {
		return user, err
	}
	return entity.User{}, err
}

func (r *userRepository) FindUserByID(ctx context.Context, userID uint64) (entity.User, error) {
	var user entity.User

	tx := r.db.WithContext(ctx).Debug().Where(("id = ?"), userID).Preload("Blog").Take(&user)
	err := tx.Error

	if err == nil {
		return user, nil
	}
	return entity.User{}, err
}
