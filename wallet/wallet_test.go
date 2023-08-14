package wallet

import "testing"

func TestWallet(t *testing.T) {
	verifyBalance := func(t *testing.T, wallet Wallet, expectedBalance Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expectedBalance {
			t.Errorf("Failed! Current Balance: %s Expected Balance %s", result, expectedBalance)
		}
	}

	verifyError := func(t *testing.T, errorMessage error, expectedMessage error) {
		t.Helper()
		if errorMessage == nil {
			t.Error("Erro not found")
		}

		if errorMessage != expectedMessage {
			t.Errorf("Result: %s Expected: %s", errorMessage, expectedMessage)

		}
	}

	verifyUnexpectedError := func(t *testing.T, errorMessage error) {
		t.Helper()
		if errorMessage != nil {
			t.Error("Error unexpected received")
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		expectedBalance := Bitcoin(10)

		wallet.Deposit(Bitcoin(10))

		verifyBalance(t, wallet, expectedBalance)
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(50)}

		result := wallet.Withdraw(Bitcoin(25))

		expectedBalance := Bitcoin(25)

		verifyBalance(t, wallet, expectedBalance)
		verifyUnexpectedError(t, result)
	})

	t.Run("Withdraw insuficient found", func(t *testing.T) {

		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}

		errorMessage := wallet.Withdraw(Bitcoin(50))

		verifyBalance(t, wallet, initialBalance)
		verifyError(t, errorMessage, ErrorInsuficientFounds)
	})

}
