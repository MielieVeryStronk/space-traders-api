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

// checks if the ShipyardShipTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipyardShipTypesInner{}

// ShipyardShipTypesInner struct for ShipyardShipTypesInner
type ShipyardShipTypesInner struct {
	Type ShipType `json:"type"`
}

type _ShipyardShipTypesInner ShipyardShipTypesInner

// NewShipyardShipTypesInner instantiates a new ShipyardShipTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipyardShipTypesInner(type_ ShipType) *ShipyardShipTypesInner {
	this := ShipyardShipTypesInner{}
	this.Type = type_
	return &this
}

// NewShipyardShipTypesInnerWithDefaults instantiates a new ShipyardShipTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipyardShipTypesInnerWithDefaults() *ShipyardShipTypesInner {
	this := ShipyardShipTypesInner{}
	return &this
}

// GetType returns the Type field value
func (o *ShipyardShipTypesInner) GetType() ShipType {
	if o == nil {
		var ret ShipType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShipyardShipTypesInner) GetTypeOk() (*ShipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipyardShipTypesInner) SetType(v ShipType) {
	o.Type = v
}

func (o ShipyardShipTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipyardShipTypesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ShipyardShipTypesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varShipyardShipTypesInner := _ShipyardShipTypesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipyardShipTypesInner)

	if err != nil {
		return err
	}

	*o = ShipyardShipTypesInner(varShipyardShipTypesInner)

	return err
}

type NullableShipyardShipTypesInner struct {
	value *ShipyardShipTypesInner
	isSet bool
}

func (v NullableShipyardShipTypesInner) Get() *ShipyardShipTypesInner {
	return v.value
}

func (v *NullableShipyardShipTypesInner) Set(val *ShipyardShipTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShipyardShipTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShipyardShipTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipyardShipTypesInner(val *ShipyardShipTypesInner) *NullableShipyardShipTypesInner {
	return &NullableShipyardShipTypesInner{value: val, isSet: true}
}

func (v NullableShipyardShipTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipyardShipTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


