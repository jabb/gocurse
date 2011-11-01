include $(GOROOT)/src/Make.inc

TARG=curses/forms
CGOFILES=\
		forms.go\
		utils.go

CGO_LDFLAGS=-lform -lncurses
CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
