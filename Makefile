all:
	make -C curses install
	make -C forms  install
	make -C menus  install
	make -C panels install
	
clean:
	make -C curses clean
	make -C forms  clean
	make -C menus  clean
	make -C panels clean
	
nuke:
	make -C curses nuke
	make -C forms  nuke
	make -C menus  nuke
	make -C panels nuke