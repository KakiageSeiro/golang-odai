package firebase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.opencensus.io/plugin/ochttp"
	"io"
	"net/http"
	"strings"
)

var apiURI = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/%s?key=%s"
var tokenURI = "https://securetoken.googleapis.com/v1/token?key=%s"

type Config struct {
	ApiKEY string `required:"true" env:"FIREBASE_API_KEY"`
}

// Firebase Firebase認証操作に必要な情報を保持
type Firebase struct {
	apiKey string
}

// New returns a Authentication
func New(apiKey string) *Firebase {
	return &Firebase{
		apiKey: apiKey,
	}
}

// Post Firebase APIにPOSTして、レスポンスを型に変換
func (f *Firebase) Post(ctx context.Context, service string, data interface{}, resp interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	r, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(apiURI, service, f.apiKey),
		strings.NewReader(string(b)),
	)
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", "application/json")


	client := &http.Client{Transport: &ochttp.Transport{}}
	res, err := client.Do(r.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, res.Body); err != nil {
		return err
	}

	if res.StatusCode == http.StatusBadRequest {
		var e ErrorResponse
		if err := json.Unmarshal(buf.Bytes(), &e); err != nil {
			return err
		}

		return e.Error
	}

	return json.Unmarshal(buf.Bytes(), &resp)
}