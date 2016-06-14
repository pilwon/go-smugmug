package smugmug

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AlbumsService struct {
	s *Service
}

func NewAlbumsService(s *Service) *AlbumsService {
	r := &AlbumsService{s: s}
	return r
}

func (r *AlbumsService) Get(id string) *AlbumsGetCall {
	c := &AlbumsGetCall{s: r.s, urlParams: url.Values{}}
	c.id = id
	return c
}

type AlbumsServiceResponse struct {
	Code     int
	Message  string
	Response struct {
		ServiceResponse
		Album *json.RawMessage
	}
	Expansions map[string]*json.RawMessage `json:",omitempty"`
}

type AlbumsGetCall struct {
	id string

	s         *Service
	urlParams url.Values
}

func (c *AlbumsGetCall) Expand(expansions []string) *AlbumsGetCall {
	c.urlParams.Set("_expand", strings.Join(expansions, ","))
	return c
}

func (c *AlbumsGetCall) Filter(filter []string) *AlbumsGetCall {
	c.urlParams.Set("_filter", strings.Join(filter, ","))
	return c
}

func (c *AlbumsGetCall) doRequest() (*http.Response, error) {
	urls := resolveRelative(c.s.BasePath, "album/"+c.id)
	urls += "?" + encodeURLParams(c.urlParams)
	req, _ := http.NewRequest("GET", urls, nil)
	c.s.setHeaders(req)
	debugRequest(req)
	return c.s.client.Do(req)
}

func (c *AlbumsGetCall) Do() (*AlbumsGetResponse, error) {
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	debugResponse(res)
	defer closeBody(res)
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	albumsRes := &AlbumsServiceResponse{}
	if err := json.NewDecoder(res.Body).Decode(&albumsRes); err != nil {
		return nil, err
	}
	album := &Album{}
	if err := json.Unmarshal(*albumsRes.Response.Album, &album); err != nil {
		return nil, err
	}
	exp, err := unmarshallExpansions(album.URIs, albumsRes.Expansions)
	if err != nil {
		return nil, err
	}
	ret := &AlbumsGetResponse{
		Album: album,
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	for name, v := range exp {
		switch name {
		case "Node":
			ret.Node = v.(*Node)
		case "User":
			ret.User = v.(*User)
		}
	}
	return ret, nil
}

type AlbumsGetResponse struct {
	Album *Album

	// AlbumDownload
	// AlbumGeoMedia
	// AlbumHighlightImage // deprecated
	// AlbumImages
	// AlbumPopularMedia
	// AlbumPrices
	// AlbumShareUris
	// ApplyAlbumTemplate
	// CollectImages
	// DeleteAlbumImages
	// Folder // deprecated
	// HighlightImage
	// MoveAlbumImages
	Node *Node
	// ParentFolders // deprecated
	// SortAlbumImages
	// UploadFromUri
	User *User

	ServerResponse `json:"-"`
}

func (c *AlbumsGetResponse) GetImages(s *Service) (images []*Image, err error) {
	var uri URI = (*c.Album.URIs)["AlbumImages"]
	var url string = parseURI(uri)
	var start int = 1
	var count int = 100

	// loop for all pages
	for true {
		var res *AlbumImagesGetResponse
		if res, err = s.AlbumImages.Get(url).Goto(start, count).Do(); err != nil {
			return nil, err
		}

		images = append(images, res.Images...)

		p := res.Pages
		start = p.Start + p.Count
		if start > p.Total {
			return images, nil
		}

		/*var link string = res.ServerResponse.Header.Get("Link")
		r := regexp.MustCompile("<(.*)>")
		url = r.FindStringSubmatch(link)[1]*/
	}

	return nil, nil
}

type Album struct {
	AlbumKey            string     `json:",omitempty"`
	AllowDownloads      bool       `json:",omitempty"`
	Backprinting        string     `json:",omitempty"`
	BoutiquePackaging   string     `json:",omitempty"`
	CanRank             bool       `json:",omitempty"`
	CanShare            bool       `json:",omitempty"`
	Clean               bool       `json:",omitempty"`
	Comments            bool       `json:",omitempty"`
	Date                *time.Time `json:",omitempty"`
	Description         string     `json:",omitempty"`
	EXIF                bool       `json:",omitempty"`
	External            bool       `json:",omitempty"`
	FamilyEdit          bool       `json:",omitempty"`
	Filenames           bool       `json:",omitempty"`
	FriendEdit          bool       `json:",omitempty"`
	Geography           bool       `json:",omitempty"`
	HasDownloadPassword bool       `json:",omitempty"`
	Header              string     `json:",omitempty"`
	HideOwner           bool       `json:",omitempty"`
	ImageCount          int        `json:",omitempty"`
	ImagesLastUpdated   string     `json:",omitempty"`
	InterceptShipping   string     `json:",omitempty"`
	Keywords            string     `json:",omitempty"`
	LargestSize         string     `json:",omitempty"`
	LastUpdated         string     `json:",omitempty"`
	Name                string     `json:",omitempty"`
	NiceName            string     `json:",omitempty"`
	NodeID              string     `json:",omitempty"`
	OriginalSizes       int        `json:",omitempty"`
	PackagingBranding   bool       `json:",omitempty"`
	Password            string     `json:",omitempty"`
	PasswordHint        string     `json:",omitempty"`
	Printable           bool       `json:",omitempty"`
	Privacy             string     `json:",omitempty"`
	ProofDays           int        `json:",omitempty"`
	Protected           bool       `json:",omitempty"`
	SecurityType        string     `json:",omitempty"`
	Share               bool       `json:",omitempty"`
	SmugSearchable      string     `json:",omitempty"`
	SortDirection       string     `json:",omitempty"`
	SortMethod          string     `json:",omitempty"`
	SquareThumbs        bool       `json:",omitempty"`
	TemplateURI         string     `json:"TemplateUri"`
	Title               string     `json:",omitempty"`
	TotalSizes          int        `json:",omitempty"`
	URLName             string     `json:"UrlName,omitempty"`
	URLPath             string     `json:"UrlPath,omitempty"`
	Watermark           bool       `json:",omitempty"`
	WorldSearchable     bool       `json:",omitempty"`

	ResponseLevel  string `json:",omitempty"`
	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
	URIs           *URIs  `json:"Uris,omitempty"`
	WebURI         string `json:"WebUri,omitempty"`
}
