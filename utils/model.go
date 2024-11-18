package utils

type Model struct {
	Service struct {
		Documents struct {
			Types string `yaml:"types"`
		} `yaml:"documents"`
	} `yaml:"service"`
}
