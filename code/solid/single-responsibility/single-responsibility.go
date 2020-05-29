package single_responsibility

type IAnimal interface {
	GetAnimal() string
}

type IAnimalStorage interface {
	Save(animal Animal)
	Get(animal Animal)
}

type AnimalStorage struct{}

func (storage *AnimalStorage) Save(animal Animal) {
	// impl
}

func (storage *AnimalStorage) Get(animal Animal) {
	// impl
}

type Animal struct {
	name string
}

func (animal *Animal) GetName() string {
	return animal.name
}
