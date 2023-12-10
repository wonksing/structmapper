package structmapper_test

import (
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
}

func Test_structMapper_Machine2(t *testing.T) {
	a := AnimalDto{
		Name: "asdf",
		Age:  19,
	}
	b := Animal{}
	err := structmapper.Map(&a, &b)
	assert.Nil(t, err)
	assert.EqualValues(t, a.Name, b.Name)
}
