package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	//t.Skip()
	//--------------------------deposit----------------------------------
	t.Run("deposit", func(t *testing.T) {
		//---set up-------------------------
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()

		//----exp------------------------------
		want := Bitcoin(10)
		assertWallet(t, got, want)
	})
	//-----------------withdraw with funds------------------------
	t.Run("withdraw with funds", func(t *testing.T) {
		//---set up-------------------------
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(15))
		assertNoError(t, err)
		got := wallet.Balance()

		//------------exp----------------
		assertWallet(t, got, Bitcoin(5))
	})

	//-----------------withdraw insufficient funds------------------------
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		//---set up-------------------------
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrorInsufficientFunds)
		got := wallet.Balance()
		//---------exp------------------------
		want := Bitcoin(20)

		assertWallet(t, got, want)
	})
}

func assertWallet(t testing.TB, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
