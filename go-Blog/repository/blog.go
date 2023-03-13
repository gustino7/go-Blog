package repository

import (
	"context"
	"fmt"
	"tugas4/entity"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

type BlogRepository interface {
	UploadBlog(ctx context.Context, tx *gorm.DB, blog entity.Blog) (entity.Blog, error)
	GetBlogComment(ctx context.Context, tx *gorm.DB, blog entity.Blog) ([]entity.Blog, error)
	Like(ctx context.Context, tx *gorm.DB, title string) (entity.Blog, error)
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{
		db: db,
	}
}

func (r *blogRepository) UploadBlog(ctx context.Context, tx *gorm.DB, blog entity.Blog) (entity.Blog, error) {
	var err error
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Create(&blog)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}

func (r *blogRepository) GetBlogComment(ctx context.Context, tx *gorm.DB, blog entity.Blog) ([]entity.Blog, error) {
	var err error
	var getBlog []entity.Blog
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Preload("Comment").Find(&getBlog)
		err = tx.Error
	} else {
		err = r.db.WithContext(ctx).Debug().Preload("Comment").Find(&getBlog).Error
	}
	if err != nil {
		fmt.Println("kondisi ini")
		return []entity.Blog{}, err
	}

	fmt.Println("kondisi 2")
	return getBlog, nil
}

func (r *blogRepository) Like(ctx context.Context, tx *gorm.DB, title string) (entity.Blog, error) {
	var err error
	var like entity.Blog

	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Model(entity.Blog{}).Where("title = ?", title).Take(&like)
		err = tx.Error

	} else {
		err = r.db.WithContext(ctx).Debug().Model(entity.Blog{}).Where("title = ?", title).Take(&like).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Model(entity.Blog{}).Where("title = ?", title).Update("like", like.Like+1)
		err = tx.Error

	} else {
		err = r.db.WithContext(ctx).Debug().Model(entity.Blog{}).Where("title = ?", title).Update("like", like.Like+1).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	return like, nil
}
