package utils

import (
	"fmt"
	"bytes"
	"io"
	"os/exec"
	"log"
)

func execCommand(cmd *exec.Cmd, prefix string) (bool, *string, *string) {
	out, err := cmd.StdoutPipe()
	outE, errE := cmd.StderrPipe()

	if err != nil {
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, out))
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, err))
		return false, nil, nil
	}

	if errE != nil {
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, outE))
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, errE))
		return false, nil, nil
	}

	err = cmd.Start()
	if err != nil {
		log.Println("[command] %s err: %s", prefix, err)
		return false, nil, nil
	}
	outResult := copyStream(out, prefix)
	errResult := copyStream(outE, prefix)

	err = cmd.Wait()
	if err != nil {
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, err))
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
		log.Println(fmt.Sprintf("[command] %s output: %s", prefix, tmpBuf[0:n]))
	}
	if err == io.EOF {
		err = nil
	} else {
		log.Println("ERROR: " + err.Error())
	}
	return buffer.String()
}

func RunWalter(config string) bool {
	cmd := exec.Command("../bin/walter", "-c", config)
	cmd.Dir = "."
	result, _, _ := execCommand(cmd, "exec")
	return result
}

func RunWalterWithForthOption(config string) bool {
	cmd := exec.Command("../bin/walter", "-c", config, "-f", "true")
	cmd.Dir = "."
	result, _, _ := execCommand(cmd, "exec")
	return result
}
