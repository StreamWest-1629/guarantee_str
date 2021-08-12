// guarantee_str.go
// Copyright (C) 2021 Kasai Koji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package guarantee_str

import "errors"

// The interface for checking expressions. This is used GuaranteeString.
// If checkStr is valid value, this function returns nil. If not, it returns some error.
type FilterRule func(checkStr string) error

// Make StringFilter instance.
func (fn FilterRule) MakeFilter() *StringFilter {
	return &StringFilter{
		stringFilter: &stringFilter{
			filter: fn,
		},
	}
}

// Definates rule whether string value is valid value or not.
// Uses GuaranteeStr's validate.
type StringFilter struct {
	*stringFilter
}

type stringFilter struct {
	filter FilterRule
}

// Contains string value guaranteed to be valid value.
type GuaranteeStr struct {
	*guaranteeStr
}

type guaranteeStr struct {
	*StringFilter
	guaranteed string // contains filtered string value
	isInited   bool   // whether assigned valid str
}

// Make new GuaranteeStr empty instance. If checkStr is invalid value, returns un-initialized instance.
// To check whether that, call IsInit() function of made instance.
func (filter *StringFilter) MustMakeGuarantee(checkStr string) (guaranteed *GuaranteeStr) {
	guaranteed = &GuaranteeStr{
		guaranteeStr: &guaranteeStr{
			StringFilter: filter,
		},
	}

	if err := guaranteed.AssignString(checkStr); err != nil {
		guaranteed.isInited = false
	}
	return
}

// Make new GuaranteeStr instance. If checkStr is invalid value, this function occers error.
func (filter *StringFilter) MakeGuarantee(checkStr string) (guaranteed *GuaranteeStr, err error) {
	if err := filter.filter(checkStr); err != nil {
		return nil, err
	} else {
		return &GuaranteeStr{
			guaranteeStr: &guaranteeStr{
				StringFilter: filter,
				guaranteed:   checkStr,
				isInited:     true,
			},
		}, nil
	}
}

// Make new GuaranteeStr instance. If src argument's string value is invalid for this filter, this function occers error.
func (filter *StringFilter) ChangeGuarantee(src *GuaranteeStr) (converted *GuaranteeStr, err error) {
	return filter.MakeGuarantee(src.guaranteed)
}

// Check whether instance is initialized or not.
func (str *GuaranteeStr) IsInitialized() bool {
	return str.isInited
}

// Check whether instance value is valid value and Assign string value to instance.
// If instance value's filter is equal to argument value's filter, this function skip filter's check. If not, it calls dest.AssignString().
func (dest *GuaranteeStr) Assign(src *GuaranteeStr) error {
	if dest.stringFilter == src.stringFilter {
		dest.guaranteed, dest.isInited = src.guaranteed, src.isInited
		return nil
	} else {
		return dest.AssignString(src.guaranteed)
	}
}

// Clone instance.
func (dest *GuaranteeStr) Clone() *GuaranteeStr {
	return &GuaranteeStr{
		guaranteeStr: &guaranteeStr{
			StringFilter: dest.StringFilter,
			guaranteed:   dest.guaranteed,
			isInited:     dest.isInited,
		},
	}
}

// Check whether argument string value is valid value for this func.
//
// If it is valid value, this function make instance and assign it. If not, it returns un-initialized instance,
// this instance can be assigned with this instance's filter function.
// To check whether that, call IsInit() function of made instance.
func (dest *GuaranteeStr) MustMakeGuarantee(checkStr string) (newInstance *GuaranteeStr) {
	return dest.StringFilter.MustMakeGuarantee(checkStr)
}

// Make new GuaranteeStr instance. If checkStr is invalid value, this function occers error.
func (dest *GuaranteeStr) MakeGuarantee(checkStr string) (newInstance *GuaranteeStr, err error) {
	return dest.StringFilter.MakeGuarantee(checkStr)
}

// Check whether checkStr argument is valid value and Assign string value to instance.
// If is checkStr is valid value, this function returns nil. If not, it returns error.
func (dest *GuaranteeStr) AssignString(checkStr string) error {
	if err := dest.filter(checkStr); err != nil {
		return err
	} else {
		dest.guaranteed = checkStr
		dest.isInited = true
		return nil
	}
}

// Clone string value from instance.
// returned string value is guaranteed to be valid value.
// If instance isn't initialized yet, this function returns error. If not, it returns nil error.
func (str *GuaranteeStr) CloneString() (string, error) {
	if str.isInited {
		buffer := make([]byte, len(str.guaranteed))
		copy(buffer, []byte(str.guaranteed))
		return string(buffer), nil
	} else {
		return "", errors.New("this instance isn't initialized yet")
	}
}

// Clone string value from instance.
// returned string value is guaranteed to be valid value.
// If instance isn't initialized yet, this function returns empty string.
func (str *GuaranteeStr) MustCloneString() string {
	if str.isInited {
		buffer := make([]byte, len(str.guaranteed))
		copy(buffer, []byte(str.guaranteed))
		return string(buffer)
	} else {
		return ""
	}
}
