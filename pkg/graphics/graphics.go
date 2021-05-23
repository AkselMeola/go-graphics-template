package graphics

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "image/color"
    "os"
)

type Graphics struct {
    cfg          *Config
    Paused       bool
}

type Config struct {
    Screen struct {
        Width  int
        Height int
    }
}

func NewGame(cfg *Config) *Graphics {
    return &Graphics{
        cfg:      cfg,
        Paused: false,
    }
}

// Initialize  graphics engine
func (g *Graphics) Init() error {
    return nil
}

// Get Screen size provided in config
func (g *Graphics) Layout(outsideWidth, outsideHeight int) (int, int) {
    return g.cfg.Screen.Width, g.cfg.Screen.Height
}

// Update graphics state
func (g *Graphics) Update() error {
    g.checkEscape()
    if g.isPaused() {
        return nil
    }


    return nil
}

// Draw graphics on screen
func (g *Graphics) Draw(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}

    image := ebiten.NewImage(20, 20)
    image.Fill(color.White)

    screen.DrawImage(image, op)
}

// Check if engine is paused
func (g *Graphics) isPaused() bool {
    if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
        g.Paused = !g.Paused
        fmt.Println("Paused: ", g.Paused)
    }

    return g.Paused
}

// Check if we should escape
func (g *Graphics) checkEscape() {
    if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
        fmt.Println("Exiting ... ")
        os.Exit(0)
    }
}
