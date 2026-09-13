package main

import "fmt"

var pressureLevels = []string{
	"1000", "975", "950", "925", "900", "850", "800", "700", "600", "500",
	"400", "300", "250", "200", "150", "100", "70", "50", "30",
}

type baseVar struct {
	GoBase   string // "Temperature"
	JSONBase string // "temperature"
	GoType   string // "float32"
}

var pressureBaseVars = []baseVar{
	{"Temperature", "temperature", "float32"},
	{"RelativeHumidity", "relative_humidity", "float32"},
	{"DewPoint", "dew_point", "float32"},
	{"CloudCover", "cloud_cover", "int"},
	{"WindSpeed", "wind_speed", "float32"},
	{"WindDirection", "wind_direction", "int"},
	{"GeopotentialHeight", "geopotential_height", "float32"},
}

func pressureLevelFields() []Field {
	var fields []Field
	for _, v := range pressureBaseVars {
		for _, level := range pressureLevels {
			fields = append(fields, Field{
				GoName: fmt.Sprintf("%s%shPa", v.GoBase, level),
				JSON:   fmt.Sprintf("%s_%shPa", v.JSONBase, level),
				GoType: v.GoType,
			})
		}
	}
	return fields
}
