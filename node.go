package smugmug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type NodesService struct {
	s *Service
}

func NewNodesService(s *Service) *NodesService {
	r := &NodesService{s: s}
	return r
}

func (r *NodesService) Get(id string) *NodesGetCall {
	c := &NodesGetCall{s: r.s, urlParams: url.Values{}}
	c.id = id
	return c
}

func (r *NodesService) Create(parentNodeID string, node *Node) *NodesCreateCall {
	c := &NodesCreateCall{s: r.s, urlParams: url.Values{}}
	c.parentNodeID = parentNodeID
	c.node = node
	return c
}

type NodesServiceResponse struct {
	Code     int
	Message  string
	Response struct {
		ServiceResponse
		Node *json.RawMessage
	}
	Expansions map[string]*json.RawMessage `json:",omitempty"`
}

type NodesGetCall struct {
	id string

	s         *Service
	urlParams url.Values
}

func (c *NodesGetCall) Expand(expansions []string) *NodesGetCall {
	c.urlParams.Set("_expand", strings.Join(expansions, ","))
	return c
}

func (c *NodesGetCall) Filter(filter []string) *NodesGetCall {
	c.urlParams.Set("_filter", strings.Join(filter, ","))
	return c
}

func (c *NodesGetCall) doRequest() (*http.Response, error) {
	urls := resolveRelative(c.s.BasePath, "node/"+c.id)
	urls += "?" + encodeURLParams(c.urlParams)
	req, _ := http.NewRequest("GET", urls, nil)
	c.s.setHeaders(req)
	debugRequest(req)
	return c.s.client.Do(req)
}

func (c *NodesGetCall) Do() (*NodesGetResponse, error) {
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	debugResponse(res)
	defer closeBody(res)
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	nodesRes := &NodesServiceResponse{}
	if err := json.NewDecoder(res.Body).Decode(&nodesRes); err != nil {
		return nil, err
	}
	node := &Node{}
	if err := json.Unmarshal(*nodesRes.Response.Node, &node); err != nil {
		return nil, err
	}
	exp, err := unmarshallExpansions(node.URIs, nodesRes.Expansions)
	if err != nil {
		return nil, err
	}
	ret := &NodesGetResponse{
		Node: node,
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	for name, v := range exp {
		switch name {
		case "ChildNodes":
			ret.ChildNodes = v.([]*Node)
		case "HighlightImage":
			ret.HighlightImage = v.(*Image)
		case "ParentNodes":
			ret.ParentNodes = v.([]*Node)
		case "User":
			ret.User = v.(*User)
		}
	}
	return ret, nil
}

type NodesGetResponse struct {
	Node *Node

	ChildNodes []*Node
	// FolderByID *FolderByID // Deprecated
	HighlightImage *Image
	// MoveNodes      *MoveNodes
	// NodeGrants     *NodeGrants
	ParentNodes []*Node
	User        *User

	ServerResponse `json:"-"`
}

type NodesCreateCall struct {
	parentNodeID string
	node         *Node

	s         *Service
	urlParams url.Values
}

func (c *NodesCreateCall) doRequest() (*http.Response, error) {
	urls := resolveRelative(c.s.BasePath, "node/"+c.parentNodeID) + "!children"
	urls += "?" + encodeURLParams(c.urlParams)
	body, err := json.Marshal(c.node)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", urls, bytes.NewReader(body))
	c.s.setHeaders(req)
	debugRequest(req)
	return c.s.client.Do(req)
}

func (c *NodesCreateCall) Do() (*Node, error) {
	if c.parentNodeID == "" {
		return nil, fmt.Errorf("parentNodeID is empty")
	} else if c.node == nil {
		return nil, fmt.Errorf("node is nil")
	}
	res, err := c.doRequest()
	if err != nil {
		return nil, err
		// } else if res.StatusCode != 201 && res.StatusCode != 409 {
		// 	return nil, fmt.Errorf("Failed to create Node (status code: %d)", res.StatusCode)
	}
	debugResponse(res)
	defer closeBody(res)
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	nodesRes := &NodesServiceResponse{}
	if err := json.NewDecoder(res.Body).Decode(&nodesRes); err != nil {
		return nil, err
	}
	node := &Node{}
	if err := json.Unmarshal(*nodesRes.Response.Node, &node); err != nil {
		return nil, err
	}
	return node, nil
}

type Node struct {
	DateAdded             *time.Time       `json:",omitempty"`
	DateModified          *time.Time       `json:",omitempty"`
	Description           string           `json:",omitempty"`
	EffectivePrivacy      string           `json:",omitempty"`
	EffectiveSecurityType string           `json:",omitempty"`
	FormattedValues       *FormattedValues `json:",omitempty"`
	HasChildren           bool             `json:",omitempty"`
	HideOwner             bool             `json:",omitempty"`
	HighlightImageURI     string           `json:"HighlightImageUri,omitempty"`
	IsRoot                bool             `json:",omitempty"`
	Keywords              []string         `json:",omitempty"`
	Name                  string           `json:",omitempty"`
	NodeID                string           `json:",omitempty"`
	Password              string           `json:",omitempty"`
	PasswordHint          string           `json:",omitempty"`
	Privacy               string           `json:",omitempty"`
	SecurityType          string           `json:",omitempty"`
	SmugSearchable        string           `json:",omitempty"`
	SortDirection         string           `json:",omitempty"`
	SortIndex             int              `json:",omitempty"`
	SortMethod            string           `json:",omitempty"`
	Type                  string           `json:",omitempty"`
	URLName               string           `json:"UrlName,omitempty"`
	URLPath               string           `json:"UrlPath,omitempty"`
	WorldSearchable       string           `json:",omitempty"`

	ResponseLevel string `json:",omitempty"`
	URI           string `json:"Uri,omitempty"`
	URIs          *URIs  `json:"Uris,omitempty"`
	WebURI        string `json:"WebUri,omitempty"`
}
