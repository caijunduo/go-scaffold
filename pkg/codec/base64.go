package codec

import "encoding/base64"

func Base64StdDecodeStd(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
