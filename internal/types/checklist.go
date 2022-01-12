package types

type Checklist struct {
	Version string          `yaml:"version"`
	Items   []ChecklistItem `yaml:"items"`
}

type ChecklistItem struct {
	ResourceName string               `yaml:"resource_name"`
	Conditions   []ChecklistCondition `yaml:"conditions"`
}

type ChecklistCondition struct {
	Type    string `yaml:"type"`
	Value   string `yaml:"value"`
	Message string `yaml:"message"`
}
