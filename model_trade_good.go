/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TradeGood type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradeGood{}

// TradeGood A good that can be traded for other goods or currency.
type TradeGood struct {
	Symbol TradeSymbol `json:"symbol"`
	// The name of the good.
	Name string `json:"name"`
	// The description of the good.
	Description string `json:"description"`
}

type _TradeGood TradeGood

// NewTradeGood instantiates a new TradeGood object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeGood(symbol TradeSymbol, name string, description string) *TradeGood {
	this := TradeGood{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	return &this
}

// NewTradeGoodWithDefaults instantiates a new TradeGood object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeGoodWithDefaults() *TradeGood {
	this := TradeGood{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *TradeGood) GetSymbol() TradeSymbol {
	if o == nil {
		var ret TradeSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *TradeGood) GetSymbolOk() (*TradeSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *TradeGood) SetSymbol(v TradeSymbol) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *TradeGood) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TradeGood) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TradeGood) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *TradeGood) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TradeGood) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TradeGood) SetDescription(v string) {
	o.Description = v
}

func (o TradeGood) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeGood) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *TradeGood) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"description",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTradeGood := _TradeGood{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTradeGood)

	if err != nil {
		return err
	}

	*o = TradeGood(varTradeGood)

	return err
}

type NullableTradeGood struct {
	value *TradeGood
	isSet bool
}

func (v NullableTradeGood) Get() *TradeGood {
	return v.value
}

func (v *NullableTradeGood) Set(val *TradeGood) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeGood) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeGood) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeGood(val *TradeGood) *NullableTradeGood {
	return &NullableTradeGood{value: val, isSet: true}
}

func (v NullableTradeGood) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeGood) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


