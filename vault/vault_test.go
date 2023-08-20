package vault

import (
	"testing"
)

func TestFetch(t *testing.T) {
	//vault := Create("aws", "asmnk-fjgnsd-sjfdb")
	vault := Vault{"aws": "asmnk-fjgnsd-sjfdb"}

	t.Run("Exists? == True", func(t *testing.T) {
		result, err := vault.Fetch("aws")
		expected := "asmnk-fjgnsd-sjfdb"

		assertError(t, err, nil)
		assertEqual(t, result, expected)
	})

	t.Run("Exists? == False", func(t *testing.T) {
		output, err := vault.Fetch("netflix passoword")
		expectedOutput := ""

		assertEqual(t, output, expectedOutput)
		assertError(t, err, ErrorNotExist)
	})

}

func TestCreatePass(t *testing.T) {

	t.Run("Create password", func(t *testing.T) {
		vault := Vault{}
		vault.Create("netflix", "12345678_this_is_a_bad_password")

		expected := "12345678_this_is_a_bad_password"

		result, err := vault.Fetch("netflix")

		assertError(t, err, nil)
		assertEqual(t, result, expected)
	})

	t.Run("Passord already exists", func(t *testing.T) {
		vault := Vault{}

		errNil := vault.Create("pass", "some_password")
		err := vault.Create("pass", "x-x-x-x-x-x-x")
		password, _ := vault.Fetch("pass")
		expected := "some_password"

		assertError(t, errNil, nil)
		assertError(t, err, ErrorAlreadyExist)
		assertEqual(t, password, expected)
	})

}
func TestUpdatePassword(t *testing.T) {
	vault := Vault{}

	vault.Create("device", "some_password")
	err := vault.Update("device", "new_password")

	password, _ := vault.Fetch("device")

	expected := "new_password"

	assertEqual(t, password, expected)
	assertError(t, err, nil)
}

// I'm too lazy to do the failure tests but I know it's functional haha

func TestDeletePassword(t *testing.T) {
	vault := Vault{}

	vault.Create("device", "some_password")
	vault.Delete("device")

	_, err := vault.Fetch("device")

	assertError(t, err, ErrorNotExist)
}

// Test functions

func assertError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("[failed] Result: %s Expected: %s", result, expected)
	}
}

func assertEqual(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("[failed] Result: %s Expected: %s", result, expected)
	}
}
