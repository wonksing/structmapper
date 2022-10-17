package structmapper_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wonksing/structmapper"
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

func TestDefaultTimePackage(t *testing.T) {
	type School struct {
		CreatedAt time.Time
		UpdatedAt *time.Time
		DeletedAt sql.NullTime
	}
	type SchoolE struct {
		CreatedAt time.Time
		UpdatedAt *time.Time
		DeletedAt sql.NullTime
	}
	src := reflect.TypeOf(School{})
	dest := reflect.TypeOf(SchoolE{})
	sm := structmapper.NewStructMapper(src, dest)

	createdAt, _ := time.Parse("20060102150405", "20220101121212")
	updatedAt, _ := time.Parse("20060102150405", "20221101121212")
	deletedAt, _ := time.Parse("20060102150405", "20231101121212")
	school := School{
		CreatedAt: createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: sql.NullTime{Time: deletedAt, Valid: true},
	}
	schoolE := SchoolE{}
	sm.Map(&school, &schoolE)

	assert.Equal(t, school.CreatedAt, schoolE.CreatedAt)
	assert.Equal(t, school.UpdatedAt, schoolE.UpdatedAt)
	assert.Equal(t, school.DeletedAt.Time, schoolE.DeletedAt.Time)
}

func TestDefaultTimeStruct(t *testing.T) {
	type School struct {
		CreatedAt time.Time
		UpdatedAt *time.Time
	}

	school := School{}
	srcType := reflect.TypeOf(school)
	srcValue := reflect.ValueOf(&school).Elem()

	// now := time.Now()
	for i := 0; i < srcType.NumField(); i++ {
		f := srcType.Field(i)
		fmt.Println(f, f.Type, f.Type.Name(), reflect.ValueOf(f.Type), reflect.ValueOf(f.Type).Interface())
		tt := reflect.New(f.Type).Elem().Interface()
		switch tt.(type) {
		case time.Time:
			fmt.Println("hey")
		case *time.Time:
			fmt.Println("hey pointer")
		}
		// timeType := reflect.TypeOf(now)
		// if f.Type.AssignableTo(timeType) {
		// 	fmt.Println("assignable", srcValue.CanSet())
		// 	srcValue.Addr().Set(reflect.ValueOf(now))
		// }
	}
	fmt.Println(srcValue)
}

func TestAliasedType(t *testing.T) {
	type StudentName string
	type Student struct {
		Name string
	}
	type StudentE struct {
		Name StudentName
	}
	src := reflect.TypeOf(Student{})
	dest := reflect.TypeOf(StudentE{})
	sm := structmapper.NewStructMapper(src, dest)

	student := Student{"wonk"}
	studentE := StudentE{}

	sm.Map(&student, &studentE)

	assert.Equal(t, student.Name, student.Name)
}
