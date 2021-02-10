package ansimac

import "github.com/rezataroosheh/axe/internal/extension"

//Ansi919MacAlgorithm represents an ansi 9.19 MAC algorithm.
type Ansi919MacAlgorithm struct {
	Key1, Key2 []byte
}

//GetName return result of Ansi 9.19
func (algorithm Ansi919MacAlgorithm) GetName() string {
	return "Ansi 9.19 MAC"
}

//Compute return result of ansi 9.19 MAC algorithm
func (algorithm Ansi919MacAlgorithm) Compute(data []byte) ([]byte, error) {
	if data == nil {
		return nil, extension.NewNilReferenceError("data")
	}
	if algorithm.Key1 == nil {
		return nil, extension.NewNilReferenceError("key1")
	}
	if algorithm.Key2 == nil {
		return nil, extension.NewNilReferenceError("key2")
	}
	var (
		cryted          []byte
		err             error
		retailAlgorithm = RetailMacAlgorithm{Key1: algorithm.Key1, Key2: algorithm.Key2, Key3: algorithm.Key1}
	)

	if cryted, err = retailAlgorithm.Compute(data); err != nil {
		return nil, err
	}

	return cryted, nil
}
