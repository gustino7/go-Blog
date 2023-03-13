package services

import (
	"context"
	"tugas4/dto"
	"tugas4/entity"
	"tugas4/repository"

	"github.com/jinzhu/copier"
)

type blogService struct {
	blogRepo repository.BlogRepository
}

type BlogService interface {
	UploadBlog(ctx context.Context, dtoBlog dto.UpBlog) (entity.Blog, error)
	GetBlogComment(ctx context.Context, blog entity.Blog) ([]entity.Blog, error)
	Like(ctx context.Context, title string) (entity.Blog, error)
}

func NewBlogService(blogRepo repository.BlogRepository) BlogService {
	return &blogService{
		blogRepo: blogRepo,
	}
}

func (s *blogService) UploadBlog(ctx context.Context, dtoBlog dto.UpBlog) (entity.Blog, error) {
	var blog entity.Blog
	copier.Copy(&blog, &dtoBlog)

	upBlog, err := s.blogRepo.UploadBlog(ctx, nil, blog)
	if err != nil {
		return entity.Blog{}, err
	}

	return upBlog, nil
}

func (s *blogService) GetBlogComment(ctx context.Context, blog entity.Blog) ([]entity.Blog, error) {
	var getBlog []entity.Blog
	getBlog, err := s.blogRepo.GetBlogComment(ctx, nil, blog)
	if err != nil {
		return []entity.Blog{}, err
	}

	return getBlog, nil
}

func (s *blogService) Like(ctx context.Context, title string) (entity.Blog, error) {
	var getBlog entity.Blog
	getBlog, err := s.blogRepo.Like(ctx, nil, title)
	if err != nil {
		return entity.Blog{}, err
	}

	return getBlog, nil
}

// func (s *blogService) Comment(ctx context.Context, comment string) (entity.Blog, error) {
// 	var getBlog entity.Blog
// 	getBlog, err := s.blogRepo.Comment(ctx, nil, comment)
// 	if err != nil {
// 		return entity.Blog{}, err
// 	}

// 	return getBlog, nil
// }
