package usecases

import (
	"github.com/yakushou730/microservices-with-go/section-6-2/src/users-service/entities"
	"github.com/yakushou730/microservices-with-go/section-6-2/src/users-service/repositories"
)

type UpdateUserUsecase struct {
	CacheRepo *repositories.RedisUsersRepository
	Repo      *repositories.MySQLUserRepository
}

func (uc *UpdateUserUsecase) UpdateUser(user *entities.User) error {
	// Update User DB
	err := uc.Repo.UpdateUser(user)
	if err != nil {
		return err
	}

	// Update Cache
	err = uc.CacheRepo.SetUser(user.Username, user)
	if err != nil {
		return err
	}
	return nil
}
