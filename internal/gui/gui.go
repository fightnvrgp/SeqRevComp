package gui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"seqrevcomp/internal/clipboard"
	"seqrevcomp/internal/formatter"
)

// readOnlyEntry is a multi-line entry that looks normal but rejects edits.
type readOnlyEntry struct {
	widget.Entry
}

func newReadOnlyEntry() *readOnlyEntry {
	e := &readOnlyEntry{}
	e.ExtendBaseWidget(e)
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapWord
	return e
}

func (e *readOnlyEntry) TypedRune(r rune) {}
func (e *readOnlyEntry) TypedKey(key *fyne.KeyEvent) {}
func (e *readOnlyEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*fyne.ShortcutCopy); ok {
		e.Entry.TypedShortcut(s)
	}
}

// BuildUI constructs the main application window content.
func BuildUI(app fyne.App, win fyne.Window) fyne.CanvasObject {
	// Input area
	inputEntry := widget.NewMultiLineEntry()
	inputEntry.SetPlaceHolder("在此粘贴 DNA 或 RNA 序列...\n支持 FASTA 格式、多行、带空格、带数字等")
	inputEntry.Wrapping = fyne.TextWrapWord

	inputScroll := container.NewScroll(inputEntry)

	// Output area: use readOnlyEntry to keep normal appearance
	outputEntry := newReadOnlyEntry()
	outputEntry.SetPlaceHolder("反向互补结果将显示在此处...")

	outputScroll := container.NewScroll(outputEntry)

	// Status label: single line, no wrapping to avoid vertical expansion
	statusLabel := widget.NewLabel("就绪 | 请输入序列")

	// Clipboard helper
	clip := win.Clipboard()

	// Input buttons
	pasteBtn := widget.NewButton("一键粘贴", func() {
		text := clipboard.Read(clip)
		if text != "" {
			inputEntry.SetText(text)
			statusLabel.SetText("已粘贴剪贴板内容")
		} else {
			statusLabel.SetText("剪贴板为空")
		}
	})
	pasteBtn.Importance = widget.HighImportance

	clearInputBtn := widget.NewButton("清空输入", func() {
		inputEntry.SetText("")
		statusLabel.SetText("输入已清空")
	})

	// Output buttons
	copyBtn := widget.NewButton("一键复制", func() {
		text := outputEntry.Text
		if text != "" {
			clipboard.Write(clip, text)
			statusLabel.SetText("输出已复制到剪贴板")
		} else {
			statusLabel.SetText("输出为空，无需复制")
		}
	})
	copyBtn.Importance = widget.HighImportance

	clearOutputBtn := widget.NewButton("清空输出", func() {
		outputEntry.SetText("")
		statusLabel.SetText("输出已清空")
	})

	swapBtn := widget.NewButton("交换输入/输出", func() {
		in := inputEntry.Text
		out := outputEntry.Text
		if out != "" {
			inputEntry.SetText(out)
			outputEntry.SetText(in)
			statusLabel.SetText("已交换输入与输出")
		} else {
			statusLabel.SetText("输出为空，无法交换")
		}
	})

	// Main action
	rcBtn := widget.NewButton("执行反向互补", func() {
		input := inputEntry.Text
		if strings.TrimSpace(input) == "" {
			statusLabel.SetText("错误: 输入为空")
			return
		}

		opts := formatter.DefaultFormatOptions()
		result := formatter.ProcessInput(input, opts)

		outputEntry.SetText(result.Output)
		statusLabel.SetText(formatter.BuildStatusMessage(result))
	})
	rcBtn.Importance = widget.SuccessImportance

	// Build consistent single-line headers: title on the left, buttons on the right.
	leftTitle := widget.NewLabelWithStyle("输入序列", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	leftHeaderRow := container.NewBorder(nil, nil, leftTitle, container.NewHBox(pasteBtn, clearInputBtn))

	rightTitle := widget.NewLabelWithStyle("反向互补结果", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	rightHeaderRow := container.NewBorder(nil, nil, rightTitle, container.NewHBox(copyBtn, clearOutputBtn))

	// Panels: header + separator + padded scroll area so both sides have identical structure and height.
	leftPanel := container.NewBorder(
		container.NewVBox(leftHeaderRow, widget.NewSeparator()),
		nil, nil, nil,
		container.NewPadded(container.NewMax(inputScroll)),
	)

	rightPanel := container.NewBorder(
		container.NewVBox(rightHeaderRow, widget.NewSeparator()),
		nil, nil, nil,
		container.NewPadded(container.NewMax(outputScroll)),
	)

	// Split pane: divides the middle area evenly; size is controlled only by window resize
	split := container.NewHSplit(leftPanel, rightPanel)
	split.Offset = 0.5

	// Keyboard shortcuts: Ctrl/Cmd + Enter triggers reverse complement.
	ctrlEnter := &desktop.CustomShortcut{
		KeyName:  fyne.KeyReturn,
		Modifier: fyne.KeyModifierControl,
	}
	win.Canvas().AddShortcut(ctrlEnter, func(shortcut fyne.Shortcut) {
		rcBtn.OnTapped()
	})

	cmdEnter := &desktop.CustomShortcut{
		KeyName:  fyne.KeyReturn,
		Modifier: fyne.KeyModifierSuper,
	}
	win.Canvas().AddShortcut(cmdEnter, func(shortcut fyne.Shortcut) {
		rcBtn.OnTapped()
	})

	// Top banner: all elements in a single horizontal row inside a card
	// to ensure compact height and horizontal layout only.
	toolbar := container.NewHBox(
		rcBtn,
		swapBtn,
		layout.NewSpacer(),
		statusLabel,
	)
	topCard := widget.NewCard("", "", toolbar)

	// Root layout: top card is fixed height, split fills the rest
	mainContent := container.NewBorder(topCard, nil, nil, nil, split)
	return mainContent
}
