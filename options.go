package dice

type options struct {
	operation func(dicesResult ...string) (string, error)
	frequency frequency
}

func NewOptions(operation func(dicesResult ...string) (string, error), frequency frequency) options {
	return options{operation, frequency}
}
