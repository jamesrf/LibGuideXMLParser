package libguidesxml

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
)

type Rewriter struct {
	Regex	*regexp.Regexp
	Map      map[string]string
}

func (r *Rewriter) Rewrite(s string) string {
	matches := r.Regex.FindStringSubmatch(s)
	if matches != nil {
		if len(matches) > 1 {
			m := matches[1]
			if r.Map[m] != "" {
				return r.Map[m]
			}
		}
	}
	return ""
}

func NewRewriter(regex string, mapfile string, trim bool) *Rewriter {
	re := regexp.MustCompile(regex)

	finalMap := make(map[string]string)
	if mapfile != "" {
		mapFile, err := os.Open(mapfile)
		if err != nil {
			panic(err)
		}
		reader := csv.NewReader(mapFile)
		reader.Comma = '\t'

		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			oldId := line[0]
			newId := line[1]
			if trim {
				oldId = oldId[1:len(oldId)-1]
			}
			finalMap[ oldId ] = newId
		}
	}


	return &Rewriter{
		Regex: re,
		Map: finalMap,
	}
}

