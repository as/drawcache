package drawcache

import (
	"image"
	"image/draw"
)

// cached is a drawer that remembers the clipping rectangle of each draw until flushed
type cached struct {
	cache []image.Rectangle
}

func (f *cached) memo(r image.Rectangle) {
	if len(f.cache) == 0 {
		f.cache = append(f.cache, r)
		return
	}
	c := f.cache[len(f.cache)-1]
	if r == f.cache[0] || r.In(f.cache[0]) {
		return
	}
	if r.Min.X == c.Max.X || r.Max.X == c.Min.X || r.Max.Y == c.Min.Y || r.Min.Y == c.Max.Y {
		f.cache[0] = f.cache[0].Union(r)
	} else {
		c := f.cache[0]
		if c.Dx()*c.Dy() < r.Dx()*r.Dy() {
			f.cache = append([]image.Rectangle{r}, f.cache...)
		} else {
			f.cache = append(f.cache, r)
		}
	}
}
func (f *cached) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op) {
	draw.Draw(dst, r, src, sp, op)
	f.memo(r)
}
func (f *cached) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op draw.Op) {
	draw.DrawMask(dst, r, src, sp, mask, mp, op)
	//f.memo(r)
}
func (f *cached) Cache() []image.Rectangle {
	return f.cache
}
func (f *cached) Flush() {
	f.cache = f.cache[:0]
}
func (f *cached) cacheinit() {
	f.cache = make([]image.Rectangle, 0, 1024)
}
