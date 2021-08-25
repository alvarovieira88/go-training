package entity_test

import (
	"99-project/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	b, err := entity.NewAccount("CONTA_CORRENTE","1234567")
	assert.Nil(t, err)
	assert.Equal(t, b.AccountNumber, "1234567")
	assert.NotNil(t, b.ID)
}

func TestAccountValidate(t *testing.T) {
	type test struct {
		AccountType   string
		AccountNumber string
		want           error
	}

	tests := []test{
		{
			AccountType:           "CONTA_CORRENTE",
			AccountNumber: "1234567",
			want:           nil,
		},
		{
			AccountType:           "CONTA_CORRENTE",
			AccountNumber: "",
			want:           entity.ErrAccountNumberInvalid,
		},

	}
	for _, tc := range tests {
		_, err := entity.NewAccount(tc.AccountType, tc.AccountNumber)
		assert.Equal(t, err, tc.want)
	}
}