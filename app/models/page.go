package models

import (
	"strings"
	"unicode/utf8"
)

type Page struct {
	// on visit
	LinkTitle  string `json:"link_title"`
	Href       string `json:"href"`
	IsTopLevel bool   `json:"top_level"`
	// After fulfillment
	Title           string    `json:"title"`
	MainContentHTML string    `json:"content"`
	Cabinets        []Cabinet `json:"cabinets"`
}

type Cabinet struct {
	Title string `json:"title"`
	Files []File `json:"files"`
}

type File struct {
	// Table column data.
	Data [8]string `json:"data"`
	// Links if any in order of appearance.
	Links []string `json:"links,omitempty"`
}

type Field1Data rune

const (
	// What is this?
	Field1c Field1Data = 'ċ'
)

func (f File) field1() Field1Data {
	a, _ := utf8.DecodeRuneInString(strings.TrimSpace(f.Data[1]))
	return Field1Data(a)
}

func (f File) Title() string {
	title := strings.TrimSuffix(f.Data[2], "Descargar")
	title = strings.TrimSpace(title)
	title = strings.TrimSuffix(title, "Ver")
	return title
}

func (f File) Description() string {
	return f.Data[3]
}
