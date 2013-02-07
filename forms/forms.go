package forms

// #define _Bool int
// #include <form.h>
// #cgo LDFLAGS: -lform -lncurses
import "C"

import (
	. "github.com/jabb/gocurse/curses"
	"unsafe"
)

type Field C.FIELD
type Form C.FORM
type FieldType C.FIELDTYPE
type FieldOptions C.Field_Options
type FormOptions C.Form_Options

const (
	NO_JUSTIFICATION = C.NO_JUSTIFICATION
	JUSTIFY_LEFT     = C.JUSTIFY_LEFT
	JUSTIFY_CENTER   = C.JUSTIFY_CENTER
	JUSTIFY_RIGHT    = C.JUSTIFY_RIGHT

	O_VISIBLE  = C.O_VISIBLE
	O_ACTIVE   = C.O_ACTIVE
	O_PUBLIC   = C.O_PUBLIC
	O_EDIT     = C.O_EDIT
	O_WRAP     = C.O_WRAP
	O_BLANK    = C.O_BLANK
	O_AUTOSKIP = C.O_AUTOSKIP
	O_NULLOK   = C.O_NULLOK
	O_PASSOK   = C.O_PASSOK
	O_STATIC   = C.O_STATIC

	O_NL_OVERLOAD = C.O_NL_OVERLOAD
	O_BS_OVERLOAD = C.O_BS_OVERLOAD

	REQ_NEXT_PAGE  = C.REQ_NEXT_PAGE
	REQ_PREV_PAGE  = C.REQ_PREV_PAGE
	REQ_FIRST_PAGE = C.REQ_FIRST_PAGE
	REQ_LAST_PAGE  = C.REQ_LAST_PAGE

	REQ_NEXT_FIELD   = C.REQ_NEXT_FIELD
	REQ_PREV_FIELD   = C.REQ_PREV_FIELD
	REQ_FIRST_FIELD  = C.REQ_FIRST_FIELD
	REQ_LAST_FIELD   = C.REQ_LAST_FIELD
	REQ_SNEXT_FIELD  = C.REQ_SNEXT_FIELD
	REQ_SPREV_FIELD  = C.REQ_SPREV_FIELD
	REQ_SFIRST_FIELD = C.REQ_SFIRST_FIELD
	REQ_SLAST_FIELD  = C.REQ_SLAST_FIELD
	REQ_LEFT_FIELD   = C.REQ_LEFT_FIELD
	REQ_RIGHT_FIELD  = C.REQ_RIGHT_FIELD
	REQ_UP_FIELD     = C.REQ_UP_FIELD
	REQ_DOWN_FIELD   = C.REQ_DOWN_FIELD

	REQ_NEXT_CHAR  = C.REQ_NEXT_CHAR
	REQ_PREV_CHAR  = C.REQ_PREV_CHAR
	REQ_NEXT_LINE  = C.REQ_NEXT_LINE
	REQ_PREV_LINE  = C.REQ_PREV_LINE
	REQ_NEXT_WORD  = C.REQ_NEXT_WORD
	REQ_PREV_WORD  = C.REQ_PREV_WORD
	REQ_BEG_FIELD  = C.REQ_BEG_FIELD
	REQ_END_FIELD  = C.REQ_END_FIELD
	REQ_BEG_LINE   = C.REQ_BEG_LINE
	REQ_END_LINE   = C.REQ_END_LINE
	REQ_LEFT_CHAR  = C.REQ_LEFT_CHAR
	REQ_RIGHT_CHAR = C.REQ_RIGHT_CHAR
	REQ_UP_CHAR    = C.REQ_UP_CHAR
	REQ_DOWN_CHAR  = C.REQ_DOWN_CHAR

	REQ_NEW_LINE   = C.REQ_NEW_LINE
	REQ_INS_CHAR   = C.REQ_INS_CHAR
	REQ_INS_LINE   = C.REQ_INS_LINE
	REQ_DEL_CHAR   = C.REQ_DEL_CHAR
	REQ_DEL_PREV   = C.REQ_DEL_PREV
	REQ_DEL_LINE   = C.REQ_DEL_LINE
	REQ_DEL_WORD   = C.REQ_DEL_WORD
	REQ_CLR_EOL    = C.REQ_CLR_EOL
	REQ_CLR_EOF    = C.REQ_CLR_EOF
	REQ_CLR_FIELD  = C.REQ_CLR_FIELD
	REQ_OVL_MODE   = C.REQ_OVL_MODE
	REQ_INS_MODE   = C.REQ_INS_MODE
	REQ_SCR_FLINE  = C.REQ_SCR_FLINE
	REQ_SCR_BLINE  = C.REQ_SCR_BLINE
	REQ_SCR_FPAGE  = C.REQ_SCR_FPAGE
	REQ_SCR_BPAGE  = C.REQ_SCR_BPAGE
	REQ_SCR_FHPAGE = C.REQ_SCR_FHPAGE
	REQ_SCR_BHPAGE = C.REQ_SCR_BHPAGE
	REQ_SCR_FCHAR  = C.REQ_SCR_FCHAR
	REQ_SCR_BCHAR  = C.REQ_SCR_BCHAR
	REQ_SCR_HFLINE = C.REQ_SCR_HFLINE
	REQ_SCR_HBLINE = C.REQ_SCR_HBLINE
	REQ_SCR_HFHALF = C.REQ_SCR_HFHALF
	REQ_SCR_HBHALF = C.REQ_SCR_HBHALF

	REQ_VALIDATION  = C.REQ_VALIDATION
	REQ_NEXT_CHOICE = C.REQ_NEXT_CHOICE
	REQ_PREV_CHOICE = C.REQ_PREV_CHOICE

	MIN_FORM_COMMAND = C.MIN_FORM_COMMAND
	MAX_FORM_COMMAND = C.MAX_FORM_COMMAND
)

type FormsError struct {
	message string
}

func (fe FormsError) Error() string {
	return fe.message
}

/*
* FIELD METHODS
 */

func NewField(height int, width int, top int, left int, offscreen int, nbuf int) (*Field, error) {
	field := (*Field)(C.new_field(C.int(height), C.int(width), C.int(top), C.int(left), C.int(offscreen), C.int(nbuf)))
	if field == nil {
		return nil, FormsError{"NewField failed"}
	}
	return field, nil
}

func (field *Field) DupField(top int, left int) (*Field, error) {
	dup := (*Field)(C.dup_field((*C.FIELD)(field), C.int(top), C.int(left)))
	if dup == nil {
		return nil, FormsError{"Field.Dup failed"}
	}
	return dup, nil
}

func (field *Field) Link(top int, left int) (*Field, error) {
	link := (*Field)(C.link_field((*C.FIELD)(field), C.int(top), C.int(left)))
	if link == nil {
		return nil, FormsError{"Field.Link failed"}
	}
	return link, nil
}

func (field *Field) Free() error {
	if C.free_field((*C.FIELD)(field)) != C.OK {
		return FormsError{"Field.Free failed"}
	}
	return nil
}

func (field *Field) Info() (int, int, int, int, int, int, error) {
	var (
		height    C.int
		width     C.int
		top       C.int
		left      C.int
		offscreen C.int
		nbuf      C.int
	)
	if C.field_info((*C.FIELD)(field), &height, &width, &top, &left, &offscreen, &nbuf) != C.OK {
		return 0, 0, 0, 0, 0, 0, FormsError{"Field.Info failed"}
	}
	return int(height), int(width), int(top), int(left), int(offscreen), int(nbuf), nil
}

func (field *Field) DynamicInfo() (int, int, int, error) {
	var (
		prows C.int
		pcols C.int
		pmax  C.int
	)
	if C.dynamic_field_info((*C.FIELD)(field), &prows, &pcols, &pmax) != C.OK {
		return 0, 0, 0, FormsError{"Field.DynamicInfo failed"}
	}
	return int(prows), int(pcols), int(pmax), nil
}

func (field *Field) SetMax(max int) bool {
	return isOk(C.set_max_field((*C.FIELD)(field), C.int(max)))
}

func (field *Field) Move(x int, y int) bool {
	return isOk(C.move_field((*C.FIELD)(field), C.int(x), C.int(y)))
}

func (field *Field) SetNewPage(newPage bool) bool {
	return isOk(C.set_new_page((*C.FIELD)(field), boolToInt(newPage)))
}

func (field *Field) SetJust(justMode int) bool {
	return isOk(C.set_field_just((*C.FIELD)(field), C.int(justMode)))
}

func (field *Field) Just() int {
	return (int)(C.field_just((*C.FIELD)(field)))
}

func (field *Field) SetFore(fore Chtype) bool {
	return isOk(C.set_field_fore((*C.FIELD)(field), (C.chtype)(fore)))
}

func (field *Field) SetBack(back Chtype) bool {
	return isOk(C.set_field_back((*C.FIELD)(field), (C.chtype)(back)))
}

func (field *Field) SetPad(pad int) bool {
	return isOk(C.set_field_pad((*C.FIELD)(field), (C.bool)(C.int(pad))))
}

func (field *Field) Pad() int {
	return (int)(C.field_pad((*C.FIELD)(field)))
}

func (field *Field) SetBuffer(ind int, message string) bool {
	return isOk(C.set_field_buffer((*C.FIELD)(field), C.int(ind), C.CString(message)))
}

func (field *Field) SetStatus(status bool) bool {
	return isOk(C.set_field_status((*C.FIELD)(field), boolToInt(status)))
}

func (field *Field) SetOpts(attr FieldOptions) bool {
	return isOk(C.set_field_opts((*C.FIELD)(field), (C.Field_Options)(attr)))
}

func (field *Field) OptsOn(attr FieldOptions) bool {
	return isOk(C.field_opts_on((*C.FIELD)(field), (C.Field_Options)(attr)))
}

func (field *Field) OptsOff(attr FieldOptions) bool {
	return isOk(C.field_opts_off((*C.FIELD)(field), (C.Field_Options)(attr)))
}

func (field *Field) Buffer(ind int) string {
	buf := C.field_buffer((*C.FIELD)(field), C.int(ind))
	return C.GoString(buf)
}

func (field *Field) Fore() Chtype {
	return (Chtype)(C.field_fore((*C.FIELD)(field)))
}

func (field *Field) Back() Chtype {
	return (Chtype)(C.field_back((*C.FIELD)(field)))
}

func (field *Field) NewPage() bool {
	return intToBool(C.new_page((*C.FIELD)(field)))
}

func (field *Field) Opts() FieldOptions {
	return (FieldOptions)(C.field_opts((*C.FIELD)(field)))
}

func (field *Field) Index() int {
	return (int)(C.field_index((*C.FIELD)(field)))
}

/*
* FORM METHODS
 */

func NewForm(fields []*Field) (*Form, error) {
	form := (*Form)(C.new_form((**C.FIELD)(void(&fields[0]))))
	if form == nil {
		return nil, FormsError{"NewForm failed"}
	}
	return form, nil
}

func (form *Form) CurrentField() *Field {
	return (*Field)(C.current_field((*C.FORM)(form)))
}

func (form *Form) DataAhead() bool {
	return intToBool(C.data_ahead((*C.FORM)(form)))
}

func (form *Form) DataBehind() bool {
	return intToBool(C.data_behind((*C.FORM)(form)))
}

func (form *Form) Free() bool {
	return isOk(C.free_form((*C.FORM)(form)))
}

func (form *Form) SetFields(fields []*Field) bool {
	return isOk(C.set_form_fields((*C.FORM)(form), (**C.FIELD)(void(&fields[0]))))
}

func (form *Form) FieldCount() int {
	return (int)(C.field_count((*C.FORM)(form)))
}

func (form *Form) SetCurrentField(field *Field) bool {
	return isOk(C.set_current_field((*C.FORM)(form), (*C.FIELD)(field)))
}

func (form *Form) SetPage(page int) bool {
	return isOk(C.set_form_page((*C.FORM)(form), C.int(page)))
}

func (form *Form) Page() int {
	return (int)(C.form_page((*C.FORM)(form)))
}

func (form *Form) Post() bool {
	return isOk(C.post_form((*C.FORM)(form)))
}

func (form *Form) Unpost() bool {
	return isOk(C.unpost_form((*C.FORM)(form)))
}

func (form *Form) Drive(req int) bool {
	return isOk(C.form_driver((*C.FORM)(form), C.int(req)))
}

func (form *Form) Opts() FormOptions {
	return (FormOptions)(C.form_opts((*C.FORM)(form)))
}

func (form *Form) SetOpts(attr FormOptions) bool {
	return isOk(C.set_form_opts((*C.FORM)(form), (C.Form_Options)(attr)))
}

func (form *Form) OptsOn(attr FormOptions) bool {
	return isOk(C.form_opts_on((*C.FORM)(form), (C.Form_Options)(attr)))
}

func (form *Form) OptsOff(attr FormOptions) bool {
	return isOk(C.form_opts_off((*C.FORM)(form), (C.Form_Options)(attr)))
}

func (form *Form) Scale() (int, int, error) {
	var (
		rows C.int
		cols C.int
	)
	if C.scale_form((*C.FORM)(form), &rows, &cols) != C.OK {
		return 0, 0, FormsError{"Form.Scale failed"}
	}
	return int(rows), int(cols), nil
}

func (form *Form) SetWin(window *Window) bool {
	return isOk(C.set_form_win((*C.FORM)(form), (*C.WINDOW)(unsafe.Pointer(window))))
}

func (form *Form) SetSub(window *Window) bool {
	return isOk(C.set_form_sub((*C.FORM)(form), (*C.WINDOW)(unsafe.Pointer(window))))
}

func (form *Form) Win() *Window {
	return (*Window)(unsafe.Pointer((C.form_win((*C.FORM)(form)))))
}

func (form *Form) Sub() *Window {
	return (*Window)(unsafe.Pointer((C.form_sub((*C.FORM)(form)))))
}
