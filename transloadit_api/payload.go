package transloadit_api

type Payload struct {
	Params    string `url:"params", json:"params"`
	Signature string `url:"signature", json:"params"`
}
