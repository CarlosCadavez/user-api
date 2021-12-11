package repositories

import (
	"database/sql"
	"fmt"
	"user-api/src/models"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users(name, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil

}

func (repository users) GetUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repository users) GetUserById(userId uint64) (models.User, error) {

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		userId,
	)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}

	}

	return user, nil

}

func (repository users) UpdateUser(userId uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ? ",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userId); err != nil {
		return err
	}

	return nil
}
