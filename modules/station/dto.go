package station

type Station struct {
	NID   string `json:"nid"`
	Title string `json:"title"`
}

type StationResponse struct {
	NID   string `json:"nid"`
	Title string `json:"title"`
}

type Schedule struct {
	NID                string `json:"nid"`
	StationName        string `json:"title"`
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	ScheduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}
