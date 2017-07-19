package transloadit_api

import "github.com/dghubble/sling"

type SlingDecorator func(*sling.Sling) *sling.Sling
