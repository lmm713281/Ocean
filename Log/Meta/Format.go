package Meta

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (entry Entry) Format() (result string) {

	lenSender := len(entry.Sender)
	if lenSender > 40 {
		lenSender = 40
	}

	lenMessageName := len(entry.MessageName)
	if lenMessageName > 26 {
		lenMessageName = 26
	}

	lenProject := len(entry.Project)
	if lenProject > 10 {
		lenProject = 10
	}

	messageDescription := entry.MessageDescription
	messageDescription = strings.Replace(messageDescription, "\n", ` `, -1)
	messageDescription = strings.Replace(messageDescription, "\t", ` `, -1)
	messageDescription = strings.Replace(messageDescription, "\r", ` `, -1)

	result = fmt.Sprintf(` [■] P:%10s [■] %s [■] %10s [■] %11s [■] %10s [■] %10s [■] sender: %-40s [■] name: %-26s [■] %s [■]`, entry.Project[:lenProject], formatTime(entry.Time), entry.Category.Format(), entry.Level.Format(), entry.Severity.Format(), entry.Impact.Format(), entry.Sender[:lenSender], entry.MessageName[:lenMessageName], messageDescription)

	for _, param := range entry.Parameters {

		paramText := param
		paramText = strings.Replace(paramText, "\n", ` `, -1)
		paramText = strings.Replace(paramText, "\t", ` `, -1)
		paramText = strings.Replace(paramText, "\r", ` `, -1)

		result = fmt.Sprintf(`%s %s [■]`, result, paramText)
	}

	return
}

func formatTime(t1 time.Time) (result string) {
	var year int = t1.Year()
	var month int = int(t1.Month())
	var day int = int(t1.Day())
	var minutes int = int(t1.Minute())
	var hours int = int(t1.Hour())
	var seconds int = int(t1.Second())
	var milliseconds int = int(float64(t1.Nanosecond()) / 1000000.0)

	var monthText, dayText, minutesText, hoursText, secondsText, millisecondsText string

	if month >= 1 && month <= 9 {
		monthText = fmt.Sprintf(`0%d`, month)
	} else {
		monthText = strconv.Itoa(month)
	}

	if day >= 1 && day <= 9 {
		dayText = fmt.Sprintf(`0%d`, day)
	} else {
		dayText = strconv.Itoa(day)
	}

	if minutes >= 0 && minutes <= 9 {
		minutesText = fmt.Sprintf(`0%d`, minutes)
	} else {
		minutesText = strconv.Itoa(minutes)
	}

	if hours >= 0 && hours <= 9 {
		hoursText = fmt.Sprintf(`0%d`, hours)
	} else {
		hoursText = strconv.Itoa(hours)
	}

	if seconds >= 0 && seconds <= 9 {
		secondsText = fmt.Sprintf(`0%d`, seconds)
	} else {
		secondsText = strconv.Itoa(seconds)
	}

	if milliseconds >= 0 && milliseconds <= 9 {
		millisecondsText = fmt.Sprintf(`00%d`, milliseconds)
	} else if milliseconds >= 10 && milliseconds <= 99 {
		millisecondsText = fmt.Sprintf(`0%d`, milliseconds)
	} else {
		millisecondsText = strconv.Itoa(milliseconds)
	}

	result = fmt.Sprintf(`%d%s%s %s%s%s.%s`, year, monthText, dayText, hoursText, minutesText, secondsText, millisecondsText)
	return
}
