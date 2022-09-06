package teams

type Element interface {
	IsElement() bool
}

func (a *ActionSet) IsElement() bool {
	return true
}
func (c *ColumnSet) IsElement() bool {
	return true
}
func (c *Container) IsElement() bool {
	return true
}
func (f *FactSet) IsElement() bool {
	return true
}
func (i *Image) IsElement() bool {
	return true
}
func (i *ImageSet) IsElement() bool {
	return true
}
func (i *InputChoiceSet) IsElement() bool {
	return true
}
func (i *InputDate) IsElement() bool {
	return true
}
func (i *InputNumber) IsElement() bool {
	return true
}
func (i *InputText) IsElement() bool {
	return true
}
func (i *InputTime) IsElement() bool {
	return true
}
func (i *InputToggle) IsElement() bool {
	return true
}
func (m *Media) IsElement() bool {
	return true
}
func (m *RichTextBlock) IsElement() bool {
	return true
}
func (t *Table) IsElement() bool {
	return true
}
func (t *TextBlock) IsElement() bool {
	return true
}

type Version string

const (
	Version10 Version = "1.0"
	Version11 Version = "1.1"
	Version12 Version = "1.2"
	Version13 Version = "1.3"
	Version14 Version = "1.4"
	Version15 Version = "1.5"
)

type Schema string

const (
	SchemaDefault Schema = "http://adaptivecards.io/schemas/adaptive-card.json"
)

type Type string

type ImageFillMode string

const (
	ImageFillModeCover              ImageFillMode = "cover"
	ImageFillModeRepeatHorizontally ImageFillMode = "repeatHorizontally"
	ImageFillModeRepeatVertically   ImageFillMode = "repeatVertically"
	ImageFillModeRepeat             ImageFillMode = "repeat"
)

type HorizontalAlignment string

const (
	HorizontalAlignmentLeft   HorizontalAlignment = "left"
	HorizontalAlignmentCenter HorizontalAlignment = "center"
	HorizontalAlignmentRight  HorizontalAlignment = "right"
)

type VerticalAlignment string

const (
	VerticalAlignmentTop    VerticalAlignment = "top"
	VerticalAlignmentCenter VerticalAlignment = "center"
	VerticalAlignmentBottom VerticalAlignment = "bottom"
)

type AssociatedInputs string

const (
	AssociatedInputAuto AssociatedInputs = "Auto"
	AssociatedInputNone AssociatedInputs = "None"
)

type BlockElementHeight string

const (
	BlockElementHeightAuto    BlockElementHeight = "auto"
	BlockElementHeightStretch BlockElementHeight = "stretch"
)

type Spacing string

const (
	SpacingDefault    Spacing = "default"
	SpacingNone       Spacing = "none"
	SpacingSmall      Spacing = "small"
	SpacingMedium     Spacing = "medium"
	SpacingLarge      Spacing = "large"
	SpacingExtraLarge Spacing = "extraLarge"
	SpacingPadding    Spacing = "padding"
)

type VerticalContentAlignment string

const (
	VerticalContentAlignmentTop    VerticalContentAlignment = "top"
	VerticalContentAlignmentCenter VerticalContentAlignment = "center"
	VerticalContentAlignmentBottom VerticalContentAlignment = "bottom"
)

type ImageSize string

const (
	ImageSizeAuto    ImageSize = "auto"
	ImageSizeStretch ImageSize = "stretch"
	ImageSizeSmall   ImageSize = "small"
	ImageSizeMedium  ImageSize = "medium"
	ImageSizeLarge   ImageSize = "large"
)

type ImageStyle string

const (
	ImageStyleDefault ImageStyle = "default"
	ImageStylePerson  ImageStyle = "person"
)

type Colors string

const (
	ColorDefault   Colors = "default"
	ColorDark      Colors = "dark"
	ColorLight     Colors = "light"
	ColorAccent    Colors = "accent"
	ColorGood      Colors = "good"
	ColorsWarning  Colors = "warning"
	ColorAttention Colors = "attention"
)

type FontType string

const (
	FontTypeDefault   FontType = "default"
	FontTypeMonospace FontType = "monospace"
)

type FontSize string

const (
	FontSizeDefault    FontSize = "default"
	FontSizeSmall      FontSize = "small"
	FontSizeMedium     FontSize = "medium"
	FontSizeLarge      FontSize = "large"
	FontSizeExtraLarge FontSize = "extraLarge"
)

type FontWeight string

const (
	FontWeightDefault FontWeight = "default"
	FontWeightLighter FontWeight = "lighter"
	FontWeightBolder  FontWeight = "bolder"
)

type TextBlockStyle string

const (
	TextBlockStyleDefault TextBlockStyle = "default"
	TextBlockStyleHeading TextBlockStyle = "heading"
)
