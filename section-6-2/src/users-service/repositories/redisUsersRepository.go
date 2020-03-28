package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yakushou730/microservices-with-go/section-6-2/src/users-service/entities"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisUsersRepository struct {
	pool *pool.Pool
}

const keySyntax = "user-%s"

func NewRedisUsersRepository() *RedisUsersRepository {
	pool, err := pool.New("tcp", "localhost:6380", 100)
	if err != nil {
		log.Fatal(err)
	}
	return &RedisUsersRepository{
		pool: pool,
	}
}

func (repo *RedisUsersRepository) GetUser(username string) (*entities.User, error) {
	conn, err := repo.pool.Get()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf(keySyntax, username)

	res, err := conn.Cmd("GET", key).Str()
	if err != nil {
		return nil, err
	}
	var User *entities.User
	err = json.Unmarshal([]byte(res), &User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *RedisUsersRepository) SetUser(username string, User *entities.User) error {
	conn, err := repo.pool.Get()
	if err != nil {
		return err
	}
	key := fmt.Sprintf(keySyntax, username)

	jsonBytes, err := json.Marshal(User)
	if err != nil {
		return err
	}
	_, err = conn.Cmd("SET", key, string(jsonBytes)).Str()
	if err != nil {
		return err
	}
	return nil
}
