package main

import (
	"flag"
	"log"

	"github.com/pilwon/go-smugmug"
)

var (
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
)

func Test(s *smugmug.Service) error {
	// res, err := s.Users.GetAuthUser().Expand([]string{"Node"}).Do()
	// res, err := s.Users.Get("cmac").Expand([]string{"Node"}).Do()
	// res, err := s.Nodes.Get("zx4Fx").Expand([]string{"ChildNodes", "ParentNodes", "User"}).Do()
	// res, err := s.Images.Get("SD5BL92-1").Expand([]string{"ImageAlbum", "ImageDownload", "ImageMetadata", "ImageOwner", "ImagePrices", "ImageSizeDetails", "ImageSizes", "LargestImage"}).Do()
	// res, err := s.Albums.Get("rtV8Y").Do().GetImages()
	res, err := s.Albums.Get("kQ3t8P").Expand([]string{"Node", "User"}).Do()
	if err != nil {
		return err
	}
	prettyPrint(res)
	return nil
}

func main() {
	client, err := buildOAuthHTTPClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	if err != nil {
		log.Fatal(err)
	}

	s, err := smugmug.New(client)
	if err != nil {
		log.Fatal(err)
	}

	if err := Test(s); err != nil {
		log.Fatal(err)
	}
}

func init() {
	flag.StringVar(&consumerKey, "consumer-key", "", "OAuth consumer key")
	flag.StringVar(&consumerSecret, "consumer-secret", "", "OAuth consumer secret")
	flag.StringVar(&accessToken, "access-token", "", "OAuth access token")
	flag.StringVar(&accessTokenSecret, "access-token-secret", "", "OAuth access token secret")
	flag.Parse()
}
