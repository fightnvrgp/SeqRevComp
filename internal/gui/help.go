package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const helpText = `
DNA/RNA 序列反向互补工具 — 使用说明

一、什么是反向互补（Reverse Complement）？
在分子生物学中，DNA 双链的两条链是反向互补的：
  • A 与 T 配对，C 与 G 配对
  • RNA 中用 U 代替 T，因此 A 与 U 配对
反向互补操作包括两步：
  1. 互补（Complement）：将每个碱基替换为其配对碱基
  2. 反向（Reverse）：将序列顺序颠倒

例如：
  DNA: 5'-ATGC-3'  →  3'-TACG-5'（写作 5'-GCAT-3'）
  RNA: 5'-AUGC-3'  →  3'-UACG-5'（写作 5'-GCAU-3'）

二、支持的输入格式
本工具支持多种常见输入形式，容错能力强：
  • 单行裸序列：ATGCCGTTA
  • 多行裸序列
  • FASTA 格式（保留 header）：
      >seq1
      ATGC...
  • 带空格、Tab、空行的文本
  • 混有序号或位置数字的文本（如 1 ATGC）
  • 从网页/PDF/Word 直接粘贴的不整洁文本

三、自动清洗规则
工具会自动对输入进行清洗：
  • 去除空格、制表符、换行符
  • 去除序列中混入的数字
  • 统一转换为大写输出
  • 保留 FASTA header（以 > 开头）
  • 多条 FASTA 序列会分别处理并保留对应 header

四、字符兼容性
支持标准碱基和常见 IUPAC 模糊碱基：
  • DNA/RNA: A, T, C, G, U
  • IUPAC 模糊碱基: R, Y, S, W, K, M, B, D, H, V, N
  • 小写字母同样支持

IUPAC 互补关系：
  R(A/G) ↔ Y(C/T/U)    S(C/G) ↔ S(C/G)
  W(A/T/U) ↔ W(A/T/U)  K(G/T/U) ↔ M(A/C)
  B(C/G/T/U) ↔ V(A/C/G) D(A/G/T/U) ↔ H(A/C/T/U)
  N ↔ N

五、DNA / RNA 处理规则
  • 若输入只含 U 不含 T → 按 RNA 规则处理
  • 若输入只含 T 不含 U → 按 DNA 规则处理
  • 若同时含有 T 和 U → 给出警告，默认按 DNA 规则处理（U 视为 T 处理）
  • 若既不含 T 也不含 U → 默认按 DNA 规则处理

六、输出规则
  • 输入为 FASTA 格式时，输出保持 FASTA 格式并保留 header
  • 输入为裸序列时，输出裸序列文本
  • 默认每行 60 个字符换行（FASTA 输出）
  • 非法字符默认被删除，不会导致程序崩溃

七、快捷键（部分平台支持）
  • Ctrl/Cmd + V：在输入框中粘贴
  • Ctrl/Cmd + C：在输出框中复制
  • 点击“一键粘贴”可直接从系统剪贴板读取内容到输入框
  • 点击“一键复制”可将输出框内容复制到系统剪贴板

如有问题，请检查输入是否为纯文本格式，并尽量避免混入特殊符号。
`

func OpenHelpWindow(app fyne.App) {
	win := app.NewWindow("使用说明")
	win.Resize(fyne.NewSize(700, 520))

	content := widget.NewRichTextFromMarkdown(helpText)
	content.Wrapping = fyne.TextWrapWord

	scroll := container.NewScroll(content)
	scroll.SetMinSize(fyne.NewSize(680, 480))

	closeBtn := widget.NewButton("关闭", func() {
		win.Close()
	})

	win.SetContent(container.NewBorder(nil, container.NewCenter(closeBtn), nil, nil, scroll))
	win.Show()
}

func OpenAboutWindow(app fyne.App) {
	win := app.NewWindow("关于")
	win.Resize(fyne.NewSize(400, 200))

	about := fmt.Sprintf("DNA/RNA 序列反向互补工具\n\n版本: 1.0.0\n\n跨平台桌面应用\n使用 Go + Fyne 构建")
	label := widget.NewLabel(about)
	label.Alignment = fyne.TextAlignCenter

	closeBtn := widget.NewButton("关闭", func() {
		win.Close()
	})

	win.SetContent(container.NewVBox(
		layout.NewSpacer(),
		label,
		layout.NewSpacer(),
		container.NewCenter(closeBtn),
	))
	win.Show()
}
