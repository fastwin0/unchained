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

// ConsensusParamsResponseResult struct for ConsensusParamsResponseResult
type ConsensusParamsResponseResult struct {
	BlockHeight string `json:"block_height"`
	ConsensusParams NullableConsensusParams `json:"consensus_params"`
}

// NewConsensusParamsResponseResult instantiates a new ConsensusParamsResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsensusParamsResponseResult(blockHeight string, consensusParams NullableConsensusParams) *ConsensusParamsResponseResult {
	this := ConsensusParamsResponseResult{}
	this.BlockHeight = blockHeight
	this.ConsensusParams = consensusParams
	return &this
}

// NewConsensusParamsResponseResultWithDefaults instantiates a new ConsensusParamsResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsensusParamsResponseResultWithDefaults() *ConsensusParamsResponseResult {
	this := ConsensusParamsResponseResult{}
	return &this
}

// GetBlockHeight returns the BlockHeight field value
func (o *ConsensusParamsResponseResult) GetBlockHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *ConsensusParamsResponseResult) GetBlockHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *ConsensusParamsResponseResult) SetBlockHeight(v string) {
	o.BlockHeight = v
}

// GetConsensusParams returns the ConsensusParams field value
// If the value is explicit nil, the zero value for ConsensusParams will be returned
func (o *ConsensusParamsResponseResult) GetConsensusParams() ConsensusParams {
	if o == nil || o.ConsensusParams.Get() == nil {
		var ret ConsensusParams
		return ret
	}

	return *o.ConsensusParams.Get()
}

// GetConsensusParamsOk returns a tuple with the ConsensusParams field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsensusParamsResponseResult) GetConsensusParamsOk() (*ConsensusParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsensusParams.Get(), o.ConsensusParams.IsSet()
}

// SetConsensusParams sets field value
func (o *ConsensusParamsResponseResult) SetConsensusParams(v ConsensusParams) {
	o.ConsensusParams.Set(&v)
}

func (o ConsensusParamsResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["block_height"] = o.BlockHeight
	}
	if true {
		toSerialize["consensus_params"] = o.ConsensusParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConsensusParamsResponseResult struct {
	value *ConsensusParamsResponseResult
	isSet bool
}

func (v NullableConsensusParamsResponseResult) Get() *ConsensusParamsResponseResult {
	return v.value
}

func (v *NullableConsensusParamsResponseResult) Set(val *ConsensusParamsResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConsensusParamsResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConsensusParamsResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsensusParamsResponseResult(val *ConsensusParamsResponseResult) *NullableConsensusParamsResponseResult {
	return &NullableConsensusParamsResponseResult{value: val, isSet: true}
}

func (v NullableConsensusParamsResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsensusParamsResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


