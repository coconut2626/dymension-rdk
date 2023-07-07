package rng

import (
	"bytes"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"io/ioutil"
	"net/http"
)

type CoSigReq struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type CoSigRes struct {
	Message    string `json:"message"`
	Signature  string `json:"signature"`
	Randomness string `json:"randomness"`
	PublicKey  string `json:"public_key"`
}

type CoSigClient struct {
	// http client
	client http.Client
	url    string
}

func NewCoSigClient() *CoSigClient {
	return &CoSigClient{
		client: http.Client{},
		url:    "https://cosig-dev.metahouse.casino/v1/cosign",
	}
}

func (c *CoSigClient) GetRandomness(ctx sdk.Context, msg string) []byte {
	rng, err := c.Post(ctx, msg)
	if err != nil {
		ctx.Logger().Error("failed to get randomness", "err", err)
		return nil
	}
	return rng
}

// Post sends a post request to the server
func (c *CoSigClient) Post(ctx sdk.Context, msg string) ([]byte, error) {
	ctx.Logger().Info("sending request to server", "msg", msg)
	reqBody := CoSigReq{
		Signature: "0xffe02a544e0b88c43fe84a9f054cefd0a2340948d45e3c5ffb5ba3fad73f651d154a550a63255cb21806f0a6846f1278f1be80def5375dd18373ce9d0a87b9a100",
		Message:   msg,
	}
	// json Marshall
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	// add header
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ctx.Logger().Debug("response from server", "body", string(body))

	respStruct := CoSigRes{}
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return nil, err
	}
	ctx.Logger().Info("response from server", "randomness", respStruct.Randomness)

	// return body, nil
	return []byte(respStruct.Randomness), nil
}
