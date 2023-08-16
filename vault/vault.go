package vault

func Create(key string, password string) map[string]string {
	newVault := map[string]string{key: password}

	return newVault
}

func Fetch(vault map[string]string, key string) string {
	return vault[key]
}
