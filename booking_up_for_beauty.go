package booking

import ("fmt"
		"time")

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	//panic("Please implement the Schedule function")
	fmt.Printf("date: %s\n", date)
	t, err := time.Parse("1/02/2006 15:04:05", date)

	if err != nil {
		fmt.Println(err)
	}
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	//panic("Please implement the HasPassed function")
	t := time.Now().UTC()
	
	t2, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("now: %s date: %s\n", t, t2)

	return t2.Before(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	//panic("Please implement the IsAfternoonAppointment function")
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	
	if err != nil {
		fmt.Println(err)
	}

	h := t.Hour()
	fmt.Printf("hour: %d\n", h)

	if h >= 12 && h < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	//panic("Please implement the Description function")

	t, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		fmt.Println(err)
	}

	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	//panic("Please implement the AnniversaryDate function")

	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
