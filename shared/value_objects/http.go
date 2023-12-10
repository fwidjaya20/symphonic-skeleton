package valueobjects

type Reject struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Metadata any    `json:"metadata,omitempty"`
}

type Resolve struct {
	Data     any    `json:"data"`
	Message  string `json:"message,omitempty"`
	Metadata any    `json:"metadata,omitempty"`
}
