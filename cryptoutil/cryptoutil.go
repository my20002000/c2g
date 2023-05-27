package cryptoutil

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func Hashmd5(src string) string {
	srcCode := md5.Sum([]byte(src))
	code := fmt.Sprintf("%x", srcCode)
	return code
}
func Base64EncodeBytes(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}
func Base64EncodeString(s string) string {
	return Base64EncodeBytes([]byte(s))
}
func Base64DecodeBytes(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
func Base64DecodeString(s string) string {
	d, err := Base64DecodeBytes(s)
	if err != nil {
		return ""
	}
	return string(d)
}
func Base64SafeEncodeBytes(s []byte) string {
	return base64.URLEncoding.EncodeToString(s)
}
func Base64SafeEncodeString(s string) string {
	return Base64SafeEncodeBytes([]byte(s))
}
func Base64SafeDecodeBytes(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
func Base64SafeDecodeString(s string) string {
	d, err := Base64SafeDecodeBytes(s)
	if err != nil {
		return ""
	}
	return string(d)
}

// @param     encoding        string         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
func Base64EncodeBytesBy(s []byte, encoding string) string {
	customEncoding := base64.NewEncoding(encoding)
	return customEncoding.EncodeToString(s)
}

// @param     encoding        string         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
func Base64EncodeStringBy(s string, encoding string) string {
	customEncoding := base64.NewEncoding(encoding)
	return customEncoding.EncodeToString([]byte(s))
}
func Base64DecodeBytesBy(s string, encoding string) ([]byte, error) {
	customEncoding := base64.NewEncoding(encoding)
	return customEncoding.DecodeString(s)
}

// @param     encoding        string         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
func Base64DecodeStringBy(s string, encoding string) string {
	d, err := Base64DecodeBytesBy(s, encoding)
	if err != nil {
		return ""
	}
	return string(d)
}
func BaseN(src string, alphabet string) string {
	base := len(alphabet)
	var result strings.Builder

	// Convert source string to a decimal number
	var decNum big.Int
	for _, r := range src {
		idx := strings.IndexRune(alphabet, r)
		if idx == -1 {
			// Invalid character in the source string
			return ""
		}
		decNum.Mul(&decNum, big.NewInt(int64(base)))
		decNum.Add(&decNum, big.NewInt(int64(idx)))
	}

	// Convert decimal number to the base-N representation
	for decNum.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		mod.Mod(&decNum, big.NewInt(int64(base)))
		result.WriteByte(alphabet[mod.Int64()])
		decNum.Div(&decNum, big.NewInt(int64(base)))
	}

	// Reverse the result string
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
