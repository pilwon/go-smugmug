package smugmug

import (
	"net/url"
)

type AlbumsService struct {
	s *Service
}

func NewAlbumsService(s *Service) *AlbumsService {
	r := &AlbumsService{s: s}
	return r
}

type AlbumsGetCall struct {
	s         *Service
	id        string
	urlParams url.Values
}

func (r *AlbumsService) Get(id string) *AlbumsGetCall {
	c := &AlbumsGetCall{s: r.s}
	return c
}

type Album struct {
	// TODO

	ResponseLevel string
	URI           string `json:"Uri"`
	URIs          *URIs  `json:"Uris"`
	WebURI        string `json:"WebUri"`
}
