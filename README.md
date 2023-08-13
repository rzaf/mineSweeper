# mine sweeper
mine sweeper clone made in go using <a href='https://github.com/gen2brain/raylib-go'>raylib-go</a> (go binding of <a href='https://github.com/raysan5/raylib'>raylib</a>)

# build
Download or clone repository 
```sh
git clone https://github.com/rzaf/mineSweeper.git

cd mineSweeper
go get -u github.com/gen2brain/raylib-go/raylib
```

```sh
go run .

go build -v -o minesweeper
./minesweeper
```
Binary should be in same directory as resources directory

# flags
### -b int
chance of a cell being bomb (minimum=0,maximum=100) (default 20)
### -w int
number of horizontal cells (minimum=5,maximum=100) (default 15)
### -g int
number of vertical cells (minimum=5,maximum=100) (default 20)

