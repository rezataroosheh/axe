package pinblock

// ECI1PinBlockAlgorithm represents ECI-1, that is same as ISO-0.
type ECI1PinBlockAlgorithm struct {
	Pan     string
	Padding PinBlockPadding
}

func NewECI1(pan string, padding PinBlockPadding) *ECI1PinBlockAlgorithm {
	return &ECI1PinBlockAlgorithm{
		Pan:     pan,
		Padding: padding,
	}
}

//GetName returns ECi-1
func (algorithm ECI1PinBlockAlgorithm) GetName() string {
	return "ECi-1"
}

func (algorithm ECI1PinBlockAlgorithm) Encode(pin string) (string, error) {
	alg := NewFormat0(algorithm.Pan, algorithm.Padding)
	return alg.Encode(pin)
}
func (algorithm ECI1PinBlockAlgorithm) Decode(pinblock string) (string, error) {
	alg := NewFormat0(algorithm.Pan, algorithm.Padding)
	return alg.Decode(pinblock)
}
