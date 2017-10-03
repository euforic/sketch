//go:generate stringer -type=ResizingType,LayerListExpandedType,BorderPosition,BorderLineCapStyle,BorderLineJoinStyle,FillType,PatternFillType,BlendMode,LineDecorationType,BooleanOperationType,CurveMode -output constants_strings.go constants.go
package sketch

type ResizingType int64

const (
	ResizingType_Stretch ResizingType = iota
	ResizingType_PinToCorner
	ResizingType_ResizeObject
	ResizingType_FloatInPlace
)

type LayerListExpandedType int64

const (
	LayerListExpandedType_Collapsed LayerListExpandedType = iota
	LayerListExpandedType_ExpandedTemp
	LayerListExpandedType_Expanded
)

type BorderPosition int64 // 0 | 1 | 2 | 3

const (
	BorderPosition_Center BorderPosition = iota
	BorderPosition_Inside
	BorderPosition_Outside
	BorderPosition_Both
)

type BorderLineCapStyle int64 // 0 | 1 | 2

const (
	BorderLineCapStyle_Butt BorderLineCapStyle = iota
	BorderLineCapStyle_Round
	BorderLineCapStyle_Square
)

type BorderLineJoinStyle int64 // 0 | 1 | 2

const (
	BorderLineJoinStyle_Miter BorderLineJoinStyle = iota
	BorderLineJoinStyle_Round
	BorderLineJoinStyle_Bevel
)

type FillType int64 // 0 | 1 | 4 | 5

const (
	FillType_Solid FillType = iota
	FillType_Gradient
	FillType_Pattern
	FillType_Noise
)

type PatternFillType int64 // 0 | 1 | 2 | 3

const (
	PatternFillType_Tile PatternFillType = iota
	PatternFillType_Fill
	PatternFillType_Stretch
	PatternFillType_Fit
)

// TODO: define correct blendmode enum
type BlendMode int64 // 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15

const (
	BlendMode_0 BlendMode = iota
	BlendMode_1
	BlendMode_2
	BlendMode_3
	BlendMode_4
	BlendMode_5
	BlendMode_6
	BlendMode_7
	BlendMode_8
	BlendMode_9
	BlendMode_10
	BlendMode_11
	BlendMode_12
	BlendMode_13
	BlendMode_14
	BlendMode_15
)

type LineDecorationType int64 // 0 | 1 | 2 | 3

const (
	LineDecorationType_None LineDecorationType = iota
	LineDecorationType_OpenArrow
	LineDecorationType_ClosedArrow
	LineDecorationType_Bar
)

type BooleanOperationType int64 // -1 | 0 | 1 | 2 | 3

const (
	BooleanOperation_None  BooleanOperationType = -1
	BooleanOperation_Union BooleanOperationType = iota
	BooleanOperation_Subtract
	BooleanOperation_Intersect
	BooleanOperation_Difference
)

type CurveMode int64 // 0 | 1 | 2 | 3 | 4

const (
	CurveMode_None CurveMode = iota
	CurveMode_Straight
	CurveMode_Mirrored
	CurveMode_Disconnected
	CurveMode_Asymmetric
)
