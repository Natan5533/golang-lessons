package vault

type Vault map[string]string

type ErrorsVault string

const (
	ErrorNotExist     = ErrorsVault("password does not exist")
	ErrorAlreadyExist = ErrorsVault("this password has already been created")
)

func (e ErrorsVault) Error() string {
	return string(e)
}

func (v Vault) Create(key, password string) error {

	_, err := v.Fetch(key)

	switch err {
	case ErrorNotExist:
		v[key] = password
	case nil:
		return ErrorAlreadyExist
	default:
		return err
	}

	return nil
}

func (v Vault) Fetch(key string) (string, error) {
	result, exist := v[key]

	if !exist {
		return "", ErrorNotExist
	}

	return result, nil
}

func (v Vault) Update(key, newPass string) error {
	_, err := v.Fetch(key)

	switch err {
	case ErrorNotExist:
		return ErrorNotExist
	case nil:
		v[key] = newPass
	default:
		return err
	}

	return nil
}

func (v Vault) Delete(key string) error {
	_, err := v.Fetch(key)

	switch err {
	case ErrorNotExist:
		return ErrorNotExist
	case nil:
		delete(v, key)
	default:
		return err
	}

	return nil
}

// func CheckKey(key s)
