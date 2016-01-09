package smugmug

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ImagesService struct {
	s *Service
}

func NewImagesService(s *Service) *ImagesService {
	r := &ImagesService{s: s}
	return r
}

func (r *ImagesService) Get(id string) *ImagesGetCall {
	c := &ImagesGetCall{s: r.s, urlParams: url.Values{}}
	c.id = id
	return c
}

type ImagesServiceResponse struct {
	Code     int
	Message  string
	Response struct {
		ServiceResponse
		Image *json.RawMessage
	}
	Expansions map[string]*json.RawMessage `json:",omitempty"`
}

type ImagesGetCall struct {
	id string

	s         *Service
	urlParams url.Values
}

func (c *ImagesGetCall) Expand(expansions []string) *ImagesGetCall {
	c.urlParams.Set("_expand", strings.Join(expansions, ","))
	return c
}

func (c *ImagesGetCall) Filter(filter []string) *ImagesGetCall {
	c.urlParams.Set("_filter", strings.Join(filter, ","))
	return c
}

func (c *ImagesGetCall) doRequest() (*http.Response, error) {
	urls := resolveRelative(c.s.BasePath, "image/"+c.id)
	urls += "?" + encodeURLParams(c.urlParams)
	req, _ := http.NewRequest("GET", urls, nil)
	c.s.setHeaders(req)
	debugRequest(req)
	return c.s.client.Do(req)
}

func (c *ImagesGetCall) Do() (*ImagesGetResponse, error) {
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	debugResponse(res)
	defer closeBody(res)
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	imagesRes := &ImagesServiceResponse{}
	if err := json.NewDecoder(res.Body).Decode(&imagesRes); err != nil {
		return nil, err
	}
	image := &Image{}
	if err := json.Unmarshal(*imagesRes.Response.Image, &image); err != nil {
		return nil, err
	}
	exp, err := unmarshallExpansions(image.URIs, imagesRes.Expansions)
	if err != nil {
		return nil, err
	}
	ret := &ImagesGetResponse{
		Image: image,
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	_ = exp
	for name, v := range exp {
		switch name {
		case "ImageAlbum":
			ret.ImageAlbum = v.(*Album)
		case "ImageDownload":
			ret.ImageDownload = v.(*ImageDownload)
		case "ImageMetadata":
			ret.ImageMetadata = v.(*ImageMetadata)
		case "ImageOwner":
			ret.ImageOwner = v.(*User)
		case "ImagePrices":
			ret.ImagePrices = v.([]*CatalogSkuPrice)
		case "ImageSizeDetails":
			ret.ImageSizeDetails = v.(*ImageSizeDetails)
		case "ImageSizes":
			ret.ImageSizes = v.(*ImageSizes)
		case "LargestImage":
			ret.LargestImage = v.(*LargestImage)
		}

	}
	return ret, nil
}

type ImagesGetResponse struct {
	Image *Image

	ImageAlbum *Album
	// ImageComments
	ImageDownload    *ImageDownload
	ImageMetadata    *ImageMetadata
	ImageOwner       *User
	ImagePrices      []*CatalogSkuPrice
	ImageSizeDetails *ImageSizeDetails
	ImageSizes       *ImageSizes
	LargestImage     *LargestImage

	ServerResponse `json:"-"`
}

type Image struct {
	Altitude        int              `json:",omitempty"`
	ArchivedMD5     string           `json:",omitempty"`
	ArchivedSize    int              `json:",omitempty"`
	ArchivedURI     string           `json:"ArchivedUri,omitempty"`
	CanEdit         bool             `json:",omitempty"`
	Caption         string           `json:",omitempty"`
	Collectable     bool             `json:",omitempty"`
	Date            *time.Time       `json:",omitempty"`
	EZProject       bool             `json:",omitempty"`
	FileName        string           `json:",omitempty"`
	Format          string           `json:",omitempty"`
	FormattedValues *FormattedValues `json:",omitempty"`
	Hidden          bool             `json:",omitempty"`
	ImageKey        string           `json:",omitempty"`
	IsArchive       bool             `json:",omitempty"`
	IsVideo         bool             `json:",omitempty"`
	KeywordArray    []string         `json:",omitempty"`
	Keywords        string           `json:",omitempty"`
	LastUpdated     *time.Time       `json:",omitempty"`
	Latitude        string           `json:",omitempty"`
	Longitude       string           `json:",omitempty"`
	OriginalHeight  int              `json:",omitempty"`
	OriginalSize    int              `json:",omitempty"`
	OriginalWidth   int              `json:",omitempty"`
	Processing      bool             `json:",omitempty"`
	Protected       bool             `json:",omitempty"`
	ThumbnailURL    string           `json:"ThumbnailUrl,omitempty"`
	Title           string           `json:",omitempty"`
	UploadKey       string           `json:",omitempty"`
	Watermarked     bool             `json:",omitempty"`

	ResponseLevel string
	URI           string `json:"Uri,omitempty"`
	URIs          *URIs  `json:"Uris,omitempty"`
	WebURI        string `json:"WebUri,omitempty"`
}

type ImageDownload struct {
	URL string `json:"Url,omitempty"`

	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
}

type ImageMetadata struct {
	Altitude               string  `json:",omitempty"`
	AltitudeReference      string  `json:",omitempty"`
	Aperture               float64 `json:",omitempty"`
	AudioCodec             string  `json:",omitempty"`
	Author                 string  `json:",omitempty"`
	AuthorTitle            string  `json:",omitempty"`
	Brightness             string  `json:",omitempty"`
	Caption                string  `json:",omitempty"`
	Category               string  `json:",omitempty"`
	CircleOfConfusion      string  `json:",omitempty"`
	City                   string  `json:",omitempty"`
	ColorSpace             string  `json:",omitempty"`
	CompressedBitsPerPixel string  `json:",omitempty"`
	Contrast               string  `json:",omitempty"`
	Copyright              string  `json:",omitempty"`
	CopyrightFlag          string  `json:",omitempty"`
	CopyrightURL           string  `json:"CopyrightUrl"`
	Country                string  `json:",omitempty"`
	CountryCode            string  `json:",omitempty"`
	CreatorContactInfo     string  `json:",omitempty"`
	Credit                 string  `json:",omitempty"`
	DateCreated            string  `json:",omitempty"` // *time.Time
	DateDigitized          string  `json:",omitempty"` // *time.Time
	DateTimeCreated        string  `json:",omitempty"` // *time.Time
	DateTimeModified       string  `json:",omitempty"` // *time.Time
	DepthOfField           string  `json:",omitempty"`
	DigitalZoomRatio       float64 `json:",omitempty"`
	Duration               string  `json:",omitempty"`
	Exposure               string  `json:",omitempty"`
	ExposureCompensation   string  `json:",omitempty"`
	ExposureMode           string  `json:",omitempty"`
	ExposureProgram        string  `json:",omitempty"`
	FieldOfView            string  `json:",omitempty"`
	Flash                  string  `json:",omitempty"`
	FocalLength            string  `json:",omitempty"`
	FocalLength35mm        string  `json:",omitempty"`
	GainControl            string  `json:",omitempty"`
	Headline               string  `json:",omitempty"`
	HyperfocalDistance     string  `json:",omitempty"`
	ISO                    int     `json:",omitempty"`
	Keywords               string  `json:",omitempty"`
	Latitude               float64 `json:",omitempty"`
	LatitudeReference      string  `json:",omitempty"`
	Lens                   string  `json:",omitempty"`
	LensSerialNumber       string  `json:",omitempty"`
	LightSource            string  `json:",omitempty"`
	Longitude              float64 `json:",omitempty"`
	LongitudeReference     string  `json:",omitempty"`
	Make                   string  `json:",omitempty"`
	Metering               string  `json:",omitempty"`
	MicroDateTimeCreated   string  `json:",omitempty"` // *time.Time
	MicroDateTimeDigitized string  `json:",omitempty"` // *time.Time
	Model                  string  `json:",omitempty"`
	NormalizedLightValue   float64 `json:",omitempty"`
	Rating                 string  `json:",omitempty"`
	Saturation             string  `json:",omitempty"`
	ScaleFactor            string  `json:",omitempty"`
	SceneCaptureType       string  `json:",omitempty"`
	SensingMethod          string  `json:",omitempty"`
	SerialNumber           string  `json:",omitempty"`
	Sharpness              string  `json:",omitempty"`
	Software               string  `json:",omitempty"`
	Source                 string  `json:",omitempty"`
	SpecialInstructions    string  `json:",omitempty"`
	State                  string  `json:",omitempty"`
	SubjectDistance        string  `json:",omitempty"`
	SubjectRange           string  `json:",omitempty"`
	SupplementalCategories string  `json:",omitempty"`
	TimeCreated            string  `json:",omitempty"`
	Title                  string  `json:",omitempty"`
	TransmissionReference  string  `json:",omitempty"`
	UserComment            string  `json:",omitempty"`
	VideoCodec             string  `json:",omitempty"`
	WhiteBalance           string  `json:",omitempty"`
	WriterEditor           string  `json:",omitempty"`

	ResponseLevel  string `json:",omitempty"`
	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
}

type CatalogSkuPrice struct {
	Currency string  `json:",omitempty"`
	Price    float64 `json:",omitempty"`

	ResponseLevel  string `json:",omitempty"`
	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
}

type ImageSize struct {
	URL    string `json:"Url,omitempty"`
	Ext    string `json:",omitempty"`
	Height int    `json:",omitempty"`
	Width  int    `json:",omitempty"`
	Size   int    `json:",omitempty"`
}

type ImageSizeDetails struct {
	ImageSizeLarge    *ImageSize `json:",omitempty"`
	ImageSizeMedium   *ImageSize `json:",omitempty"`
	ImageSizeOriginal *ImageSize `json:",omitempty"`
	ImageSizeSmall    *ImageSize `json:",omitempty"`
	ImageSizeThumb    *ImageSize `json:",omitempty"`
	ImageSizeTiny     *ImageSize `json:",omitempty"`
	ImageSizeX2Large  *ImageSize `json:",omitempty"`
	ImageSizeX3Large  *ImageSize `json:",omitempty"`
	ImageSizeXLarge   *ImageSize `json:",omitempty"`
	ImageURLTemplate  string     `json:"ImageUrlTemplate,omitempty"`
	UsableSizes       []string   `json:",omitempty"`

	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
}

type ImageSizes struct {
	LargeImageURL    string `json:"LargeImageUrl,omitempty"`
	LargestImageURL  string `json:"LargestImageUrl,omitempty"`
	MediumImageURL   string `json:"MediumImageUrl,omitempty"`
	OriginalImageURL string `json:"OriginalImageUrl,omitempty"`
	SmallImageURL    string `json:"SmallImageUrl,omitempty"`
	ThumbImageURL    string `json:"ThumbImageUrl,omitempty"`
	TinyImageURL     string `json:"TinyImageUrl,omitempty"`
	X2LargeImageURL  string `json:"X2LargeImageUrl,omitempty"`
	X3LargeImageURL  string `json:"X3LargeImageUrl,omitempty"`
	XLargeImageURL   string `json:"XLargeImageUrl,omitempty"`

	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
}

type LargestImage struct {
	Ext         string `json:",omitempty"`
	Height      int    `json:",omitempty"`
	Size        int    `json:",omitempty"`
	URL         string `json:"Url,omitempty"`
	Usable      bool   `json:",omitempty"`
	Watermarked bool   `json:",omitempty"`
	Width       int    `json:",omitempty"`

	URI            string `json:"Uri,omitempty"`
	URIDescription string `json:"UriDescription,omitempty"`
}
