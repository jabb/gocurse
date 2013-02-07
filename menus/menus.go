package menus

// #define _Bool int
// #include <menu.h>
// #cgo LDFLAGS: -lmenu -lncurses
import "C"

import (
	. "github.com/jabb/gocurse/curses"
	"unsafe"
)

type Text C.TEXT
type Item C.ITEM
type Menu C.MENU
type ItemOptions C.Item_Options
type MenuOptions C.Menu_Options

const (
	O_ONEVALUE        = C.O_ONEVALUE
	O_SHOWDESC        = C.O_SHOWDESC
	O_ROWMAJOR        = C.O_ROWMAJOR
	O_IGNORECASE      = C.O_IGNORECASE
	O_SHOWMATCH       = C.O_SHOWMATCH
	O_NONCYCLIC       = C.O_NONCYCLIC
	O_SELECTABLE      = C.O_SELECTABLE
	REQ_LEFT_ITEM     = C.REQ_LEFT_ITEM
	REQ_RIGHT_ITEM    = C.REQ_RIGHT_ITEM
	REQ_UP_ITEM       = C.REQ_UP_ITEM
	REQ_DOWN_ITEM     = C.REQ_DOWN_ITEM
	REQ_SCR_ULINE     = C.REQ_SCR_ULINE
	REQ_SCR_DLINE     = C.REQ_SCR_DLINE
	REQ_SCR_DPAGE     = C.REQ_SCR_DPAGE
	REQ_SCR_UPAGE     = C.REQ_SCR_UPAGE
	REQ_FIRST_ITEM    = C.REQ_FIRST_ITEM
	REQ_LAST_ITEM     = C.REQ_LAST_ITEM
	REQ_NEXT_ITEM     = C.REQ_NEXT_ITEM
	REQ_PREV_ITEM     = C.REQ_PREV_ITEM
	REQ_TOGGLE_ITEM   = C.REQ_TOGGLE_ITEM
	REQ_CLEAR_PATTERN = C.REQ_CLEAR_PATTERN
	REQ_BACK_PATTERN  = C.REQ_BACK_PATTERN
	REQ_NEXT_MATCH    = C.REQ_NEXT_MATCH
	REQ_PREV_MATCH    = C.REQ_PREV_MATCH

	MIN_MENU_COMMAND = C.MIN_MENU_COMMAND
	MAX_MENU_COMMAND = C.MAX_MENU_COMMAND
)

type MenusError struct {
	message string
}

func (e MenusError) Error() string {
	return e.message
}

/*
 * Menu functions
 */

func (menu *Menu) CurrentItem() *Item {
	return (*Item)(C.current_item((*C.MENU)(menu)))
}

func NewItem(name string, desc string) *Item {
	return (*Item)(C.new_item(C.CString(name), C.CString(desc)))
}

func NewMenu(items []*Item) (*Menu, error) {
	menu := (*Menu)(C.new_menu((**C.ITEM)(void(&items[0]))))
	if menu == nil {
		return nil, MenusError{"NewMenu failed"}
	}
	return menu, nil
}

func (item *Item) Opts() ItemOptions {
	return ItemOptions(C.item_opts((*C.ITEM)(item)))
}

func (menu *Menu) Opts() MenuOptions {
	return MenuOptions(C.menu_opts((*C.MENU)(menu)))
}

func (item *Item) Description() string {
	return C.GoString(C.item_description((*C.ITEM)(item)))
}

func (item *Item) Name() string {
	return C.GoString(C.item_name((*C.ITEM)(item)))
}

func (menu *Menu) Mark() string {
	return C.GoString(C.menu_mark((*C.MENU)(menu)))
}

func (menu *Menu) SetMark(mark string) bool {
	return isOk(C.set_menu_mark((*C.MENU)(menu), C.CString(mark)))
}

func (menu *Menu) Pattern() string {
	return C.GoString(C.menu_pattern((*C.MENU)(menu)))
}

func (menu *Menu) Back() Chtype {
	return Chtype(C.menu_back((*C.MENU)(menu)))
}

func (menu *Menu) Fore() Chtype {
	return Chtype(C.menu_fore((*C.MENU)(menu)))
}

func (menu *Menu) Grey() Chtype {
	return Chtype(C.menu_grey((*C.MENU)(menu)))
}

func (item *Item) Free() bool {
	return isOk(C.free_item((*C.ITEM)(item)))
}

func (menu *Menu) Free() bool {
	return isOk(C.free_menu((*C.MENU)(menu)))
}

func (menu *Menu) ItemCount() int {
	return int(C.item_count((*C.MENU)(menu)))
}

func (item *Item) Index() int {
	return int(C.item_index((*C.ITEM)(item)))
}

func (item *Item) OptsOn(opt ItemOptions) bool {
	return isOk(C.item_opts_on((*C.ITEM)(item), (C.Item_Options)(opt)))
}

func (item *Item) OptsOff(opt ItemOptions) bool {
	return isOk(C.item_opts_off((*C.ITEM)(item), (C.Item_Options)(opt)))
}

func (menu *Menu) Drive(req int) bool {
	return isOk(C.menu_driver((*C.MENU)(menu), C.int(req)))
}

func (menu *Menu) OptsOn(opt MenuOptions) bool {
	return isOk(C.menu_opts_on((*C.MENU)(menu), (C.Menu_Options)(opt)))
}

func (menu *Menu) OptsOff(opt MenuOptions) bool {
	return isOk(C.menu_opts_off((*C.MENU)(menu), (C.Menu_Options)(opt)))
}

func (menu *Menu) Pad() int {
	return int(C.menu_pad((*C.MENU)(menu)))
}

func (menu *Menu) Post() bool {
	return isOk(C.post_menu((*C.MENU)(menu)))
}

func (menu *Menu) Unpost() bool {
	return isOk(C.unpost_menu((*C.MENU)(menu)))
}

func (menu *Menu) SetCurrentItem(item *Item) bool {
	return isOk(C.set_current_item((*C.MENU)(menu), (*C.ITEM)(item)))
}

func (menu *Menu) SetWin(win *Window) bool {
	return isOk(C.set_menu_win((*C.MENU)(menu), (*C.WINDOW)(unsafe.Pointer(win))))
}

func (menu *Menu) Win() *Window {
	return (*Window)(unsafe.Pointer((C.menu_win((*C.MENU)(menu)))))
}

func (menu *Menu) SetSub(win *Window) bool {
	return isOk(C.set_menu_sub((*C.MENU)(menu), (*C.WINDOW)(unsafe.Pointer(win))))
}

func (menu *Menu) Sub() *Window {
	return (*Window)(unsafe.Pointer((C.menu_sub((*C.MENU)(menu)))))
}

func (item *Item) Value() bool {
	return intToBool(C.item_value((*C.ITEM)(item)))
}

func (menu *Menu) Scale() (int, int, error) {
	var (
		rows C.int
		cols C.int
	)
	if C.scale_menu((*C.MENU)(menu), &rows, &cols) != C.OK {
		return 0, 0, MenusError{"Form.Scale failed"}
	}
	return int(rows), int(cols), nil
}

func (item *Item) Visible() bool {
	return intToBool(C.item_visible((*C.ITEM)(item)))
}

func (menu *Menu) Format(rows int, cols int) {
	cRows := C.int(rows)
	cCols := C.int(cols)
	C.menu_format((*C.MENU)(menu), &cRows, &cCols)
}

func (item *Item) SetUserPtr(ptr unsafe.Pointer) bool {
	return isOk(C.set_item_userptr((*C.ITEM)(item), ptr))
}

func (item *Item) UserPtr() unsafe.Pointer {
	return unsafe.Pointer(C.item_userptr((*C.ITEM)(item)))
}
