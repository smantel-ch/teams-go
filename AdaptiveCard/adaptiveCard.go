package teams

const (
	TypeAdaptiveCard Type = "AdaptiveCard"
)

// An Adaptive Card, containing a free-form body of card elements, and an optional set of actions
//
// Source: https://adaptivecards.io/explorer/AdaptiveCard.html
type AdaptiveCard struct {
	// Must be  TypeAdaptiveCard ("AdaptiveCard")
	Type Type `json:"type,omitempty"`
	// Schema version that this card requires. If a client is lower than this version, the fallbackText will be rendered. NOTE: Version is not required for cards within an Action.ShowCard. However, it is required for the top-level card
	Version Version `json:"version,omitempty"`
	// Defines how the card can be refreshed by making a request to the target Bot
	Refresh *Refresh `json:"refresh,omitempty"`
	// Defines authentication information to enable on-behalf-of single sign on or just-in-time OAuth
	Authentication *Authentication `json:"authentication,omitempty"`
	// The card elements to show in the primary card region
	Body []Element `json:"body,omitempty"`
	// 	The Actions to show in the card’s action bar
	Actions []Action `json:"actions,omitempty"`
	// An Action that will be invoked when the card is tapped or selected. Action.ShowCard is not supported
	SelectAction []ISelectAction `json:"selectAction,omitempty"`
	// Text shown when the client doesn’t support the version specified (may contain markdown)
	FallbackText string `json:"fallbackText,omitempty"`
	// Specifies the background image of the card
	BackgroundImage *BackgroundImage `json:"backgroundImage,omitempty"`
	// Specifies the minimum height of the card
	MinHeight string `json:"minHeight,omitempty"`
	// When true content in this Adaptive Card should be presented right to left. When ‘false’ content in this Adaptive Card should be presented left to right. If unset, the default platform behavior will apply
	Rtl bool `json:"rtl,omitempty"`
	// Specifies what should be spoken for this entire card. This is simple text or SSML fragment
	Speak string `json:"speak,omitempty"`
	// The 2-letter ISO-639-1 language used in the card. Used to localize any date/time functions
	Lang string `json:"lang,omitempty"`
	// Defines how the content should be aligned vertically within the container. Only relevant for fixed-height cards, or cards with a minHeight specified
	VerticalContentAlignment *VerticalContentAlignment `json:"verticalContentAlignment,omitempty"`
	// The Adaptive Card schema
	Schema Schema `json:"$schema,omitempty"`
}

func NewAdaptiveCard() *AdaptiveCard {
	return &AdaptiveCard{
		Schema:  SchemaDefault,
		Type:    TypeAdaptiveCard,
		Version: Version13,
	}
}
