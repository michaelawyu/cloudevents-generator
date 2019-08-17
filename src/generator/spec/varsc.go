package spec

// AllowableValue is
type AllowableValue struct {
	Name    string
	Value   string
	HasMore bool
}

// VarSc is
type VarSc struct {
	Name             string
	DefaultValue     string
	DataType         string
	HasMore          bool
	BaseName         string
	AllowableValues  []AllowableValue
	IsEnum           bool
	Required         bool
	IsContainer      bool
	IsListContainer  bool
	HasValidation    bool
	MaxLength        string
	MinLength        string
	Maximum          string
	ExclusiveMaximum bool
	Minimum          string
	ExclusiveMinimum bool
	Pattern          string
	MaxItems         string
	MinItems         string
	Format           string
	Auto             *AutoFormat
	Description      string
}
