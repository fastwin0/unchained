/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BlockHeaderVersion struct for BlockHeaderVersion
type BlockHeaderVersion struct {
	Block string `json:"block"`
	App string `json:"app"`
}

// NewBlockHeaderVersion instantiates a new BlockHeaderVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockHeaderVersion(block string, app string) *BlockHeaderVersion {
	this := BlockHeaderVersion{}
	this.Block = block
	this.App = app
	return &this
}

// NewBlockHeaderVersionWithDefaults instantiates a new BlockHeaderVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockHeaderVersionWithDefaults() *BlockHeaderVersion {
	this := BlockHeaderVersion{}
	return &this
}

// GetBlock returns the Block field value
func (o *BlockHeaderVersion) GetBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderVersion) GetBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *BlockHeaderVersion) SetBlock(v string) {
	o.Block = v
}

// GetApp returns the App field value
func (o *BlockHeaderVersion) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderVersion) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *BlockHeaderVersion) SetApp(v string) {
	o.App = v
}

func (o BlockHeaderVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["block"] = o.Block
	}
	if true {
		toSerialize["app"] = o.App
	}
	return json.Marshal(toSerialize)
}

type NullableBlockHeaderVersion struct {
	value *BlockHeaderVersion
	isSet bool
}

func (v NullableBlockHeaderVersion) Get() *BlockHeaderVersion {
	return v.value
}

func (v *NullableBlockHeaderVersion) Set(val *BlockHeaderVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeaderVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeaderVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeaderVersion(val *BlockHeaderVersion) *NullableBlockHeaderVersion {
	return &NullableBlockHeaderVersion{value: val, isSet: true}
}

func (v NullableBlockHeaderVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeaderVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


