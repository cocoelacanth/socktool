package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
	aic "github.com/TheZoraiz/ascii-image-converter/aic_package"
)

type cliFlags struct {
	jsonPath string
	imgDir string
	color bool
	charset string
}

type image struct {
	Frames []string
	Delay *int
	Loop *int
}

type model struct {
	flags cliFlags
	imgs map[string]image
	curImg *image
	ascii string
	frame int
	width int
	height int
	animID int
}

type AsciiGetMsg string

type TickMsg struct{
	ID int
}

func (m model) getAsciiCmd() tea.Cmd {
	return func() tea.Msg {
		return AsciiGetMsg(m.getAscii())
	}
}

func tickCmd(d time.Duration, id int) tea.Cmd {
	return tea.Tick(time.Millisecond*d, func(t time.Time) tea.Msg {
		return TickMsg{ID: id}
	})
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags] <json> <imgs>\n", os.Args[0])

	fmt.Fprintln(os.Stderr, "Required arguments:")
	fmt.Fprintln(os.Stderr, "  json")
	fmt.Fprintln(os.Stderr, "\tthe JSON file to parse")
	fmt.Fprintln(os.Stderr, "  imgs")
	fmt.Fprintln(os.Stderr, "\tthe location to search for image files")

	fmt.Fprintln(os.Stderr, "Optional flags:")
	flag.PrintDefaults()
}

func initialModel() model {
	var m model

	flag.BoolVar(&m.flags.color, "color", false, "enable colored ASCII art")
	flag.StringVar(&m.flags.charset, "chars", "", "a custom set of characters use in the ASCII art")

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(2)
	}

	m.flags.jsonPath = flag.Arg(0)
	m.flags.imgDir = flag.Arg(1)

	if imgs, err := m.getImgs(); err == nil {
		m.imgs = imgs
	}

	return m
}

func (m model) getImgs() (map[string]image, error) {
	jsonFile, err := os.Open(m.flags.jsonPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var imgs map[string]image

	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&imgs); err != nil {
		return nil, err
	}

	for key, img := range imgs {
		if (img.Loop == nil) {
			img.Loop = img.Delay
			imgs[key] = img
		}
	}

	return imgs, nil
}

func (m model) menuView() string {
	var imgs []string
	for key, img := range m.imgs {
		imgs = append(imgs, fmt.Sprintf("%s - %s", key, img.Frames[0]))
	}
	sort.Strings(imgs)
	list := strings.Join(imgs, "\n")

	logo :=
` Press a keybind to display an image.
                _    _              _
 ___  ___   ___| | _| |_ ___   ___ | |
/ __|/ _ \ / __| |/ / __/ _ \ / _ \| |
\__ \ (_) | (__|   <| || (_) | (_) | |
|___/\___/ \___|_|\_\\__\___/ \___/|_|

    Press Esc at any time to exit.`

	list = gloss.Place(
		m.width,
		m.height,
		gloss.Left,
		gloss.Top,
		list,
	)
    logo = gloss.Place(
   		38,
    	8,
      	gloss.Center,
       	gloss.Center,
        logo,
    )

	a := gloss.NewLayer(list).X(0).Y(0)
	b := gloss.NewLayer(logo).X((m.width-38)/2).Y((m.height-8)/2)

	return gloss.NewCompositor(a, b).Render()
}

func (m model) asciiView() string {
	return gloss.Place(
		m.width,
		m.height,
		gloss.Center,
		gloss.Center,
		m.ascii,
	)
}

func (m model) getAscii() string {
	imgPath := filepath.Join(m.flags.imgDir, m.curImg.Frames[m.frame])

	flags := aic.DefaultFlags()
	flags.Colored = m.flags.color
	if m.flags.charset != "" {
		flags.CustomMap = m.flags.charset
	}

	ascii, err := aic.Convert(imgPath, flags)
	if err != nil {
		return fmt.Sprintf("Error: could not convert %s to ASCII", imgPath)
	}

	return ascii
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.width = msg.Width
			m.height = msg.Height

		case tea.KeyPressMsg:
			switch msg.String() {
				case "ctrl+c", "esc":
					return m, tea.Quit
				default:
					img, ok := m.imgs[msg.String()]
					if ok {
						m.curImg = &img
						m.frame = 0
						m.animID++
						return m, m.getAsciiCmd()
					}
					return m, nil
			}

		case AsciiGetMsg:
			m.ascii = string(msg)
			if m.curImg.Delay != nil {
				if (m.frame == len(m.curImg.Frames)-1) {
					return m, tickCmd(time.Duration(*m.curImg.Loop), m.animID)
				}
				return m, tickCmd(time.Duration(*m.curImg.Delay), m.animID)
			}

		case TickMsg:
			if m.curImg == nil || msg.ID != m.animID {
				return m, nil
			}
			m.frame++
			if (m.frame >= len(m.curImg.Frames)) {
				m.frame = 0
			}
			return m, m.getAsciiCmd()
	}

	return m, nil
}

func (m model) View() tea.View {
	var canvas string
	if m.curImg == nil {
		canvas = m.menuView()
	} else {
		canvas = m.asciiView()
	}

    v := tea.NewView(canvas)
    v.AltScreen = true
	return v
}

func main() {
	flag.Usage = usage
	m := initialModel();
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Program error: %v\n", err)
		os.Exit(1)
	}
}
