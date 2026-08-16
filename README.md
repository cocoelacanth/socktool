```
                _    _              _
 ___  ___   ___| | _| |_ ___   ___ | |
/ __|/ _ \ / __| |/ / __/ _ \ / _ \| |
\__ \ (_) | (__|   <| || (_) | (_) | |
|___/\___/ \___|_|\_\\__\___/ \___/|_|
```
a tool to create keybinds for ascii art images

## getting socktool
there are three main methods of obtaining socktool.

### prebuilt releases
prebuilt binaries for linux, macOS, and windows are available on the [releases page](https://github.com/cocoelacanth/socktool/releases).

note that at the time of writing, the macOS and arm64 binaries are untested because i do not have a machine that can run them.

### installing from source
make sure that you have go (>=1.26) installed, then run:
```sh
go install github.com/cocoelacanth/socktool@latest
```
this will build the latest release of socktool and install it to your `GOPATH`.

### building from source
make sure that you have go (>=1.26) installed.

enter the source folder, then run:
```sh
mkdir -p bin
go build -o ./bin .
```
this will build socktool and place it in the bin/ folder.

## usage
1. create a folder with all your images
2. create a json file with entries for each image/animation
3. run socktool
4. press a keybind to display an image!

### command line
```
Usage: socktool [flags] <json> <imgs>
Required arguments:
  json
        the JSON file to parse
  imgs
        the location to search for image files
Optional flags:
  -chars string
        a custom set of characters use in the ASCII art
  -color
        enable colored ASCII art
```
you can run `socktool -h` to see this information in your terminal.

### supported image formats
* .jpeg, .jpg
* .png
* .bmp
* .webp
* .tiff, .tif

looping image sequences are supported, but animated gifs are not.

### json format

#### example
```json
{
  "q": {
    "frames": ["foo.png"]
  },
  "w": {
    "frames": [
      "bar_1.png",
      "bar_2.png",
      "bar_3.png"
    ],
    "delay": 100,
    "loop": 1000
  }
}
```

* `"q"` / `"w"` - the desired key as a string
  * some key combinations, like `"ctrl+w"`, `"ctrl+y"`, etc. are also supported
* `"frames"` - chronological array of frames to display, separated with commas
* `"delay"` - millisecond gap between frames (`100` = 0.1s = 10 FPS)
  * omit to disable animation
* `"loop"` - millisecond pause between loops
  * omit to use `"delay"` value

## licenses
this project is licensed under the GNU GPLv3. see LICENSE for details.

### third-party
this project includes software from:

* [https://github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)
  * licensed under the MIT License
* [https://github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)
  * licensed under the MIT License
* [https://github.com/TheZoraiz/ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter)
  * licensed under the Apache License 2.0

see LICENSES/ for details.
