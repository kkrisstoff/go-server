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

	expectedLenOfIDs := 10
	if len(ids) != expectedLenOfIDs {
		t.Errorf("got len of ids: %v, wanted: %v", len(ids), expectedLenOfIDs)
	}

	expectedForthID := 5
	if ids[4] != expectedForthID {
		t.Errorf("got 4th id equals to: %v, wanted: %v", ids[4], expectedForthID)
	}

}

func TestIdGenerator_Current(t *testing.T) {
	Generator.counter = 0

	currentId := Generator.Current()

	expectedCurrentID := 0
	if currentId != expectedCurrentID {
		t.Errorf("got current id equals to: %v, wanted: %v", currentId, expectedCurrentID)
	}

	Generator.Generate()
	currentIdAfterGenerate := Generator.Current()

	expectedCurrentIDAfterGenerate := 1
	if currentIdAfterGenerate != expectedCurrentIDAfterGenerate {
		t.Errorf("got current id equals to: %v, wanted: %v", currentIdAfterGenerate, expectedCurrentIDAfterGenerate)
	}
}
