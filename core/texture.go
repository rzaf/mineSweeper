package core

import (
	ray "github.com/gen2brain/raylib-go/raylib"
)

type Texture struct {
	Texture  ray.Texture2D
	Dest     ray.Rectangle
	Src      ray.Rectangle
	Rotation float32
	Tint     ray.Color
	Filter   ray.TextureFilterMode
}

func NewTexture(path string, src ray.Rectangle) *Texture {
	var t Texture
	t.Texture = ray.LoadTexture(path)
	if !ray.IsTextureReady(t.Texture) {
		panic("texture loading failed")
	}
	t.Src = src
	t.Tint = ray.White
	t.Filter = ray.FilterBilinear
	ray.SetTextureFilter(t.Texture, ray.FilterBilinear)
	return &t
}
func (t *Texture) SetTextureFilter(f ray.TextureFilterMode) {
	t.Filter = f
	ray.SetTextureFilter(t.Texture, f)
}

func (t *Texture) Unload() {
	ray.UnloadTexture(t.Texture)
}

func (t *Texture) Draw() {
	ray.DrawTexturePro(t.Texture, t.Src, t.Dest, ray.NewVector2(0, 0), t.Rotation, t.Tint)
}

func (t *Texture) DrawAt(dst ray.Rectangle) {
	ray.DrawTexturePro(t.Texture, t.Src, dst, ray.NewVector2(0, 0), t.Rotation, t.Tint)
	// ray.DrawTexturePro(t.Texture, t.Src, ray.NewRectangle(x, y, t.Dest.Width, t.Dest.Height), ray.NewVector2(0, 0), t.Rotation, t.Tint)
}
