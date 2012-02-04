// Seven Days (7d) is a simple calendar software.
//
// Main point is that quite often I find myself being only interested in the stuff that is 
// happening during the next seven days.
// I also tend to have weekly entries and I dislike complex syntax or having to click through
// pile of menus before I can add simple entry.
//
// 7d reads a simple plaintext file that has calendar entries formatted in simple syntax:
//       each|every <weekday> <description>
//   or
//       <date> <description>
//
//  Everything else is discarded so you can store whatever else you want in the same file.
//
//  <date> can be formatted in three ways:
//      25.3.2012
//      2012-03-25
//      25.3    <- Uses the _current_ year.
//
// One thing to note: If current day is Wednesday for example,
// the displayed Monday and Tuesday will be the _next_ weeks, not the previous days.
// You will never see past dates. (Unless you offset the week that is)
//
//author: Erno Hopearuoho ( https://github.com/errnoh )
package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var filename *string = flag.String("file", "", "file to read the data from")

// A single calendar entry
//
// text: the description
// priority: 0 if entry is weekly, 2 if entry is marked within 2 days
type entry struct {
	text     string
	priority int
}

func (e entry) String() string {
	return e.text
}

// A single day. :)
//
// number: int identifier of the weekday, starting from Sunday being 0.
// entries: slice of entries
type Day struct {
	number  int
	entries []*entry
}

func (d Day) String() (s string) {
	if d.number == today {
		s += "\x1b[1;37m" + time.Weekday(d.number%7).String() + "\x1b[0m:\n"
	} else {
		s += time.Weekday(d.number%7).String() + ":\n"
	}
	for _, e := range d.entries {
		switch e.priority {
		case 0:
			s += "\x1b[1;32m"
		case 1:
		case 2:
			s += "\x1b[1;31m"
		}
		s += " " + e.text + "\x1b[0m\n"
	}
	return
}

// Week contains days. Usually seven.
type Week []Day

var week Week = make([]Day, 7)

var today int

// Text output, displays a week, starting from Sunday.
func (w Week) String() (s string) {
	for _, d := range week {
		s += (d.String())
	}
	return
}

// Text output, displays a week forward, starting from the current day.
func (w Week) StringFromToday() (s string) {
	for i := 0; i < len(week); i++ {
		s += week[(today+i)%7].String()
	}
	return
}

func main() {
	flag.Parse()
	now := time.Now()
	today = int(now.Weekday())
	initWeek()
	parseData(*filename)
	//	println(week.String())
	println(week.StringFromToday())
}

func chkerr(err error) {
	if err != nil {
		println("<ERROR> " + err.Error())
		os.Exit(1)
	}
}

func initWeek() {
	for i := 0; i < len(week); i++ {
		d := new(Day)
		d.entries = make([]*entry, 0, 5)
		d.number = i
		week[i] = *d
	}
}

/*
*   Reads a plaintext file with calendar entries.
*   Entries should be formatted as either:
*       each|every <weekday> <description>
*   or
*       <date> <description>
 */
func parseData(target string) {
	var day int

	contents, err := ioutil.ReadFile(target)
	chkerr(err)
	rows := strings.Split(string(contents), "\n")
	for _, r := range rows {
		row := strings.Split(r, " ")
		if len(row) < 2 {
			continue
		}
		if strings.ToLower(row[0]) == "each" || strings.ToLower(row[0]) == "every" {
			day = getDay(row[1])
			if day == 7 {
				continue
			}
			addEntry(row[2:], day, 0)
		} else {
			date, err := parseTime(row[0])
			if err != nil {
				continue
			}

			since := time.Since(date)
			// no need for entries in the past
			if since > 0 {
				continue
			}

			untildate := (time.Hour * 24 * 7) + since
			if untildate > (time.Hour * 24 * 4) {
				addEntry(row[1:], int(date.Weekday()), 2)
			} else if untildate > 0 {
				addEntry(row[1:], int(date.Weekday()), 1)
			}
		}
	}
}

func addEntry(text []string, day int, priority int) {
	s := ""
	for _, word := range text {
		s += word + " "
	}
	week[day].entries = append(week[day].entries, &entry{text: s, priority: priority})
}

/*
*   Parses the <date>, supported formatting:
*   25.3.2012
*   2012-03-25
*   25.3    <- Uses the _current_ year.
 */
func parseTime(s string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", s)
	if err == nil {
		return
	}
	date, err = time.Parse("2.1.2006", s)
	if err == nil {
		return
	}
	date, err = time.Parse("2.1", s)
	if err == nil {
		date = date.AddDate(2012, 0, 0)
	}
	return
}

/*
*   Returns the <weekday> as integer, Sunday being 0
 */
func getDay(day string) int {
	day = strings.ToLower(day)
	switch day {
	case "monday":
		return 1
	case "tuesday":
		return 2
	case "wednesday":
		return 3
	case "thursday":
		return 4
	case "friday":
		return 5
	case "saturday":
		return 6
	case "sunday":
		return 0
	}
	return 7
}
