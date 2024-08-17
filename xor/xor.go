package xor

func XorStrings(text, key string) []byte {
	data := []byte(text)
	k := []byte(key)

	return Xor(data, k)
}

func XorByte(data []byte, key byte) []byte {
	var bytes []byte

	for _, c := range data {
		bytes = append(bytes, c^key)
	}

	return bytes
}

func Xor(data, key []byte) []byte {
	var bytes []byte
	i := 0
	for _, c := range data {
		bytes = append(bytes, c^key[i])
		i++

		if i == len(key) {
			i = 0
		}
	}

	return bytes
}
