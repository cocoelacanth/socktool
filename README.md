# socktool
a tool to create keybinds for ascii art images

> [!NOTE]
> this tool is still unfinished and in development. it is functional, but not particularly portable and likely not yet fit for general use.

## requirements
ensure you have go (>=1.26.3) installed. other than that, there are no external dependencies to install.

## installing
to install socktool as a binary to your `GOPATH`, run:
```sh
go install github.com/cocoelacanth/socktool@latest
```

## building
to build the socktool binary, run:
```sh
git clone https://github.com/cocoelacanth/socktool
cd socktool
mkdir -p bin
go build -o ./bin .
```

## usage
```sh
$ socktool -h
```
```
Usage of socktool:
  -chars string
    	a custom set of characters use in the ASCII art
  -color
    	whether the ASCII art should have color
  -imgs string
    	(required) the location to search for image files
  -json string
    	(required) the JSON file containing images
```

  1. create a folder with all your images
  2. create a json file with entries for each image/animation
  3. run socktool
  4. press a keybind to display an image!

### supported formats
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
      "bar/frame_1.png",
      "bar/frame_2.png",
      "bar/frame_3.png"
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
