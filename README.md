# SeqRevComp

A simple cross-platform desktop GUI tool for DNA/RNA reverse complement.

## Download

Pre-built binaries are available on the [Releases](https://github.com/fightnvrgp/SeqRevComp/releases) page.

| Platform | Package |
|---|---|
| macOS | `.dmg` |
| Windows | `.exe` |
| Linux | `.tar.gz` |

## Features

- Paste DNA/RNA sequences and get reverse complement results instantly
- Supports FASTA format, raw sequences, and dirty inputs (spaces, numbers, mixed case)
- Auto-detects DNA vs RNA
- Supports IUPAC ambiguity codes
- One-click clipboard paste/copy
- Keyboard shortcut: `Ctrl/Cmd + Enter` to run

## Build from source

Requires [Go](https://go.dev/dl/) 1.21+.

```bash
go mod download
go build -o seqrevcomp .
```

Or package with [Fyne](https://fyne.io/):

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon Icon.png   # macOS
fyne package -os windows -icon Icon.png  # Windows
fyne package -os linux -icon Icon.png    # Linux
```

## Usage

1. Paste or type your sequence into the left panel
2. Click **执行反向互补** or press `Ctrl/Cmd + Enter`
3. Copy the result from the right panel

For detailed help, open **Help → 使用说明** from the menu bar.

---

# SeqRevComp（中文）

DNA/RNA 序列反向互补桌面工具。

## 下载

已编译版本见 [Releases](https://github.com/fightnvrgp/SeqRevComp/releases)。

| 平台 | 安装包 |
|---|---|
| macOS | `.dmg` |
| Windows | `.exe` |
| Linux | `.tar.gz` |

## 功能

- 粘贴序列，一键获得反向互补结果
- 支持 FASTA 格式、裸序列、带空格/数字的不规范输入
- 自动识别 DNA / RNA
- 支持 IUPAC 模糊碱基
- 一键剪贴板粘贴/复制
- 快捷键：`Ctrl/Cmd + Enter` 执行

## 自行编译

需要 [Go](https://go.dev/dl/) 1.21+。

```bash
go mod download
go build -o seqrevcomp .
```

使用 [Fyne](https://fyne.io/) 打包：

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon Icon.png   # macOS
fyne package -os windows -icon Icon.png  # Windows
fyne package -os linux -icon Icon.png    # Linux
```

## 使用方式

1. 在左侧面板粘贴或输入序列
2. 点击「执行反向互补」或按 `Ctrl/Cmd + Enter`
3. 在右侧面板复制结果

详细说明请从菜单栏打开 **帮助 → 使用说明**。

## License

MIT
