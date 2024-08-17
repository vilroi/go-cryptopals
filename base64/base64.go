package base64

var base64Table = [64]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
}

func Decode(bytes []byte) string {
	var s string
	var padlen int

	if len(bytes)%3 != 0 {
		padlen = 3 - (len(bytes) % 3)
		for i := 0; i < padlen; i++ {
			bytes = append(bytes, 0)
		}
	}

	for i := 0; i < len(bytes); i += 3 {
		tmp := uint32(bytes[i])<<16 | uint32(bytes[i+1])<<8 | uint32(bytes[i+2])

		s += string(base64Table[(tmp&0xfc0000)>>18])
		s += string(base64Table[(tmp&0x3f000)>>12])
		s += string(base64Table[(tmp&0xfc0)>>6])
		s += string(base64Table[tmp&0x3f])
	}

	s = s[:len(s)-padlen]
	for i := 0; i < padlen; i++ {
		s += "="
	}

	return s
}
