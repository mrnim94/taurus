package request

type ReqAutoScalingGroup struct {
	Name    string `json:"name,omitempty"` // tags
	Profile string `json:"profile,omitempty"`
}
