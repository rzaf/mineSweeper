package game

import (
	"fmt"
	"math/rand"
	"time"

	ray "github.com/gen2brain/raylib-go/raylib"
	"github.com/rzaf/mineSweeper/core"
)

type gameState uint8

const (
	STATE_RUNNING gameState = iota
	STATE_GAMEOVER
	STATE_WIN
	STATE_MENU
)

var (
	gs            gameState = STATE_RUNNING
	boardWidth    int       = 20
	boardHeight   int       = 15
	bombChance    int       = 20
	flagsCount    int
	selectedCount int

	gameBoard  board
	gameCamera ray.Camera2D

	mineTexture      *core.Texture
	falseMineTexture *core.Texture
	radMineTexture   *core.Texture
	flagTexture      *core.Texture
	cellUpTexture    *core.Texture
	cellDownTexture  *core.Texture
	cell1Texture     *core.Texture
	cell2Texture     *core.Texture
	cell3Texture     *core.Texture
	cell4Texture     *core.Texture
	cell5Texture     *core.Texture
	cell6Texture     *core.Texture
	cell7Texture     *core.Texture
	cell8Texture     *core.Texture
	flag0Texture     *core.Texture
	smileFaceTexture *core.Texture
	clickFaceTexture *core.Texture
	winFaceTexture   *core.Texture
	lostFaceTexture  *core.Texture
	flagsText        *core.Text

	isWindowMaximized                 bool = false
	lastWindowWidth, lastWindowHeight int
)

func Load() {
	fmt.Println("loading !!!")
	gameBoard = make([][]tile, boardWidth)
	for i := 0; i < boardWidth; i++ {
		gameBoard[i] = make([]tile, boardHeight)
	}
	gameCamera = ray.NewCamera2D(ray.NewVector2(0, 0), ray.NewVector2(0, 0), 0, 1)
	gameCamera.Target.X = float32(length / 2 * int32(boardWidth))
	gameCamera.Target.Y = float32(length / 2 * int32(boardHeight))
	fitCamera()
	mineTexture = core.NewTexture("resources/mine1.png", ray.NewRectangle(0, 0, 234, 234))
	falseMineTexture = core.NewTexture("resources/falsemine.png", ray.NewRectangle(0, 0, 234, 234))
	radMineTexture = core.NewTexture("resources/mine2.png", ray.NewRectangle(0, 0, 234, 234))
	flagTexture = core.NewTexture("resources/flag.png", ray.NewRectangle(0, 0, 250, 250))
	cellUpTexture = core.NewTexture("resources/cellup.png", ray.NewRectangle(0, 0, 250, 250))
	cellDownTexture = core.NewTexture("resources/celldown.png", ray.NewRectangle(0, 0, 250, 250))
	cell1Texture = core.NewTexture("resources/cell1.png", ray.NewRectangle(0, 0, 250, 250))
	cell2Texture = core.NewTexture("resources/cell2.png", ray.NewRectangle(0, 0, 250, 250))
	cell3Texture = core.NewTexture("resources/cell3.png", ray.NewRectangle(0, 0, 250, 250))
	cell4Texture = core.NewTexture("resources/cell4.png", ray.NewRectangle(0, 0, 250, 250))
	cell5Texture = core.NewTexture("resources/cell5.png", ray.NewRectangle(0, 0, 250, 250))
	cell6Texture = core.NewTexture("resources/cell6.png", ray.NewRectangle(0, 0, 250, 250))
	cell7Texture = core.NewTexture("resources/cell7.png", ray.NewRectangle(0, 0, 250, 250))
	cell8Texture = core.NewTexture("resources/cell8.png", ray.NewRectangle(0, 0, 250, 250))

	flag0Texture = core.NewTexture("resources/flag0.png", ray.NewRectangle(0, 0, 283, 311))
	flag0Texture.Dest = ray.NewRectangle(13, 12, 35, 38)
	flagsText = core.NewText(fmt.Sprint(flagsCount), ray.GetFontDefault(), ray.NewVector2(55, 10), 50, 8, ray.NewColor(255, 0, 0, 255))
	smileFaceTexture = core.NewTexture("resources/smileface.png", ray.NewRectangle(0, 0, 321, 321))
	clickFaceTexture = core.NewTexture("resources/clickface.png", ray.NewRectangle(0, 0, 321, 321))
	winFaceTexture = core.NewTexture("resources/winface.png", ray.NewRectangle(0, 0, 321, 321))
	lostFaceTexture = core.NewTexture("resources/lostface.png", ray.NewRectangle(0, 0, 321, 321))
	smileFaceTexture.Dest = ray.NewRectangle(float32(ray.GetScreenWidth())-13-45, 12, 38, 38)
	clickFaceTexture.Dest = ray.NewRectangle(float32(ray.GetScreenWidth())-13-45, 12, 38, 38)
	winFaceTexture.Dest = ray.NewRectangle(float32(ray.GetScreenWidth())-13-45, 12, 38, 38)
	lostFaceTexture.Dest = ray.NewRectangle(float32(ray.GetScreenWidth())-13-45, 12, 38, 38)

	image := ray.LoadImageFromTexture(mineTexture.Texture)
	ray.SetWindowIcon(*image)
	ray.UnloadImage(image)
	startGame()
	rand.Seed(time.Now().Unix())
}

func Unload() {
	mineTexture.Unload()
	radMineTexture.Unload()
	flagTexture.Unload()
	cellUpTexture.Unload()
	cellDownTexture.Unload()
	cell1Texture.Unload()
	cell2Texture.Unload()
	cell3Texture.Unload()
	cell4Texture.Unload()
	cell5Texture.Unload()
	cell6Texture.Unload()
	cell7Texture.Unload()
	cell8Texture.Unload()
	flag0Texture.Unload()
	smileFaceTexture.Unload()
	clickFaceTexture.Unload()
	winFaceTexture.Unload()
	lostFaceTexture.Unload()
}

func startGame() {
	fmt.Println("Starting")
	selectedCount = 0
	gameBoard.randomize()
	gs = STATE_RUNNING
	// flagsCount = (boardWidth * boardHeight) / 5
	flagsText.SetText(fmt.Sprint(flagsCount))
	highlightTile(-1, -1)
}

func fitCamera() {
	gameCamera.Offset.X = float32(ray.GetScreenWidth() / 2)
	gameCamera.Offset.Y = float32(ray.GetScreenHeight() / 2)

	screenRatio := float32(ray.GetScreenHeight()) / float32(ray.GetScreenWidth())
	gridRatio := float32(length*int32(boardHeight)) / float32(length*int32(boardWidth))
	if screenRatio < gridRatio {
		gameCamera.Zoom = float32(ray.GetScreenHeight()) / float32(int(length)*boardHeight)
	} else {
		gameCamera.Zoom = float32(ray.GetScreenWidth()) / float32(int(length)*boardWidth)
	}
	// fmt.Printf("camera fitted zoom:%f ofsset:(%.2f,%.2f)\n", gameCamera.Zoom, gameCamera.Offset.X, gameCamera.Offset.Y)
}

func UpdateGame() {
	if ray.IsWindowResized() {
		faceNewDest := ray.NewRectangle(float32(ray.GetScreenWidth())-13-45, 12, 38, 38)
		smileFaceTexture.Dest = faceNewDest
		clickFaceTexture.Dest = faceNewDest
		winFaceTexture.Dest = faceNewDest
		lostFaceTexture.Dest = faceNewDest
		fitCamera()
	}
	if ray.IsKeyDown(ray.KeyLeftAlt) || ray.IsKeyDown(ray.KeyRightAlt) {
		if ray.IsKeyPressed(ray.KeyEnter) {
			if isWindowMaximized {
				// fmt.Println("back to ", lastWindowWidth, ", ", lastWindowHeight)
				ray.ClearWindowState(ray.FlagWindowMaximized)
				ray.SetWindowSize(lastWindowWidth, lastWindowHeight)
			} else {
				lastWindowWidth, lastWindowHeight = ray.GetScreenWidth(), ray.GetScreenHeight()
				ray.MaximizeWindow()
			}
			isWindowMaximized = !isWindowMaximized
		}
	}
	if gs == STATE_RUNNING && selectedCount == boardWidth*boardHeight {
		checkFullBoard()
	}
	if ray.IsMouseButtonDown(ray.MouseLeftButton) || ray.IsMouseButtonReleased(ray.MouseLeftButton) {
		// fmt.Println("mouse position: ", ray.GetMousePosition())
		// fmt.Println("world position: ", ray.GetScreenToWorld2D(ray.GetMousePosition(), gameCamera))
		if gs == STATE_RUNNING {
			pos := ray.GetScreenToWorld2D(ray.GetMousePosition(), gameCamera)
			i, j, error := findTile(pos.X, pos.Y)
			if error == nil {
				if ray.IsMouseButtonReleased(ray.MouseLeftButton) {
					gameBoard[i][j].selectTile(i, j)
				} else {
					highlightTile(i, j)
				}
			} else {
				highlightTile(-1, -1)
			}
		} else {
			if ray.IsMouseButtonReleased(ray.MouseLeftButton) &&
				ray.CheckCollisionPointRec(ray.GetMousePosition(), smileFaceTexture.Dest) {
				startGame()
			}
		}
	} else if ray.IsMouseButtonPressed(ray.MouseRightButton) {
		if gs == STATE_RUNNING {
			pos := ray.GetScreenToWorld2D(ray.GetMousePosition(), gameCamera)
			i, j, error := findTile(pos.X, pos.Y)
			if error == nil {
				gameBoard[i][j].flagTile(i, j)
			}
		}
	}

	if ray.IsKeyPressed(ray.KeyR) {
		startGame()
	}

}

func DrawGame() {
	gameBoard.draw()
	flag0Texture.Draw()
	flagsText.Draw()
	drawFace()
}

func drawFace() {
	switch gs {
	case STATE_RUNNING:
		if ray.IsMouseButtonDown(ray.MouseLeftButton) {
			clickFaceTexture.Draw()
		} else {
			smileFaceTexture.Draw()
		}
	case STATE_GAMEOVER:
		if ray.IsMouseButtonDown(ray.MouseLeftButton) && ray.CheckCollisionPointRec(ray.GetMousePosition(), smileFaceTexture.Dest) {
			smileFaceTexture.Draw()
		} else {
			lostFaceTexture.Draw()
		}
	case STATE_WIN:
		if ray.IsMouseButtonDown(ray.MouseLeftButton) && ray.CheckCollisionPointRec(ray.GetMousePosition(), smileFaceTexture.Dest) {
			smileFaceTexture.Draw()
		} else {
			winFaceTexture.Draw()
		}
	}
}

func gameOver() {
	gs = STATE_GAMEOVER
	fmt.Println("Game over")
}

func win() {
	gs = STATE_WIN
	fmt.Println("Win")
}

func checkFullBoard() {
	for i := 0; i < len(gameBoard); i++ {
		for j := 0; j < len(gameBoard[i]); j++ {
			switch gameBoard[i][j].state {
			case Hidden:
				panic("bug")
			case Flagged:
				if !gameBoard[i][j].isBomb {
					gameOver()
				}
			}
		}
	}
	win()
}
