package teams

// Specifies a background image. Acceptable formats are PNG, JPEG, and GIF
type BackgroundImage struct {
	// The URL (or data url) of the image. Acceptable formats are PNG, JPEG, and GIF
	URL string `json:"url,omitempty"`
	// Describes how the image should fill the area.
	FillMode ImageFillMode `json:"fillMode,omitempty"`
	// Describes how the image should be aligned if it must be cropped or if using repeat fill mode.
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	// Describes how the image should be aligned if it must be cropped or if using repeat fill mode.
	VerticalAlignment VerticalAlignment `json:"verticalAlignment,omitempty"`
}

// Defines how a card can be refreshed by making a request to the target Bot.
//
// Source: https://adaptivecards.io/explorer/Refresh.html
type Refresh struct {
	// The action to be executed to refresh the card. Clients can run this refresh action automatically or can provide an affordance for users to trigger it manually
	Action ActionExecute `json:"action,omitempty"`
	// A list of user Ids informing the client for which users should the refresh action should be run automatically. Some clients will not run the refresh action automatically unless this property is specified. Some clients may ignore this property and always run the refresh action automatically
	UserIds []string `json:"userIds,omitempty"`
}

// Defines authentication information associated with a card. This maps to the OAuthCard type defined by the Bot Framework (https://docs.microsoft.com/dotnet/api/microsoft.bot.schema.oauthcard)
//
// Source: https://adaptivecards.io/explorer/Authentication.html
type Authentication struct {
	// Text that can be displayed to the end user when prompting them to authenticate
	Text string `json:"text,omitempty"`
	// The identifier for registered OAuth connection setting information
	ConnectionName string `json:"connectionName,omitempty"`
	// Provides information required to enable on-behalf-of single sign-on user authentication
	TokenExchangeResource TokenExchangeResource `json:"tokenExchangeResource,omitempty"`
	// Buttons that should be displayed to the user when prompting for authentication. The array MUST contain one button of type “signin”. Other button types are not currently supported
	Buttons []AuthCardButton `json:"buttons,omitempty"`
}

// Defines information required to enable on-behalf-of single sign-on user authentication. Maps to the TokenExchangeResource type defined by the Bot Framework (https://docs.microsoft.com/dotnet/api/microsoft.bot.schema.tokenexchangeresource)
//
// Souce: https://adaptivecards.io/explorer/TokenExchangeResource.html
type TokenExchangeResource struct {
	// The unique identified of this token exchange instance
	Id string `json:"id,omitempty"`
	// An application ID or resource identifier with which to exchange a token on behalf of. This property is identity provider- and application-specific
	Uri string `json:"uri,omitempty"`
	// An identifier for the identity provider with which to attempt a token exchange
	ProviderId string `json:"providerId,omitempty"`
}

// Defines a button as displayed when prompting a user to authenticate. This maps to the cardAction type defined by the Bot Framework (https://docs.microsoft.com/dotnet/api/microsoft.bot.schema.cardaction)
//
// Source: https://adaptivecards.io/explorer/AuthCardButton.html
type AuthCardButton struct {
	// The type of the button
	Type string `json:"type,omitempty"`
	// The value associated with the button. The meaning of value depends on the button’s type
	Value string `json:"value,omitempty"`
	// The caption of the button
	Title string `json:"title,omitempty"`
	// A URL to an image to display alongside the button’s caption
	Image string `json:"image,omitempty"`
}
