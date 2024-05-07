package readtoml

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Config represents the structure of the TOML configuration file.
type Config struct {
	StringValue  string     `toml:"string_value"`    // StringValue represents a string value.
	IntValue     int        `toml:"int_value"`       // IntValue represents an integer value.
	BooleanValue bool       `toml:"boolean_value"`   // BooleanValue represents a boolean value.
	FloatValue   float64    `toml:"float_value"`     // FloatValue represents a floating-point value.
	StringArray  []string   `toml:"string_array"`    // StringArray represents an array of strings.
	IntArray     []int      `toml:"int_array"`       // IntArray represents an array of integers.
	FloatArray   []float64  `toml:"float_array"`     // FloatArray represents an array of floating-point values.
	BoolArray    []bool     `toml:"bool_array"`      // BoolArray represents an array of boolean values.
	TwoDIntArray [][]int    `toml:"two_d_int_array"` // TwoDIntArray represents a 2D array of integers.
	Structure    MyStruct   `toml:"structure"`       // Structure represents a structure.
	StructArray  []MyStruct `toml:"struct_array"`    // StructArray represents an array of structures.
}

// MyStruct represents a custom structure.
type MyStruct struct {
	Name string `toml:"name"` // Name represents the name field of MyStruct.
	Age  int    `toml:"age"`  // Age represents the age field of MyStruct.
}

// GetTomlValues reads values from a TOML file and prints them.
func GetTomlValues() {
	var config Config
	if _, err := toml.DecodeFile("./toml/config.toml", &config); err != nil {
		fmt.Println("Error decoding TOML:", err)
		return
	}

	fmt.Println("String Value:", config.StringValue)
	fmt.Println("Integer Value:", config.IntValue)
	fmt.Println("Boolean Value:", config.BooleanValue)
	fmt.Println("Float Value:", config.FloatValue)
	fmt.Println("String Array:", config.StringArray)
	fmt.Println("Integer Array:", config.IntArray)
	fmt.Println("Float Array:", config.FloatArray)
	fmt.Println("Boolean Array:", config.BoolArray)
	fmt.Println("2D Integer Array:", config.TwoDIntArray)
	fmt.Println("Structure Name:", config.Structure.Name)
	fmt.Println("Structure Age:", config.Structure.Age)
	fmt.Println("Array of Structures:")
	for i, s := range config.StructArray {
		fmt.Printf("  Structure %d - Name: %s, Age: %d\n", i+1, s.Name, s.Age)
	}
}
