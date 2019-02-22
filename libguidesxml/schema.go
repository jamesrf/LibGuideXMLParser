package libguidesxml


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