package services

import (
	"context"
	"tugas4/entity"
	"tugas4/repository"
)

type commentService struct {
	commentRepo repository.CommentRepository
}

type CommentService interface {
	AddComment(ctx context.Context, title string, comment string) (entity.BlogComment, error)
}

func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (s *commentService) AddComment(ctx context.Context, title string, comment string) (entity.BlogComment, error) {
	blogComment, err := s.commentRepo.AddComment(ctx, nil, title, comment)
	if err != nil {
		return entity.BlogComment{}, err
	}

	return blogComment, nil
}
