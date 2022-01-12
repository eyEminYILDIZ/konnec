package types

// Inventory

type InventoryFile struct {
	Resources []Resource `yaml:"resources"`
}

type Resource struct {
	Name        string `yaml:"name"`
	Ip          string `yaml:"ip"`
	Domain      string `yaml:"domain"`
	Description string `yaml:"description"`
}

// Checklist

type Checklist struct {
	Version   string          `yaml:"version"`
	Checklist []ChecklistItem `yaml:"checklist"`
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
