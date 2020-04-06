package routes

// We need to make this a "generic" Object since sometimes it might get different parameters.
type Symbol struct {
	Name              string      `json:"01. symbol"`
	Open              float64     `json:"02. open,string"`
	High              float64     `json:"03. high,string"`
	Low               float64     `json:"04. low,string"`
	Price             float64     `json:"05. price,string"`
	Volume            int32       `json:"06. volume,string"`
	LatestTradingDate string      `json:"07. latest trading day"`
	PreviousClose     float64     `json:"08. previous close,string"`
	Change            string      `json:"09. change"`
	ChangePercent     string      `json:"10. change percent"`
	XData             interface{} `json:"-"`
}
