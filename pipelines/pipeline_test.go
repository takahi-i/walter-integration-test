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
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takahi-i/walter-integration-test/utils"
)

func TestRunWalter(t *testing.T) {
	assert.Equal(t, true, utils.RunWalter("pipeline.yml").IsSucceed)
}

func TestRunWalterWithPipelineWithFail(t *testing.T) {
	result := utils.RunWalter("pipeline_fail.yml")
	assert.Equal(t, false, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("Execution is skipped: command_stage_2"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("Execution is skipped: command_stage_3"), *result.ErrResult)
}

func TestRunWalterWithForceOptionWithFailedPipeline(t *testing.T) {
	result := utils.RunWalterWithForthOption("pipeline_fail.yml")
	assert.Equal(t, false, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_2"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_3"), *result.ErrResult)
}

func TestRunWalterWithCleanup(t *testing.T) {
	result := utils.RunWalterWithForthOption("pipeline_with_cleanup.yml")
	assert.Equal(t, true, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_1"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_2"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_3"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec output: hello cleanup"), *result.ErrResult)
}

func TestRunWalterWithFailedPipelineWithCleanup(t *testing.T) {
	result := utils.RunWalterWithForthOption("failed_pipeline_with_cleanup.yml")
	assert.Equal(t, false, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("exec output: hello cleanup"), *result.ErrResult)
}

func TestRunWalterWithConcurrentPipeline(t *testing.T) {
	result := utils.RunWalter("pipeline_with_concurrent_block.yml")
	assert.Equal(t, true, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_1"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_2_group_1"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("exec: command_stage_3_group_1"), *result.ErrResult)
}

func TestRunWalterWithEnvironmentVariables(t *testing.T) {
	result := utils.RunWalter("pipeline_env.yml")
	assert.Equal(t, true, result.IsSucceed)
	assert.Regexp(t, regexp.MustCompile("walter-integration-test/gopath"), *result.ErrResult)
	assert.Regexp(t, regexp.MustCompile("walter-integration-test/bin"), *result.ErrResult)
}
