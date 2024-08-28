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

package task

import (
	"fmt"
	"path"
	"path/filepath"
)

func expandPath(d, p string) ([]string, error) {
	return filepath.Glob(path.Join(d, p))
}

type Copy struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
}

func (e *Copy) exec(c *Config) error {
	source, err := expandPath(c.Root, e.Source)
	if err != nil {
		return err
	}

	for _, s := range source {
		fmt.Println(s)
	}

	return nil
}

type Move struct {
	Source2     string `json:"src"`
	Destination string `json:"dest"`
}

func (e *Move) exec(c *Config) error {
	return nil
}
