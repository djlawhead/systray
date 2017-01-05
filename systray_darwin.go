//+build darwin,!linux,!windows
package systray

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa

#include "systray.h"
*/
import "C"

import (
	"github.com/djlawhead/gobridgecocoa"
	"github.com/djlawhead/gococoaurlscheme"
	"unsafe"
)

func nativeLoop() {
	gococoaurlscheme.Setup()
	gobridgecocoa.Run()
	systrayReady()
}

func quit() {
	C.quit()
}

// SetIcon sets the systray icon.
// iconBytes should be the content of .ico for windows and .ico/.jpg/.png
// for other platforms.
func SetIcon(iconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&iconBytes[0]))
	C.setIcon(cstr, (C.int)(len(iconBytes)))
}

// SetTitle sets the systray title, only available on Mac.
func SetTitle(title string) {
	C.setTitle(C.CString(title))
}

// SetTooltip sets the systray tooltip to display on mouse hover of the tray icon,
// only available on Mac and Windows.
func SetTooltip(tooltip string) {
	C.setTooltip(C.CString(tooltip))
}

func addOrUpdateMenuItem(item *MenuItem) {
	var disabled C.short
	if item.disabled {
		disabled = 1
	}
	var checked C.short
	if item.checked {
		checked = 1
	}
	C.add_or_update_menu_item(
		C.int(item.id),
		C.CString(item.title),
		C.CString(item.tooltip),
		disabled,
		checked,
	)
}

//export systray_ready
func systray_ready() {
	systrayReady()
}

//export systray_menu_item_selected
func systray_menu_item_selected(cID C.int) {
	systrayMenuItemSelected(int32(cID))
}