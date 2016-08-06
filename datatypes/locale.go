/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

//
type Locale struct {
	Entity

	//
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Internal identification number of a locale
	Id *int `json:"id,omitempty"`

	//
	LanguageTag *string `json:"languageTag,omitempty"`

	// Locale name
	Name *string `json:"name,omitempty"`
}

//
type Locale_Country struct {
	Entity

	// Binary flag denoting if this country is part of the European Union
	IsEuropeanUnionFlag *int `json:"isEuropeanUnionFlag,omitempty"`

	//
	LongName *string `json:"longName,omitempty"`

	//
	ShortName *string `json:"shortName,omitempty"`

	// A count of states that belong to this country.
	StateCount *uint `json:"stateCount,omitempty"`

	// States that belong to this country.
	States []Locale_StateProvince `json:"states,omitempty"`
}

// This object represents a state or province for a country.
type Locale_StateProvince struct {
	Entity

	//
	LongName *string `json:"longName,omitempty"`

	//
	ShortName *string `json:"shortName,omitempty"`
}

// Each User is assigned a timezone allowing for a precise local timestamp.
type Locale_Timezone struct {
	Entity

	// A timezone's identifying number.
	Id *int `json:"id,omitempty"`

	// A timezone's long name. For example, "(GMT-06:00) America/Dallas - CST".
	LongName *string `json:"longName,omitempty"`

	// A timezone's name. For example, "America/Dallas".
	Name *string `json:"name,omitempty"`

	// A timezone's offset based on the GMT standard. For example, Central Standard Time's offset is "-0600" from GMT=0000.
	Offset *string `json:"offset,omitempty"`

	// A timezone's common abbreviation. For example, Central Standard Time's abbreviation is "CST".
	ShortName *string `json:"shortName,omitempty"`
}
