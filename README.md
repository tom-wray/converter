# Obsidian Markdown to Jekyll just-the-docs Converter

## Overview

This Go program processes a directory of markdown files, adding front matter and creating index files for each subdirectory. It's designed to prepare markdown files for use in static site generators or other documentation systems that require specific front matter and directory structures. Specifically, it's designed to prepare obsidian vault files for use in the [just-the-docs](https://github.com/just-the-docs/just-the-docs) project.

## Requirements

- Go 1.x (replace x with the minimum version required)

## Usage

1. Clone the repository:
   ```
   git clone https://github.com/tom-wray/obsidian-jekyll-converter.git
   cd obsidian-jekyll-converter
   ```

2. Modify the constants in `main.go` if needed:
   - `sourceDir`: The directory containing your markdown files (use the relative path)
   - `sourceParent`: The same as above but just the name of the folder, rather than the relative path
   - `outputDir`: The directory where processed files will be saved (use the relative path)

3. Place your markdown files in the `sourceDir`

4. Run the program:
   ```
   go run .
   ```

5. Check the `output` directory for the processed files and new directory structure.

## How It Works

1. The program starts by cleaning the output directory.
2. It then walks through the source directory, processing each file and subdirectory:
   - For each subdirectory (except the root), it creates an `index.md` file with a title.
   - For each markdown file, it adds front matter including the title, parent (if applicable), and layout.
3. The processed files are saved in the output directory, maintaining the original directory structure.

## Customization

You can customize the behavior of the program by modifying the following:

- Modify the `createIndex` function to change the content of the generated index files.
- Adjust the `updateFile` function to alter the front matter added to each markdown file.

## License

This project is licensed under the MIT License. See the LICENSE file for details.