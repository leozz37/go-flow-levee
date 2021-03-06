// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sourcetest

// type alias
type Alias = Source

// type definition where the underlying type is a Source
type Definition Source

// This function allows us to consume multiple arguments in a single line so this file can compile
func noop(args ...interface{}) {}

func TestSourceDeclarations() {
	var varZeroVal Source                         // want "source identified"
	declZeroVal := Source{}                       // want "source identified"
	populatedVal := Source{Data: "secret", ID: 0} // want "source identified"

	// We do not want a "source identified" here, since this is nil
	// and gets optimized out when the SSA is built.
	var constPtr *Source

	var ptr *Source
	// We do want a "source identified" here.
	// ptr does not get optimized out because it gets assigned.
	ptr = &Source{}                                       // want "source identified"
	newPtr := new(Source)                                 // want "source identified"
	ptrToDeclZero := &Source{}                            // want "source identified"
	ptrToDeclPopulataed := &Source{Data: "secret", ID: 1} // want "source identified"

	alias := Alias{} // want "source identified"
	def := Definition{}

	noop(varZeroVal, declZeroVal, populatedVal, constPtr, ptr, newPtr, ptrToDeclZero, ptrToDeclPopulataed, alias, def)
}

// A report should be emitted for each parameter, as well as the (implicit) Alloc for val.
func TestSourceParameters(val Source, ptr *Source) { // want "source identified" "source identified" "source identified"

}

func TestSourceExtracts() {
	s, err := CreateSource() // want "source identified"
	sptr, err := NewSource() // want "source identified"

	// we expect two reports for the following cases, since the map is a Source
	mapSource, ok := map[string]Source{}[""]     // want "source identified" "source identified"
	mapSourcePtr, ok := map[string]*Source{}[""] // want "source identified" "source identified"

	// we expect two reports for the following cases, since the chan is a Source
	chanSource, ok := <-(make(chan Source))     // want "source identified" "source identified"
	chanSourcePtr, ok := <-(make(chan *Source)) // want "source identified" "source identified"

	_, _, _, _, _, _, _, _ = s, sptr, mapSource, chanSource, mapSourcePtr, chanSourcePtr, err, ok
}

func TestCollections(ss []Source) { // want "source identified"
	_ = map[Source]string{} // want "source identified"
	_ = map[string]Source{} // want "source identified"
	_ = map[Source]Source{} // want "source identified"
	_ = [1]Source{}         // want "source identified"
	_ = []Source{}          // want "source identified"
	_ = make(chan Source)   // want "source identified"
	_ = []*Source{}         // want "source identified"
	_ = SourceSlice{}       // want "source identified"
	_ = DeeplyNested{}      // want "source identified"
}

type SourceSlice []Source
type DeeplyNested map[string][]map[string]map[string][][]map[string]Source
