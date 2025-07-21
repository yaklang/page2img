<div align="center">
  <h1>page2img</h1>
</div>

<div align="center">
  <a href="https://github.com/yaklang/page2img/actions/workflows/build.yml">
    <img src="https://github.com/yaklang/page2img/actions/workflows/build.yml/badge.svg" alt="Build Status">
  </a>
  <a href="https://www.gnu.org/licenses/agpl-3.0">
    <img src="https://img.shields.io/badge/License-AGPL%20v3-blue.svg" alt="License: AGPL v3">
  </a>
  <a href="https://goreportcard.com/report/github.com/yaklang/page2img">
    <img src="https://goreportcard.com/badge/github.com/yaklang/page2img" alt="Go Report Card">
  </a>
  <img src="https://img.shields.io/badge/go-1.22-blue.svg" alt="Go version">
</div>

<p align="center">
  <b>A simple, fast, and static command-line tool for converting document pages to images.</b>
  <br />
  <br />
  <a href="#page2img">English</a> | <a href="#page2img-中文版">中文</a>
</p>

---

## page2img

`page2img` is a simple and efficient command-line tool for converting each page of a document into a separate image file (PNG and JPEG are supported). It is built on top of [MuPDF](https://mupdf.com/), ensuring excellent compatibility, high performance, and making it a powerful utility for any "PDF to image" or "document conversion" tasks.

This project is **statically compiled**, meaning it has no external dynamic library dependencies. The resulting binary is self-contained, making it easy to deploy and use across various environments without worrying about installing libraries like `mupdf`.

### ✨ Features

- **Multi-format Support**: Natively supports all document formats handled by MuPDF.
- **High-Quality Output**: Can save pages as `.png` or `.jpeg`/`.jpg` files with configurable quality.
- **Simple CLI**: Intuitive command-line arguments for easy integration and automation.
- **Statically Linked**: The binary is dependency-free, making it portable and easy to distribute.
- **Cross-Platform**: Builds and runs on Linux, macOS, and Windows.

### supported-formats Supported Formats

The tool supports a wide range of document formats, thanks to its MuPDF core. Here are some of the most common ones:

| Format | Description |
| :--- | :--- |
| **PDF** | Portable Document Format |
| **XPS** | XML Paper Specification |
| **EPUB** | Electronic Publication |
| **CBZ** | Comic Book Archive |
| **FB2** | FictionBook 2.0 |
| *and more*| Any other format supported by MuPDF. |

### 🚀 Usage

Specify the input file and the output file pattern via command-line arguments.

```bash
./page2img -i <input-file-path> -o <output-image-pattern>
```

- `-i`: Specify the input document path.
- `-o`: Specify the output image naming pattern. Use `%d` as a placeholder for the page number. The file extension (`.png`, `.jpg`, `.jpeg`) determines the output format.

**Examples:**

```bash
# Convert each page of my_document.pdf to page-1.png, page-2.png, ...
./page2img -i my_document.pdf -o "page-%d.png"

# Convert each page of report.xps to report-1.jpg with 95% quality
./page2img -i report.xps -o "report-%d.jpg"
```

### 📦 Installation and Building

#### Option 1: Using `go install` (Recommended)

If you have a Go environment set up, you can install it directly with:

```bash
go install github.com/yaklang/page2img@latest
```

#### Option 2: Build from Source (For Static Binary)

For a fully static binary that you can copy to any machine, build it from the source.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/yaklang/page2img.git
    cd page2img
    ```

2.  **Run the build command:**
    ```bash
    CGO_ENABLED=1 go build -tags "static" -o page2img page2img.go
    ```
    - `CGO_ENABLED=1`: Enables CGO, which is necessary to call the underlying C library `mupdf`.
    - `-tags "static"`: Forces static linking, ensuring the `mupdf` library is compiled into the final binary. This eliminates the need to have `mupdf` installed on the runtime environment.

### 📜 License

This project relies on the [go-fitz](https://github.com/gen2brain/go-fitz) library, which is licensed under the **GNU Affero General Public License v3.0 (AGPLv3)**.

According to the terms of the AGPLv3, any derivative work that links to the library must also be released under the same license. Therefore, the `page2img` project is also licensed under **AGPLv3**.

If you use this tool in your services (especially if provided over a network), you are obligated to provide users with the complete source code of this tool.

---

## page2img (中文版)

`page2img` 是一个简单、高效的命令行工具，用于将文档的每一页转换为独立的图片文件（支持 PNG 和 JPEG 格式）。它基于 [MuPDF](https://mupdf.com/)，具有出色的兼容性、高性能和强大的“PDF转图片”或“文档转换”能力。

本项目为**纯静态编译**，不依赖任何外部动态链接库。生成的可执行文件是自包含的，方便在各种环境中部署和使用，无需担心安装 `mupdf` 等库。

### ✨ 功能特性

- **多格式支持**: 原生支持所有 MuPDF 能处理的文档格式。
- **高质量输出**: 可将页面保存为 `.png` 或 `.jpeg`/`.jpg` 格式，并可配置图片质量。
- **简单易用**: 简单直观的命令行参数，易于集成和自动化。
- **纯静态链接**: 二进制文件无外部依赖，易于分发和移植。
- **跨平台**: 可在 Linux、macOS 和 Windows 上编译和运行。

### 格式支持

得益于其 MuPDF 内核，本工具支持多种文档格式。以下是一些最常见的格式：

| 格式 | 描述 |
| :--- | :--- |
| **PDF** | 便携式文档格式 |
| **XPS** | XML 纸张规范 |
| **EPUB** | 电子出版物 |
| **CBZ** | 漫画书存档 |
| **FB2** | FictionBook 2.0 |
| *及更多*| MuPDF 支持的任何其他格式。 |

### 🚀 使用方法

通过命令行参数指定输入文件和输出文件命名模式。

```bash
./page2img -i <输入文件路径> -o <输出图片命名模式>
```

- `-i`: 指定输入的文档路径。
- `-o`: 指定输出图片的命名格式。请使用 `%d` 作为页码占位符。文件扩展名 `.png`, `.jpg`, `.jpeg` 用于决定输出格式。

**示例:**

```bash
# 将 my_document.pdf 的每一页转换为 page-1.png, page-2.png, ...
./page2img -i my_document.pdf -o "page-%d.png"

# 将 report.xps 的每一页转换为 report-1.jpg
./page2img -i report.xps -o "report-%d.jpg"
```

### 📦 安装和编译

#### 方式一：使用 `go install` (推荐)

如果您已经配置好 Go 环境，可以使用以下命令直接安装：

```bash
go install github.com/yaklang/page2img@latest
```

#### 方式二：从源码编译 (生成静态文件)

为了获得可以复制到任何机器上运行的纯静态二进制文件，请从源代码编译。

1.  **克隆本仓库：**
    ```bash
    git clone https://github.com/yaklang/page2img.git
    cd page2img
    ```

2.  **执行编译：**
    ```bash
    CGO_ENABLED=1 go build -tags "static" -o page2img page2img.go
    ```
    - `CGO_ENABLED=1`: 启用 CGO，这是调用底层 C 库 `mupdf` 所必需的。
    - `-tags "static"`: 强制进行静态链接，确保 `mupdf` 库被编译进最终的二进制文件中，从而无需在运行环境中安装 `mupdf`。

### 📜 许可证 (License)

本项目基于 [go-fitz](https://github.com/gen2brain/go-fitz) 库，而 `go-fitz` 是在 **GNU Affero General Public License v3.0 (AGPLv3)** 下授权的。

根据 AGPLv3 的条款，任何链接到该库的衍生作品也必须在该许可证下发布。因此，`page2img` 项目同样采用 **AGPLv3** 许可证。

如果您在您的服务中使用了本工具（尤其是通过网络提供服务），您有义务向用户提供本工具的完整源代码。 