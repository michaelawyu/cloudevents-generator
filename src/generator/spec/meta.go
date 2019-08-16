package spec

// Metadata is
type Metadata struct {
	PackageName string `json:"packageName" yaml:"packageName"`
	Version     string `json:"version" yaml:"version"`
	Description string `json:"description" yaml:"description"`
	Contact     string `json:"contact" yaml:"contact"`
	URL         string `json:"url" yaml:"url"`
}
