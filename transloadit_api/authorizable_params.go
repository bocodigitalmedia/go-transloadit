package transloadit_api

type AuthorizableParams interface {
	SetAuth(*Auth)
}
