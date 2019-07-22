package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

const (
	tagName    = "env"
	tagDefault = "default"
)

type varInfo struct {
	Name  string
	Alt   string
	Key   string
	Field reflect.Value
	Tags  reflect.StructTag
}

// Load the environment variables into the provided struct
func Load(t interface{}) {
	if err := process(t); err != nil {
		log.Panicf("config: unable to load config for %T: %s", t, err)
	}
}

func process(t interface{}) error {
	infos, err := gatherInfo(t)
	fmt.Println(infos)
	return err
	// for i := 0; i < st.NumField(); i++ {
	// 	field := st.Field(i)
	// 	tag := field.Tag.Get(tagName)
	// 	dft := field.Tag.Get(tagDefault)
	// 	val := os.Getenv(tag)
	// 	if val == "" {
	// 		val = dft
	// 	}
	// }
}

func gatherInfo(t interface{}) ([]varInfo, error) {
	s := reflect.ValueOf(t)
	if s.Kind() != reflect.Ptr {
		return nil, errors.New("specification must be a struct pointer")
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, errors.New("specification must be a struct pointer")
	}

	typeOfS := s.Type()
	infos := make([]varInfo, 0, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fType := typeOfS.Field(i)
		if !f.CanSet() {
			continue
		}
		for !f.IsNil() {
			f.Set(reflect.New(f.Type().Elem()))
			f = f.Elem()
		}
		info := varInfo{
			Name:  fType.Name,
			Field: f,
			Tags:  fType.Tag,
			Alt:   fType.Tag.Get(tagName),
		}
		info.Key = info.Name
		if info.Alt != "" {
			info.Key = info.Alt
		}
		infos = append(infos, info)
	}
	return infos, nil
}
