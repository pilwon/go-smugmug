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
	Altitude        int
	ArchivedMD5     string
	ArchivedSize    int
	ArchivedURI     string `json:"ArchivedUri"`
	CanEdit         bool
	Caption         string
	Collectable     bool
	Date            *time.Time
	EZProject       bool
	FileName        string
	Format          string
	FormattedValues *FormattedValues
	Hidden          bool
	ImageKey        string
	IsArchive       bool
	IsVideo         bool
	KeywordArray    []string
	Keywords        string
	LastUpdated     *time.Time
	Latitude        string
	Longitude       string
	OriginalHeight  int
	OriginalSize    int
	OriginalWidth   int
	Processing      bool
	Protected       bool
	ThumbnailURL    string `json:"ThumbnailUrl"`
	Title           string
	UploadKey       string
	Watermarked     bool

	ResponseLevel string
	URI           string `json:"Uri"`
	URIs          *URIs  `json:"Uris"`
	WebURI        string `json:"WebUri"`
}

type ImageDownload struct {
	URL string `json:"Url"`

	URI            string `json:"Uri"`
	URIDescription string `json:"UriDescription"`
}

type ImageMetadata struct {
	Altitude               string
	AltitudeReference      string
	Aperture               float64
	AudioCodec             string
	Author                 string
	AuthorTitle            string
	Brightness             string
	Caption                string
	Category               string
	CircleOfConfusion      string
	City                   string
	ColorSpace             string
	CompressedBitsPerPixel string
	Contrast               string
	Copyright              string
	CopyrightFlag          string
	CopyrightURL           string `json:"CopyrightUrl"`
	Country                string
	CountryCode            string
	CreatorContactInfo     string
	Credit                 string
	DateCreated            string // *time.Time
	DateDigitized          string // *time.Time
	DateTimeCreated        string // *time.Time
	DateTimeModified       string // *time.Time
	DepthOfField           string
	DigitalZoomRatio       float64
	Duration               string
	Exposure               string
	ExposureCompensation   string
	ExposureMode           string
	ExposureProgram        string
	FieldOfView            string
	Flash                  string
	FocalLength            string
	FocalLength35mm        string
	GainControl            string
	Headline               string
	HyperfocalDistance     string
	ISO                    int
	Keywords               string
	Latitude               float64
	LatitudeReference      string
	Lens                   string
	LensSerialNumber       string
	LightSource            string
	Longitude              float64
	LongitudeReference     string
	Make                   string
	Metering               string
	MicroDateTimeCreated   string // *time.Time
	MicroDateTimeDigitized string // *time.Time
	Model                  string
	NormalizedLightValue   float64
	Rating                 string
	Saturation             string
	ScaleFactor            string
	SceneCaptureType       string
	SensingMethod          string
	SerialNumber           string
	Sharpness              string
	Software               string
	Source                 string
	SpecialInstructions    string
	State                  string
	SubjectDistance        string
	SubjectRange           string
	SupplementalCategories string
	TimeCreated            string
	Title                  string
	TransmissionReference  string
	UserComment            string
	VideoCodec             string
	WhiteBalance           string
	WriterEditor           string

	ResponseLevel  string
	URI            string `json:"Uri"`
	URIDescription string `json:"UriDescription"`
}

type CatalogSkuPrice struct {
	Currency string
	Price    float64

	ResponseLevel  string
	URI            string `json:"Uri"`
	URIDescription string `json:"UriDescription"`
}

type ImageSize struct {
	URL    string `json:"Url"`
	Ext    string
	Height int
	Width  int
	Size   int
}

type ImageSizeDetails struct {
	ImageSizeLarge    *ImageSize
	ImageSizeMedium   *ImageSize
	ImageSizeOriginal *ImageSize
	ImageSizeSmall    *ImageSize
	ImageSizeThumb    *ImageSize
	ImageSizeTiny     *ImageSize
	ImageSizeX2Large  *ImageSize
	ImageSizeX3Large  *ImageSize
	ImageSizeXLarge   *ImageSize
	ImageURLTemplate  string `json:"ImageUrlTemplate"`
	UsableSizes       []string

	URI            string `json:"Uri"`
	URIDescription string `json:"UriDescription"`
}

type ImageSizes struct {
	LargeImageURL    string `json:"LargeImageUrl"`
	LargestImageURL  string `json:"LargestImageUrl"`
	MediumImageURL   string `json:"MediumImageUrl"`
	OriginalImageURL string `json:"OriginalImageUrl"`
	SmallImageURL    string `json:"SmallImageUrl"`
	ThumbImageURL    string `json:"ThumbImageUrl"`
	TinyImageURL     string `json:"TinyImageUrl"`
	X2LargeImageURL  string `json:"X2LargeImageUrl"`
	X3LargeImageURL  string `json:"X3LargeImageUrl"`
	XLargeImageURL   string `json:"XLargeImageUrl"`

	URI            string `json:"Uri"`
	URIDescription string `json:"UriDescription"`
}

type LargestImage struct {
	Ext         string
	Height      int
	Size        int
	URL         string `json:"Url"`
	Usable      bool
	Watermarked bool
	Width       int

	URI            string `json:"Uri"`
	URIDescription string `json:"UriDescription"`
}
