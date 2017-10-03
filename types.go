package sketch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"howett.net/plist"
)

type Document struct {
	Class                  string                    `json:"_class"`
	DoObjectID             string                    `json:"do_objectID"`
	Assets                 *AssetsCollection         `json:"assets"`
	CurrentPageIndex       json.Number               `json:"currentPageIndex"`
	EnableLayerInteraction bool                      `json:"enableLayerInteraction"`
	EnableSliceInteraction bool                      `json:"enableSliceInteraction"`
	ForeignSymbols         []interface{}             `json:"foreignSymbols"`
	LayerStyles            *SharedStyleContainer     `json:"layerStyles"`
	LayerSymbols           *SharedSymbolContainer    `json:"layerSymbols"`
	LayerTextStyles        *SharedTextStyleContainer `json:"layerTextStyles"`
	Pages                  []*MSJSONFileReference    `json:"pages"`
}

type Page struct {
	Class                 string         `json:"_class"`
	DoObjectID            string         `json:"do_objectID"`
	ExportOptions         *ExportOptions `json:"exportOptions"`
	Frame                 *Rect          `json:"frame"`
	HasClickThrough       bool           `json:"hasClickThrough"`
	HorizontalRulerData   *RulerData     `json:"horizontalRulerData"`
	IncludeInCloudUpload  bool           `json:"includeInCloudUpload"`
	IsFlippedHorizontal   bool           `json:"isFlippedHorizontal"`
	IsFlippedVertical     bool           `json:"isFlippedVertical"`
	IsLocked              bool           `json:"isLocked"`
	IsVisible             bool           `json:"isVisible"`
	LayerListExpandedType json.Number    `json:"layerListExpandedType"`
	Layers                []*Layer       `json:"layers"`
	Name                  string         `json:"name"`
	NameIsFixed           bool           `json:"nameIsFixed"`
	ResizingType          int64          `json:"resizingType"`
	Rotation              json.Number    `json:"rotation"`
	ShouldBreakMaskChain  bool           `json:"shouldBreakMaskChain"`
	Style                 *Style         `json:"style"`
	VerticalRulerData     *RulerData     `json:"verticalRulerData"`
}

type Layer struct {
	DoObjectID                        string                     `json:"do_objectID"`
	Name                              string                     `json:"name"`
	Class                             string                     `json:"_class"`
	Frame                             Rect                       `json:"frame"`
	IsFlippedHorizontal               bool                       `json:"isFlippedHorizontal"`
	IsFlippedVertical                 bool                       `json:"isFlippedVertical"`
	IsLocked                          bool                       `json:"isLocked"`
	IsVisible                         bool                       `json:"isVisible"`
	ResizesContent                    bool                       `json:"resizesContent"`
	ResizingConstraint                json.Number                `json:"resizingConstraint"`
	ResizingType                      int64                      `json:"resizingType"`
	Radius                            json.Number                `json:"radius"`
	Rotation                          json.Number                `json:"rotation"`
	ShouldBreakMaskChain              bool                       `json:"shouldBreakMaskChain"`
	Style                             *Style                     `json:"style"`
	Layers                            []*Layer                   `json:"layers"`
	AttributedString                  *MSAttributedString        `json:"attributedString"`
	AutomaticallyDrawOnUnderlyingPath bool                       `json:"automaticallyDrawOnUnderlyingPath"`
	BackgroundColor                   *Color                     `json:"backgroundColor"`
	BooleanOperation                  int64                      `json:"booleanOperation"`
	ClippingMask                      *NestedPositionCoordinates `json:"clippingMask"`
	ClippingMaskMode                  json.Number                `json:"clippingMaskMode"`
	DontSynchroniseWithSymbol         bool                       `json:"dontSynchroniseWithSymbol"`
	Edited                            bool                       `json:"edited"`
	ExportOptions                     *ExportOptions             `json:"exportOptions"`
	FillReplacesImage                 bool                       `json:"fillReplacesImage"`
	FixedRadius                       json.Number                `json:"fixedRadius"`
	GlyphBounds                       *NestedPositionCoordinates `json:"glyphBounds"`
	HasBackgroundColor                bool                       `json:"hasBackgroundColor"`
	HasClickThrough                   bool                       `json:"hasClickThrough"`
	HasClippingMask                   bool                       `json:"hasClippingMask"`
	HasConvertedToNewRoundCorners     bool                       `json:"hasConvertedToNewRoundCorners"`
	HeightIsClipped                   bool                       `json:"heightIsClipped"`
	HorizontalRulerData               *RulerData                 `json:"horizontalRulerData"`
	HorizontalSpacing                 json.Number                `json:"horizontalSpacing"`
	Image                             *MSJSONFileReference       `json:"image"`
	IncludeBackgroundColorInExport    bool                       `json:"includeBackgroundColorInExport"`
	IncludeBackgroundColorInInstance  bool                       `json:"includeBackgroundColorInInstance"`
	IncludeInCloudUpload              bool                       `json:"includeInCloudUpload"`
	IsEquilateral                     bool                       `json:"isEquilateral"`
	Layout                            *LayoutGrid                `json:"layout"`
	LayerListExpandedType             json.Number                `json:"layerListExpandedType"`
	LineSpacingBehaviour              json.Number                `json:"lineSpacingBehaviour"`
	MasterInfluenceEdgeMaxXPadding    json.Number                `json:"masterInfluenceEdgeMaxXPadding"`
	MasterInfluenceEdgeMaxYPadding    json.Number                `json:"masterInfluenceEdgeMaxYPadding"`
	MasterInfluenceEdgeMinXPadding    json.Number                `json:"masterInfluenceEdgeMinXPadding"`
	MasterInfluenceEdgeMinYPadding    json.Number                `json:"masterInfluenceEdgeMinYPadding"`
	NumberOfPoints                    int64                      `json:"numberOfPoints"`
	NameIsFixed                       bool                       `json:"nameIsFixed"`
	NineSliceCenterRect               *NestedPositionCoordinates `json:"nineSliceCenterRect"`
	NineSliceScale                    *PositionCoordinates       `json:"nineSliceScale"`
	OriginalObjectID                  string                     `json:"originalObjectID"`
	Overrides                         *Overrides                 `json:"overrides"`
	Path                              *Path                      `json:"path"`
	SymbolID                          string                     `json:"symbolID"`
	TextBehaviour                     json.Number                `json:"textBehaviour,omitempty"`
	VerticalRulerData                 *RulerData                 `json:"verticalRulerData,omitempty"`
	VerticalSpacing                   json.Number                `json:"verticalSpacing"`
	WindingRule                       json.Number                `json:"windingRule"`
}

type Meta struct {
	Commit           string            `json:"commit"`
	AppVersion       string            `json:"appVersion"`
	Build            json.Number       `json:"build"`
	App              string            `json:"app"`
	PagesAndArtboard map[string]string `json:"pagesAndArtboards"`
	Fonts            []string          `json:"fonts"`
	Version          json.Number       `json:"version"`
	SaveHistory      []string          `json:"saveHistory"`
	Autosaved        json.Number       `json:"autosaved"`
	Variant          string            `json:"variant"`
}

type LayoutGrid struct {
	Class                   string      `json:"_class"`
	DoObjectID              string      `json:"do_objectID"`
	ColumnWidth             json.Number `json:"columnWidth"`
	DrawHorizontal          bool        `json:"drawHorizontal"`
	DrawHorizontalLines     bool        `json:"drawHorizontalLines"`
	DrawVertical            bool        `json:"drawVertical"`
	GutterHeight            json.Number `json:"gutterHeight"`
	GutterWidth             json.Number `json:"gutterWidth"`
	GutterOutside           bool        `json:"cutterOutside"`
	HorizontalOffset        json.Number `json:"horizontalOffset"`
	IsEnabled               bool        `json:"isEnabled"`
	NumberOfColumns         int64       `json:"numberOfColumns"`
	RowHeightMultiplication json.Number `json:"rowHeightMultiplication"`
	TotalWidth              json.Number `json:"totalWidth"`
}

type ExportFormat struct {
	Class            string      `json:"_class"`
	DoObjectID       string      `json:"do_objectID"`
	AbsoluteSize     json.Number `json:"absoluteSize"`
	FileFormat       string      `json:"fileFormat"`
	Name             string      `json:"name"`
	NamingScheme     json.Number `json:"namingScheme"`
	Scale            json.Number `json:"scale"`
	VisibleScaleType json.Number `json:"visibleScaleType"`
}

type Color struct {
	Class string      `json:"_class"`
	Alpha json.Number `json:"alpha"`
	Blue  json.Number `json:"blue"`
	Green json.Number `json:"green"`
	Red   json.Number `json:"red"`
}

type RulerData struct {
	Class  string        `json:"_class"`
	Base   json.Number   `json:"base"`
	Guides []json.Number `json:"guides"`
}

type GraphicContextSettings struct {
	Class     string      `json:"_class"`
	BlendMode int64       `json:"blendMode"`
	Opacity   json.Number `json:"opacity"`
}

type Shadow struct {
	Class           string                  `json:"_class"`
	BlurRadius      json.Number             `json:"blurRadius"`
	DoObjectID      string                  `json:"do_objectID"`
	Color           *Color                  `json:"color"`
	ContextSettings *GraphicContextSettings `json:"contextSettings"`
	IsEnabled       bool                    `json:"isEnabled"`
	OffsetX         json.Number             `json:"offsetX"`
	OffsetY         json.Number             `json:"offsetY"`
	Spread          json.Number             `json:"spread"`
}

type InnerShadow struct {
	Shadow
}

type ColorControls struct {
	Class      string      `json:"_class"`
	Brightness json.Number `json:"brightness"`
	Contrast   json.Number `json:"contrast"`
	Hue        json.Number `json:"hue"`
	IsEnabled  bool        `json:"isEnabled"`
	Saturation json.Number `json:"saturation"`
}

type Fill struct {
	Class            string      `json:"_class"`
	DoObjectID       string      `json:"do_objectID"`
	Color            *Color      `json:"color"`
	FillType         int64       `json:"fillType"`
	Gradient         *Gradient   `json:"gradient"`
	IsEnabled        bool        `json:"isEnabled"`
	NoiseIndex       json.Number `json:"noiseIndex"`
	NoiseIntensity   json.Number `json:"noiseIntensity"`
	PatternFillType  int64       `json:"patternFillType"`
	PatternTileScale json.Number `json:"patternTileScale"`
}

type Border struct {
	Class     string      `json:"_class"`
	Color     Color       `json:"color"`
	FillType  json.Number `json:"fillType"`
	IsEnabled bool        `json:"isEnabled"`
	Position  int         `json:"position"`
	Thickness json.Number `json:"thickness"`
}

type GradientStop struct {
	Class      string      `json:"_class"`
	DoObjectID string      `json:"do_objectID"`
	Color      Color       `json:"color"`
	Position   json.Number `json:"position"`
}

type Rect struct {
	Class                string      `json:"_class"`
	DoObjectID           string      `json:"do_objectID"`
	ConstrainProportions bool        `json:"constrainProportions"`
	Height               json.Number `json:"height"`
	Width                json.Number `json:"width"`
	X                    json.Number `json:"x"`
	Y                    json.Number `json:"y"`
}

type CurvePoint struct {
	Class        string               `json:"_class"`
	CornerRadius json.Number          `json:"cornerRadius"`
	CurveFrom    *PositionCoordinates `json:"curveFrom"`
	CurveMode    int64                `json:"curveMode"`
	CurveTo      PositionCoordinates  `json:"curveTo"`
	HasCurveFrom bool                 `json:"hasCurveFrom"`
	HasCurveTo   bool                 `json:"hasCurveTo"`
	Point        *PositionCoordinates `json:"point"`
}

type BorderOptions struct {
	Class         string  `json:"_class"`
	DoObjectID    string  `json:"do_objectID"`
	DashPattern   []int64 `json:"dashPattern"`
	IsEnabled     bool    `json:"isEnabled"`
	LineCapStyle  int64   `json:"lineCapStyle"`
	LineJoinStyle int64   `json:"lineJoinStyle"`
}

type Gradient struct {
	Class                 string               `json:"_class"`
	DoObjectID            string               `json:"do_objectID"`
	ElipseLength          json.Number          `json:"elipseLength"`
	From                  *PositionCoordinates `json:"from"`
	GradientType          json.Number          `json:"gradientType"`
	ShouldSmoothenOpacity bool                 `json:"shouldSmoothenOpacity"`
	Stops                 []*GradientStop      `json:"stops"`
	To                    *PositionCoordinates `json:"to"`
}

type TextStyle struct {
	Class             string             `json:"_class"`
	EncodedAttributes *EncodedAttributes `json:"encodedAttributes"`
	VerticalAlignment json.Number        `json:"verticalAlignment"`
}

type Style struct {
	Class               string                  `json:"_class,omitempty"`
	DoObjectID          string                  `json:"do_objectID,omitempty"`
	Blur                *Blur                   `json:"blur,omitempty"`
	Borders             []*Border               `json:"borders,omitempty"`
	BorderOptions       *BorderOptions          `json:"borderOptions,omitempty"`
	ContextSettings     *GraphicContextSettings `json:"contextSettings,omitempty"`
	ColorControls       *ColorControls          `json:"colorControls,omitempty"`
	StartDecorationType int64                   `json:"startDecorationType,omitempty"`
	EndDecorationType   int64                   `json:"endDecorationType,omitempty"`
	Fills               []*Fill                 `json:"fills,omitempty"`
	InnerShadows        []*InnerShadow          `json:"innerShadows,omitempty"`
	MiterLimit          json.Number             `json:"miterLimit,omitempty"`
	Shadows             []*Shadow               `json:"shadows,omitempty"`
	SharedObjectID      string                  `json:"sharedObjectID,omitempty"`
	TextStyle           *TextStyle              `json:"textStyle,omitempty"`
}

type SharedStyle struct {
	Class      string `json:"_class"`
	DoObjectID string `json:"do_objectID"`
	Name       string `json:"name"`
	Value      *Style `json:"value"`
}

type SharedStyleContainer struct {
	Class   string         `json:"_class"`
	Objects []*SharedStyle `json:"object"`
}

type SharedSymbolContainer struct {
	Class   string        `json:"_class"`
	Objects []interface{} `json:"object"`
}

type SharedTextStyleContainer struct {
	Class   string         `json:"_class"`
	Objects []*SharedStyle `json:"objects"`
}

type SharedAssetsCollection struct {
	Class           string          `json:"_class"`
	Colors          []*Color        `json:"colors"`
	Gradients       []*Gradient     `json:"gradients"`
	ImageCollection ImageCollection `json:"imageCollection"`
	Images          []Image         `json:"images"`
}

type Image interface{}

type Blur struct {
	Class       string               `json:"_class"`
	Center      *PositionCoordinates `json:"center"`
	IsEnabled   bool                 `json:"isEnabled"`
	MotionAngle json.Number          `json:"motionAngle"`
	Radius      json.Number          `json:"radius"`
	Type        json.Number          `json:"type"`
}

type ExportOptions struct {
	Class            string          `json:"_class"`
	ExportFormats    []*ExportFormat `json:"exportFormats"`
	IncludedLayerIds []string        `json:"includedLayerIds"`
	LayerOptions     json.Number     `json:"layerOptions"`
	ShouldTrim       bool            `json:"shouldTrim"`
}

type Path struct {
	Class                string        `json:"_class"`
	PointRadiusBehaviour int64         `json:"pointRadiusBehaviour"`
	IsClosed             bool          `json:"isClosed"`
	Points               []*CurvePoint `json:"points"`
}

type EncodedAttributes struct {
	NSKern                                   json.Number               `json:"NSKern,omitempty"`
	NSStrokeWidth                            json.Number               `json:"NSStrokeWidth,omitempty"`
	NSStrokeColor                            *ArchivedAttributedString `json:"NSStrokeColor,omitempty"`
	NSStrikethrough                          json.Number               `json:"NSStrikethrough,omitempty"`
	NSUnderline                              json.Number               `json:"NSUnderline,omitempty"`
	MSAttributedStringFontAttribute          *ArchivedAttributedString `json:"MSAttributedStringFontAttribute,omitempty"`
	MSAttributedStringTextTransformAttribute json.Number               `json:"MSAttributedStringTextTransformAttribute,omitempty"`
	NSColor                                  *ArchivedAttributedString `json:"NSColor,omitempty"`
	NSParagraphStyle                         *ArchivedAttributedString `json:"NSParagraphStyle,omitempty"`
}

type MSJSONFileReference struct {
	Class    string     `json:"_class"`
	Ref      string     `json:"_ref"`
	RefClass string     `json:"_ref_class"`
	Data     DataString `json:"data"`
	Sha1     DataString `json:"sha1"`
}

type DataString struct {
	Data string `json:"_data"`
}

type MSAttributedString struct {
	Class                    string                   `json:"_class"`
	ArchivedAttributedString ArchivedAttributedString `json:"archivedAttributedString"`
}

type AssetsCollection struct {
	Class           string                 `json:"_class"`
	Colors          []*Color               `json:"colors"`
	Gradients       []*Gradient            `json:"gradients"`
	ImageCollection *ImageCollection       `json:"imageCollection"`
	Images          []*MSJSONFileReference `json:"images"`
}

type ImageCollection struct {
	Class  string               `json:"_class"`
	Images *MSJSONFileReference `json:"images"`
}

type Overrides map[string]interface{}

// Example `{0.5, 0.67135115527602085}`
type PositionCoordinates struct {
	X    json.Number
	Y    json.Number
	data []byte
}

func (p *PositionCoordinates) MarshalJSON() ([]byte, error) {
	// TODO: fix marshal function to properly encode
	out := fmt.Sprintf("{%v, %v}", p.X, p.Y)
	out = strconv.Quote(out)
	return []byte(out), nil
}

func (p *PositionCoordinates) UnmarshalJSON(b []byte) error {
	p.data = b
	b = bytes.Replace(b, []byte("{"), []byte("["), -1)
	b = bytes.Replace(b, []byte("}"), []byte("]"), -1)

	s, err := strconv.Unquote(string(b))
	b = []byte(s)
	if err != nil {
		return err
	}

	data := []json.Number{}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	p.X = data[0]
	p.Y = data[1]

	return nil
}

// Example `{{0.5, 0.67135115527602085}, {0.5, 0.67135115527602085}}`
type NestedPositionCoordinates struct {
	Data []*PositionCoordinates
}

func (n *NestedPositionCoordinates) MarshalJSON() ([]byte, error) {
	// TODO: fix marshal function to properly encode
	out := ""
	for i, p := range n.Data {
		out += fmt.Sprintf("{%v, %v}", p.X, p.Y)
		if i != (len(n.Data) - 1) {
			out += ","
		}
	}
	out = "\"{" + out + "}\""
	return []byte(out), nil
}

func (n *NestedPositionCoordinates) UnmarshalJSON(b []byte) error {
	b = bytes.Replace(b, []byte("{"), []byte("["), -1)
	b = bytes.Replace(b, []byte("}"), []byte("]"), -1)

	s, err := strconv.Unquote(string(b))
	b = []byte(s)
	if err != nil {
		return err
	}

	data := [][]json.Number{}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for _, posAr := range data {
		n.Data = append(n.Data, &PositionCoordinates{X: posAr[0], Y: posAr[1]})
	}
	return nil
}

type ArchivedAttributedString struct {
	Archive Archive `json:"_archive"`
}

type Archive struct {
	Data map[string]interface{}
	data []byte
}

func (n Archive) MarshalJSON() ([]byte, error) {
	// TODO: fix marshal function to properly encode
	out, err := plist.Marshal(n, plist.BinaryFormat)
	if err != nil {
		return out, errors.Wrap(err, "Archive.Marshal.JSON")
	}

	encOut := base64.StdEncoding.EncodeToString([]byte(out))

	return []byte(encOut), nil
}

func (a *Archive) UnmarshalJSON(b []byte) error {
	a.data = b
	ad := map[string]interface{}{}
	a.data = b
	d := []byte{}

	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}

	_, err = plist.Unmarshal(d, &ad)
	if err != nil {
		return err
	}

	a.Data = ad

	return nil
}
