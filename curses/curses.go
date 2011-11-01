package curses

// struct ldat{};
// struct _win_st{};
// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
import "C"

import (
	"fmt"
	"os"
    "unsafe"
) 

type void unsafe.Pointer
type Window C.WINDOW

type CursesError struct {
	message string
}

func (ce CursesError) String() string {
	return ce.message
}

// Cursor options.
const (
	CURS_HIDE = iota
	CURS_NORM
	CURS_HIGH
)

// Pointers to the values in curses, which may change values.
var Cols *int = nil;
var Rows *int = nil;

var Colors *int = nil;
var ColorPairs *int = nil;

var Tabsize *int = nil;

// The window returned from C.initscr()
var Stdwin *Window = nil;

// Initializes gocurses
func init() {
	Cols = (*int)(void(&C.COLS));
	Rows = (*int)(void(&C.LINES));
	
	Colors = (*int)(void(&C.COLORS));
	ColorPairs = (*int)(void(&C.COLOR_PAIRS));
	
	Tabsize = (*int)(void(&C.TABSIZE));
}

func Initscr() (*Window, os.Error) {
	Stdwin = (*Window)(C.initscr());
	
	if Stdwin == nil {
		return nil, CursesError{"Initscr failed"};
	}
	
	return Stdwin, nil;
}

func Newwin(rows int, cols int, starty int, startx int) (*Window, os.Error) {
	nw := (*Window)(C.newwin(C.int(rows), C.int(cols), C.int(starty), C.int(startx)));

	if nw == nil {
		return nil, CursesError{"Failed to create window"};
	}

	return nw, nil;
}

func (win *Window) Del() os.Error {
	if int(C.delwin((*C.WINDOW)(win))) == 0 {
		return CursesError{"delete failed"};
	}
	return nil
}

func (win *Window) Subwin(rows int, cols int, starty int, startx int) (*Window, os.Error)  {
	sw := (*Window)(C.subwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)));

	if sw == nil {
		return nil, CursesError{"Failed to create window"};
	}

	return sw, nil;
}

func (win *Window) Derwin(rows int, cols int, starty int, startx int) (*Window, os.Error)  {
	dw := (*Window)(C.derwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)));

	if dw == nil {
		return nil, CursesError{"Failed to create window"};
	}

	return dw, nil;
}

func Start_color() os.Error {
	if int(C.has_colors()) == 0 {
		return CursesError{"terminal does not support color"};
	}
	C.start_color();
	
	return nil;
}

func Init_pair(pair int, fg int, bg int) os.Error {
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == 0 {
		return CursesError{"Init_pair failed"};
	}
	return nil;
}

func Color_pair(pair int) int32 {
	return int32(C.COLOR_PAIR(C.int(pair)));
}

func Noecho() os.Error {
	if int(C.noecho()) == 0 {
		return CursesError{"Noecho failed"};
	}
	return nil;
}

func DoUpdate() os.Error {
		if int(C.doupdate()) == 0 {
		return CursesError{"Doupdate failed"};
	}
	return nil;
}

func Echo() os.Error {
	if int(C.noecho()) == 0 {
		return CursesError{"Echo failed"};
	}
	return nil;
}

func Curs_set(c int) os.Error {
	if C.curs_set(C.int(c)) == 0 {
		return CursesError{"Curs_set failed"};
	}
	return nil;
}

func Nocbreak() os.Error {
	if C.nocbreak() == 0 {
		return CursesError{"Nocbreak failed"};
	}
	return nil;
}

func Cbreak() os.Error {
	if C.cbreak() == 0 {
		return CursesError{"Cbreak failed"};
	}
	return nil;
}

func Endwin() os.Error {
	if C.endwin() == 0 {
		return CursesError{"Endwin failed"};
	}
	return nil;
}

func (win *Window) Getch() int {
	return int(C.wgetch((*C.WINDOW)(win)));
}

func (win *Window) Addch(x, y int, c int32, flags int32) {
	C.mvwaddch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c) | C.chtype(flags));
}

// Since CGO currently can't handle varg C functions we'll mimic the
// ncurses addstr functions.
func (win *Window) Addstr(x, y int, str string, flags int32, v ...interface{}) {
	var newstr string
    if v != nil {
        newstr = fmt.Sprintf(str, v);
    } else {
        newstr = str
    }
	
	win.Move(x, y);
	
	for i := 0; i < len(newstr); i++ {
		C.waddch((*C.WINDOW)(win), C.chtype(newstr[i]) | C.chtype(flags));
	}
}

// Normally Y is the first parameter passed in curses.
func (win *Window) Move(x, y int) {
	C.wmove((*C.WINDOW)(win), C.int(y), C.int(x));
}

func (win *Window) Resize(rows, cols int) {
	C.wresize((*C.WINDOW)(win), C.int(rows), C.int(cols));
}


func (w *Window) Keypad(tf bool) os.Error {
	var outint int;
	if tf == true {outint = 1;}
	if tf == false {outint = 0;}
	if C.keypad((*C.WINDOW)(w), C.int(outint)) == 0 {
		return CursesError{"Keypad failed"};
	}
	return nil;
}

func (win *Window) Refresh() os.Error {
	if C.wrefresh((*C.WINDOW)(win)) == 0 {
		return CursesError{"refresh failed"};
	}
	return nil;
}

func (win *Window) Redrawln(beg_line, num_lines int) {
	C.wredrawln((*C.WINDOW)(win), C.int(beg_line), C.int(num_lines));
}

func (win *Window) Redraw() {
	C.redrawwin((*C.WINDOW)(win));
}

func (win *Window) Clear() {
	C.wclear((*C.WINDOW)(win));
}

func (win *Window) Erase() {
	C.werase((*C.WINDOW)(win));
}

func (win *Window) Clrtobot() {
	C.wclrtobot((*C.WINDOW)(win));
}

func (win *Window) Clrtoeol() {
	C.wclrtoeol((*C.WINDOW)(win));
}

func (win *Window) Box(verch, horch int) {
	C.box((*C.WINDOW)(win), C.chtype(verch), C.chtype(horch));
}

func (win *Window) Background(colour int32) {
	C.wbkgd((*C.WINDOW)(win), C.chtype(colour));
}
