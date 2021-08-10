// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package il

import (
	"strconv"

	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
)

// Schemas bundles a property's Terraform and Pulumi schema information into a single type. This information is then
// used to determine type and name information for the property. If the Terraform property is of a composite type--a
// map, list, or set--the property's schemas may also be used to access child schemas.
type Schemas struct {
	TF     shim.Schema
	TFRes  shim.Resource
	Pulumi *tfbridge.SchemaInfo
}

// PropertySchemas returns the Schemas for the child property with the given name. If the name is an integer, this
// function returns the value of a call to ElemSchemas.
func (s Schemas) PropertySchemas(key string) Schemas {
	var propSch Schemas

	if _, err := strconv.ParseInt(key, 0, 0); err == nil {
		return s.ElemSchemas()
	}

	if s.TFRes != nil && s.TFRes.Schema() != nil {
		propSch.TF = s.TFRes.Schema().Get(key)
	}

	if propSch.TF != nil {
		if propResource, ok := propSch.TF.Elem().(shim.Resource); ok {
			propSch.TFRes = propResource
		}
	}

	if s.Pulumi != nil && s.Pulumi.Fields != nil {
		propSch.Pulumi = s.Pulumi.Fields[key]
	}

	return propSch
}

// ElemSchemas returns the element Schemas for a list property.
func (s Schemas) ElemSchemas() Schemas {
	var elemSch Schemas

	if s.TF != nil {
		switch e := s.TF.Elem().(type) {
		case shim.Schema:
			elemSch.TF = e
		case shim.Resource:
			elemSch.TFRes = e
		}
	}

	if s.Pulumi != nil {
		elemSch.Pulumi = s.Pulumi.Elem
	}

	return elemSch
}

// Type returns the appropriate bound type for the property associated with these Schemas.
func (s Schemas) Type() Type {
	if s.TF != nil {
		switch s.TF.Type() {
		case shim.TypeBool:
			return TypeBool
		case shim.TypeInt, shim.TypeFloat:
			return TypeNumber
		case shim.TypeString:
			return TypeString
		case shim.TypeList, shim.TypeSet:
			return s.ElemSchemas().Type().ListOf()
		case shim.TypeMap:
			return TypeMap
		default:
			return TypeUnknown
		}
	}

	return TypeUnknown
}

// ModelType returns the appropriate model type for the property associated with these Schemas.
func (s Schemas) ModelType() model.Type {
	if s.TF != nil {
		switch s.TF.Type() {
		case shim.TypeBool:
			return model.BoolType
		case shim.TypeInt, shim.TypeFloat:
			return model.NumberType
		case shim.TypeString:
			return model.StringType
		case shim.TypeList, shim.TypeSet:
			return model.NewListType(s.ElemSchemas().ModelType())
		case shim.TypeMap:
			if s.TFRes == nil {
				return model.NewMapType(model.StringType)
			}
		default:
			if s.TFRes == nil {
				return model.DynamicType
			}
		}
	}

	if s.TFRes != nil {
		properties := map[string]model.Type{}
		s.TFRes.Schema().Range(func(prop string, _ shim.Schema) bool {
			properties[prop] = s.PropertySchemas(prop).ModelType()
			return true
		})
		return model.NewObjectType(properties)
	}

	return model.DynamicType
}
