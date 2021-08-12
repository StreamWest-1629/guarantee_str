// regexp_guarantee.go
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

import (
	"errors"
	"regexp"
)

type regex regexp.Regexp

func (r *regex) Filter(checkStr string) error {
	if found := (*regexp.Regexp)(r).FindStringIndex(checkStr); len(found) != 2 {
		return errors.New("string value isn't match with regular expression")
	} else if found[0] != 0 || found[1] != len(checkStr) {
		return errors.New("string value isn't match with regular expression")
	} else {
		return nil
	}
}

// Make new filter with regular expression.
func MakeRegexpFilter(re *regexp.Regexp) *StringFilter {
	return (FilterRule)((*regex)(re).Filter).MakeFilter()
}
