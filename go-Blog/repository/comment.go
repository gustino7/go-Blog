package repository

import (
	"context"
	"tugas4/entity"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

type CommentRepository interface {
	AddComment(ctx context.Context, tx *gorm.DB, title string, comment string) (entity.BlogComment, error)
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

// fungsi comment dengan parameter title
func (r *commentRepository) AddComment(ctx context.Context, tx *gorm.DB, title string, comment string) (entity.BlogComment, error) {
	var blogComment entity.BlogComment
	var blog entity.Blog

	if tx == nil {
		//find id from title
		tx = r.db.WithContext(ctx).Debug().Model(entity.Blog{}).Where("title = ?", title).Take(&blog)
		if tx.Error != nil {
			return entity.BlogComment{}, tx.Error
		}

		//assign to blog comment
		blogComment.BlogID = blog.ID
		blogComment.Comment = comment

		tx = r.db.WithContext(ctx).Debug().Create(&blogComment)
		if tx.Error != nil {
			return entity.BlogComment{}, tx.Error
		}

	} else {
		tx.Error = r.db.WithContext(ctx).Debug().Model(entity.Blog{}).Where("title = ?", title).Take(&blog).Error
	}

	return blogComment, nil
}
