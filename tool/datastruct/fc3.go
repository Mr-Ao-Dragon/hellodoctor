package datastruct

type UniversalRepose struct {
	StatusCode      int16   `json:"statusCode"`
	Headers         Headers `json:"headers"`
	IsBase64Encoded string  `json:"isBase64Encoded"`
	Body            string  `json:"body"`
}
type Headers struct {
	AccessControlAllowOrigin      string `json:"Access-Control-Allow-Origin"`
	AccessControlAllowHeaders     string `json:"Access-Control-Allow-Headers"`
	AccessControlAllowMethods     string `json:"Access-Control-Allow-Methods"`
	AccessControlMaxAge           string `json:"Access-Control-Max-Age"`
	AccessControlAllowCredentials string `json:"Access-Control-Allow-Credentials"`
	AccessControlExposeHeaders    string `json:"Access-Control-Expose-Headers"`
	AccessControlRequestHeaders   string `json:"Access-Control-Request-Headers"`
	AccessControlRequestMethod    string `json:"Access-Control-Request-Method"`
	ContentType                   string `json:"Content-Type"`
}
