package station

type Station struct {
	NID   string `json:"nid"`
	Title string `json:"title"`
}

type StationResponse struct {
	NID   string `json:"nid"`
	Title string `json:"title"`
}
