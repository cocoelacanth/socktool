# socktool
a tool to create keybinds for ascii art images

> [!NOTE]
> this tool is still unfinished and in development. it is functional, but not particularly portable and likely not yet fit for general use.

## building
ensure you have go (>=1.26.3) installed, then:

```sh
mkdir -p bin
go build -o ./bin .
```

## usage
  1. put your images in bin/img/
  2. add an entry for each image/animation in bin/images.json
  3. run socktool
  4. press a keybind to display an image!

### supported formats
* .jpeg, .jpg
* .png
* .bmp
* .webp
* .tiff, .tif

looping image sequences are supported, but animated gifs are not.

### images.json format

#### example: still image
```json
"q": {
  "frames": ["foo.png"]
}
```

#### example: animation
```json
"w": {
  "frames": [
    "bar/frame_1.png",
    "bar/frame_2.png",
    "bar/frame_3.png"
  ],
  "delay": 100,
  "loop": 1000
}
```

* `"q"` / `"w"` - the desired key as a string
  * some key combinations, like `"ctrl+w"`, `"ctrl+y"`, etc. are also supported
* `"frames"` - chronological array of frames to display, separated with commas
* `"delay"` - millisecond gap between frames (`100` = 0.1s = 10 FPS)
  * omit to disable animation
* `"loop"` - millisecond pause between loops
  * defaults to `"delay"` value
