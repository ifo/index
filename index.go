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

func Normalize(s string) string {
	lower := strings.ToLower(s)
	reg, err := regexp.Compile("[^a-z0-9 \n]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(lower, "")
}
