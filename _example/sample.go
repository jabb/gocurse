// +build example
// Sample code for Go ncurses bindings
package main

import (
	. "github.com/jabb/gocurse/curses"
	. "github.com/jabb/gocurse/forms"
	m "github.com/jabb/gocurse/menus" // chtype imported in both forms/utils.go and menus/utils.go (and for some reason usage of C.chtype is inconsistent)
	. "github.com/jabb/gocurse/panels"
)

var screen *Window
var currentColor int = 1
var current int

func main() {
	// Initscr() initializes the terminal in curses mode.
	screen, _ = Initscr()
	// Endwin must be called when done.
	defer Endwin()

	// Initialize colors with Start_color() if needed.
	Start_color()
	// Add some color pairs. Color_pair(1) is now Yellow text with black background.
	Init_pair(1, COLOR_YELLOW, COLOR_BLACK)
	Init_pair(2, COLOR_RED, COLOR_BLACK)

	menutext := `
    Arrow keys:     Move window
    Tab:            Change focus
    Pgup/Pgdown:    Move selection (form/menu)
    Backspace/Del:  Delete characters (form)
    Home/End:       Move within field (form)
    Ctrl-d:         Exit`
	screen.Addstr(0, 0, menutext, Color_pair(currentColor))

	createPanels()
	createForm()
	defer form.Free()
	createMenu()
	defer menu.Free()
	listenKeys()
}

// Container for a Window with some additional information regarding the dimensions and position.
type AdjustableWindow struct {
	startx, starty int
	height, width  int
	w              *Window
}

func (win *AdjustableWindow) Pos() (int, int) {
	return win.starty, win.startx
}

func createWindow(height, width, starty, startx int) *AdjustableWindow {
	window, _ := Newwin(height, width, starty, startx)
	return &AdjustableWindow{w: window, height: height, width: width, starty: starty, startx: startx}
}

var windows []*AdjustableWindow
var panels []*Panel

// Test ncurses Panels library by creating three panels.
//
// The library maintains information about the order of windows, their overlapping and update the screen properly.
//    Create the windows (with Newwin()) to be attached to the panels.
//    The function NewPanel(window) is used to create panel with the window attached.
//    Call UpdatePanels() to write the panels to the virtual screen in correct visibility order. Do a DoUpdate() to show it on the screen.
func createPanels() {
	win1 := createWindow(10, 30, 21, 9)
	win2 := createWindow(9, 26, 7, 2)
	win3 := createWindow(27, 42, 10, 24)

	windows = []*AdjustableWindow{win1, win2, win3}
	panels = []*Panel{NewPanel(win1.w), NewPanel(win2.w), NewPanel(win3.w)}

	win3.w.Addstr(1, 1, octo, 0)

	// Add borders around the panels
	win1.w.Box(0, 0)
	win2.w.Box(0, 0)
	win3.w.Box(0, 0)

	panels[0].Top()
	UpdatePanels()
	DoUpdate()
}

var menu *m.Menu

// Test ncurses Menu library. Create menu with two options and fill a panel window with it.
//    Create items using new_item(). You can specify a name and description for the items.
//    Create the menu with new_menu() by specifying the items to be attached with.
//    Post the menu with menu_post() and refresh the screen.
//    Process the user requests with a loop and do necessary updates to menu with menu.Driver().
//    Unpost the menu with menu.Unpost()
//    Free the memory allocated to menu by menu.Free()
//    Free the memory allocated to the items with item.Free()
func createMenu() {
	c1 := m.NewItem("Black & Yellow", "BY")
	c2 := m.NewItem("Black & Red", "BR")

	// When using NewForm & NewMenu (passing a pointer to array of pointers in C) remember to add a nil entry to the end of your slice.
	items := []*m.Item{c1, c2, nil}
	menu, _ = m.NewMenu(items)

	menu.SetWin(windows[1].w)
	// Derwin creates a subwindow with a position that is relative to the parent window.
	dw, _ := windows[1].w.Derwin(5, 19, 4, 3)
	menu.SetSub(dw)

	windows[1].w.Addstr(6, 2, "Border color", 0)

	menu.Post()
	windows[1].w.Refresh()
}

var form *Form

// Test ncurses Form library. Create simple form with two fields.
// A form is a collection of fields; each field can be either a label(static text) or a data-entry location. The forms also library provides functions to divide forms into multiple pages. 
//    Create fields using NewField(). You can specify the height and width of the field, and its position on the form.
//    Create the forms with NewForm() by specifying the fields to be attached with.
//    Post the form with form.Post() and refresh the screen.
//    Process the user requests with a loop and do necessary updates to form with form.Driver().
//    Unpost the menu with form.Unpost()
//    Free the memory allocated to menu by form.Free()
//    Free the memory allocated to the items with field.Free()
func createForm() {
	f1, _ := NewField(1, 10, 4, 10, 0, 0)
	f2, _ := NewField(1, 10, 6, 10, 0, 0)

	// When using NewForm & NewMenu (passing a pointer to array of pointers in C) remember to add a nil entry to the end of your slice.
	fields := []*Field{f1, f2, nil}
	form, _ = NewForm(fields)

	form.SetWin(windows[0].w)
	form.Post()

	windows[0].w.Attron(Color_pair(currentColor))
	windows[0].w.Box(0, 0)
	windows[0].w.Attroff(Color_pair(currentColor))

	f1.SetBack(A_UNDERLINE)
	f1.OptsOff(O_AUTOSKIP)
	f2.SetBack(A_UNDERLINE)
	f2.OptsOff(O_AUTOSKIP)

	windows[0].w.Addstr(8, 2, "Sample form", 0)
	windows[0].w.Addstr(3, 4, "1:", 0)
	windows[0].w.Addstr(3, 6, "2:", 0)
	windows[0].w.Refresh()
}

// Switch active panel.
func nextPanel() {
	// Remove borders from previous panel
	windows[current].w.Box(0, 0)

	current++
	current %= len(panels)

	// Show/Hide cursor depending on which panel is active
	if current == 0 {
		Curs_set(1)
		form.Drive(REQ_END_LINE)
	} else {
		Curs_set(0)
	}

	// Add colored borders to new panel
	windows[current].w.Attron(Color_pair(currentColor))
	windows[current].w.Box(0, 0)
	windows[current].w.Attroff(Color_pair(currentColor))

	panels[current].Top()
}

func listenKeys() {
	var ch int

	// Suppress unnecessary echoing while taking input from the user
	Noecho()
	// Enables the reading of function keys like F1, F2, arrow keys etc
	screen.Keypad(true)

	// Since we start with the Form panel active, move cursor to the first field.
	form.Drive(REQ_END_LINE)
	UpdatePanels()
	DoUpdate()

forloop:
	for {
		ch = screen.Getch()
		switch ch {

		// -- Move windows --
		case KEY_UP:
			if windows[current].starty > 0 {
				windows[current].starty--
				// Move window to new position.
				panels[current].Move(windows[current].Pos())
			}
		case KEY_DOWN:
			if windows[current].starty < *Rows-windows[current].height {
				windows[current].starty++
				panels[current].Move(windows[current].Pos())
			}
		case KEY_RIGHT:
			if windows[current].startx < *Cols-windows[current].width {
				windows[current].startx++
				panels[current].Move(windows[current].Pos())
			}
		case KEY_LEFT:
			if windows[current].startx > 0 {
				windows[current].startx--
				panels[current].Move(windows[current].Pos())
			}

		// -- Change selection --
		case 339: // pgup
			if current == 0 {
				form.Drive(REQ_PREV_FIELD)
				form.Drive(REQ_END_LINE)
			} else if current == 1 {
				menu.Drive(m.REQ_PREV_ITEM)
				currentColor = menu.CurrentItem().Index() + 1
			}
		case 338: // pgdown
			if current == 0 {
				form.Drive(REQ_NEXT_FIELD)
				form.Drive(REQ_END_LINE)
			} else if current == 1 {
				menu.Drive(m.REQ_NEXT_ITEM)
				currentColor = menu.CurrentItem().Index() + 1
			}

		// -- Switch active window --
		case 9: // tab
			nextPanel()

		// -- Erase characters in a form --
		case 330: // delete
			if current == 0 {
				form.Drive(REQ_DEL_CHAR)
			}
		case KEY_BACKSPACE:
			if form.Drive(REQ_PREV_CHAR) {
				form.Drive(REQ_DEL_CHAR)
			}

		// -- Move inside a form --
		case KEY_HOME:
			if current == 0 {
				form.Drive(REQ_BEG_LINE)
			}
		case KEY_END:
			if current == 0 {
				form.Drive(REQ_END_LINE)
			}

		// -- Exit --
		case 4: // EOT (ctrl-d)
			break forloop

		// -- Type text into a form --
		default:
			if current == 0 {
				form.Drive(ch)
			}
		}

		// Draw panels in correct order and update screen.
		UpdatePanels()
		DoUpdate()
	}
}

// Octocat, base image from http://octodex.github.com/original/
var octo string = `
                                        
         MMMM            .MMM           
         MMMMMMMMMMMMMMMMMMMM           
         MMMMMMMMMMMMMMMMMMMM           
         MMMMMMMMMMMMMMMMMMMMM          
        MMMMMMMMMMMMMMMMMMMMMMM         
       MMMMMMMMMMMMMMMMMMMMMMMM         
       MMMMM:::::::::::::::MMMM         
       MMMM::.7.:::::::.7.::MMM         
        MM~:~777~::::::777~:MMM         
   .  MMMMM:: . :::+::: . ::MM7MM ..    
         .MM::::::7:?::::::MM.          
            MMMM~::::::MMMM             
        MM      MMMMMMM                 
         M+    MMMMMMMMM                
          MMMMMMMMM MMMM                
               MMMM MMMM                
               MMMM MMMM                
            .~~MMMM~MMMM~~.             
         ~~~~MM:~MM~MM~:MM~~~~          
        ~~~~~~====~~~====~~~~~~         
         :~~~~~====~====~~~~~~          
             :~====~====~~              
                                       `
