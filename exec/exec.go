package exec

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// 封装一个函数来执行命令
func ExecNoCommand(commandName string, params []string) (string, error) {
	return ExecInfoCommand(commandName, params, false)
}

func ExecCommand(commandName string, params []string) (string, error) {

	return ExecInfoCommand(commandName, params, true)
}

// 封装一个函数来执行命令
func ExecInfoCommand(commandName string, params []string, info bool) (string, error) {
	//执行命令
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, _ := cmd.StdoutPipe()
	errReader, _ := cmd.StderrPipe()
	e := cmd.Start()
	if e != nil {
		return "", e
	}
	infoBuf := bytes.NewBufferString("")
	go func(buf *bytes.Buffer) {
		in := bufio.NewScanner(stdout)
		for in.Scan() {
			cmdRe := ConvertByte2String(in.Bytes(), "UTF8")
			if info {
				fmt.Println(cmdRe)
			}
			buf.WriteString(cmdRe)
		}
	}(infoBuf)
	// 错误日志
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(errReader)
	for scan.Scan() {
		s := scan.Text()
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}
	// 等待命令执行完
	e = cmd.Wait()
	if !cmd.ProcessState.Success() {
		// 执行失败，返回错误信息
		if errBuf.Len() > 0 {
			return "", errors.New(errBuf.String())
		}
		return "", e
	}
	return infoBuf.String(), nil
}

// 开启一个协程来输出错误
func handlerErr(errReader io.ReadCloser) {
	in := bufio.NewScanner(errReader)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "UTF8")
		fmt.Errorf(cmdRe)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 对字符进行转码
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
