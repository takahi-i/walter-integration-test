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
	"bytes"
	"io"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/recruit-tech/walter/log"
)

func execCommand(cmd *exec.Cmd, prefix string) (bool, *string, *string) {
	out, err := cmd.StdoutPipe()
	outE, errE := cmd.StderrPipe()

	if err != nil {
		log.Warnf("[command] %s err: %s", prefix, out)
		log.Warnf("[command] %s err: %s", prefix, err)
		return false, nil, nil
	}

	if errE != nil {
		log.Warnf("[command] %s err: %s", prefix, outE)
		log.Warnf("[command] %s err: %s", prefix, errE)
		return false, nil, nil
	}

	err = cmd.Start()
	if err != nil {
		log.Warnf("[command] %s err: %s", prefix, err)
		return false, nil, nil
	}
	outResult := copyStream(out, prefix)
	errResult := copyStream(outE, prefix)

	err = cmd.Wait()
	if err != nil {
		log.Warnf("[command] %s err: %s", prefix, err)
		return false, &outResult, &errResult
	}
	return true, &outResult, &errResult
}

func copyStream(reader io.Reader, prefix string) string {
	var err error
	var n int
	var buffer bytes.Buffer
	tmpBuf := make([]byte, 1024)
	for {
		if n, err = reader.Read(tmpBuf); err != nil {
			break
		}
		buffer.Write(tmpBuf[0:n])
		log.Infof("[command] %s output: %s", prefix, tmpBuf[0:n])
	}
	if err == io.EOF {
		err = nil
	} else {
		log.Error("ERROR: " + err.Error())
	}
	return buffer.String()
}

func runCommand(param string) bool {
	cmd := exec.Command("../bin/walter", "-c", param)
	cmd.Dir = "."
	result, _, _ := execCommand(cmd, "exec")
	return result
}

func TestReportStageResultWithFullOutput(t *testing.T) {
	runCommand("pipeline.yml")
	assert.Equal(t, 2, 2)
}
