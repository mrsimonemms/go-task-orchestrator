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
	"bytes"
	"fmt"
	"os"
	"path"
	"reflect"
	"text/template"

	"sigs.k8s.io/yaml"
)

var ErrConflict = fmt.Errorf("conflicting commands received")

type Command interface {
	exec(c *Config) error
}

type TaskFile struct {
	Commands []TaskCommand `json:"commands"`
}

type TaskCommand struct {
	Copy *Copy `json:"copy,omitempty"`
	Move *Move `json:"move,omitempty"`
}

func (t *TaskCommand) exec(c *Config) error {
	v := reflect.ValueOf(*t)

	// Ensure only one command set
	var command Command
	for i := 0; i < v.NumField(); i++ {
		isNil := v.Field(i).IsNil()
		if !isNil {
			if command != nil {
				return ErrConflict
			}
			command = v.Field(i).Interface().(Command)
		}
	}

	return command.exec(c)
}

type Config struct {
	Root      string
	TaskFile  string
	Variables map[string]string

	parsedTaskFile bytes.Buffer
}

func (c *Config) Exec() error {
	// Load the task file
	f, err := os.ReadFile(path.Join(c.Root, c.TaskFile))
	if err != nil {
		return err
	}

	t := template.New("tpl")
	t, err = t.Parse(string(f))
	if err != nil {
		return err
	}

	if err := t.Execute(&c.parsedTaskFile, c.Variables); err != nil {
		return err
	}

	taskFile := &TaskFile{}
	if err := yaml.Unmarshal(c.parsedTaskFile.Bytes(), taskFile); err != nil {
		return err
	}

	for _, cmd := range taskFile.Commands {
		if err := cmd.exec(c); err != nil {
			return err
		}
	}

	return nil
}
