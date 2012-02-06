package main

import (
    "fmt"
    "time"
)

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

// Urf. Ugly one. Temporary version, to be replaced :)
// modularize, make it scale. not display all lines and display more when it should..
//
func (w Week) StringCalendar() (s string) {
    s = "╔═════════╤═════════╤═════════╤═════════╤═════════╤═════════╤═════════╗\n║"
    for i := 0; i < len(w); i++ {
	if i == today {
		s += "\x1b[1;37m" + fmt.Sprintf("%9s", time.Weekday(i).String()) + "\x1b[0m"
	} else {
		s += fmt.Sprintf("%9s", time.Weekday(i).String())
	}
        if i != len(w)-1 {
            s += "│"
        }
    }
    s += "║\n╟═════════╪═════════╪═════════╪═════════╪═════════╪═════════╪═════════╢\n║"
    for i := 0; i < 5; i++ {
        for j := 0; j < len(w); j++ {
            if len(w[j].entries) > i {
		switch w[j].entries[i].priority {
		case 0:
			s += "\x1b[1;32m"
		case 1:
		case 2:
			s += "\x1b[1;31m"
		}
                if len(w[j].entries[i].text) > 9 {
                    s += fmt.Sprintf("%9s", w[j].entries[i].text[:9])
                } else {
                    s += fmt.Sprintf("%9s", w[j].entries[i].text)
                }
                s += "\x1b[0m"
            } else {
                s += "         "
            }
            if j != len(w)-1 {
                s += "│"
            }
       }
       if i < 4 {
           s += "║\n║"
       } else {
           s += "║\n╚═════════╧═════════╧═════════╧═════════╧═════════╧═════════╧═════════╝"
       }
    }
    return
}

