package ansimac

//AnsiMacAlgorithm represents an interface for computing ansi MAC (Message Authentication Code) algorithms.
type AnsiMacAlgorithm interface {
	Compute(data []byte) ([]byte, error)
	GetName() string
}
