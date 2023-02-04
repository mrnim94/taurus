package model

type Schedule struct {
	Autoscalings []struct {
		GroupName string `yaml:"group_name"`
		Profile   string `yaml:"profile"`
		Schedule  string `yaml:"schedule"`
		Config    struct {
			Min     string `yaml:"min"`
			Max     string `yaml:"max"`
			Desired string `yaml:"desired"`
		} `yaml:"config"`
	} `yaml:"autoscalings"`
}
