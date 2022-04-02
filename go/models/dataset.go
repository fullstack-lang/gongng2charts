package models

type Dataset struct {
	Name       string
	DataPoints []*DataPoint
	Label      string
}
