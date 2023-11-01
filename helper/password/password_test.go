package password

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := "correct password"
	salt := "salt"
	pepper := "pepper"

	hashedPassword, err := HashPassword(password, salt, pepper)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword, salt, pepper)
	require.NoError(t, err)

	wrongPassword := "wrong password"
	err = CheckPassword(wrongPassword, hashedPassword, salt, pepper)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
