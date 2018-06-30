package index

import (
	"reflect"
	"testing"
)

func TestIndexAdd(t *testing.T) {
	type Case struct {
		Index     Index
		Documents []string
		Terms     map[string]map[int]struct{}
	}

	cases := []Case{
		Case{
			Index:     Index{Terms: map[string]map[int]struct{}{}},
			Documents: []string{"This is the beginning of the test."},
			Terms: map[string]map[int]struct{}{
				"this":      map[int]struct{}{0: struct{}{}},
				"is":        map[int]struct{}{0: struct{}{}},
				"the":       map[int]struct{}{0: struct{}{}},
				"beginning": map[int]struct{}{0: struct{}{}},
				"of":        map[int]struct{}{0: struct{}{}},
				"test":      map[int]struct{}{0: struct{}{}},
			},
		},
	}

	for _, c := range cases {
		i := c.Index
		for _, doc := range c.Documents {
			i.Add(doc)
		}

		if !reflect.DeepEqual(i.Documents, c.Documents) {
			t.Errorf("Got: %+v, Expected: %q", i.Documents, c.Documents)
		}
		if !reflect.DeepEqual(i.Terms, c.Terms) {
			t.Errorf("Got: %+v, Expected: %q", i.Terms, c.Terms)
		}
	}
}

func TestIndexFindOne(t *testing.T) {
	t.Skip()
}

func TestIndexFindAll(t *testing.T) {
	t.Skip()
}

func TestIndexFindAny(t *testing.T) {
	t.Skip()
}

func TestNormalize(t *testing.T) {
	t.Skip()
}
