package transloadit_api

type Payload struct {
	Params    string `json:"params" url:"params"`
	Signature string `json:"signature" url:"signature"`
}
