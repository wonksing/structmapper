package structmapper_test

import (
	"reflect"
	"testing"

	"github.com/wonksing/structmapper"
)

func BenchmarkMapper(b *testing.B) {
	src := reflect.TypeOf(Person{})
	dest := reflect.TypeOf(PersonEntity{})

	mobiles := make([]*Mobile, 0)
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})

	person := Person{
		Name:      "wonk",
		MobilePtr: &Mobile{Number: "010"},
		Mobile:    Mobile{Number: "20202"},
		Mobiles:   mobiles,
	}

	for i := 0; i < b.N; i++ {
		personEntity := PersonEntity{}
		sm := structmapper.NewStructMapper(src, dest)
		sm.Map(&person, &personEntity)
	}
}

func BenchmarkMapperCached(b *testing.B) {
	mobiles := make([]*Mobile, 0)
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})

	person := Person{
		Name:      "wonk",
		MobilePtr: &Mobile{Number: "010"},
		Mobile:    Mobile{Number: "20202"},
		Mobiles:   mobiles,
	}

	for i := 0; i < b.N; i++ {
		personEntity := PersonEntity{}
		structmapper.Map(&person, &personEntity)
	}
}
