package request

type Repo interface {
	Create(request Request) error
	Aprove(id string) error
	List() ([]Request, error)
	GetByID(id string) (Request, error)
}
