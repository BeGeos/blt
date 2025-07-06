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

type Utils interface {
	Stepper(start *int) stepper
	CheckEnvVariables() error
}
type utils struct{}

func (u *utils) CheckEnvVariables() error {
	// these are the variables that must be set in the environment
	return nil
}

func (u *utils) Stepper(start *int) stepper {
	return stepper{start: start}
}

func NewUtils() Utils {
	return &utils{}
}
