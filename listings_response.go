package client

type Listings struct {
	Results []Listing `json:"results"`
}

type Listing struct {
	Type  string `json:"type"`
	MlNum string `json:"ml_number"`
	//Price []float64 `json:"price"`
}
