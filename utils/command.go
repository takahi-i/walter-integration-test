package utils

import (
	"fmt"
	"bytes"
	"io"
	"os/exec"
	"log"
)

type Result struct {
	IsSucceed bool
	OutResult *string
	ErrResult *string
}

func execCommand(cmd *exec.Cmd, prefix string) (*Result) {
	out, err := cmd.StdoutPipe()
	outE, errE := cmd.StderrPipe()

	if err != nil {
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, out))
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, err))
		return &Result{
			IsSucceed: false,
			OutResult: nil,
			ErrResult: nil}
	}

	if errE != nil {
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, outE))
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, errE))
		return &Result{IsSucceed: false, OutResult: nil, ErrResult: nil}
	}

	err = cmd.Start()
	if err != nil {
		log.Println("[command] %s err: %s", prefix, err)
		return &Result{IsSucceed: false, OutResult: nil, ErrResult: nil}
	}
	outResult := copyStream(out, prefix)
	errResult := copyStream(outE, prefix)

	err = cmd.Wait()
	if err != nil {
		log.Println(fmt.Sprintf("[command] %s err: %s", prefix, err))
		return &Result{IsSucceed: false, OutResult: &outResult, ErrResult: &errResult}
	}
	log.Println("Succeeded to run command")
	return &Result{IsSucceed: true, OutResult: &outResult, ErrResult: &errResult}
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

func RunWalter(config string) *Result {
	cmd := exec.Command("../bin/walter", "-c", config)
	cmd.Dir = "."
	result := execCommand(cmd, "exec")
	return result
}

func RunWalterWithForthOption(config string) *Result {
	cmd := exec.Command("../bin/walter", "-c", config, "-f", "true")
	cmd.Dir = "."
	result := execCommand(cmd, "exec")
	return result
}

func RunWalterWithServiceMode(config string) *Result {
	cmd := exec.Command("../bin/walter", "-c", config, "-mode", "service")
	cmd.Dir = "."
	result := execCommand(cmd, "exec")
	return result
}

