package venuesuite_test

import (
	"encoding/json"
	"log"
	"testing"

	venuesuite "github.com/omniboost/go-venuesuite"
)

func TestValueString(t *testing.T) {
	// t1 := venuesuite.ValueString("")
	// b, err := json.Marshal(t1)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// var t2 *venuesuite.ValueString
	// b, err = json.Marshal(t2)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// b = []byte(`
	// 	{"T": null}
	// `)
	// t3 := struct {
	// 	T *venuesuite.ValueString
	// }{}
	// err = json.Unmarshal(b, &t3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// b, err = json.Marshal(t3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	t1 := venuesuite.ValueNullString{nil}
	b, err := json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))

	var s venuesuite.ValueString
	s = ""
	t1 = venuesuite.ValueNullString{&s}
	b, err = json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
}
