package teams

import (
	"errors"
	"fmt"
)

type ISelectAction interface {
	IsISelectAction() bool
}

func (a *ActionOpenUrl) IsISelectAction() bool {
	return true
}
func (a *ActionSubmit) IsISelectAction() bool {
	return true
}
func (a *ActionToggleVisibility) IsISelectAction() bool {
	return true
}
func (a *ActionExecute) IsISelectAction() bool {
	return true
}

type Action interface {
	IsAction() bool
}

func (a *ActionOpenUrl) IsAction() bool {
	return true
}
func (a *ActionSubmit) IsAction() bool {
	return true
}
func (a *ActionShowCard) IsAction() bool {
	return true
}
func (a *ActionToggleVisibility) IsAction() bool {
	return true
}
func (a *ActionExecute) IsAction() bool {
	return true
}

const (
	TypeActionOpenUrl          Type = "Action.OpenUrl"
	TypeActionSubmit           Type = "Action.Submit"
	TypeActionToggleVisibility Type = "Action.ToggleVisibility"
	TypeActionExecute          Type = "Action.Execute"
)

type ActionStyle string

const (
	// Action is displayed as normal
	ActionStyleDefault ActionStyle = "default"
	// Action is displayed with a positive style (typically the button becomes accent color)
	ActionStylePositive ActionStyle = "positive"
	// Action is displayed with a destructive style (typically the button becomes red)
	ActionStyleDestructive ActionStyle = "destructive"
)

type ActionMode string

const (
	// Action is displayed as a button
	ActionModePrimary ActionMode = "primary"
	// Action is placed in an overflow menu (typically a popup menu under a ... button)
	ActionModeSecondary ActionMode = "secondary"
)

// When invoked, show the given url either by launching it in an external web browser or showing within an embedded web browser
//
// Source: https://adaptivecards.io/explorer/Action.OpenUrl.html
type ActionOpenUrl struct {
	// Must be  TypeActionOpenUrl ("Action.OpenUrl")
	Type Type `json:"type,omitempty"`
	// 	The URL to open
	Url string `json:"url,omitempty"`
	// Label for button or link that represents this action
	Title string `json:"title,omitempty"`
	// Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+
	IconUrl string `json:"iconUrl,omitempty"`
	// A unique identifier associated with this Action
	Id string `json:"id,omitempty"`
	// Controls the style of an Action, which influences how the action is displayed, spoken, etc.
	Style ActionStyle `json:"style,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software
	Tooltip string `json:"tooltip,omitempty"`
	// Determines whether the action should be enabled
	IsEnabled bool `json:"isEnabled,omitempty"`
	// Determines whether the action should be displayed as a button or in the overflow menu
	Mode ActionMode `json:"mode,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

func NewActionOpenUrl() *ActionOpenUrl {
	return &ActionOpenUrl{
		Type: TypeActionOpenUrl,
	}
}

func (a *ActionOpenUrl) isValid() (bool, error) {
	if a.Type != TypeActionOpenUrl {
		return false, errors.New(fmt.Sprintf("Type is invalid; expected: %s, got %s", a.Type, TypeActionOpenUrl))
	}
	if a.Url == "" {
		return false, errors.New("Url is missing")
	} else if !isValidUri(a.Url) {
		return false, errors.New(fmt.Sprintf("Url is invlid: %s", a.Url))
	}

	return true, nil
}

// Gathers input fields, merges with optional data field, and sends an event to the client. It is up to the client to determine how this data is processed. For example: With BotFramework bots, the client would send an activity through the messaging medium to the bot. The inputs that are gathered are those on the current card, and in the case of a show card those on any parent cards. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more details
//
// Source: https://adaptivecards.io/explorer/Action.Submit.html
type ActionSubmit struct {
	// Must be  TypeActionSubmit ("Action.Submit")
	Type Type `json:"type,omitempty"`
	// Initial data that input fields will be combined with. These are essentially ‘hidden’ properties
	Data interface{} `json:"data,omitempty"`
	// Controls which inputs are associated with the submit action
	AssociatedInputs AssociatedInputs `json:"associatedInputs,omitempty"`
	// Label for button or link that represents this action
	Title string `json:"title,omitempty"`
	// Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+
	IconUrl string `json:"iconUrl,omitempty"`
	// A unique identifier associated with this Action
	Id string `json:"id,omitempty"`
	// Controls the style of an Action, which influences how the action is displayed, spoken, etc.
	Style ActionStyle `json:"style,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software
	Tooltip string `json:"tooltip,omitempty"`
	// Determines whether the action should be enabled
	IsEnabled bool `json:"isEnabled,omitempty"`
	// Determines whether the action should be displayed as a button or in the overflow menu
	Mode ActionMode `json:"mode,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Defines an AdaptiveCard which is shown to the user when the button or link is clicked
//
// Source: https://adaptivecards.io/explorer/Action.ShowCard.html
type ActionShowCard struct {
	// Must be  TypeActionShowCard ("Action.ShowCard")
	Type Type `json:"type,omitempty"`
	// The Adaptive Card to show. Inputs in ShowCards will not be submitted if the submit button is located on a parent card. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more details
	Card AdaptiveCard `json:"card,omitempty"`
	// Label for button or link that represents this action
	Title string `json:"title,omitempty"`
	// Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+
	IconUrl string `json:"iconUrl,omitempty"`
	// A unique identifier associated with this Action
	Id string `json:"id,omitempty"`
	// Controls the style of an Action, which influences how the action is displayed, spoken, etc.
	Style ActionStyle `json:"style,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software
	Tooltip string `json:"tooltip,omitempty"`
	// Determines whether the action should be enabled
	IsEnabled bool `json:"isEnabled,omitempty"`
	// Determines whether the action should be displayed as a button or in the overflow menu
	Mode ActionMode `json:"mode,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// An action that toggles the visibility of associated card elements
//
// Source: https://adaptivecards.io/explorer/Action.ToggleVisibility.html
type ActionToggleVisibility struct {
	// Must be  TypeActionToggleVisibility ("Action.ToggleVisibility")
	Type Type `json:"type,omitempty"`
	// 	The array of TargetElements. It is not recommended to include Input elements with validation under Action.Toggle due to confusion that can arise from invalid inputs that are not currently visible. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more information
	TargetElements []TargetElement `json:"targetElements,omitempty"`
	// Label for button or link that represents this action
	Title string `json:"title,omitempty"`
	// Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+
	IconUrl string `json:"iconUrl,omitempty"`
	// A unique identifier associated with this Action
	Id string `json:"id,omitempty"`
	// Controls the style of an Action, which influences how the action is displayed, spoken, etc.
	Style ActionStyle `json:"style,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software
	Tooltip string `json:"tooltip,omitempty"`
	// Determines whether the action should be enabled
	IsEnabled bool `json:"isEnabled,omitempty"`
	// Determines whether the action should be displayed as a button or in the overflow menu
	Mode ActionMode `json:"mode,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}

// Represents an entry for Action.ToggleVisibility’s targetElements property
//
// Source: https://adaptivecards.io/explorer/TargetElement.html
type TargetElement struct {
	// Element ID of element to toggle
	ElementId string `json:"elementId,omitempty"`
	// If true, always show target element. If false, always hide target element. If not supplied, toggle target element’s visibility
	IsVisible bool `json:"isVisible,omitempty"`
}

// Gathers input fields, merges with optional data field, and sends an event to the client. Clients process the event by sending an Invoke activity of type adaptiveCard/action to the target Bot. The inputs that are gathered are those on the current card, and in the case of a show card those on any parent cards. See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/universal-action-model documentation for more details
//
// Source: https://adaptivecards.io/explorer/Action.Execute.html
type ActionExecute struct {
	// Must be  TypeActionExecute ("Action.Execute")
	Type Type `json:"type,omitempty"`
	// The card author-defined verb associated with this action
	Verb string `json:"verb,omitempty"`
	// Initial data that input fields will be combined with. These are essentially ‘hidden’ properties
	Data interface{} `json:"data,omitempty"`
	// Controls which inputs are associated with the action
	AssociatedInputs AssociatedInputs `json:"associatedInputs,omitempty"`
	// Label for button or link that represents this action
	Title string `json:"title,omitempty"`
	// Optional icon to be shown on the action in conjunction with the title. Supports data URI in version 1.2+
	IconUrl string `json:"iconUrl,omitempty"`
	// A unique identifier associated with this Action
	Id string `json:"id,omitempty"`
	// Controls the style of an Action, which influences how the action is displayed, spoken, etc.
	Style ActionStyle `json:"style,omitempty"`
	// 	Describes what to do when an unknown element is encountered or the requires of this or any children can’t be met
	Fallback interface{} `json:"fallback,omitempty"`
	// Defines text that should be displayed to the end user as they hover the mouse over the action, and read when using narration software
	Tooltip string `json:"tooltip,omitempty"`
	// Determines whether the action should be enabled
	IsEnabled bool `json:"isEnabled,omitempty"`
	// Determines whether the action should be displayed as a button or in the overflow menu
	Mode ActionMode `json:"mode,omitempty"`
	// A series of key/value pairs indicating features that the item requires with corresponding minimum version. When a feature is missing or of insufficient version, fallback is triggered
	Requires interface{} `json:"requires,omitempty"`
}
