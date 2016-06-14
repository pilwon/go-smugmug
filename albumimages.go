package smugmug

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"strconv"
)

type AlbumImagesService struct {
	s *Service
}

func NewAlbumImagesService(s *Service) *AlbumImagesService {
	r := &AlbumImagesService{s: s}
	return r
}

func (r *AlbumImagesService) Get(u string) *AlbumImagesGetCall {
	c := &AlbumImagesGetCall{s: r.s, urlParams: url.Values{}}
	c.u = u
	return c
}

type AlbumImagesServiceResponse struct {
	Code     int
	Message  string
	Response struct {
		ServiceResponse
		AlbumImage []*Image
	}
}

type AlbumImagesGetCall struct {
	u string

	s         *Service
	urlParams url.Values
}

func (c *AlbumImagesGetCall) Expand(expansions []string) *AlbumImagesGetCall {
	c.urlParams.Set("_expand", strings.Join(expansions, ","))
	return c
}

func (c *AlbumImagesGetCall) Filter(filter []string) *AlbumImagesGetCall {
	c.urlParams.Set("_filter", strings.Join(filter, ","))
	return c
}

func (c *AlbumImagesGetCall) Goto(start int, count int) *AlbumImagesGetCall {
	c.urlParams.Set("start", strconv.Itoa(start))
	c.urlParams.Set("count", strconv.Itoa(count))
	return c
}

func (c *AlbumImagesGetCall) doRequest() (*http.Response, error) {
	if u, err := url.Parse(c.s.BasePath); err != nil {
		return nil, err
	} else {
		urls := c.s.BasePath[:len(c.s.BasePath)-len(u.Path)] + c.u
		urls += "?" + encodeURLParams(c.urlParams)
		req, _ := http.NewRequest("GET", urls, nil)
		c.s.setHeaders(req)
		debugRequest(req)
		return c.s.client.Do(req)
	}
}

func (c *AlbumImagesGetCall) Do() (*AlbumImagesGetResponse, error) {
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	debugResponse(res)
	defer closeBody(res)
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	albumsRes := &AlbumImagesServiceResponse{}
	if err := json.NewDecoder(res.Body).Decode(&albumsRes); err != nil {
		return nil, err
	}
	ServiceResponse := albumsRes.Response
	fmt.Println(ServiceResponse.Pages.Count)
	ret := &AlbumImagesGetResponse{
		Images: albumsRes.Response.AlbumImage,
		Pages: albumsRes.Response.Pages,
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	return ret, nil
}

type AlbumImagesGetResponse struct {
	Images []*Image
	Pages Pages
	ServerResponse `json:"-"`
}
