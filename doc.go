// Package drawcache implements a drawer that remembers the clipping rectangles used in calls to Draw since the
// last call to Flush. Rectangles are combined if their vertices are connected on the x or y-axis. This
// allows coalesed subimages to be uploaded to high resolution windows and buffers instead of the entire image.
package drawcache
