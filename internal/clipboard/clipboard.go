// Package clipboard provides cross-platform clipboard access.
package clipboard

import (
	"fyne.io/fyne/v2"
)

// Read reads text from the system clipboard.
func Read(clipboard fyne.Clipboard) string {
	return clipboard.Content()
}

// Write writes text to the system clipboard.
func Write(clipboard fyne.Clipboard, text string) {
	clipboard.SetContent(text)
}
