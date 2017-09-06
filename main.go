package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"unicode/utf8"
	"text/template"
	"bytes"
	"github.com/spf13/viper"
)

type LibGuides struct {
	Customer Customer  `xml:"customer"`
	Site     Site      `xml:"site"`
	Accounts []Account `xml:"accounts>account"`
	Subjects []Subject `xml:"subjects>subject"`
	Tags     []Tag     `xml:"tags>tag"`
	Vendors  []Vendor  `xml:"vendors>vendor"`
	Guides   []Guide   `xml:"guides>guide"`
}

type Customer struct {
	ID       string `xml:"id"`
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	URL      string `xml:"url"`
	City     string `xml:"city"`
	State    string `xml:"state"`
	Country  string `xml:"country"`
	TimeZone string `xml:"time_zone"`
	Created  string `xml:"created"`
	Updated  string `xml:"updated"`
}
type Site struct {
	ID      string `xml:"id"`
	Name    string `xml:"name"`
	Domain  string `xml:"domain"`
	Admin   string `xml:"admin"`
	Created string `xml:"created"`
	Updated string `xml:"updated"`
}

type Account struct {
	ID        string `xml:"id"`
	Email     string `xml:"email"`
	FirstName string `xml:"first_name"`
	LastName  string `xml:"last_name"`
	Nickname  string `xml:"nickname"`
	Signature string `xml:"signature"`
	Image     string `xml:"image"`
	Address   string `xml:"address"`
	Phone     string `xml:"phone"`
	Skype     string `xml:"skype"`
	Website   string `xml:"website"`
	Created   string `xml:"created"`
	Updated   string `xml:"updated"`
}

type Subject struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type Tag struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
}

type Vendor struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
}

type Group struct{}

type Guide struct {
	ID          string    `xml:"id"`
	Type        string    `xml:"type"`
	Name        string    `xml:"name"`
	Description string    `xml:"description"`
	URL         string    `xml:"url"`
	Owner       Account   `xml:"owner"`
	Group       Group     `xml:"group"`
	Redirect    string    `xml:"redirect"`
	Status      string    `xml:"status"`
	Password    string    `xml:"password"`
	Created     string    `xml:"created"`
	Updated     string    `xml:"updated"`
	Published   string    `xml:"published"`
	Subjects    []Subject `xml:"subjects>subject"`
	Tags        []Tag     `xml:"tags>tag"`
	Pages       []Page    `xml:"pages>page"`
}

type Page struct {
	ID           string `xml:"id"`
	Name         string `xml:"name"`
	Description  string `xml:"description"`
	URL          string `xml:"url"`
	Redirect     string `xml:"redirect"`
	SourcePageID string `xml:"source_page_id"`
	ParentPageID string `xml:"parent_page_id"`
	Created      string `xml:"created"`
	Updated      string `xml:"updated"`
	Boxes        []Box  `xml:"boxes>box"`
}

type Box struct {
	ID       string  `xml:"id"`
	Type     string  `xml:"type"`
	Name     string  `xml:"name"`
	MapID    string  `xml:"map_id"`
	Column   string  `xml:"column"`
	Position string  `xml:"position"`
	Hidden   string  `xml:"hidden"`
	Created  string  `xml:"created"`
	Updated  string  `xml:"updated"`
	Assets   []Asset `xml:"assets>asset"`
}

type Asset struct {
	ID          string  `xml:"id"`
	Type        string  `xml:"type"`
	Name        string  `xml:"name"`
	Description string  `xml:"description"`
	URL         string  `xml:"url"`
	Owner       Account `xml:"owner"`
	MapID       string  `xml:"map_id"`
	Position    string  `xml:"position"`
	Created     string  `xml:"created"`
	Updated     string  `xml:"updated"`
	MoreInfo    string  `xml:"more_info"`
	EnableProxy string  `xml:"enable_proxy"`
}

type BadCharCleaner struct {
	buffer *bufio.Reader
}

func (c BadCharCleaner) Read(b []byte) (n int, err error) {
	for {
		var r rune
		var s int
		r, s, err = c.buffer.ReadRune()
		if err != nil {
			return
		}
		if (r == '\u0001' || r == '\u0014' || r == '\u0019') && s == 1 {
			continue
		} else if n+s < len(b) {
			utf8.EncodeRune(b[n:], r)
			n += s
		} else {
			c.buffer.UnreadRune()
			break
		}
	}
	return
}



type URLRewriter struct {
	Matcher		*regexp.Regexp
	Output		*template.Template
	matchIndex     int
}
func NewURLRewriter(in string, out string, idx int) *URLRewriter{
	r := regexp.MustCompile(in)
	t :=  template.Must(template.New("name").Parse(out))
	u := &URLRewriter{Matcher: r, Output:t, matchIndex:idx}
	return u
}

func (u *URLRewriter) RewriteURL (a *Asset){
	m := u.Matcher.FindStringSubmatch(a.URL)
	if len(m) > 0 {
		buf := new(bytes.Buffer)
		s := m[ u.matchIndex ]
		u.Output.Execute(buf,s)
		a.URL = buf.String()
	}
	return
}



func main() {
	xmlFile, err := os.Open("export.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	cr := BadCharCleaner{bufio.NewReader(xmlFile)}
	dec := xml.NewDecoder(cr)

	var lg LibGuides
	fmt.Printf("Reading XML file: ...")
	err = dec.Decode(&lg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Done! (%d bytes)\n\n", dec.InputOffset())

	encoreBibURL := "http://encore.vcc.ca/iii/encore/record/C__R(b\\d{7})__S"
	webpacOutputURL :=  "http://webpac.vcc.ca/record={{.}}~S1/"
	ur := NewURLRewriter(encoreBibURL, webpacOutputURL, 1)
	
	for _, guide := range lg.Guides {
		fmt.Printf("GUIDE:%s\n", guide.Name)
		for _, page := range guide.Pages {
			for _, box := range page.Boxes {
				for _, asset := range box.Assets {
					oldURL := asset.URL
					ur.RewriteURL(&asset)
					if oldURL == asset.URL {
						fmt.Printf(asset.URL)
						fmt.Printf("\n")
					}

				}

			}
		}

	}
}
