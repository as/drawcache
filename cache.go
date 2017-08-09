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
	if r == f.cache[0] || r.In(f.cache[0]) {
		return
	}
	if c := f.cache[len(f.cache)-1]; r.Min.X == c.Max.X || r.Max.X == c.Min.X || r.Max.Y == c.Min.Y || r.Min.Y == c.Max.Y {
		f.cache[0] = f.cache[0].Union(r)
		return
	}
	f.cache = append(f.cache, r)
	if c := f.cache[0]; c.Dx()*c.Dy() < r.Dx()*r.Dy() {
		l := len(f.cache)-1
		f.cache[0], f.cache[l] = f.cache[l], f.cache[0]
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
