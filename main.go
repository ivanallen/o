/*
* Copyright (c) 2018, 1007729991@qq.com.
* All rights reserved.
*
* Redistribution and use in source and binary forms, with or without modification,
* are permitted provided that the following conditions are met:
*
* * Redistributions of source code must retain the above copyright notice, this
*   list of conditions and the following disclaimer.
*
* * Redistributions in binary form must reproduce the above copyright notice, this
*   list of conditions and the following disclaimer in the documentation and/or
*   other materials provided with the distribution.
*
* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
* ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
* WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
* ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
* (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
* LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
* ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
* (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
* SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var once sync.Once
var buf = &bytes.Buffer{}
var port = 8080
var fileType = "json"
var contentType = map[string]string{
	"json": "application/json",
	"html": "text/html",
	"md":   "text/markdown",
	"svg":  "image/svg+xml",
}

func fillBuf() {
	io.Copy(buf, os.Stdin)
}

func o(w http.ResponseWriter, r *http.Request) {
	once.Do(fillBuf)

	t, ok := contentType[fileType]
	if !ok {
		t = contentType["json"]
	}
	w.Header().Set("Content-Type", t)
	w.Write(buf.Bytes())
}

func help() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "    o [file type] [port]\n")
	fmt.Fprintf(os.Stderr, "    o [port] [file type]\n\n")
	fmt.Fprintf(os.Stderr, "example:\n")
	fmt.Fprintf(os.Stderr, "    echo \"<h1>hello world</h1>\" | o 8000 html\n\n")
	fmt.Fprintf(os.Stderr, "file type:\n")
	for k, v := range contentType {
		fmt.Fprintf(os.Stderr, "    %v (%v)\n", k, v)
	}
}

func main() {
	var err error
	port := 8010
	if len(os.Args) == 2 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fileType = os.Args[1]
			port = 8010
		}
	}

	if fileType == "-h" || fileType == "--help" {
		help()
		return
	}

	if len(os.Args) == 3 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fileType = os.Args[1]
			port, err = strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fileType = os.Args[2]
		}
	}

	hostname, _ := os.Hostname()
	fmt.Printf("\x1b[32mClick me\x1b[0m: http://%s:%d/o\n", hostname, port)

	http.HandleFunc("/o", o)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
