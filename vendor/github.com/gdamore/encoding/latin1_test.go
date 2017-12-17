// Copyright 2015 Garrett D'Amore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encoding

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestISO8859_1(t *testing.T) {
	Convey("8859-1 identity transforms", t, func() {
		for i := 0; i < 256; i++ {
			verifyMap(ISO8859_1, byte(i), rune(i))
		}
	})

	Convey("Large UTF maps to ASCIISub", t, func() {
		verifyFromUTF(ISO8859_1, ASCIISub, '㿿')
	})
}
