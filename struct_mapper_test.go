package structmapper_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	structmapper.StoreMapper(&Person{}, &PersonEntity{})
}

func TestStructMapperCached(t *testing.T) {
	mobiles := make([]*Mobile, 0)
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})

	person := Person{
		Name:      "wonk",
		MobilePtr: &Mobile{Number: "010"},
		Mobile:    Mobile{Number: "20202"},
		Mobiles:   mobiles,
	}
	personEntity := PersonEntity{}
	structmapper.Map(&person, &personEntity)

	// fmt.Println(personEntity.String())
	assert.Equal(t, person.String(), personEntity.String())
}

func TestStructMapper(t *testing.T) {

	src := reflect.TypeOf(Person{})
	dest := reflect.TypeOf(PersonEntity{})
	sm := structmapper.NewStructMapper(src, dest)

	mobiles := make([]*Mobile, 0)
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})

	person := Person{
		Name:      "wonk",
		MobilePtr: &Mobile{Number: "010"},
		Mobile:    Mobile{Number: "20202"},
		Mobiles:   mobiles,
	}
	personEntity := PersonEntity{}
	sm.Map(&person, &personEntity)

	// fmt.Println(personEntity.String())
	assert.Equal(t, person.String(), personEntity.String())
}

type Person struct {
	Name      string
	MobilePtr *Mobile
	Mobile    Mobile
	Mobiles   []*Mobile
}

func (e *Person) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type PersonEntity struct {
	Name      string
	MobilePtr *MobileEntity
	Mobile    MobileEntity
	Mobiles   []*MobileEntity
}

func (e *PersonEntity) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Mobile struct {
	Number      string
	CountryCode string
	Provider    *Provider
}

type Provider struct {
	Name string
}

type MobileEntity struct {
	Number      string
	CountryCode string
	Provider    *Provider
}

func TestStruct(t *testing.T) {
	type Student struct {
		Name string
	}
	type School struct {
		Student Student
	}
	type StudentE struct {
		Name string
	}
	type SchoolE struct {
		Student StudentE
	}

	src := reflect.TypeOf(School{})
	dest := reflect.TypeOf(SchoolE{})
	sm := structmapper.NewStructMapper(src, dest)

	school := School{
		Student: Student{"wonk"},
	}
	schoolE := SchoolE{}
	sm.Map(&school, &schoolE)

	assert.Equal(t, school.Student.Name, schoolE.Student.Name)
}

func TestStructPointer(t *testing.T) {
	type Student struct {
		Name string
	}
	type School struct {
		Student *Student
	}
	type StudentE struct {
		Name string
	}
	type SchoolE struct {
		Student *StudentE
	}
	src := reflect.TypeOf(School{})
	dest := reflect.TypeOf(SchoolE{})
	sm := structmapper.NewStructMapper(src, dest)

	school := School{
		Student: &Student{"wonk"},
	}
	schoolE := SchoolE{}
	sm.Map(&school, &schoolE)

	assert.Equal(t, school.Student.Name, schoolE.Student.Name)
}
