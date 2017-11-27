// Copyright (c) 2017 Oliver Schneider
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// FileStore is an abstraction layer for file access
type FileStore struct {
	inputFile *os.File
}

func NewFileStore(inputFile string) (*FileStore, error) {
	file, err := os.Open("./result.json")
	return &FileStore{file}, err
}

func (fileStore *FileStore) readIPs() <-chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			var grab Grab

			jsonBlob := scanner.Bytes()
			err := json.Unmarshal(jsonBlob, &grab)
			if err != nil {
				fmt.Println("error:", err)
			}

			if !isRouter(grab.Data.Banner) {
				continue
			}
			out <- grab.IP
		}
		close(out)
	}()
	return out
}
