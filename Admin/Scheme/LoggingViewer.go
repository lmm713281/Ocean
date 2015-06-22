package Scheme

import (
	"html/template"
)

// The type for the web logger viewer template
type LoggingViewer struct {
	Title              string
	SetLiveView        bool
	CurrentLevel       string
	CurrentTimeRange   string
	CurrentCategory    string
	CurrentImpact      string
	CurrentSeverity    string
	CurrentMessageName string
	CurrentSender      string
	CurrentPage        string
	MessageNames       []MessageNames
	Sender             []Sender
	Events             []LogEvent
}

// The type for the message names is necessary to be able to define a function on it.
type MessageNames string

// The type for the senders is necessary to be able to define a function on it.
type Sender string

// This function is used from the template to mark selected values. This is for the type MessageNames.
func (lv MessageNames) IsSelected(field MessageNames, currentValue string) string {
	if string(field) == currentValue {
		return ` selected`
	} else {
		return ``
	}
}

// This function is necessary to mark the HTML attribute as safe. Only then it is possible
// to change plain HTML code.
func (lv MessageNames) Safe(element string) template.HTMLAttr {
	return template.HTMLAttr(element)
}

// This function is used from the template to mark selected values. This is for the type Sender.
func (lv Sender) IsSelected(field Sender, currentValue string) string {
	if string(field) == currentValue {
		return ` selected`
	} else {
		return ``
	}
}

// This function is necessary to mark the HTML attribute as safe. Only then it is possible
// to change plain HTML code.
func (lv Sender) Safe(element string) template.HTMLAttr {
	return template.HTMLAttr(element)
}

// This function is used from the template to mark selected values.
func (lv LoggingViewer) IsSelected(field, currentValue string) string {
	if field == currentValue {
		return ` selected`
	} else {
		return ``
	}
}

// This function is necessary to mark the HTML attribute as safe. Only then it is possible
// to change plain HTML code.
func (lv LoggingViewer) Safe(element string) template.HTMLAttr {
	return template.HTMLAttr(element)
}
