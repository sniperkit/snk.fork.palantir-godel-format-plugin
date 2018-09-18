/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2016 Palantir Technologies, Inc.
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

package formatplugin

import (
	"io"
)

type Formatter interface {
	// TypeName returns the type of this Formatter.
	TypeName() (string, error)

	// Format runs the format operation on the provided files. If "list" is true, then the files that would be changed
	// are printed to stdout rather than formatting the files.
	Format(files []string, list bool, projectDir string, stdout io.Writer) error
}

type Factory interface {
	Types() []string
	NewFormatter(typeName string, cfgYMLBytes []byte) (Formatter, error)
	ConfigUpgrader(typeName string) (ConfigUpgrader, error)
}
