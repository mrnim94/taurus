package request

type ReqAutoScalingGroup struct {
	Name    string `json:"name,omitempty"` // tags
	Profile string `json:"profile,omitempty"`
}

type ReqUpdateAutoScalingGroup struct {
	Name    string `json:"name,omitempty"` // tags
	Profile string `json:"profile,omitempty"`
	Min     string `json:"min,omitempty"`
	Max     string `json:"max,omitempty"`
	Desired string `json:"desired,omitempty"`
}
