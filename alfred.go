package alfred

import (
	"encoding/json"
	. "fmt"
)

// Indent specifies the indent string used for the JSON in String() and Run(). If set to "", no indentation will be used.
var Indent = "    "

// Rerun specifies the "rerun" value.
var Rerun float64

// Variables specifies the "variables" object.
var Variables = map[string]string{}

// Items specifies the "items" array. It can be accessed and iterated directly. It can also be appended to directly or appended to using the convenience function Add(item).
var Items = []Item{}

// Icon specifies the "icon" field of Item.
type Icon struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}

// Item specifies the members of the "items" array.
type Item struct {
	UID          string `json:"uid,omitempty"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle,omitempty"`
	Arg          string `json:"arg,omitempty"`
	Icon         *Icon  `json:"icon,omitempty"`
	Autocomplete string `json:"autocomplete,omitempty"`
	Type         string `json:"type,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
	Match        string `json:"match,omitempty"`
}

type output struct {
	Rerun     float64           `json:"rerun,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
	Items     []Item            `json:"items"`
}

// Add is a convenience function for adding new Item instances to Items.
func Add(item Item) {
	Items = append(Items, item)
}

// String returns the JSON for currently populated values or the minimum required values.
func String() string {
	output := output{
		Rerun:     Rerun,
		Variables: Variables,
		Items:     Items,
	}
	var err error
	var b []byte
	if Indent == "" {
		b, err = json.Marshal(output)
	} else {
		b, err = json.MarshalIndent(output, "", Indent)
	}
	if err != nil {
		messageErr := Errorf("Error in parser. Please report this output to https://github.com/drgrib/alfred/issues: %v", err)
		panic(messageErr)
	}
	s := string(b)
	return s
}

// Run prints the result of String() to standard output for debugging or direct consumption by an Alfred script filter.
func Run() {
	Println(String())
}
