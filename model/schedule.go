package model

type Schedule struct {
	Autoscalings []struct {
		GroupName interface{} `yaml:"group_name"`
		Schedule  interface{} `yaml:"schedule"`
		Config    struct {
			Min     interface{} `yaml:"min"`
			Max     interface{} `yaml:"max"`
			Desired interface{} `yaml:"desired"`
		} `yaml:"config"`
	} `yaml:"autoscalings"`
}
