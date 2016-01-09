package smugmug

import (
	"net/url"
	"time"
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

type AlbumsGetResponse struct {
	Album *Album

	// AlbumDownload
	// AlbumGeoMedia
	// AlbumHighlightImage
	// AlbumImages
	// AlbumPopularMedia
	// AlbumPrices
	// AlbumShareUris
	// ApplyAlbumTemplate
	// CollectImages
	// DeleteAlbumImages
	// Folder
	// HighlightImage
	// MoveAlbumImages
	// Node
	// ParentFolders
	// SortAlbumImages
	// UploadFromUri
	// User

	ServerResponse `json:"-"`
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
