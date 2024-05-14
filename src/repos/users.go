package repos

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into users (name, nick, email, password, role_id) values (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password, user.Role)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastId), nil
}

func (repo users) Search(urlQuery string) ([]models.User, error) {
	urlQuery = fmt.Sprintf("%%%s%%", urlQuery)

	lines, err := repo.db.Query(
		"select id, name, nick, email, role_id, created from users where name like ? or nick like ?",
		urlQuery, urlQuery,
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
			&user.Role,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repo users) SearchById(Id uint64) (models.User, error) {
	lines, err := repo.db.Query(
		"select id, name, nick, email, created from users where id = ?",
		Id,
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

func (repo users) Update(ID uint64, user models.User) error {
	statement, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repo users) Delete(ID uint64) error {
	statement, err := repo.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repo users) SearchByEmail(email string) (models.User, error) {
	lines, err := repo.db.Query(
		"select id, password, role_id from users where email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Password, &user.Role); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo users) SearchPass(userId uint64) (string, error) {
	line, err := repo.db.Query(
		"select password from users where id = ?",
		userId,
	)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repo users) UpdatePass(userId uint64, hashedPass string) error {
	statement, err := repo.db.Prepare(
		"update users set password = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(hashedPass, userId); err != nil {
		return err
	}

	return nil
}
