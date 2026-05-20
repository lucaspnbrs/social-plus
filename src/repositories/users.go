package repositories

import (
	"database/sql"
	"fmt"
	"social-plus/src/models"
)

type users struct {
	database *sql.DB
}

//Create a User Repository
func NewRepositoryFromUsers ( database *sql.DB) *users {
	return &users{database}
}

//Insert a User in the database
func (repository users ) Create( user models.User) (uint64, error) {
	statement, erro := repository.database.Prepare(
		"INSERT INTO users (nome, nick, email, pass ) values($1, $2, $3, $4) RETURNING id",
	)
	
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	var lastInserted uint64
	erro = statement.QueryRow(user.Nome, user.Nick, user.Email, user.Pass).Scan(&lastInserted)
	if erro != nil {
		return 0, erro
	}

	return lastInserted, nil

}

//Fetch users in the database
func (repository users) Search(nomeOuNick string) ([]models.User, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	lines, erro := repository.database.Query(
		"select id, nome, nick, email, created_at from users where nome LIKE $1 or nick LIKE $2",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

//Fetch User by ID in database 
func (repository users) FetchUserID(ID uint64) (models.User, error) {
	lines, erro := repository.database.Query(
		"SELECT id, nome, nick, email, created_at from users where id = $1",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

//Update a existing User with your fields
func (repository users) UpdateUser(ID uint64, user models.User) error {
	statement, erro := repository.database.Prepare(
		"UPDATE users SET nome = $1, nick = $2, email = $3 WHERE id = $4",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(user.Nome, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

//Delete user by ID
func(repository users) DeleteUser( ID uint64) error {
	statement, erro := repository.database.Prepare(
		"DELETE FROM users WHERE id = $1")
		if erro != nil {
			return erro
		}

		defer statement.Close()

		if _, erro = statement.Exec(ID); erro != nil {
			return erro
		}

		return nil 
}