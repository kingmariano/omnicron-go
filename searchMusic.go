package omnicron

type MusicSearchRequest struct {
	Song  string `json:"song"`
	Limit int    `json:"limit,omitempty"`
	Proxy string `json:"proxy,omitempty"`
}
