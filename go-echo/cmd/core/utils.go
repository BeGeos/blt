package core

type _Utils struct{}

type stepper struct {
	start *int
}

func (s *stepper) increment() {
	*s.start++
}

func (s *stepper) decrement() {
	*s.start--
}

func (u *_Utils) CheckEnvVariables() error {
	// these are the variables that must be set in the environment
	return nil
}

func (u *_Utils) Stepper(start *int) stepper {
	return stepper{start: start}
}

var Utils = &_Utils{}
