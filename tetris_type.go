package main

// Tetromino type
type Tetromino string

const (
	// I var
	I Tetromino = "I"
	// O var
	O = "O"
	// T var
	T = "T"
	// S var
	S = "S"
	// Z var
	Z = "Z"
	// J var
	J = "J"
	// L var
	L = "L"
)

// Command type
type Command string

const (
	// LEFT var
	LEFT Command = "L"
	// RIGHT var
	RIGHT = "R"
	// ROTATE var
	ROTATE = "RO"
	// DOWN var
	DOWN = "D"
)

// Orientation type
type Orientation string

const (
	// POINTUP var
	POINTUP Orientation = "POINT_UP"
	// POINTRIGHT var
	POINTRIGHT = "POINT_RIGHT"
	// POINTDOWN var
	POINTDOWN = "POINT_DOWN"
	// POINTLEFT var
	POINTLEFT = "POINT_LEFT"
)

// Brick type
type Brick struct {
	tetromino   Tetromino
	x, y        int
	orientation Orientation
}

// Map type
type Map string

// MoveLeft func
func (b *Brick) MoveLeft() {
	b.x--
}

// MoveRight func
func (b *Brick) MoveRight() {
	b.x++
}

// MoveDown func
func (b *Brick) MoveDown() {
	b.y++
}

// Rotate func
func (b *Brick) Rotate() {
	switch b.orientation {
	case POINTUP:
		b.orientation = POINTRIGHT
	case POINTRIGHT:
		b.orientation = POINTDOWN
	case POINTDOWN:
		b.orientation = POINTLEFT
	case POINTLEFT:
		b.orientation = POINTUP
	}
}

// RunCommand func
func RunCommand(brick *Brick, command Command) {
	switch command {
	case LEFT:
		brick.MoveLeft()
	case RIGHT:
		brick.MoveRight()
	case ROTATE:
		brick.Rotate()
	case DOWN:
		brick.MoveDown()
	}
}

// ValidateCommand func
func ValidateCommand(m Map, brick *Brick, command Command) bool {
	return false
}
