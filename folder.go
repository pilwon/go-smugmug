package smugmug

import (
	"net/url"
)

type FoldersService struct {
	s *Service
}

func NewFoldersService(s *Service) *FoldersService {
	r := &FoldersService{s: s}
	return r
}

type FoldersGetCall struct {
	s         *Service
	id        string
	urlParams url.Values
}

func (r *FoldersService) Get(id string) *FoldersGetCall {
	c := &FoldersGetCall{s: r.s}
	return c
}
