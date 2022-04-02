package api

import (
	db "goProject/db/sqlc"
	"goProject/until"
	"testing"

	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = until.RandomString(6)
	hashedPassword, err := until.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       until.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       until.RandomOwner(),
		Email:          until.RandomEmail(),
	}

	return
}
