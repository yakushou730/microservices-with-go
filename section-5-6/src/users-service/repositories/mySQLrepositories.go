package repositories

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yakushou730/microservices-with-go/section-5-4/src/users-service/entities"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository() *MySQLUserRepository {

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "packt:packt@tcp(localhost:3307)/users?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	repo := &MySQLUserRepository{
		db,
	}

	return repo
}

func (repo *MySQLUserRepository) Close() {
	repo.db.Close()
}

func (repo *MySQLUserRepository) GetUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}
	row := repo.db.QueryRow("Select id, username, first_name, last_name, email, birthdate, added, account, password from users where username=?", username)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.Added, &user.Account, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MySQLUserRepository) GetUserByID(userID uint32) (*entities.User, error) {
	user := &entities.User{}
	row := repo.db.QueryRow("Select id, username, first_name, last_name, email, birthdate, added, account, password from users where id=?", userID)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.Added, &user.Account, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MySQLUserRepository) UpdateUser(user *entities.User) error {
	row := repo.db.QueryRow("Update users set first_name=?, last_name=?, email=?, birthdate=?, account=?, password=? from users where id=?", user.FirstName, user.LastName, user.Email, user.BirthDate, user.Account, user.Password)
	err := row.Scan()
	if err != nil {
		return err
	}
	return nil
}
