package libguidesxml

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

type EncoreRewriter struct {
	InputBase string
	BibURL    *regexp.Regexp
	SubjURL   *regexp.Regexp
}

func NewEncoreRewriter(ib string) (e *EncoreRewriter) {
	e = &EncoreRewriter{InputBase: ib}

	EncoreBibPath := "/iii/encore/record/C__R(b\\d{7})__S"
	baseBibUrl := []string{"https?://", e.InputBase, EncoreBibPath}
	fullURL := strings.Join(baseBibUrl, "")
	e.BibURL = regexp.MustCompile(fullURL)

	EncoreSubjectPath := "/iii/encore/search/C__Sd%3A%28(.*)%29__O(.*)"
	baseSubjUrl := []string{"https?://", e.InputBase, EncoreSubjectPath}
	e.SubjURL = regexp.MustCompile(strings.Join(baseSubjUrl, ""))
	return e
}
func (e *EncoreRewriter) MatchBib(input string) string {
	m := e.BibURL.FindStringSubmatch(input)
	if len(m) > 0 {
		return m[1]
	}
	return ""
}
func (e *EncoreRewriter) MatchSubj(input string) string {
	m := e.SubjURL.FindStringSubmatch(input)
	if len(m) > 0 {
		return m[1]
	}
	return ""
}

func (e EncoreRewriter) RewriteAssetToWebPac(w *WebPacOutputter, a *Asset) {
	updated := false
	oldURL := a.URL
	urlbit := e.MatchBib(a.URL)
	if urlbit != "" {
		a.URL = w.RewriteBib(urlbit)
		updated = true
	} else {

		urlbit = e.MatchSubj(a.URL)
		if urlbit != "" {
			a.URL = w.RewriteSubj(urlbit)
			updated = true
		}
	}

	if updated {
		fmt.Printf(a.ID)
		fmt.Printf("\t")
		fmt.Printf("%s", oldURL)
		fmt.Printf("\t")
		fmt.Printf("%s", a.URL)
		fmt.Printf("\n")
	}
}

type WebPacOutputter struct {
	BibOutput  *template.Template
	SubjOutput *template.Template
	OutputBase string
}

func NewWebPacOutputter(base string) *WebPacOutputter {
	w := &WebPacOutputter{OutputBase: base}
	bTemplate := []string{"https://", w.OutputBase, "/record={{.}}~S1/"}
	w.BibOutput = template.Must(template.New("bibUrl").Parse(strings.Join(bTemplate, "")))

	funcMap := template.FuncMap{
		"clean": func(s string) string {
			s1 := strings.ToLower(s)
			s2 := strings.Replace(s1, "%20", "+", -1)
			return s2
		},
	}

	sTemplate := []string{"https://", w.OutputBase, "/search~S1?/d/d/1%2C5%2C141%2CB/exact&FF=d{{clean .}}&1%2C82%2C"}
	w.SubjOutput = template.Must(template.New("bibUrl").Funcs(funcMap).Parse(strings.Join(sTemplate, "")))
	return w
}
func (w *WebPacOutputter) RewriteBib(in string) string {
	var buf bytes.Buffer
	w.BibOutput.Execute(&buf, in)
	return buf.String()
}
func (w *WebPacOutputter) RewriteSubj(in string) string {
	var buf bytes.Buffer
	w.SubjOutput.Execute(&buf, in)
	return buf.String()
}

// rewriters:
// # - input: "http://encore.vcc.ca/iii/encore/record/C__R(b\\d{7})__S"
// #   output: "http://webpac.vcc.ca/record={{.}}~S1/"
// #   index: 1
// - input: "http://encore.vcc.ca/iii/encore/search/C__Sd%3A%28(.*)%29__O(.*)"
//   output: "http://webpac.vcc.ca/"
//   index: 1
//   replace:
// 	 "%20": "+"
//   lower: true
