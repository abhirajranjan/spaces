package constants

// TODO: write schemas of topics to decode into
type Community struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}
