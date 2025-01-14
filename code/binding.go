// Copyright 2016 CoreOS, Inc.
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

package code

import (
	"fmt"
	"io"
)

type Binding struct {
	pkg    string
	fppath string
	fps    []*Failpoint
}

func NewBinding(pkg string, fppath string, fps []*Failpoint) *Binding {
	return &Binding{pkg, fppath, fps}
}

// Write writes the fp.generated.go file for a package.
func (b *Binding) Write(dst io.Writer) error {
	hdr := "// GENERATED BY GOFAIL. DO NOT EDIT.\n\n" +
		"package " + b.pkg +
		"\n\nimport \"go.etcd.io/gofail/runtime\"\n\n"
	if _, err := fmt.Fprint(dst, hdr); err != nil {
		return err
	}
	for _, fp := range b.fps {
		_, err := fmt.Fprintf(
			dst,
			"var %s *runtime.Failpoint = runtime.NewFailpoint(%q, %q)\n",
			fp.Runtime(),
			b.fppath,
			fp.Name(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
