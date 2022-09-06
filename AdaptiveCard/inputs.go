package teams

const (
	TypeInputText      Type = "Input.Text"
	TypeInputNumber    Type = "Input.Number"
	TypeInputDate      Type = "Input.Date"
	TypeInputTime      Type = "Input.Time"
	TypeInputToggle    Type = "Input.Toggle"
	TypeInputChoiceSet Type = "Input.ChoiceSet"
)

type TextInputStyle string

const (
	TextInputStyleText     TextInputStyle = "text"
	TextInputStyleTel      TextInputStyle = "tel"
	TextInputStyleUrl      TextInputStyle = "url"
	TextInputStyleEmail    TextInputStyle = "email"
	TextInputStylePassword TextInputStyle = "password"
)

type ChoiceInputStyle string

const (
	ChoiceInputStyleCompact  ChoiceInputStyle = "compact"
	ChoiceInputStyleExpanded ChoiceInputStyle = "expanded"
	ChoiceInputStyleFiltered ChoiceInputStyle = "filtered"
)

// Lets a user enter text
//
// Source: https://adaptivecards.io/explorer/Input.Text.html
type InputText struct {
	// Must be  TypeInputText ("Input.Text")
	Type Type `json:"type"`
	// Unique identifier for the value. Used to identify collected input when the Submit action is performed
	Id string `json:"id"`
	// If true, allow multiple lines of input
	IsMultiLine bool `json:"isMultiline,omitempty"`
	// Hint of maximum length characters to collect (may be ignored by some clients)
	MaxLength int `json:"maxLength,omitempty"`
	// Description of the input desired. Displayed when no text has been input
	Placeholder string `json:"placeholder,omitempty"`
	// Regular expression indicating the required format of this text input
	Regex string `json:"regex,omitempty"`
	// Style hint for text input
	Style TextInputStyle `json:"style,omitempty"`
	// The inline action for the input. Typically displayed to the right of the input. It is strongly recommended to provide an icon on the action (which will be displayed instead of the title of the action)
	InlineAction ISelectAction `json:"inlineAction,omitempty"`
	// The initial value for this field
	Value string `json:"value,omitempty"`
	// Error message to display when entered input is invalid
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether or not this input is required
	IsRequired bool `json:"isRequireds,omitempty"`
	// 	Label for this input
	Label string `json:"label,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Allows a user to enter a number
//
// Source: https://adaptivecards.io/explorer/Input.Number.html
type InputNumber struct {
	// Must be  TypeInputNumber ("Input.Number")
	Type Type `json:"type"`
	// Unique identifier for the value. Used to identify collected input when the Submit action is performed
	Id string `json:"id"`
	// Hint of maximum value (may be ignored by some clients)
	Max int `json:"max,omitempty"`
	// Hint of minimum value (may be ignored by some clients)
	Min int `json:"min,omitempty"`
	// Description of the input desired. Displayed when no text has been input
	Placeholder string `json:"placeholder,omitempty"`
	// The initial value for this field
	Value int `json:"value,omitempty"`
	// Error message to display when entered input is invalid
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether or not this input is required
	IsRequired bool `json:"isRequireds,omitempty"`
	// 	Label for this input
	Label string `json:"label,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Lets a user choose a date
//
// Source: https://adaptivecards.io/explorer/Input.Date.html
type InputDate struct {
	// Must be  TypeInputDate ("Input.Date")
	Type Type `json:"type"`
	// Unique identifier for the value. Used to identify collected input when the Submit action is performed
	Id string `json:"id"`
	// Hint of maximum value expressed in YYYY-MM-DD(may be ignored by some clients)
	Max string `json:"max,omitempty"`
	// Hint of minimum value expressed in YYYY-MM-DD(may be ignored by some clients)
	Min string `json:"min,omitempty"`
	// Description of the input desired. Displayed when no text has been input
	Placeholder string `json:"placeholder,omitempty"`
	// The initial value for this field expressed in YYYY-MM-DD
	Value string `json:"value,omitempty"`
	// Error message to display when entered input is invalid
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether or not this input is required
	IsRequired bool `json:"isRequireds,omitempty"`
	// 	Label for this input
	Label string `json:"label,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Lets a user select a time
//
// Source: https://adaptivecards.io/explorer/Input.Time.html
type InputTime struct {
	// Must be  TypeInputTime ("Input.Time")
	Type Type `json:"type"`
	// Unique identifier for the value. Used to identify collected input when the Submit action is performed
	Id string `json:"id"`
	// Hint of maximum value expressed in HH:MM (may be ignored by some clients)
	Max string `json:"max,omitempty"`
	// Hint of minimum value expressed in HH:MM (may be ignored by some clients)
	Min string `json:"min,omitempty"`
	// Description of the input desired. Displayed when no text has been input
	Placeholder string `json:"placeholder,omitempty"`
	// The initial value for this field expressed in HH:MM
	Value string `json:"value,omitempty"`
	// Error message to display when entered input is invalid
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether or not this input is required
	IsRequired bool `json:"isRequireds,omitempty"`
	// 	Label for this input
	Label string `json:"label,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Lets a user choose between two options
//
// Source: https://adaptivecards.io/explorer/Input.Toggle.html
type InputToggle struct {
	// Must be  TypeInputToggle ("Input.Toggle")
	Type Type `json:"type"`
	// Title for the toggle
	Title string `json:"title"`
	// Unique identifier for the value. Used to identify collected input when the Submit action is performed
	Id string `json:"id"`
	// The initial selected value. If you want the toggle to be initially on, set this to the value of valueOn‘s value
	Value string `json:"value"`
	// The value when toggle is off
	ValueOff string `json:"valueOff"`
	// The value when toggle is on
	ValueOn string `json:"valueOn"`
	// If true, allow text to wrap. Otherwise, text is clipped
	Wrap bool `json:"wrap"`
	// Error message to display when entered input is invalid
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether or not this input is required
	IsRequired bool `json:"isRequireds,omitempty"`
	// 	Label for this input
	Label string `json:"label,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Specifies the height of the element
	Height BlockElementHeight `json:"height,omitempty"`
	// When true, draw a separating line at the top of the element
	Separator bool `json:"separator,omitempty"`
	// Controls the amount of spacing between this element and the preceding element
	Spacing Spacing `json:"spacing,omitempty"`
	// If false, this item will be removed from the visual tree
	IsVisible bool `json:"isVisible,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Allows a user to input a Choice
//
// Source: https://adaptivecards.io/explorer/Input.ChoiceSet.html
type InputChoiceSet struct {
	// Must be  TypeInputText ("Input.Text")
	Type Type `json:"type"`
	// Unique identifier for the value. Used to identify collected input when the Submit action is performed
	Id string `json:"id"`
	// Choice options
	Choices []InputChoice `json:"choices"`
	// Allow multiple choices to be selected
	IsMultiSelect bool `json:"isMultiSelect"`
	// Description of the input desired. Only visible when no selection has been made, the style is compact and isMultiSelect is false
	Style ChoiceInputStyle `json:"style"`
	// The initial choice (or set of choices) that should be selected. For multi-select, specify a comma-separated string of values
	Value string `json:"value"`
	// Description of the input desired. Displayed when no text has been input
	Placeholder string `json:"placeholder"`
	// If true, allow text to wrap. Otherwise, text is clipped
	Wrap bool `json:"wrap"`
}

// Describes a choice for use in a ChoiceSet
//
// Source: https://adaptivecards.io/explorer/Input.Choice.html
type InputChoice struct {
	// Text to display.
	Title string `json:"title"`
	// The raw value for the choice. NOTE: do not use a , in the value, since a ChoiceSet with isMultiSelect set to true returns a comma-delimited string of choice values.
	Value string `json:"value"`
}
