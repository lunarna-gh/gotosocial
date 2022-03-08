// Code generated by astool. DO NOT EDIT.

package propertysignaturealgorithm

import (
	"fmt"
	string1 "github.com/superseriousbusiness/activity/streams/values/string"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
	"net/url"
)

// TootSignatureAlgorithmProperty is the functional property "signatureAlgorithm".
// It is permitted to be a single default-valued value type.
type TootSignatureAlgorithmProperty struct {
	xmlschemaStringMember string
	hasStringMember       bool
	unknown               interface{}
	iri                   *url.URL
	alias                 string
}

// DeserializeSignatureAlgorithmProperty creates a "signatureAlgorithm" property
// from an interface representation that has been unmarshalled from a text or
// binary format.
func DeserializeSignatureAlgorithmProperty(m map[string]interface{}, aliasMap map[string]string) (*TootSignatureAlgorithmProperty, error) {
	alias := ""
	if a, ok := aliasMap["http://joinmastodon.org/ns"]; ok {
		alias = a
	}
	propName := "signatureAlgorithm"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "signatureAlgorithm")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &TootSignatureAlgorithmProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := string1.DeserializeString(i); err == nil {
			this := &TootSignatureAlgorithmProperty{
				alias:                 alias,
				hasStringMember:       true,
				xmlschemaStringMember: v,
			}
			return this, nil
		}
		this := &TootSignatureAlgorithmProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewTootSignatureAlgorithmProperty creates a new signatureAlgorithm property.
func NewTootSignatureAlgorithmProperty() *TootSignatureAlgorithmProperty {
	return &TootSignatureAlgorithmProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsXMLSchemaString
// afterwards will return false.
func (this *TootSignatureAlgorithmProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.hasStringMember = false
}

// Get returns the value of this property. When IsXMLSchemaString returns false,
// Get will return any arbitrary value.
func (this TootSignatureAlgorithmProperty) Get() string {
	return this.xmlschemaStringMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this TootSignatureAlgorithmProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this TootSignatureAlgorithmProperty) HasAny() bool {
	return this.IsXMLSchemaString() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this TootSignatureAlgorithmProperty) IsIRI() bool {
	return this.iri != nil
}

// IsXMLSchemaString returns true if this property is set and not an IRI.
func (this TootSignatureAlgorithmProperty) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this TootSignatureAlgorithmProperty) JSONLDContext() map[string]string {
	m := map[string]string{"http://joinmastodon.org/ns": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this TootSignatureAlgorithmProperty) KindIndex() int {
	if this.IsXMLSchemaString() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this TootSignatureAlgorithmProperty) LessThan(o vocab.TootSignatureAlgorithmProperty) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsXMLSchemaString() && !o.IsXMLSchemaString() {
		// Both are unknowns.
		return false
	} else if this.IsXMLSchemaString() && !o.IsXMLSchemaString() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsXMLSchemaString() && o.IsXMLSchemaString() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return string1.LessString(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "signatureAlgorithm".
func (this TootSignatureAlgorithmProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "signatureAlgorithm"
	} else {
		return "signatureAlgorithm"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this TootSignatureAlgorithmProperty) Serialize() (interface{}, error) {
	if this.IsXMLSchemaString() {
		return string1.SerializeString(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsXMLSchemaString afterwards will
// return true.
func (this *TootSignatureAlgorithmProperty) Set(v string) {
	this.Clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *TootSignatureAlgorithmProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}