package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Role      uint64    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório!")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatório!")
	}

	if user.Email == "" {
		return errors.New("O email é obrigatório!")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O email inserido é inválido!")
	}

	if step == "register" && user.Password == "" {
		return errors.New("A senha é obrigatória!")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passHash)
	}

	return nil
}
