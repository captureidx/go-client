package client

type Listings struct {
	Total float64 `json:"total"`
	Start float64 `json:"start"`
	Limit float64 `json:"limit"`
	//Query    []string  `json:"query"`
	Sort     bool      `json:"sort"`
	Links    Links     `json:"links"`
	Results  []Listing `json:"results"`
	AddedAt  Date      `json:"added_at"`
	Source   string    `json:"source"`
	SelfLink string    `json:"self_link"`
}
type Listing struct {
	Type         string            `json:"type"`
	Status       Status            `json:"status"`
	MlNum        string            `json:"ml_number"`
	Price        Prices            `json:"price"`
	Description  string            `json:"description"`
	Location     Location          `json:"location"`
	Building     Building          `json:"building"`
	Surroundings Surroundings      `json:"surroundings"`
	LastUpdated  Date              `json:"last_updated"`
	Office       Office            `json:"office"`
	Agents       []Agents          `json:"agents"`
	Media        Media             `json:"media"`
	Key          string            `json:"key"`
	Src          map[string]string `json:"src"`
}

type Links struct {
	Self string      `json:"self"`
	Prev interface{} `json:"prev"`
	Next interface{} `json:"next"`
}

type Status struct {
	Current string `json:"current"`
	Changed Date   `json:"changed"`
}

type Prices struct {
	Amount        float64 `json:"amount"`
	Original      float64 `json:"original"`
	LastChangedOn Date    `json:"last_changed_on"`
}

type Location struct {
	City        string      `json:"city"`
	State       string      `json:"state"`
	PostalCode  string      `json:"postal_code"`
	County      string      `json:"county"`
	Subdivision string      `json:"subdivision"`
	Street      Street      `json:"street"`
	Directions  string      `json:"directions"`
	Coordinates Coordinates `json:"coordinates"`
}

type Street struct {
	Number float64 `json:"number"`
	Name   string  `json:"name"`
	Unit   string  `json:"unit"`
	Full   string  `json:"full"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	GeoJson   gj      `json:"geo_json"`
}

//Geo Json
type gj struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Building struct {
	Style        []string   `json:"style"`
	Construction []string   `json:"construction"`
	SquareFeet   sf         `json:"square_feet"`
	Stories      string     `json:"stories"`
	BuiltIn      float64    `json:"built_in"`
	Rooms        Rooms      `json:"rooms"`
	Pool         []string   `json:"pool"`
	Features     []string   `json:"features"`
	Appliances   []string   `json:"appliances"`
	Hvac         []string   `json:"hvac"`
	Fireplaces   Fireplaces `json:"fireplaces"`
}

//Square Feet
type sf struct {
	Size   float64 `json:"size"`
	Source string  `json:"source"`
}

type Rooms struct {
	Bed         Bed      `json:"bed"`
	Baths       Baths    `json:"baths"`
	Kitchen     []string `json:"kitchen"`
	Description []string `json:"description"`
	Basement    []string `json:"basement"`
}

type Bed struct {
	Description []string `json:"description"`
	Count       float64  `json:"count"`
}

type Baths struct {
	Master []string `json:"master"`
	Full   float64  `json:"full"`
	Half   float64  `json:"half"`
}

type Fireplaces struct {
	Count       float64  `json:"count"`
	Description []string `json:"description"`
	Exists      bool     `json:"exists"`
}

type Surroundings struct {
	Water     Water    `json:"water"`
	Sewer     []string `json:"sewer"`
	Services  []string `json:"services"`
	Utilities []string `json:"utilities"`
	Parking   []string `json:"parking"`
	Features  []string `json:"features"`
	Land      Land     `json:"land"`
	School    School   `json:"school"`
}

type Water struct {
	Type []string `json:"type"`
}

type Land struct {
	Size        Size     `json:"size"`
	Description []string `json:"description"`
}

type Size struct {
	Acres float64 `json:"acres"`
}

type School struct {
	Elementary string `json:"elementary"`
	Junior     string `json:"junior"`
	High       string `json:"high"`
}

type Office struct {
	Name        string `json:"name"`
	BrokerId    string `json:"broker_id"`
	BrokerPhone string `json:"broker_phone"`
}

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Phone struct {
	Preferred string `json:"preferred"`
	Mobile    string `json:"mobile"`
	Fax       string `json:"fax"`
}

type Media struct {
	VirtualTours []string `json:"virtual_tours"`
	Photos       []Photos `json:"photos"`
}

type Photos struct {
	Position float64 `json:"position"`
	Urls     Urls    `json:"urls"`
	S3Key    string  `json:"S3Key"`
}

type Urls struct {
	Original      string `json:"original"`
	TwoOneFiveX   string `json:"215x"`
	FourFourZeroX string `json:"440x"`
}

type Date struct {
	Date string `json:"$date"`
}
