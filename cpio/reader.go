/*
 * Copyright (c) SAS Institute, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cpio

import "io"

type Reader struct {
	stream  *CpioStream
	cur_ent *CpioEntry
}

func NewReader(stream io.Reader) *Reader {
	return &Reader{
		stream:  NewCpioStream(stream),
		cur_ent: nil,
	}
}

func (r *Reader) Next() (*cpio_newc_header, error) {
	ent, err := r.stream.ReadNextEntry()
	if err != nil {
		return nil, err
	}
	r.cur_ent = ent
	return r.cur_ent.header, nil
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.cur_ent.payload.Read(p)
}
