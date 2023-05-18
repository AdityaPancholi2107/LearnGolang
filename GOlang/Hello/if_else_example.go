package main

import "fmt"

// The existing program is used to restrict access to a resos based on day of the week and user role. currently, the pr allows any user to access the resource. Implement the Fi
// needed to grant resource access using any combination and 'else OR if or elseif'
// Requirements:
// Use the accessGranted() and accessDenied() functions informational messages
// Access at any time: Admin, Manager
// Access weekends: Contractor
//Access weekdays: Member
//Access Mondays, Wednesdays, and Fridays: Guest
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// user roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func weekday(day int) bool {
	return day <= 4
}
func accessGranted() {
	fmt.Println("Granted")

}
func accessDenied() {
	fmt.Println("Denied")
}
func main() {
	today, role := Monday, Guest
	// Access at any time: Admin, Manager
	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		//Access weekends= Contractor
		accessGranted()
	} else if role == Member && weekday(today) {
		//Access weekdays: Member
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		// Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
	} else {
		accessDenied()
	}
}

