package types

type Config struct {
	Name     string `yaml:"name"`
	Database struct {
		Provider string `yaml:"provider"`
		Path     string `yaml:"path"`
	} `yaml:"database"`

	Webserver struct {
		Port string `yaml:"port"`
	} `yaml:"webserver"`
}
