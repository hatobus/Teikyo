package models

type Landmark struct {
	EyeRight Righteye
	EyeLeft  Lefteye
}

type Righteye struct {
	TopX    float64
	TopY    float64
	BottomX float64
	BottomY float64
}

type Lefteye struct {
	TopX    float64
	TopY    float64
	BottomX float64
	BottomY float64
}
