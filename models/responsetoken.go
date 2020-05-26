package models

// ResponseToken is used only to give a json response.
// It is optional to use it, but it makes easier
// and it's a more organized way to handled it
type ResponseToken struct {
	Token string `json:"token"`
}
