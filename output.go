package main

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
