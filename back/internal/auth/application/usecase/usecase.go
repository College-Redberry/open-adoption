package command

type Usecase[In, Out any] interface {
	Execute(input In) (Out, error)
}
