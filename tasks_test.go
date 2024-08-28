/*
 * Copyright 2024 Simon Emms <simon@simonemms.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package task_test

import (
	"os"
	"path"
	"testing"

	task "github.com/mrsimonemms/goto"
	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	tests := []struct {
		Name   string
		Config task.Config
	}{
		{
			Name: "Basic",
			Config: task.Config{
				Root:     path.Join(dir, "testdata/basic"),
				TaskFile: "task.yaml",
				Variables: map[string]string{
					"dir": "files",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert := assert.New(t)

			err := test.Config.Exec()

			assert.NoError(err)
		})
	}
}
