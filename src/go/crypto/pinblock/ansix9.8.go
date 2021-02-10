package pinblock

// AnsiX98PinBlockAlgorithm represents Ansi X9.8, that is same as ISO-0.
type AnsiX98PinBlockAlgorithm struct {
	Pan     string
	Padding PinBlockPadding
}

func NewAnsiX98(pan string, padding PinBlockPadding) *AnsiX98PinBlockAlgorithm {
	return &AnsiX98PinBlockAlgorithm{
		Pan:     pan,
		Padding: padding,
	}
}

//GetName returns Ansi X9.8
func (algorithm AnsiX98PinBlockAlgorithm) GetName() string {
	return "Ansi X9.8"
}

func (algorithm AnsiX98PinBlockAlgorithm) Encode(pin string) (string, error) {
	alg := NewFormat0(algorithm.Pan, algorithm.Padding)
	return alg.Encode(pin)
}
func (algorithm AnsiX98PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	alg := NewFormat0(algorithm.Pan, algorithm.Padding)
	return alg.Decode(pinblock)
}
