package ansimac

import (
	"crypto/cipher"

	"github.com/rezataroosheh/axe/internal/cryptoutils"
	"github.com/rezataroosheh/axe/internal/extension"
)

//Ansi99MacAlgorithm represents an ansi 9.9 MAC algorithm.
type Ansi99MacAlgorithm struct {
	Key []byte
}

//GetName return result of Ansi 9.9
func (algorithm Ansi99MacAlgorithm) GetName() string {
	return "Ansi 9.9 MAC"
}

//Compute return result of ansi 9.9 MAC algorithm
func (algorithm Ansi99MacAlgorithm) Compute(data []byte) ([]byte, error) {

	if data == nil {
		return nil, extension.NewNilReferenceError("data")
	}
	if algorithm.Key == nil {
		return nil, extension.NewNilReferenceError("key")
	}
	var (
		blockMode cipher.BlockMode
		blockSize = 8
		err       error
	)
	cryted := make([]byte, blockSize)
	data = cryptoutils.AnsiZeroPadding(data, blockSize)

	for index := 0; index < len(data); index += blockSize {
		extension.Xor(cryted, data[index:index+blockSize])

		if blockMode, err = cryptoutils.CreateDesEncryptBlocks(algorithm.Key); err != nil {
			return nil, err
		}

		blockMode.CryptBlocks(cryted, cryted)
	}
	return cryted, nil
}
