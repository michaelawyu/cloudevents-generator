package config

// GenConfig is the configuration for generating artifacts
type GenConfig struct {
	Input    string
	Output   string
	Language Language
	Binding  Binding
}

// GetConfig takes the names of language and binding
// and returns the configuration for generating artifacts
func GetConfig(i string, o string, lang string, bind string) GenConfig {
	language := GetLanguage(lang)

	binding := GetBinding(bind)

	return GenConfig{
		Input:    i,
		Output:   o,
		Language: language,
		Binding:  binding,
	}
}
