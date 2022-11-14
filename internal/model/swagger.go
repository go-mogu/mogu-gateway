package model

type SwaggerResources struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	SwaggerVersion string `json:"swaggerVersion"`
	Location       string `json:"location"`
}

type SwaggerConfig struct {
	ConfigURL         string `json:"configUrl"`
	Oauth2RedirectURL string `json:"oauth2RedirectUrl"`
	Urls              []URL  `json:"urls"`
	ValidatorURL      string `json:"validatorUrl"`
}

type URL struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
