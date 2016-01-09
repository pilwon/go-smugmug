package smugmug

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type UsersService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	r := &UsersService{s: s}
	return r
}

func (r *UsersService) Get(id string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, urlParams: url.Values{}}
	c.id = id
	return c
}

func (r *UsersService) GetAuthUser() *UsersGetCall {
	c := &UsersGetCall{s: r.s, urlParams: url.Values{}}
	c.useAuthUser = true
	return c
}

type UsersServiceResponse struct {
	Code     int
	Message  string
	Response struct {
		ServiceResponse
		User *json.RawMessage
	}
	Expansions map[string]*json.RawMessage `json:",omitempty"`
}

type UsersGetCall struct {
	id          string
	useAuthUser bool

	s         *Service
	urlParams url.Values
}

func (c *UsersGetCall) Expand(expansions []string) *UsersGetCall {
	c.urlParams.Set("_expand", strings.Join(expansions, ","))
	return c
}

func (c *UsersGetCall) Filter(filter []string) *UsersGetCall {
	c.urlParams.Set("_filter", strings.Join(filter, ","))
	return c
}

func (c *UsersGetCall) doRequest() (*http.Response, error) {
	urls := resolveRelative(c.s.BasePath, "user/"+c.id)
	if c.useAuthUser {
		urls = strings.TrimRight(c.s.BasePath, "/") + "!authuser"
	}
	urls += "?" + encodeURLParams(c.urlParams)
	req, _ := http.NewRequest("GET", urls, nil)
	c.s.setHeaders(req)
	debugRequest(req)
	return c.s.client.Do(req)
}

func (c *UsersGetCall) Do() (*UsersGetResponse, error) {
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	debugResponse(res)
	defer closeBody(res)
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	usersRes := &UsersServiceResponse{}
	if err := json.NewDecoder(res.Body).Decode(&usersRes); err != nil {
		return nil, err
	}
	user := &User{}
	if err := json.Unmarshal(*usersRes.Response.User, &user); err != nil {
		return nil, err
	}
	exp, err := unmarshallExpansions(user.URIs, usersRes.Expansions)
	if err != nil {
		return nil, err
	}
	ret := &UsersGetResponse{
		User: user,
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	for name, v := range exp {
		switch name {
		case "Node":
			ret.Node = v.(*Node)
		}
	}
	return ret, nil
}

type UsersGetResponse struct {
	User *User

	// BioImage               *BioImage
	// CoverImage             *CoverImage
	// DuplicateImageSearch   *DuplicateImageSearch
	// Features               *Features
	// Folder                 *Folder
	Node *Node
	// SortUserFeaturedAlbums *SortUserFeaturedAlbums
	// UrlPathLookup          *UrlPathLookup
	// UserAlbumTemplates     *UserAlbumTemplates
	// UserAlbums             []*Album
	// UserContacts           *UserContacts
	// UserCoupons            *UserCoupons
	// UserDeletedAlbums      *UserDeletedAlbums
	// UserDeletedFolders     *UserDeletedFolders
	// UserDeletedPages       *UserDeletedPages
	// UserFeaturedAlbums     *UserFeaturedAlbums
	// UserGeoMedia           *UserGeoMedia
	// UserGrants             *UserGrants
	// UserGuideStates        *UserGuideStates
	// UserHideGuides         *UserHideGuides
	// UserImageSearch        *UserImageSearch
	// UserLatestQuickNews    *UserLatestQuickNews
	// UserPopularMedia       *UserPopularMedia
	// UserPrintmarks         *UserPrintmarks
	// UserProfile            *UserProfile
	// UserRecentImages       *UserRecentImages
	// UserTasks              *UserTasks
	// UserTopKeywords        *UserTopKeywords
	// UserUploadLimits       *UserUploadLimits
	// UserWatermarks         *UserWatermarks

	ServerResponse `json:"-"`
}

type User struct {
	AccountStatus     string `json:",omitempty"`
	Domain            string `json:",omitempty"`
	DomainOnly        string `json:",omitempty"`
	FirstName         string `json:",omitempty"`
	FriendsView       bool   `json:",omitempty"`
	ImageCount        int    `json:",omitempty"`
	IsTrial           bool   `json:",omitempty"`
	LastName          string `json:",omitempty"`
	Name              string `json:",omitempty"`
	NickName          string `json:",omitempty"`
	Plan              string `json:",omitempty"`
	QuickShare        bool   `json:",omitempty"`
	RefTag            string `json:",omitempty"`
	SortBy            string `json:",omitempty"`
	TotalAccountSize  string `json:",omitempty"`
	TotalUploadedSize string `json:",omitempty"`
	ViewPassHint      string `json:",omitempty"`
	ViewPassword      string `json:",omitempty"`

	ResponseLevel string `json:",omitempty"`
	URI           string `json:"Uri,omitempty"`
	URIs          *URIs  `json:"Uris,omitempty"`
	WebURI        string `json:"WebUri,omitempty"`
}
