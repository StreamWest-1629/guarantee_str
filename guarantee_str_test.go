// guarantee_str_test.go
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

package guarantee_str_test

import (
	"fmt"
	"regexp"

	"github.com/streamwest-1629/guarantee_str"
)

func Example() {

	// number regular expression
	regex := regexp.MustCompile("([0-9]|[1-9][0-9]+)")
	filter := guarantee_str.MakeRegexpFilter(regex)

	numberStr := filter.MustMakeGuarantee("19")

	// numberStr is guaranteed to match with
	// regular expression `([0-9]|[1-9][0-9]+)`.

	fmt.Println("guaranteed to match with regex: " + numberStr.MustCloneString())
	// guaranteed to match with regex: 19

	if numberStr, err := filter.MakeGuarantee("7"); err != nil { // success
		fmt.Println(err.Error())
	} else {
		fmt.Println("guaranteed to match with regex: " + numberStr.MustCloneString())
	}
	// guaranteed to match with regex: 7

	if numberStr, err := filter.MakeGuarantee("092"); err != nil { // fail
		fmt.Println(err.Error())
	} else {
		fmt.Println("guaranteed to match with regex: " + numberStr.MustCloneString())
	}
	// string value isn't match with regular expression

}
