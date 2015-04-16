/* walter: a deployment pipeline template
 * Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
 * (see CONTRIBUTORS.md)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package pipelines

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahi-i/walter-integration-test/utils"
)

func TestRunWalterForService(t *testing.T) {
	result := utils.RunWalterWithServiceMode("pipeline.yml", "../walter-github-sample-for-test")
	assert.Equal(t, false, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("exec output: sh: ho: command not found"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec output: sh: cho: command not found"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec output: fixed, world"), *result.ErrResult)
	os.Remove("../walter-github-sample-for-test/.walter");
}
