package netAPI

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"strings"
)

type APIError struct {
	ErrorCode    int    `json:"errorcode"`
	APIErrorCode int    `json:"cserrorcode"`
	ErrorText    string `json:"errortext"`
}

func (e *APIError) Error() error {
	return fmt.Errorf("API error %d (CSExceptionErrorCode: %d): %s", e.ErrorCode, e.APIErrorCode, e.ErrorText)
}

type NetAPIClient struct {
	client  *http.Client // The http client for communicating
	baseURL string       // The base URL of the API
	apiKey  string       // Root api key
	secret  string       // Root secret key
	acronym string       // Account acronym for the account that we will deploy into

	Network              *NetworkService
	PrivateDirectConnect *PrivateDirectConnectService
	PublicDirectConnect  *PublicDirectConnectService
	DirectConnectGroup   *DirectConnectGroupService
}

// Creates a new client for communicating with Net API
func NewClient(apiurl, apikey, secret, acronym string, verifyssl bool) *NetAPIClient {
	cs := &NetAPIClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifyssl}, // If verifyssl is true, skipping the verify should be false and vice versa
			},
		},
		baseURL: apiurl,
		apiKey:  apikey,
		secret:  secret,
		acronym: acronym,
	}
	cs.Network = NewNetworkService(cs)
	cs.PrivateDirectConnect = NewPrivateDirectConnectService(cs)
	cs.PublicDirectConnect = NewPublicDirectConnectService(cs)
	cs.DirectConnectGroup = NewDirectConnectGroupService(cs)
	return cs
}

// we must return the mock server so that we are able to shut it down after we finished using it
func NewMockServerAndClient(code int, body string) (*httptest.Server, *NetAPIClient) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	cs := &NetAPIClient{
		client:  httpClient,
		baseURL: server.URL,
		apiKey:  "mockKey",
		secret:  "mockSecret",
		acronym: "MOCKA",
	}
	cs.Network = NewNetworkService(cs)
	cs.PrivateDirectConnect = NewPrivateDirectConnectService(cs)
	cs.PublicDirectConnect = NewPublicDirectConnectService(cs)
	cs.DirectConnectGroup = NewDirectConnectGroupService(cs)
	return server, cs

}

// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occured. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *NetAPIClient) newRequest(api string, params url.Values) (json.RawMessage, error) {
	params.Set("apiKey", cs.apiKey)
	params.Set("command", api)
	params.Set("response", "json")
	// acronym is optional - req for root key but not for myservices key
	if cs.acronym != "" {
		params.Set("acronym", cs.acronym)
	}

	// Generate signature for API call
	// * Serialize parameters, URL encoding only values and sort them by key, done by encodeValues
	// * Convert the entire argument string to lowercase
	// * Replace all instances of '+' to '%20'
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	s := encodeValues(params)
	s2 := strings.ToLower(s)
	s3 := strings.Replace(s2, "+", "%20", -1)
	mac := hmac.New(sha1.New, []byte(cs.secret))
	mac.Write([]byte(s3))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	var err error
	var resp *http.Response

	// Create the final URL before we issue the request
	url := cs.baseURL + "?" + s + "&signature=" + url.QueryEscape(signature)

	// Make a GET call
	resp, err = cs.client.Get(url)

	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	//log.Printf("[DEBUG] Raw JSON response: %s \n", b)

	// Need to get the raw value to make the result play nice
	b, err = getRawValue(b)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e APIError
		if err := json.Unmarshal(b, &e); err != nil {
			return nil, err
		}
		return nil, e.Error()
	}
	return b, nil
}

// Custom version of net/url Encode that only URL escapes values
// Unmodified portions here remain under BSD license of The Go Authors: https://go.googlesource.com/go/+/master/LICENSE
func encodeValues(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

// Generic function to get the first raw value from a response as json.RawMessage
func getRawValue(b json.RawMessage) (json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for _, v := range m {
		return v, nil
	}
	return nil, fmt.Errorf("Unable to extract the raw value from:\n\n%s\n\n", string(b))
}

type NetworkService struct {
	cs *NetAPIClient
}

func NewNetworkService(cs *NetAPIClient) *NetworkService {
	return &NetworkService{cs: cs}
}

type PrivateDirectConnectService struct {
	cs *NetAPIClient
}

func NewPrivateDirectConnectService(cs *NetAPIClient) *PrivateDirectConnectService {
	return &PrivateDirectConnectService{cs: cs}
}

type PublicDirectConnectService struct {
	cs *NetAPIClient
}

func NewPublicDirectConnectService(cs *NetAPIClient) *PublicDirectConnectService {
	return &PublicDirectConnectService{cs: cs}
}

type DirectConnectGroupService struct {
	cs *NetAPIClient
}

func NewDirectConnectGroupService(cs *NetAPIClient) *DirectConnectGroupService {
	return &DirectConnectGroupService{cs: cs}
}
