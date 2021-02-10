package ansimac

import (
	"crypto/cipher"

	"github.com/rezataroosheh/axe/internal/cryptoutils"
	"github.com/rezataroosheh/axe/internal/extension"
)

//RetailMacAlgorithm represents an Retail MAC algorithm.
type RetailMacAlgorithm struct {
	Key1, Key2, Key3 []byte
}

//GetName return result of retail MAC
func (algorithm RetailMacAlgorithm) GetName() string {
	return "Retail MAC"
}

//Compute return result of retail MAC algorithm
func (algorithm RetailMacAlgorithm) Compute(data []byte) ([]byte, error) {
	if data == nil {
		return nil, extension.NewNilReferenceError("data")
	}
	if algorithm.Key1 == nil {
		return nil, extension.NewNilReferenceError("key1")
	}
	if algorithm.Key2 == nil {
		return nil, extension.NewNilReferenceError("key2")
	}
	if algorithm.Key3 == nil {
		return nil, extension.NewNilReferenceError("key3")
	}
	var (
		cryted         []byte
		blockMode      cipher.BlockMode
		err            error
		mac99Algorithm = Ansi99MacAlgorithm{Key: algorithm.Key1}
	)

	if cryted, err = mac99Algorithm.Compute(data); err != nil {
		return nil, err
	}
	if blockMode, err = cryptoutils.CreateDesDecryptBlocks(algorithm.Key2); err != nil {
		return nil, err
	}
	blockMode.CryptBlocks(cryted, cryted)
	if blockMode, err = cryptoutils.CreateDesEncryptBlocks(algorithm.Key3); err != nil {
		return nil, err
	}
	blockMode.CryptBlocks(cryted, cryted)

	return cryted, nil
}
