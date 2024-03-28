package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func stringHash(source string) string {
	sum := sha256.Sum256([]byte(source))
	var sb strings.Builder
	// 遍历每个字节，转换为二进制字符串
	for _, b := range sum {
		sb.WriteString(fmt.Sprintf("%08b", b))
	}
	// 取前64位
	return sb.String()[:64]
}
func simhash(token string) string {
	v := make([]int, 64)
	tokens := strings.Split(token, " ")
	for _, t := range tokens {
		tHash := stringHash(t)
		for i := 0; i < 64; i++ {
			if tHash[i] == '1' {
				v[i]++
			} else {
				v[i]--
			}
		}
	}
	var fingerprint strings.Builder
	for i := 0; i < 64; i++ {
		if v[i] >= 0 {
			fingerprint.WriteString("1")
		} else {
			fingerprint.WriteString("0")
		}
	}
	return binaryToHex(fingerprint.String())
}

func hammingDistance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("strings must be of the same length")
	}
	distance := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance, nil
}
func hexToBinary(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	binaryStr := ""
	for _, b := range bytes {
		binaryStr += fmt.Sprintf("%08b", b)
	}
	return binaryStr, nil
}
func similarity(token1, token2 string) (float64, error) {
	hex1 := simhash(token1)
	hex2 := simhash(token2)
	fmt.Println(hex1)
	fmt.Println(hex2)
	bin1, err := hexToBinary(hex1)
	if err != nil {
		return 0, err
	}
	bin2, err := hexToBinary(hex2)
	if err != nil {
		return 0, err
	}
	// fmt.Println(bin1, bin2)
	distance, err := hammingDistance(bin1, bin2)
	if err != nil {
		return 0, err
	}
	return 1 - float64(distance)/float64(len(bin1)), nil
}
func binaryToHex(binaryStr string) string {
	num, _ := new(big.Int).SetString(binaryStr, 2)
	return hex.EncodeToString(num.Bytes())
}

func main() {
	tokens1 := "this is a test phrase"
	tokens2 := "this is a test phrasi"
	fmt.Println(similarity(tokens1, tokens2))
}
