package game

import (
	"errors"
	"fmt"
	"math/rand"

	ray "github.com/gen2brain/raylib-go/raylib"
)

var length int32 = 50
var hoveringTileX, hoveringTileY int32 = -1, -1

type TileState uint8

const (
	Hidden TileState = iota
	Found
	Flagged
)

type tile struct {
	isBomb             bool
	state              TileState
	neighborBombsCount uint8
}

type board [][]tile

func (b board) randomize() {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j].state = Hidden
			gameBoard[i][j].neighborBombsCount = 0
			if rand.Intn(100) < bombChance {
				b[i][j].isBomb = true
			} else {
				b[i][j].isBomb = false
			}
		}
	}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j].isBomb {
				for ii := i - 1; ii < i+2; ii++ {
					if ii == -1 || ii == boardWidth {
						continue
					}
					for jj := j - 1; jj < j+2; jj++ {
						if jj == -1 || jj == boardHeight {
							continue
						}
						if ii == i && jj == j {
							continue
						}
						gameBoard[ii][jj].neighborBombsCount++
					}
				}
			}
		}
	}
}

func (b board) draw() {
	ray.BeginMode2D(gameCamera)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			dst := ray.NewRectangle(float32(int32(i)*length+1), float32(int32(j)*length+1), float32(length-1), float32(length-1))
			switch b[i][j].state {
			case Hidden:
				if gs == STATE_GAMEOVER && b[i][j].isBomb {
					mineTexture.DrawAt(dst)
				} else {
					if int32(i) == hoveringTileX && int32(j) == hoveringTileY {
						cellDownTexture.DrawAt(dst)
					} else {
						cellUpTexture.DrawAt(dst)
					}
				}
			case Found:
				if b[i][j].isBomb {
					radMineTexture.DrawAt(dst)
				} else {
					cellDownTexture.DrawAt(dst)
					switch b[i][j].neighborBombsCount {
					case 1:
						cell1Texture.DrawAt(dst)
					case 2:
						cell2Texture.DrawAt(dst)
					case 3:
						cell3Texture.DrawAt(dst)
					case 4:
						cell4Texture.DrawAt(dst)
					case 5:
						cell5Texture.DrawAt(dst)
					case 6:
						cell6Texture.DrawAt(dst)
					case 7:
						cell7Texture.DrawAt(dst)
					case 8:
						cell8Texture.DrawAt(dst)
					}
				}
			case Flagged:
				if gs == STATE_GAMEOVER && b[i][j].isBomb {
					falseMineTexture.DrawAt(dst)
				} else {
					cellDownTexture.DrawAt(dst)
					flagTexture.DrawAt(dst)
				}
			}
		}
	}
	ray.EndMode2D()
}

func findTile(x float32, y float32) (int32, int32, error) {
	i := int32(x) / length
	j := int32(y) / length

	if i < 0 || j < 0 || i >= int32(boardWidth) || j >= int32(boardHeight) {
		return -1, -1, errors.New("out of range index")
	}
	return i, j, nil
}

func highlightTile(i int32, j int32) {
	hoveringTileX = i
	hoveringTileY = j
}

func (t *tile) flagTile(i int32, j int32) {
	if gameBoard[i][j].state == Hidden {
		if flagsCount == 0 {
			return
		}
		gameBoard[i][j].state = Flagged
		flagsCount--
		flagsText.SetText(fmt.Sprint(flagsCount))
	} else if gameBoard[i][j].state == Flagged {
		gameBoard[i][j].state = Hidden
		flagsCount++
		flagsText.SetText(fmt.Sprint(flagsCount))
	}
}

func (t *tile) selectTile(i int32, j int32) {
	if gameBoard[i][j].state == Hidden {
		gameBoard[i][j].state = Found
		if gameBoard[i][j].isBomb {
			gameOver()
		} else {
			// show non bomb neighbors
			if gameBoard[i][j].neighborBombsCount != 0 {
				return
			}

			for ii := i - 1; ii < i+2; ii++ {
				if ii == -1 || ii == int32(boardWidth) {
					continue
				}
				for jj := j - 1; jj < j+2; jj++ {
					if jj == -1 || jj == int32(boardHeight) {
						continue
					}
					if ii == i && jj == j {
						continue
					}
					gameBoard[ii][jj].selectTile(ii, jj)
					gameBoard[ii][jj].state = Found

				}
			}
		}
	}
}
