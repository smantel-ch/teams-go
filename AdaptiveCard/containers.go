package teams

const (
	TypeActionSet Type = "ActionSet"
	TypeContainer Type = "Container"
	TypeColumnSet Type = "ColumnSet"
	TypeFactSet   Type = "FactSet"
	TypeImageSet  Type = "ImageSet"
)

type ContainerStyle string

const (
	ContainerStyleDefault   ContainerStyle = "default"
	ContainerStyleEmphasis  ContainerStyle = "emphasis"
	ContainerStyleGood      ContainerStyle = "good"
	ContainerStyleAttention ContainerStyle = "attention"
	ContainerStyleWarning   ContainerStyle = "warning"
	ContainerStyleAccent    ContainerStyle = "accent"
)

// Displays a set of actions
//
// Source: https://adaptivecards.io/explorer/ActionSet.html
type ActionSet struct {
	// Must be  TypeActionSet ("ActionSet")
	Type Type `json:"type"`
	// The array of Action elements to show
	Actions []Action `json:"actions"`
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

// Containers group items together
//
// Source: https://adaptivecards.io/explorer/Container.html
type Container struct {
	// Must be  TypeContainer ("Container")
	Type Type `json:"type"`
	// The card elements to render inside the Container
	Items []Element `json:"items"`
	// An Action that will be invoked when the Container is tapped or selected. Action.ShowCard is not supported
	SelectAction ISelectAction `json:"selectAction,omitempty"`
	// Style hint for Container
	Style ContainerStyle `json:"style,omitempty"`
	// Defines how the content should be aligned vertically within the container. When not specified, the value of verticalContentAlignment is inherited from the parent container. If no parent container has verticalContentAlignment set, it defaults to Top
	VerticalContentAlignment VerticalContentAlignment `json:"verticalContentAlignment,omitempty"`
	// Determines whether the element should bleed through its parent’s padding.
	Bleed bool `json:"bleed,omitempty"`
	// Specifies the background image. Acceptable formats are PNG, JPEG, and GIF
	BackgroundImage BackgroundImage `json:"backgroundImage,omitempty"`
	// Specifies the minimum height of the container in pixels, like "80px"
	MinHeight string `json:"minHeight,omitempty"`
	// When true content in this container should be presented right to left. When ‘false’ content in this container should be presented left to right. When unset layout direction will inherit from parent container or column. If unset in all ancestors, the default platform behavior will apply
	Rtl bool `json:"rtl,omitempty"`
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

// ColumnSet divides a region into Columns, allowing elements to sit side-by-side
//
// Source: https://adaptivecards.io/explorer/ColumnSet.html
type ColumnSet struct {
	// Must be  TypeColumnSet ("ColumnSet")
	Type Type `json:"type"`
	//The array of Columns to divide the region into
	Columns []Column `json:"columns,omitempty"`
	// An Action that will be invoked when the Container is tapped or selected. Action.ShowCard is not supported
	SelectAction ISelectAction `json:"selectAction,omitempty"`
	// Style hint for Container
	Style ContainerStyle `json:"style,omitempty"`
	// Determines whether the element should bleed through its parent’s padding
	Bleed bool `json:"bleed,omitempty"`
	// Specifies the minimum height of the container in pixels, like "80px"
	MinHeight string `json:"minHeight,omitempty"`
	// Controls the horizontal alignment of the ColumnSet. When not specified, the value of horizontalAlignment is inherited from the parent container. If no parent container has horizontalAlignment set, it defaults to Left
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment"`
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

// Defines a container that is part of a ColumnSet
//
// Source: https://adaptivecards.io/explorer/Column.html
type Column struct {
	// The card elements to render inside the Column
	Items []Element `json:"items,omitempty"`
	// Specifies the background image. Acceptable formats are PNG, JPEG, and GIF
	BackgroundImage BackgroundImage `json:"backgroundImage,omitempty"`
	// Determines whether the element should bleed through its parent’s padding
	Bleed bool `json:"bleed,omitempty"`
	// Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the minimum height of the container in pixels, like "80px"
	MinHeight string `json:"minHeight,omitempty"`
	// When true content in this container should be presented right to left. When ‘false’ content in this container should be presented left to right. When unset layout direction will inherit from parent container or column. If unset in all ancestors, the default platform behavior will apply
	Rtl bool `json:"rtl,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// An Action that will be invoked when the Container is tapped or selected. Action.ShowCard is not supported
	SelectAction ISelectAction `json:"selectAction,omitempty"`
	// Style hint for Container
	Style ContainerStyle `json:"style,omitempty"`
	// Defines how the content should be aligned vertically within the container. When not specified, the value of verticalContentAlignment is inherited from the parent container. If no parent container has verticalContentAlignment set, it defaults to Top
	VerticalContentAlignment VerticalContentAlignment `json:"verticalContentAlignment,omitempty"`
	// "auto", "stretch", a number representing relative width of the column in the column group, or in version 1.1 and higher, a specific pixel width, like "50px"
	Width interface{} `json:"width"`
	// A unique identifier associated with the item
	Id string `json:"id,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// The FactSet element displays a series of facts (i.e. name/value pairs) in a tabular form
//
// Source: https://adaptivecards.io/explorer/FactSet.html
type FactSet struct {
	// Must be  TypeFactSet ("FactSet")
	Type Type `json:"type"`
	// The array of Fact‘s
	Facts []Fact `json:"facts"`
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

// Describes a Fact in a FactSet as a key/value pair
//
// Source: https://adaptivecards.io/explorer/Fact.html
type Fact struct {
	// The title of the fact
	Title string `json:"title"`
	// The value of the fact
	Value string `json:"value"`
}

// The ImageSet displays a collection of Images similar to a gallery. Acceptable formats are PNG, JPEG, and GIF
//
// Source: https://adaptivecards.io/explorer/ImageSet.html
type ImageSet struct {
	// Must be  TypeImageSet ("ImageSet")
	Type Type `json:"type"`
	// The array of Image elements to show
	Images []Image `json:"images"`
	// Controls the approximate size of each image. The physical dimensions will vary per host. Auto and stretch are not supported for ImageSet. The size will default to medium if those values are set.
	ImageSize ImageSize `json:"imageSize"`
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

//
//
// Source:
type Table struct{}

//
//
// Source:
type TableCell struct{}
