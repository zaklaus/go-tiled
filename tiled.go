/*
Copyright (c) 2017 Lauris Bukšis-Haberkorns <lauris@nix.lv>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package tiled

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
)

// LoadFromReader function loads tiled map in TMX format from io.Reader
// baseDir is used for loading additional tile data, current directory is used if empty
func LoadFromReader(baseDir string, r io.Reader) (*Map, error) {
	d := xml.NewDecoder(r)

	m := &Map{baseDir: baseDir}
	if err := d.Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}

// LoadFromFile function loads tiled map in TMX format from file
func LoadFromFile(fileName string) (*Map, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dir, err := filepath.Abs(filepath.Dir(fileName))
	if err != nil {
		return nil, err
	}

	return LoadFromReader(dir, f)
}
