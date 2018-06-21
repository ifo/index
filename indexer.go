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
	str := RemoveNonAlphaNum(doc)
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

func RemoveNonAlphaNum(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 \n]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}
