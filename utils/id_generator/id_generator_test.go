package id_generator

import (
	"testing"
)

func TestIdGenerator_Generate(t *testing.T) {

	ids := make([]int, 0)

	for i := 0; i < 10; i++ {
		id := Generator.Generate()
		ids = append(ids, id)
	}

	if len(ids) != 10 {
		t.Errorf("got len of ids: %v, wanted: %v", len(ids), 10)
	}

	if ids[4] != 5 {
		t.Errorf("got 4th id equals to: %v, wanted: %v", ids[4], 10)
	}

}

func TestIdGenerator_Current(t *testing.T) {
	currentId := Generator.Current()

	if currentId != 0 {
		t.Errorf("got current id equals to: %v, wanted: %v", currentId, 0)
	}

	Generator.Generate()
	currentIdAfterGenerate := Generator.Current()

	if currentIdAfterGenerate != 1 {
		t.Errorf("got current id equals to: %v, wanted: %v", currentIdAfterGenerate, 1)
	}
}
