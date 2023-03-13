package services

import (
	"context"
	"errors"
	"fmt"
	"tugas4/dto"
	"tugas4/entity"
	"tugas4/repository"
	"tugas4/utils"

	"github.com/jinzhu/copier"
)

type userService struct {
	userRepo repository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, userDTO dto.RegisterUser) (entity.User, error)
	Login(ctx context.Context, email string, password string) (bool, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, userDTO dto.LoginUser) (entity.User, error)
	FindUser(ctx context.Context, userDTO dto.LoginUser) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint64) (entity.User, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, userDTO dto.RegisterUser) (entity.User, error) {
	var user entity.User
	copier.Copy(&user, &userDTO)

	fmt.Println(user)
	userRegist, err := s.userRepo.FindUser(ctx, nil, user.Email)
	if err != nil {
		fmt.Println("kondisi ini")
		return entity.User{}, err
	}
	fmt.Println(userRegist.Email)
	if userRegist.Email == user.Email {
		return entity.User{}, errors.New("email or username already use")
	}
	//fmt.Println(user)

	createdUser, err := s.userRepo.CreateUser(ctx, nil, user) //stop disini
	if err != nil {
		return entity.User{}, err
	}
	return createdUser, nil
}

func (s *userService) Login(ctx context.Context, email string, password string) (bool, error) {
	var user entity.User

	user, err := s.userRepo.FindUser(ctx, nil, email)
	if err != nil {
		return false, err
	}

	checkPass, err := utils.PasswordCompare(user.Password, []byte(password))
	if err != nil {
		return false, err
	}
	if user.Email == email && checkPass {
		return true, nil
	}

	return false, err
}

func (s *userService) Update(ctx context.Context, user entity.User) (entity.User, error) {
	updUser, err := s.userRepo.Update(ctx, nil, user)
	if err != nil {
		return entity.User{}, err
	}

	return updUser, nil
}

func (s *userService) Delete(ctx context.Context, userDTO dto.LoginUser) (entity.User, error) {
	var user entity.User
	copier.Copy(&user, &userDTO)

	delUser, err := s.userRepo.Delete(ctx, nil, user)
	if err != nil {
		return entity.User{}, err
	}

	return delUser, nil
}

func (s *userService) FindUser(ctx context.Context, userDTO dto.LoginUser) (entity.User, error) {
	var user entity.User
	copier.Copy(&user, &userDTO)

	user, err := s.userRepo.FindUser(ctx, nil, user.Email)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, userID uint64) (entity.User, error) {
	var user entity.User
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
