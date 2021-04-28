package p

type Changelog struct {
	LogId        int    `json:"log_id"`
	Text         string `json:"text"`
	TSCreated    int    `json:"ts_created"`
	TSCreatedRaw int    `json:"string"`
	Type         string `json:"type"`
}

type Changelogs struct {
	Limiter int          `json:"limiter"`
	State   bool         `json:"state"`
	Changes []*Changelog `json:"changes"`
}
