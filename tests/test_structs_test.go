package structmapper_test

import "encoding/json"

type Animal struct {
	Name string `json:"name,omitempty"`
}

type AnimalDto struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Machine struct {
	Animal  Animal   `json:"animal,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Age     *int     `json:"age,omitempty"`
	Bullets []string `json:"bullets,omitempty"`
	Animals []Animal `json:"animals,omitempty"`
}

func (e *Machine) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type MachineDto struct {
	Animal  AnimalDto   `json:"animal,omitempty"`
	Name    *string     `json:"name,omitempty"`
	Age     *int        `json:"age,omitempty"`
	Bullets []string    `json:"bullets,omitempty"`
	Animals []AnimalDto `json:"animals,omitempty"`
}

func (e *MachineDto) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type WithMap struct {
	Map map[interface{}]interface{}
}

type WithMapDto struct {
	Map map[interface{}]interface{}
}
