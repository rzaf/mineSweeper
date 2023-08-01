package core

import (
	ray "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	Text     string
	Font     ray.Font
	Pos      ray.Vector2
	FontSize float32
	Spacing  float32
	Color    ray.Color
	size     ray.Vector2
}

func NewText(text string, Font ray.Font, Pos ray.Vector2, FontSize float32, Spacing float32, Color ray.Color) *Text {
	var txt *Text = &Text{
		text,
		Font,
		Pos,
		FontSize,
		Spacing,
		Color,
		ray.NewVector2(0, 0),
	}

	txt.Measure()
	return txt
}

func (t *Text) SetText(s string) {
	t.Text = s
	t.Measure()
}

func (t *Text) Measure() {
	if t.Font.Chars == nil {
		panic("nil font in initialized Text")
	}
	t.size = ray.MeasureTextEx(t.Font, t.Text, t.FontSize, t.Spacing)
}

func (t *Text) DrawCentered() {
	if t.size.X == 0.0 {
		t.Measure()
	}
	ray.DrawTextEx(
		t.Font,
		t.Text,
		ray.Vector2{X: t.Pos.X - t.size.X/2, Y: t.Pos.Y - t.size.Y/2},
		t.FontSize,
		t.Spacing,
		t.Color,
	)
}

func (t *Text) Draw() {
	ray.DrawTextEx(t.Font, t.Text, t.Pos, t.FontSize, t.Spacing, t.Color)
}
