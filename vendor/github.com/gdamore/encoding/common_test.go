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
	"bytes"
	"golang.org/x/text/encoding"
	"unicode/utf8"

	. "github.com/smartystreets/goconvey/convey"
)

func verifyMap(enc encoding.Encoding, b byte, r rune) {
	verifyFromUTF(enc, b, r)
	verifyToUTF(enc, b, r)
}

func verifyFromUTF(enc encoding.Encoding, b byte, r rune) {

	encoder := enc.NewEncoder()

	out := make([]byte, 6)
	utf := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(utf, r)

	ndst, nsrc, err := encoder.Transform(out, utf, true)
	So(err, ShouldBeNil)
	So(nsrc, ShouldEqual, len(utf))
	So(ndst, ShouldEqual, 1)
	So(b, ShouldEqual, out[0])
}

func verifyToUTF(enc encoding.Encoding, b byte, r rune) {
	decoder := enc.NewDecoder()

	out := make([]byte, 6)
	nat := []byte{b}
	utf := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(utf, r)

	ndst, nsrc, err := decoder.Transform(out, nat, true)
	So(err, ShouldBeNil)
	So(nsrc, ShouldEqual, 1)
	if !bytes.Equal(utf, out[:ndst]) {
		Printf("UTF expected %v, but got %v for %x\n", utf, out, b)
	}
	So(bytes.Equal(utf, out[:ndst]), ShouldBeTrue)
}
