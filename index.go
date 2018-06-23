package index

import (
	"log"
	"regexp"
	"strings"
)

type Index struct {
	Terms     map[string]map[int]struct{}
	Documents []string
}

func (i *Index) Add(doc string) {
	str := Normalize(doc)
	fields := strings.Fields(str)
	i.Documents = append(i.Documents, doc)
	idx := len(i.Documents) - 1
	for _, f := range fields {
		if _, exist := i.Terms[f]; !exist {
			i.Terms[f] = map[int]struct{}{}
		}
		i.Terms[f][idx] = struct{}{}
	}
}

// FindOne returns a map of all documents which contain the given term.
// The term is normalized before searching the index.
func (i *Index) FindOne(term string) map[int]string {
	norm := Normalize(term)
	out := map[int]string{}
	for idx := range i.Terms[norm] {
		out[idx] = i.Documents[idx]
	}
	return out
}

// FindAll returns a map of all documents containing every given term.
func (i *Index) FindAll(terms []string) map[int]string {
	shortest := map[int]string{}
	allMaps := make([]map[int]string, len(terms))
	for idx, term := range terms {
		allMaps[idx] = i.FindOne(term)
		if len(allMaps[idx]) < len(shortest) || len(shortest) == 0 {
			shortest = allMaps[idx]
		}
	}

Shortest:
	for k := range shortest {
		for _, m := range allMaps {
			if _, exist := m[k]; !exist {
				delete(shortest, k)
				continue Shortest
			}
		}
	}
	return shortest
}

// FindAny returns a map of all documents containing any of the given terms.
func (i *Index) FindAny(terms []string) map[int]string {
	out := map[int]string{}
	for _, term := range terms {
		for k, v := range i.FindOne(term) {
			out[k] = v
		}
	}
	return out
}

func Normalize(s string) string {
	lower := strings.ToLower(s)
	reg, err := regexp.Compile("[^a-z0-9 \n]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(lower, "")
}
