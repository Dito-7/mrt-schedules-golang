package station

type Station struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StationResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
