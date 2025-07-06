package core

type stepper struct {
	start *int
}

func (s *stepper) increment() {
	*s.start++
}

func (s *stepper) decrement() {
	*s.start--
}

func CheckEnvVariables() error {
	// these are the variables that must be set in the environment
	return nil
}

func Stepper(start *int) stepper {
	return stepper{start: start}
}
