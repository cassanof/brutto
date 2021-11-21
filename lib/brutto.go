package brutto

import "net/http"

type Config struct {
	Request          *http.Request
	Threads          int               // per second
	FieldMethods     map[string]method // the method for generating values for each field
	FieldDictPaths   map[string]string // the path to the dictionary for all Dictionary fields
	FieldDictPresets map[string]string // the presets for all Preset fields
}

// field value generating methods
type method int8

const (
	Preset                method = iota // Preset a static value, eg: ["admin"]
	Dictionary                          // Loads each line from the given text file as values
	NumericGenerate                     // generates numerically, eg: [1, 2, 3, ..., 01, 02, ...] (includes 0 to the front)
	NumericGenerateNoZero               // generates numerically, eg: [1, 2, 3, ..., 10, 11, ...]
)

type bruter struct {
	Config      *Config
	FieldValues map[string][]string // holds all the generated/preset values for each field
}
