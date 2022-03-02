package configurator

type collection struct {
	Description  string     `yaml:"description"`
	Bootstrapper string     `yaml:"bootstrapper"`
	Variables    []variable `yaml:"vars"`
}

type variable struct {
	Key     string `yaml:"key"`
	Default string `yaml:"default"`
	IsEnv   bool   `yaml:"env"`
}
