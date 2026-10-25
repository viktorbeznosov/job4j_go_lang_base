package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newAccount() *Account {
    return &Account{Balance: 1000}
}

func TestDeposit(t *testing.T) {
	t.Parallel()

    tests := []struct {
        name     string
        raw      int64
        expected int64
        wantErr  bool
    }{
        {
            name:     "correct",
            raw:      100,
            expected: 1100,
            wantErr:  false,
        },
        {
            name:     "amount_less_than_zero",
            raw:      -100,
            wantErr:  true,
        },
        {
            name:     "more_than_max_int_64",
            raw:      9223372036854775807,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

             account := newAccount()
             err := account.Deposit(tt.raw)

            if tt.wantErr {
                assert.Error(t, err)
                return
            }

            assert.NoError(t, err)
            assert.Equal(t, tt.expected, account.Balance)
        })
    }
}

func TestWithdraw(t *testing.T) {
	t.Parallel()

    tests := []struct {
        name     string
        raw      int64
        expected int64
        wantErr  bool
    }{
        {
            name:     "correct",
            raw:      100,
            expected: 900,
            wantErr:  false,
        },
        {
            name:     "amount_less_than_zero",
            raw:      -100,
            wantErr:  true,
        },
        {
            name:     "amount_more_than_balance",
            raw:      2000,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

             account := newAccount()
             err := account.Withdraw(tt.raw)

            if tt.wantErr {
                assert.Error(t, err)
                return
            }

            assert.NoError(t, err)
            assert.Equal(t, tt.expected, account.Balance)
        })
    }
}

