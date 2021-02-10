package cryptoutils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

//CreateDesEncryptBlocks represent a function that create des encrypt block
func CreateDesEncryptBlocks(key []byte) (cipher.BlockMode, error) {

	if key == nil || len(key) != 8 {
		return nil, errors.New("Key size must be 8")
	}
	var (
		block cipher.Block
		iv    [8]byte
		err   error
	)

	if block, err = des.NewCipher(key[:]); err != nil {
		return nil, err
	}

	return cipher.NewCBCEncrypter(block, iv[:]), nil
}

//CreateDesDecryptBlocks represent a function that create des decrypt block
func CreateDesDecryptBlocks(key []byte) (cipher.BlockMode, error) {

	if key == nil || len(key) != 8 {
		return nil, errors.New("Key size must be 8")
	}
	var (
		block cipher.Block
		iv    [8]byte
		err   error
	)

	if block, err = des.NewCipher(key[:]); err != nil {
		return nil, err
	}

	return cipher.NewCBCDecrypter(block, iv[:]), nil
}

//AnsiZeroPadding make cipherValue right pad with zero til blockSize
func AnsiZeroPadding(cipherValue []byte, blockSize int) []byte {
	if len(cipherValue)%blockSize == 0 {
		return cipherValue
	}
	padding := blockSize - len(cipherValue)%blockSize
	paddingValue := bytes.Repeat([]byte{0}, padding)
	return append(cipherValue, paddingValue...)
}
