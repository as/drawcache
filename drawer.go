package drawcache

import (
	"image"
	"image/draw"

	"github.com/as/frame/font"
)

// A Drawer implements the drawing primitives in image/draw along methods
// for accessing and deleting a slice of possibly-coalesced clipping rectangles.
type Drawer interface {
	Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op)
	DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op draw.Op)
	Flush(r ...image.Rectangle) error
	StringBG(dst draw.Image, p image.Point, src image.Image, sp image.Point, ft *font.Font, s []byte, bg image.Image, bgp image.Point) int
}

// New returns a Drawer that coalesces rectangles
func New() Drawer {
	return &cached{make([]image.Rectangle, 0, 1024)}
}
