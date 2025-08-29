package command

type Usecase[In, Out any] interface {
	Execute(input In) (Out, error)
}

type UsecaseWithNoReturn[In any] interface {
	Execute(input In) error
}
