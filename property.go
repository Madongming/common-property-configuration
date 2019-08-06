package property

func Create(obj InterStorage, key, data string) error {
	buffer := []byte(data)
	return obj.Create(key, buffer)
}

func Update(obj InterStorage, key, data string) error {
	buffer := []byte(data)
	return obj.Update(key, buffer)
}

func Read(obj InterStorage, key string) (string, error) {
	data, err := obj.Read(key)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Delete(obj InterStorage, key string) (string, error) {
	data, err := obj.Delete(key)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
