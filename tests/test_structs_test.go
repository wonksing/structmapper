package structmapper_test

type Animal struct {
	Name string
}

type AnimalDto struct {
	Name string
	Age  int
}

type Machine struct {
	Animal  Animal
	Name    *string
	Age     *int
	Bullets []string
	Animals []Animal
}

type MachineDto struct {
	Animal  AnimalDto
	Name    *string
	Age     *int
	Bullets []string
	Animals []AnimalDto
}
