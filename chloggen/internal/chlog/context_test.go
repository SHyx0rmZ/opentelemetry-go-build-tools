// Copyright The OpenTelemetry Authors
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

package chlog

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	root := "/tmp"
	ctx := New(root)
	require.Equal(t, root, ctx.rootDir)
	require.Equal(t, filepath.Join(root, unreleasedDir), ctx.UnreleasedDir)
	require.Equal(t, filepath.Join(root, changelogMD), ctx.ChangelogMD)
	require.Equal(t, filepath.Join(root, unreleasedDir, templateYAML), ctx.TemplateYAML)
}

func TestWithUnreleasedDir(t *testing.T) {
	root := "/tmp"
	unreleased := ".test"
	ctx := New(root, WithUnreleasedDir(unreleased))
	require.Equal(t, root, ctx.rootDir)
	require.Equal(t, filepath.Join(root, unreleased), ctx.UnreleasedDir)
	require.Equal(t, filepath.Join(root, changelogMD), ctx.ChangelogMD)
	require.Equal(t, filepath.Join(root, unreleased, templateYAML), ctx.TemplateYAML)
}
