package models

//RespuestaLogin tiene el token que devuelve el login
type RespuestaLogin struct {
	Token string `json:"token, omitempty"`
}
