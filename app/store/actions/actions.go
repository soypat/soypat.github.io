package actions

// Context contains all information to completely define a page layout.
// It also has information on last page visited to enable the use of a "back" button.
// It is not an "action" item strictly speaking.
type Context struct {
	// Defines the current page.
	Page View
	// Referrer contains information on last page visited.
	Referrer *Context
	// Action can contain the executed action struct along with all
	// data contained. Can be very useful though author is on the
	// fence on whether it is good practice.
	Action interface{} // Uncomment for use.
}

type View int

const (
	ViewLanding View = iota
	ViewPage
)

// PageSelect Navigates view to new page.
type PageSelect struct {
	View    View
	PageIdx int
}

// Back button pressed. Navigate to previous page.
type Back struct {
}
