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

func (i *InventoryFile) GetResourceWithName(resourceName string) *Resource {
	for _, resource := range i.Resources {
		if resource.Name == resourceName {
			return &resource
		}
	}

	return nil
}

// Checklist

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
