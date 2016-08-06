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

package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

// Supported hardware raid modes
type Configuration_Storage_Group_Array_Type struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationStorageGroupArrayTypeService() Configuration_Storage_Group_Array_Type {
	return Configuration_Storage_Group_Array_Type{Session: r}
}

//
func (r *Configuration_Storage_Group_Array_Type) GetAllObjects() (resp []datatypes.Configuration_Storage_Group_Array_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Storage_Group_Array_Type) GetHardwareComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Storage_Group_Array_Type) GetObject() (resp datatypes.Configuration_Storage_Group_Array_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Configuration_Template data type contains general information of an arbitrary resource.
type Configuration_Template struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateService() Configuration_Template {
	return Configuration_Template{Session: r}
}

// Copy a configuration template and returns a newly created template copy
func (r *Configuration_Template) CopyTemplate(templateObject *datatypes.Configuration_Template) (resp datatypes.Configuration_Template, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Deletes a customer configuration template.
func (r *Configuration_Template) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Edit the object by passing in a modified instance of the object. Use this method to modify configuration template name or description.
func (r *Configuration_Template) EditObject(templateObject *datatypes.Configuration_Template) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves all available configuration templates
func (r *Configuration_Template) GetAllObjects() (resp []datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetConfigurationSections() (resp []datatypes.Configuration_Template_Section, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetConfigurationTemplateReference() (resp []datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetDefaultValues() (resp []datatypes.Configuration_Template_Section_Definition_Value, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetDefinitions() (resp []datatypes.Configuration_Template_Section_Definition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetLinkedSectionReferences() (resp datatypes.Configuration_Template_Section_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Template) GetObject() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetParent() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Updates default configuration values.
func (r *Configuration_Template) UpdateDefaultValues(configurationValues []datatypes.Configuration_Template_Section_Definition_Value) (resp bool, err error) {
	params := []interface{}{
		configurationValues,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Configuration_Template_Section data type contains information of a configuration section.
//
// Configuration can contain sub-sections.
type Configuration_Template_Section struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionService() Configuration_Template_Section {
	return Configuration_Template_Section{Session: r}
}

// Retrieve
func (r *Configuration_Template_Section) GetDefinitions() (resp []datatypes.Configuration_Template_Section_Definition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetDisallowedDeletionFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetLinkedTemplate() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetLinkedTemplateReference() (resp datatypes.Configuration_Template_Section_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Template_Section) GetObject() (resp datatypes.Configuration_Template_Section, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetProfiles() (resp []datatypes.Configuration_Template_Section_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetSectionType() (resp datatypes.Configuration_Template_Section_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetSectionTypeName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetSubSections() (resp []datatypes.Configuration_Template_Section, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section) GetTemplate() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns true if the object has sub-sections
func (r *Configuration_Template_Section) HasSubSections() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Configuration definition gives you details of the value that you're setting.
//
// Some monitoring agents requires values unique to your system. If value type is defined as "Resource Specific Values", you will have to make an additional API call to retrieve your system specific values.
//
// See [[SoftLayer_Monitoring_Agent::getAvailableConfigurationValues|Monitoring Agent]] service to retrieve your system specific values.
type Configuration_Template_Section_Definition struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionDefinitionService() Configuration_Template_Section_Definition {
	return Configuration_Template_Section_Definition{Session: r}
}

// Retrieve
func (r *Configuration_Template_Section_Definition) GetAttributes() (resp []datatypes.Configuration_Template_Section_Definition_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition) GetDefaultValue() (resp datatypes.Configuration_Template_Section_Definition_Value, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition) GetGroup() (resp datatypes.Configuration_Template_Section_Definition_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition) GetMonitoringDataFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Template_Section_Definition) GetObject() (resp datatypes.Configuration_Template_Section_Definition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition) GetSection() (resp datatypes.Configuration_Template_Section, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition) GetValueType() (resp datatypes.Configuration_Template_Section_Definition_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Configuration definition group gives you details of the definition and allows extra functionality.
//
//
type Configuration_Template_Section_Definition_Group struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionDefinitionGroupService() Configuration_Template_Section_Definition_Group {
	return Configuration_Template_Section_Definition_Group{Session: r}
}

// Get all configuration definition group objects.
//
// ''getAllGroups'' returns an array of SoftLayer_Configuration_Template_Section_Definition_Group objects upon success.
func (r *Configuration_Template_Section_Definition_Group) GetAllGroups() (resp []datatypes.Configuration_Template_Section_Definition_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Template_Section_Definition_Group) GetObject() (resp datatypes.Configuration_Template_Section_Definition_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition_Group) GetParent() (resp datatypes.Configuration_Template_Section_Definition_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Configuration_Template_Section_Definition_Type further defines the value of a configuration definition.
type Configuration_Template_Section_Definition_Type struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionDefinitionTypeService() Configuration_Template_Section_Definition_Type {
	return Configuration_Template_Section_Definition_Type{Session: r}
}

//
func (r *Configuration_Template_Section_Definition_Type) GetObject() (resp datatypes.Configuration_Template_Section_Definition_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Configuration_Section_Value is used to set the value for a configuration definition
type Configuration_Template_Section_Definition_Value struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionDefinitionValueService() Configuration_Template_Section_Definition_Value {
	return Configuration_Template_Section_Definition_Value{Session: r}
}

// Retrieve
func (r *Configuration_Template_Section_Definition_Value) GetDefinition() (resp datatypes.Configuration_Template_Section_Definition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Template_Section_Definition_Value) GetObject() (resp datatypes.Configuration_Template_Section_Definition_Value, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Definition_Value) GetTemplate() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Some configuration templates let you create a unique configuration profiles.
//
// For example, you can create multiple configuration profiles to monitor multiple hard drives with "CPU/Memory/Disk Monitoring Agent". SoftLayer_Configuration_Template_Section_Profile help you keep track of custom configuration profiles.
type Configuration_Template_Section_Profile struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionProfileService() Configuration_Template_Section_Profile {
	return Configuration_Template_Section_Profile{Session: r}
}

// Retrieve
func (r *Configuration_Template_Section_Profile) GetConfigurationSection() (resp datatypes.Configuration_Template_Section, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Profile) GetMonitoringAgent() (resp datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Configuration_Template_Section_Profile) GetObject() (resp datatypes.Configuration_Template_Section_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Configuration_Template_Section_Reference data type contains information of a configuration section and its associated configuration template.
type Configuration_Template_Section_Reference struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionReferenceService() Configuration_Template_Section_Reference {
	return Configuration_Template_Section_Reference{Session: r}
}

//
func (r *Configuration_Template_Section_Reference) GetObject() (resp datatypes.Configuration_Template_Section_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Reference) GetSection() (resp datatypes.Configuration_Template_Section, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Configuration_Template_Section_Reference) GetTemplate() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Configuration_Template_Section_Type data type contains information of a configuration section type.
//
// Configuration can contain sub-sections.
type Configuration_Template_Section_Type struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateSectionTypeService() Configuration_Template_Section_Type {
	return Configuration_Template_Section_Type{Session: r}
}

//
func (r *Configuration_Template_Section_Type) GetObject() (resp datatypes.Configuration_Template_Section_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Configuration_Template_Type data type contains configuration template type information.
type Configuration_Template_Type struct {
	Session *Session
	Options
}

func (r *Session) GetConfigurationTemplateTypeService() Configuration_Template_Type {
	return Configuration_Template_Type{Session: r}
}

//
func (r *Configuration_Template_Type) GetObject() (resp datatypes.Configuration_Template_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
