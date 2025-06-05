package entities

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type User struct {
	ID        string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"notnull"`
	Email     string    `json:"email" valid:"email"`
	CreatedAt time.Time `json:"created_at" `
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewUser() *User {
	return &User{}
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}
