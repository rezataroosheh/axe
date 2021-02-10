package pinblock

// ECI1PinBlockAlgorithm represents ECI-4, that is same as ISO-0.
type ECI4PinBlockAlgorithm struct {
	Pan     string
	Padding PinBlockPadding
}

func NewECI4(pan string, padding PinBlockPadding) *ECI4PinBlockAlgorithm {
	return &ECI4PinBlockAlgorithm{
		Pan:     pan,
		Padding: padding,
	}
}

//GetName returns ECI4
func (algorithm ECI4PinBlockAlgorithm) GetName() string {
	return "ECI-4"
}

func (algorithm ECI4PinBlockAlgorithm) Encode(pin string) (string, error) {
	alg := NewFormat0(algorithm.Pan, algorithm.Padding)
	return alg.Encode(pin)
}
func (algorithm ECI4PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	alg := NewFormat0(algorithm.Pan, algorithm.Padding)
	return alg.Decode(pinblock)
}
