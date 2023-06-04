/*
   Copyright 2022 Go-Flow Authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package flow

import (
	"context"
	"github.com/basenana/go-flow/utils"
)

type Context struct {
	context.Context
	utils.Logger

	FlowId    string
	TaskName  string
	Message   string
	MaxRetry  int
	IsSucceed bool
}

func (c *Context) Succeed() {
	c.IsSucceed = true
}

func (c *Context) Fail(message string, maxRetry int) {
	c.Message = message
	c.MaxRetry = maxRetry
}
