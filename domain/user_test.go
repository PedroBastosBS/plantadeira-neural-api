package domain_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"plantadeira-neural-api/domain"
	"testing"
	"time"
)

func TestNewUser_Valid(t *testing.T) {
	user := domain.NewUser()
	user.ID = uuid.New().String()
	user.Name = "Pedro Bastos"
	user.Email = "pedro@example.com"
	user.CreatedAt = time.Now()

	err := user.Validate()
	require.Nil(t, err)
}

func TestNewUser_InvalidUUID(t *testing.T) {
	user := domain.NewUser()
	user.ID = "invalid-uuid"
	user.Name = "Pedro Bastos"
	user.Email = "pedro@example.com"
	user.CreatedAt = time.Now()

	err := user.Validate()
	require.NotNil(t, err)
}

func TestNewUser_EmptyName(t *testing.T) {
	user := domain.NewUser()
	user.ID = uuid.New().String()
	user.Name = ""
	user.Email = "pedro@example.com"
	user.CreatedAt = time.Now()

	err := user.Validate()
	require.NotNil(t, err)
}

func TestNewUser_InvalidEmail(t *testing.T) {
	user := domain.NewUser()
	user.ID = uuid.New().String()
	user.Name = "Pedro Bastos"
	user.Email = "invalid-email"
	user.CreatedAt = time.Now()

	err := user.Validate()
	require.NotNil(t, err)
}
