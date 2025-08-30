package pet

type PetRepo interface {
	Create(pet Pet) error
	Update(id string, pet PetProps) error
	AdoptById(id string) error
	GetByID(id string) (Pet, error)
	List() ([]Pet, error)
	ListImagesById(id string) ([]string, error)
	SaveImagesById(id string, images []string) error
}
