package collections

type Collection struct {
	Slug         string
	Description  string
	Bootstrapper string
	Path         string
	Variables    []*Variable
}

type Variable struct {
	Key     string
	Value   string
	IsEnv   bool
	Default string
}
