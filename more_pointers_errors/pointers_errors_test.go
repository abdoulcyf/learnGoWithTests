package more_pointers_errors

import (
	"testing"
)

func TestAccount(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Account{}
		wallet.Deposit(GBP(10))

		assertBalance(t, wallet, GBP(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Account{GBP(20)}
		//err := wallet.Withdraw(GBP(10))
		wallet.Withdraw(GBP(10))

		//assertNoError(t, err)
		assertBalance(t, wallet, GBP(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Account{GBP(20)}
		//err := wallet.Withdraw(GBP(10))
		wallet.Withdraw(GBP(10))

		//assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, GBP(10))
	})
}

func assertBalance(t testing.TB, wallet Account, want GBP) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
