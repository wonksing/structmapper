package structmapper_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&Animal{}, &AnimalDto{})
	structmapper.StoreMapper(&AnimalDto{}, &Animal{})

	structmapper.StoreMapper(&Machine{}, &MachineDto{})
	structmapper.StoreMapper(&MachineDto{}, &Machine{})

	structmapper.StoreMapper(&WithMap{}, &WithMapDto{})
	structmapper.StoreMapper(&WithMapDto{}, &WithMap{})
}

func Test_structMapper_Animal1(t *testing.T) {
	a := Animal{
		Name: "asdf",
	}
	b := AnimalDto{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a.Name, b.Name)
	assert.EqualValues(t, 0, b.Age)
}

func Test_structMapper_Animal2(t *testing.T) {
	a := AnimalDto{
		Name: "asdf",
		Age:  19,
	}
	b := Animal{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a.Name, b.Name)
}

func Test_structMapper_Machine1(t *testing.T) {
	str := uuid.NewString()
	i := rand.Intn(100)
	a := Machine{
		Animal: Animal{
			Name: "asdf",
		},
		Name:    &str,
		Age:     &i,
		Bullets: []string{"awefi", "zxcvd", "feifeoe9"},
		Animals: []Animal{
			Animal{
				Name: "1asdf",
			},
			Animal{
				Name: "33asdf",
			},
			Animal{
				Name: "33asdf44",
			},
			Animal{
				Name: "33asdf44",
			},
		},
	}
	b := MachineDto{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a.Name, b.Name)
	assert.EqualValues(t, a.Age, b.Age)
	assert.EqualValues(t, a.String(), b.String())
	fmt.Println(a.String())
	fmt.Println(b.String())
}

func Test_structMapper_Machine2(t *testing.T) {
	str := uuid.NewString()
	i := rand.Intn(100)
	a := MachineDto{
		Animal: AnimalDto{
			Name: "asdf",
		},
		Name:    &str,
		Age:     &i,
		Bullets: []string{"awefi", "zxcvd", "feifeoe9"},
		Animals: []AnimalDto{
			AnimalDto{
				Name: "1asdf",
				Age:  0,
			},
			AnimalDto{
				Name: "33asdf",
			},
			AnimalDto{
				Name: "33asdf44",
			},
			AnimalDto{
				Name: "33asdf44",
			},
		},
	}
	b := Machine{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a.Name, b.Name)
	assert.EqualValues(t, a.Age, b.Age)
	assert.EqualValues(t, a.String(), b.String())
	fmt.Println(a.String())
	fmt.Println(b.String())
}

func Test_structMapper_WithMap1(t *testing.T) {
	m := make(map[interface{}]interface{})
	m[1] = "1"
	a := WithMap{
		Map: m,
	}
	b := WithMapDto{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a, b)
	// m[1] = "2"
	// assert.NotEqualValues(t, a, b)
}

func Test_structMapper_WithMap2(t *testing.T) {
	m := make(map[interface{}]interface{})
	m[1] = "1"
	a := WithMapDto{
		Map: m,
	}
	b := WithMap{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a, b)
	// m[1] = "2"
	// assert.NotEqualValues(t, a, b)
}
