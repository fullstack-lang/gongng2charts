package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case ChartConfiguration:
		fieldCoder := ChartConfiguration{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.ChartType = "3"
		fieldCoder.Width = 4
		fieldCoder.Heigth = 5
		return (any)(fieldCoder).(Type)
	case DataPoint:
		fieldCoder := DataPoint{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Value = 1.000000
		return (any)(fieldCoder).(Type)
	case Dataset:
		fieldCoder := Dataset{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Label = "2"
		return (any)(fieldCoder).(Type)
	case Label:
		fieldCoder := Label{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *ChartConfiguration | []*ChartConfiguration | *DataPoint | []*DataPoint | *Dataset | []*Dataset | *Label | []*Label
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *ChartConfiguration:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "ChartType"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 4 {
				return "Width"
			}
			if field == 5 {
				return "Heigth"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *DataPoint:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "Value"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Dataset:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "2" {
				return "Label"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Label:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
