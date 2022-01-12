package types

type InventoryFile struct {
	Resources []Resource `yaml:"resources"`
}

type Resource struct {
	Name        string `yaml:"name"`
	Ip          string `yaml:"ip"`
	Domain      string `yaml:"domain"`
	Description string `yaml:"description"`
}

func (i *InventoryFile) GetResourceByName(resourceName string) *Resource {
	for _, resource := range i.Resources {
		if resource.Name == resourceName {
			return &resource
		}
	}

	return nil
}
