package teams

const (
	TypeTextBlock     Type = "TextBlock"
	TypeImage         Type = "Image"
	TypeMedia         Type = "Media"
	TypeRichTextBlock Type = "RichTextBlock"
	TypeTextRun       Type = "TextRun"
)

// Displays text, allowing control over font sizes, weight, and color
//
// Source: https://adaptivecards.io/explorer/TextBlock.html
type TextBlock struct {
	// Must be  TypeTextBlock ("TextBlock")
	Type Type `json:"type"`
	// Text to display. A subset of markdown is supported (https://aka.ms/ACTextFeatures)
	Text string `json:"text"`
	// Controls the color of TextBlock elements
	Color Colors `json:"color,omitempty"`
	// Type of font to use for rendering
	FontType FontType `json:"fontType,omitempty"`
	// Controls the horizontal text alignment. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	// If true, displays text slightly toned down to appear less prominent
	IsSubtle bool `json:"isSubtle,omitempty"`
	// Specifies the maximum number of lines to display
	MaxLines int `json:"maxLines,omitempty"`
	// Controls size of text
	Size FontSize `json:"size,omitempty"`
	// Controls the weight of TextBlock elements
	Weight FontWeight `json:"weight,omitempty"`
	// If true, allow text to wrap. Otherwise, text is clipped
	Wrap bool `json:"wrap,omitempty"`
	// The style of this TextBlock for accessibility purposes
	Style TextBlockStyle `json:"style,omitempty"`
	// Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// A unique identifier associated with the item
	Id string `json:"id,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

func NewTextBlock(text string) *TextBlock {
	return &TextBlock{
		Type: TypeTextBlock,
		Text: text,
	}
}

// Displays an image. Acceptable formats are PNG, JPEG, and GIF
//
// Source: https://adaptivecards.io/explorer/Image.html
type Image struct {
	// Must be  TypeImage ("Image")
	Type Type `json:"type"`
	// The URL to the image. Supports data URI in version 1.2+
	Url string `json:"url"`
	// Alternate text describing the image
	AltText string `json:"altText,omitempty"`
	// Applies a background to a transparent image. This property will respect the image style
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// The desired height of the image. If specified as a pixel value, ending in ‘px’, E.g., 50px, the image will distort to fit that exact height. This overrides the size property
	Height BlockElementHeight `json:"height,omitempty"`
	// Controls how this element is horizontally positioned within its parent. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	// An Action that will be invoked when the Image is tapped or selected. Action.ShowCard is not supported
	SelectAction ISelectAction `json:"selectAction,omitempty"`
	// Controls the approximate size of the image. The physical dimensions will vary per host
	Size ImageSize `json:"size,omitempty"`
	// Controls how this Image is displayed
	Style ImageStyle `json:"style,omitempty"`
	// The desired on-screen width of the image, ending in ‘px’. E.g., 50px. This overrides the size property
	Width string `json:"width,omitempty"`
	// Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// A unique identifier associated with the item
	Id string `json:"id,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Displays a media player for audio or video content
//
// Source: https://adaptivecards.io/explorer/Media.html
type Media struct {
	// Must be  TypeMedia ("Media")
	Type Type `json:"type"`
	// Array of media sources to attempt to play
	Sources []MediaSource `json:"sources"`
	// URL of an image to display before playing. Supports data URI in version 1.2+
	Poster string `json:"poster,omitempty"`
	// Alternate text describing the audio or video.
	AltText string `json:"altText,omitempty"`
	// Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// A unique identifier associated with the item
	Id string `json:"id,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Defines a source for a Media element
//
// Source: https://adaptivecards.io/explorer/MediaSource.html
type MediaSource struct {
	// Mime type of associated media (e.g. "video/mp4")
	MimeType string `json:"mimeType"`
	// URL to media. Supports data URI in version 1.2+
	Url string `json:"url"`
}

// Defines an array of inlines, allowing for inline text formatting
//
// Source: https://adaptivecards.io/explorer/RichTextBlock.html
type RichTextBlock struct {
	// Must be  TypeRichTextBlock ("RichTextBlock")
	Type Type `json:"type"`
	// The array of inlines
	Inlines []TextRun `json:"inlines"`
	// Specifies the height of the element
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment"`
	// Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// A unique identifier associated with the item
	Id string `json:"id,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Defines a single run of formatted text. A TextRun with no properties set can be represented in the json as string containing the text as a shorthand for the json object. These two representations are equivalent
//
// Source: https://adaptivecards.io/explorer/TextRun.html
type TextRun struct {
	// Must be  TypeTextRun ("TextRun")
	Type Type `json:"type"`
	// Text to display. Markdown is not supported
	Text string `json:"text"`
	// Controls the color of the text
	Color Colors `json:"color,omitempty"`
	// The type of font to use
	FontType FontType `json:"fontType,omitempty"`
	// If true, displays the text highlighted
	Highlight bool `json:"highlight,omitempty"`
	// If true, displays text slightly toned down to appear less prominent
	IsSubtle bool `json:"isSubtle,omitempty"`
	// If true, displays the text using italic font
	Italic bool `json:"italic,omitempty"`
	// Action to invoke when this text run is clicked. Visually changes the text run into a hyperlink. Action.ShowCard is not supported
	SelectAction ISelectAction `json:"selectAction,omitempty"`
	// Controls size of text
	Size FontSize `json:"size,omitempty"`
	// If true, displays the text with strikethrough
	Strikethrough bool `json:"strikethrough,omitempty"`
	// If true, displays the text with an underline
	Underline bool `json:"underline,omitempty"`
	// Controls the weight of the text
	Weight FontWeight `json:"weight,omitempty"`
}
