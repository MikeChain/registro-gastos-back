package models

// RespuestaLogin es la estrucutara con la que se devuelve el JWT al hacer login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
