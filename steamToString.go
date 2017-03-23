/*
*
* Author: Hui Ye - <bonjovis@163.com>
*
* Last modified: 2017-03-22 01:57
*
* Filename: steamToString.go
*
* Copyright (c) 2016 JOVI
*
 */
package awsservice

import "bytes"
import "io"

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
