// Code generated by go-bindata.
// sources:
// tmpl/golang/echo_module.gogo
// tmpl/golang/include.gogo
// tmpl/golang/service.gogo
// tmpl/golang/services_file.gogo
// tmpl/golang/struct.gogo
// tmpl/golang/structs_file.gogo
// tmpl/swift/struct.goswift
// DO NOT EDIT!

package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _tmplGolangEcho_moduleGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x5d\x4f\xdc\x3a\x10\x7d\xcf\xaf\xf0\x8d\x96\xab\x04\x2d\xde\xfb\xcc\x15\x0f\x40\xab\x42\x5b\x96\x15\xf0\x58\x09\x99\x64\x36\xeb\x92\xd8\xa9\xed\x50\x50\xe4\xff\x5e\x8f\xed\x64\x3f\xca\x96\xaa\xdd\x17\x1c\x7b\x7c\x7c\xe6\xcc\x99\xa1\xef\x4b\x58\x72\x01\x24\x85\x62\x25\xef\x1b\x59\x76\x35\xa4\xd6\xb6\xac\x78\x64\x15\x90\xbe\xa7\x8b\xb0\xb4\xb6\xef\x27\x1a\xd4\x13\x2f\x60\xce\x1a\x20\xc7\x27\x84\xde\x86\x6f\x8a\x1b\xd6\x26\x09\x6f\x5a\xa9\x0c\xc9\x12\xe2\x7e\x29\x88\x42\x96\x5c\x54\xb3\xaf\x5a\x8a\x34\xec\x2d\x1b\x13\x57\x02\xcc\x6c\x65\x4c\x9b\x26\xe1\xbb\xe2\x66\xd5\x3d\xd0\x42\x36\xb3\x9a\x3d\x68\xe3\x9e\x9d\x21\xa9\x34\xc9\x93\x64\xd9\x89\x82\xdc\x40\xc5\xb5\x01\xb5\x4d\xc4\xda\xc8\xe2\x74\x71\x79\xe5\xf9\x67\x4d\xf7\x4c\x0e\xf1\x2e\xfd\xa0\x64\xd7\x4e\x49\x0c\x27\x7b\x6e\xe6\xa4\xf7\x1c\x5c\xfa\x98\xd6\xbf\x6f\x3d\x10\xa2\xf1\x17\xc3\x8e\x87\xc5\xd4\x9f\xd8\x90\x52\x85\x8f\x23\xa0\xe3\x13\x98\x64\xe9\x6c\x17\x3b\xcd\x7d\x6c\xdf\x1f\x11\xc5\x84\x93\x7c\xd2\x80\x59\x05\x22\xa3\xbe\x57\x7e\x4b\x3b\x89\x47\x5c\xba\x90\xda\x78\xbc\xae\x6d\x41\x5d\x00\x2b\x87\xab\xb1\x1c\xe9\x14\x13\xa2\xfb\x03\xd6\x2f\x83\x28\x1d\xb6\xa3\x6d\x5e\xda\xbd\x2a\x8d\xe9\x13\x6d\x54\x57\x98\x28\xda\x1b\xda\x22\x6a\xdf\xff\x46\x6a\xee\x7e\x38\x1e\xdc\xb5\x8f\x76\x30\x43\x66\x56\x5c\x93\xc3\xb7\xb8\xe6\x64\x0b\xd7\xda\xac\x88\xd6\x38\x97\xc2\xc0\xb3\xc9\x09\x28\x25\x55\x4c\x06\xc5\x98\x30\x55\xe9\x73\xd9\x09\x83\x2c\x6a\x10\xe3\xfb\xa7\xaa\xea\x1a\x10\x66\x28\x04\x46\xf3\x25\x71\xed\xb3\x71\xe7\xbf\x78\xd8\x32\xc5\x1a\xed\xfd\x74\x03\xdf\x3a\xd0\x66\x0f\xd7\x1d\x82\xe3\x23\x7d\xb4\x51\x09\xae\x8d\x40\x21\x12\x36\x12\x9d\xc3\xf7\x77\x61\x2b\x2b\x68\x84\xce\x72\x7a\x26\xcb\x97\x50\x51\x47\xc9\xe5\x84\xf1\xf1\x2a\x0d\xf1\x59\xa0\x94\xff\xef\x8f\xff\x39\x21\x82\xd7\x64\xed\x65\xd7\x9a\x74\xa1\xb8\x30\xcb\x57\x6c\x3a\xdb\x61\x79\x4c\xb8\x78\x62\x35\x2f\x89\x0a\x0c\xc8\x83\x7b\x9f\x1c\xe8\x2f\xc2\xf9\xce\xe1\xe7\x23\xae\x02\xd3\x29\x41\xbc\xe8\x8e\xfb\xc5\xdd\xdd\xe2\x3d\x4a\x9e\x61\xf7\xd3\x5b\xc3\x4c\xa7\xcf\x58\x19\x33\x99\x92\x74\x07\x39\x76\xc8\x20\x79\xf0\xea\x58\x2c\x21\x6f\x02\xbe\x4b\x97\xeb\x39\xaf\xef\xd0\xc2\x43\xc5\xc2\x19\x6e\x8d\x77\xb0\x60\xd2\xac\x2f\x5a\xab\x40\x77\xb5\x7b\x38\x62\x47\xed\xd0\x5f\x34\x6a\x40\x77\x4d\xd4\xf7\x06\x9a\xb6\x66\x66\x9c\x9c\xfe\xf4\x1e\x6d\x90\xbe\xe6\x97\x90\xc3\x66\x79\xfe\x56\xff\x03\xfd\x27\x42\x5f\x3a\xd3\x2b\xc1\x6a\xf4\x1e\x28\x7f\xec\x15\x0f\xbb\xa1\x17\x46\xc1\x93\x0d\xd4\x82\x7e\xbc\xbd\x9e\x6f\x42\x5d\x7f\x9a\xfe\x42\x4e\x27\x66\xad\x71\x0a\xc9\xc7\x34\x0a\x9b\xbb\x59\x30\xd4\x6f\xf8\x9b\x60\x11\xb7\xff\x07\x6d\x28\xe9\xa7\xc2\x4f\xed\x48\xc9\x91\x47\x88\x43\x85\x97\xcf\x53\xdf\x80\x7e\xae\xc4\xb3\x49\xcd\xb4\x09\x9e\xf8\x8c\x2b\x8c\xda\xec\x52\x8c\x0a\xed\xb0\x3d\x20\x5d\x44\x1c\x33\xeb\xd4\x10\xca\xda\xd1\x1f\x49\x9c\x98\xf1\x25\xbf\xf7\x23\x00\x00\xff\xff\xe8\xd2\x22\x95\x48\x07\x00\x00")

func tmplGolangEcho_moduleGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangEcho_moduleGogo,
		"tmpl/golang/echo_module.gogo",
	)
}

func tmplGolangEcho_moduleGogo() (*asset, error) {
	bytes, err := tmplGolangEcho_moduleGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/echo_module.gogo", size: 1864, mode: os.FileMode(420), modTime: time.Unix(1456544366, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangIncludeGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\x4b\xce\x29\x4d\x49\x55\xaa\xad\xe5\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\xe0\x52\x00\x82\xea\x6a\x5d\x85\xa2\xc4\xbc\xf4\x54\x05\x15\xa8\x9a\x14\x05\x2b\x5b\x05\x3d\xa0\x42\x88\x74\x66\x5e\x4a\x6a\x05\x92\xa4\x41\x6d\xad\x82\x12\xa6\xb0\x61\x6d\xad\x12\xdc\xc0\xd4\xbc\x14\xa0\x7e\x4d\x2e\x38\x1b\x10\x00\x00\xff\xff\xca\x18\x03\xdd\x89\x00\x00\x00")

func tmplGolangIncludeGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangIncludeGogo,
		"tmpl/golang/include.gogo",
	)
}

func tmplGolangIncludeGogo() (*asset, error) {
	bytes, err := tmplGolangIncludeGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/include.gogo", size: 137, mode: os.FileMode(420), modTime: time.Unix(1456544165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangServiceGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x52\xcb\x8e\xdb\x30\x0c\xbc\xe7\x2b\x08\x23\x87\x5d\x20\xeb\x0f\x58\xa0\x87\xa2\x3d\xb4\x40\xbb\x05\x92\xde\x1b\x21\xa6\x1d\x15\xb6\xec\x52\x74\xd0\x40\xd0\xbf\x57\x4f\xc7\x4e\xdc\xdb\xe6\x44\x85\x9c\x19\xce\x98\xc6\x54\x58\x4b\x85\x50\x68\xa4\x8b\x3c\x61\x61\xad\x31\xdb\xf4\x78\x13\x1d\xc2\xeb\x07\x28\x7d\x61\xed\xc6\x98\x17\x20\xa1\x1a\x84\x6d\x87\x7c\xee\xab\xd0\xfc\x1e\x4a\xed\xfa\x7c\x1d\x10\xf6\xf8\x67\x44\xcd\x4b\x16\x6b\x0f\xf1\x61\xcc\x38\x0c\x48\x5f\x50\x54\x99\x24\xb1\x7f\xa4\x66\xec\x50\xb1\x06\xcd\x34\x9e\x18\xcc\x06\xdc\x6f\xa6\x29\xa8\xf1\x82\x19\x36\x01\x9c\x72\x9c\x9c\x51\xbb\xd9\xc4\xeb\xfe\x6f\x50\xfd\x74\xab\x1d\x98\xa4\x6a\x6e\xbd\x58\xf9\x4e\xac\x7e\x0c\x2c\x7b\x25\x5a\xa8\x45\xab\x3d\xf2\xc8\x67\x92\x35\xbf\x16\xce\x8c\x1f\xf8\xfa\xd9\xa7\x23\x6b\x50\x3d\x2f\x21\xd6\xee\xc8\xf9\x96\x84\x95\x31\xa8\x2a\x6b\x0b\xf8\xad\x7b\x35\x41\xe3\x2e\x01\x7c\x0f\xec\x3b\xc9\xd8\x0d\x7c\xcd\xc8\xe3\x64\x3c\xbc\x37\x3e\xf8\x58\x85\x80\x8d\x29\x17\x91\x82\x54\x8c\x54\x0b\x57\x3d\x44\xb6\xfa\x99\x1e\xc2\x5a\x7c\x87\x27\x63\xfc\x3a\xad\x60\x77\x15\xb1\xf3\xcb\x6d\xac\x8b\xb5\xe0\x9f\x61\x65\x9a\x90\x47\x52\xb7\xf9\x7d\x78\xfb\x9c\x27\xf1\x47\x6b\xe1\xb8\xf2\x31\xce\x65\xc3\x41\xfa\xea\x53\x3f\x2a\xf6\x56\x5a\x54\x50\xc2\x4b\x38\xc8\xe4\x53\x56\x7f\x77\xd3\x81\xe4\xde\xb6\x15\x3a\x00\xa4\xfe\xe6\x2b\x3f\x05\x33\xaa\x34\xf5\x1e\xa7\x72\xbb\x0a\xaf\xe9\x3e\x2a\x64\x5f\xc9\x6b\x12\xfb\xbf\xd7\x1c\x5a\x9c\x73\x6c\x52\xbf\xc9\x36\x28\x96\xd6\x22\x51\x4f\x0e\x1e\xc4\x9e\xee\xd7\x4c\x1c\x10\x39\xe0\x22\xda\x11\x0b\x17\x43\x58\x2e\xaf\xb8\x83\x40\xf2\xbc\xb2\xd8\xbf\x00\x00\x00\xff\xff\x34\xbb\xb2\x76\x06\x04\x00\x00")

func tmplGolangServiceGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangServiceGogo,
		"tmpl/golang/service.gogo",
	)
}

func tmplGolangServiceGogo() (*asset, error) {
	bytes, err := tmplGolangServiceGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/service.gogo", size: 1030, mode: os.FileMode(420), modTime: time.Unix(1456544165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangServices_fileGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8f\xcd\x0a\xc2\x30\x10\x84\xef\x7d\x8a\x21\xf4\xda\xd6\xb3\xe0\xc9\x93\x37\xc1\x07\x90\xd8\x6c\x4b\x30\x6e\x4b\x9b\x0a\xb2\xe4\xdd\x8d\x69\x50\x8a\xa7\xc9\xcf\xb7\x33\xb3\x22\x86\x3a\xcb\x04\x35\xd3\xf4\xb4\x2d\xcd\xd7\xce\x3a\x52\x21\x34\x0d\xf4\xe2\x07\xf4\xc4\x34\x69\x4f\x06\xb7\x17\x7c\xbc\x15\xa3\x6e\xef\xba\x27\x88\xd4\xe7\xf5\x18\x42\x51\x88\x54\x28\x2d\xb7\x6e\x31\x74\x1c\x16\xf6\xd8\x1f\xe0\x88\x51\x9f\xd6\xc7\x19\x55\xe4\x44\x6c\x87\x98\xb7\x45\x77\x21\x88\x78\x7a\x8c\x2e\x26\x41\xe5\x3f\xf5\x37\x4b\x6c\x92\x56\x98\x34\xc7\x0a\x65\x6e\xfd\xc9\xaa\x2f\x79\x83\x44\xfc\xcc\x32\xa2\xbe\xf0\xc6\x29\xe9\x3b\x00\x00\xff\xff\xbc\xee\x53\xd1\x06\x01\x00\x00")

func tmplGolangServices_fileGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangServices_fileGogo,
		"tmpl/golang/services_file.gogo",
	)
}

func tmplGolangServices_fileGogo() (*asset, error) {
	bytes, err := tmplGolangServices_fileGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/services_file.gogo", size: 262, mode: os.FileMode(420), modTime: time.Unix(1456544165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangStructGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x31\x4e\x03\x31\x10\x45\x7b\x9f\x62\x64\x51\x92\x1c\x20\x12\x1d\x42\xd0\x40\x01\x07\x88\x85\xc7\xcb\xa0\xdd\xb1\xb1\x67\x8b\x68\x34\x77\x67\x37\x1b\x47\x11\xb8\x7a\xd6\xff\xef\xcb\xb2\x6a\xc4\x44\x8c\xe0\x9b\xd4\xf9\x53\x3c\xec\xcc\x9c\x9c\x0a\x82\xea\xfe\x35\x4c\x68\x06\x5b\x04\xea\x60\x39\xaa\x3b\xa8\x81\x07\x84\xbb\x44\x38\x46\x38\x3c\xc0\xfe\x69\xa5\xb6\x98\x5b\x63\x2e\x05\xeb\x33\x86\x78\xe9\xf4\x21\xd5\x01\xf9\x63\x19\x7f\x97\x4a\x3c\xdc\xa6\x9d\xd7\xb4\xf3\x5b\x11\xca\x1c\x46\x48\x61\x6c\xab\x7f\x94\xaf\x4a\x49\x0e\x5e\xf5\x52\x79\x79\x34\x53\xa5\x04\x9c\xe5\xaf\x66\x76\x5f\xf1\x67\xa6\x8a\x51\x15\x39\x9a\x79\xf8\x6e\x99\x6f\xf4\xed\x5d\xe7\x81\xff\x72\x9e\x48\x70\x2a\x72\xea\xf6\xf1\xfa\x01\xe7\xbb\x33\x77\xe5\xdf\x00\x00\x00\xff\xff\x62\x86\x05\xa2\x48\x01\x00\x00")

func tmplGolangStructGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangStructGogo,
		"tmpl/golang/struct.gogo",
	)
}

func tmplGolangStructGogo() (*asset, error) {
	bytes, err := tmplGolangStructGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/struct.gogo", size: 328, mode: os.FileMode(420), modTime: time.Unix(1456544165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangStructs_fileGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8f\x31\x0f\xc2\x20\x10\x85\xf7\xfe\x8a\x0b\xe9\xda\xd6\xd9\xc4\xc9\xc9\xcd\xc4\x1f\x60\xb0\x5c\x1b\x22\x5e\x1b\x38\x06\x73\xe1\xbf\x8b\x60\xa2\x8d\xd3\x3d\xe0\xe3\xbd\x77\x22\x06\x27\x4b\x08\x2a\xb0\x8f\x23\x87\xeb\x64\x1d\xaa\x94\x86\x01\x74\xe4\x05\x66\x24\xf4\x9a\xd1\xc0\xed\x09\x9c\x4f\xcd\xaa\xc7\xbb\x9e\x11\x44\xfa\x73\x95\x29\x35\x8d\x48\x07\xad\xa5\xd1\x45\x83\xc7\x25\x12\xc3\xfe\x00\x0e\x09\xfa\x53\xbd\x0c\xd0\x65\x4e\xc4\x4e\x90\xe3\xb6\xe8\x2e\x25\x11\xc6\xc7\xea\x72\x12\xa8\xcf\x9b\xfa\xfb\x8b\x64\xca\xf4\x9a\x72\x81\x76\x21\xbc\x94\xd6\xef\xac\xbe\xca\x50\x80\xaf\x57\x5d\x4b\xfd\xc0\x1b\xa7\x32\x5f\x01\x00\x00\xff\xff\xa0\x00\x02\xe7\x05\x01\x00\x00")

func tmplGolangStructs_fileGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangStructs_fileGogo,
		"tmpl/golang/structs_file.gogo",
	)
}

func tmplGolangStructs_fileGogo() (*asset, error) {
	bytes, err := tmplGolangStructs_fileGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/structs_file.gogo", size: 261, mode: os.FileMode(420), modTime: time.Unix(1456544165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplSwiftStructGoswift = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\xcd\x6e\xe3\x36\x10\x80\xef\x7e\x8a\xa9\x4e\x12\xa0\x1a\x3d\x2c\x72\x70\x9b\x35\xd2\x34\x05\xba\xdd\x3a\x41\xbc\xed\xa1\x86\x0f\x8a\x34\x49\xb4\x95\x45\x97\xa4\xd2\x06\x86\xde\xbd\xfc\x13\x45\x52\xda\x6c\xb2\x08\x50\x01\x6b\x1d\x0c\x89\x9a\x3f\x72\xbe\x19\x52\x3e\x1c\x0a\xbc\x2d\x6b\x84\x88\x71\xda\xe4\x3c\x6a\xdb\xd9\xac\xdc\xed\x09\xe5\xf0\x33\x69\xea\x22\xe3\x25\xa9\x67\xb3\xbc\xca\x18\x83\x0f\xd7\x1f\x70\xb7\xaf\x32\x8e\x0b\xb8\xf8\xf3\xf2\xe6\x23\xe6\x1c\x0e\xb3\x19\x88\x8b\x3c\x20\xa5\x65\x81\xf0\x90\x51\xc8\xaa\xea\x57\x7c\x64\x0b\x58\x23\xff\x61\xcd\x69\x59\xdf\xbd\x15\x82\x60\x2e\x8a\xbc\xa1\x35\x6c\xa2\x3a\xdb\x61\x94\x42\x94\x0b\x57\x5c\xde\x3c\x22\xbb\xa4\x2b\x22\x6f\xb3\x5d\x37\x58\x60\x5e\xee\xb2\x4a\x0d\xd6\x84\xdf\x23\xd5\xae\xe5\x00\x51\x77\x2c\xda\x2a\xe3\xad\xfa\x55\x3f\x32\x0c\x69\x5e\xc4\xa0\xfc\x2f\xfd\x57\xca\xe3\x02\x7e\xa9\x39\x9c\xc2\x77\xfe\x3b\x13\xc4\x02\x7e\x24\xa4\x12\xaf\x6f\xb3\x8a\xa1\x2f\xa2\x83\x5b\xc0\x4f\xa4\xb9\xa9\x50\x9a\x98\x07\x46\x4c\xd0\xca\xc5\xc9\x9b\xa1\x13\x6f\x26\x0b\xb1\xb4\x67\xee\x40\x10\xad\x99\xe5\x02\x36\x81\xdc\xd6\x11\xb4\x19\xb8\x6d\xea\x1c\x6e\x29\xd9\xbd\x5b\x5f\xae\xe2\x8f\x8c\xd4\x9d\x97\xb3\xfa\xd1\xd8\x4f\xe0\xdb\xb7\x7a\x7e\x7d\x5a\xec\xcd\x5d\x93\xd1\x02\x58\xb3\x47\x3a\x1f\xb1\x93\x00\x8a\x05\x81\x43\x97\x47\xb5\x3e\x66\xed\x47\xec\x54\xc8\xa1\x28\x73\xb9\xd0\xbd\x0d\xc8\xd8\x12\x36\x3a\x35\x4e\x5c\xdb\xe7\x9a\x96\xa9\x05\xe7\x3a\x55\x2e\x0c\x51\x5b\x65\x5d\x1b\xb7\x1a\x2a\xe3\x23\x1a\x9a\x3d\xad\x22\x71\x58\x2e\x4d\xaa\xe4\x65\x50\x18\x28\x75\x9c\x6a\x35\xb5\x8e\x42\xaf\x07\x45\x5e\x9a\x91\xa1\x43\x03\xb6\x56\x35\x00\x49\xa7\xf3\xde\xad\x81\x67\xa0\xda\x55\x82\x8d\x56\x90\xe5\xc5\xeb\x51\xa5\x14\x03\x60\x3c\x1c\x4c\x3c\x5e\x4d\x6d\x13\x6b\xcc\x50\xe7\x44\x31\xc0\x6f\xc4\x9c\xad\xc8\x64\x98\x35\x93\x55\xd1\x6a\x30\xac\x57\x9f\x5e\x4e\x14\x73\x8a\x52\x0b\x87\x83\xaa\xaa\x30\xcd\xd4\x08\x44\xf1\x88\x6b\x0f\x0f\x0f\x03\x39\x16\x88\x75\x4c\xb8\x62\x6a\x2c\x90\xeb\x31\xe8\xe5\xcc\x58\x20\x69\x93\xee\x48\xea\xb1\x40\xb0\x4f\x71\x2f\xb8\x5a\xaf\x9a\xdd\x0d\xd2\xb8\x22\xf5\xdd\x7b\x22\x27\x6b\xc4\x92\xd0\x8d\x9f\x4b\xed\xc6\x6b\x2c\xf3\x6e\x69\x03\x4d\x9b\x36\xc7\xaf\x19\x1b\xd1\x09\x33\x2a\x6d\x98\x8c\xb6\xfd\x6e\x71\xe6\xf7\xb8\xd7\xdb\x32\x5e\xd2\xed\xbf\xbe\xb6\xf8\xf9\x5e\xf8\xfa\x05\xf9\x5a\x95\x19\x94\xe3\x73\x38\xbb\xf8\x37\x13\x87\x12\x91\xfa\xd5\xda\x0b\x4b\xfd\x68\x11\x35\x85\xdf\x19\xd2\x2b\xc2\x78\x77\x86\x89\xb9\x3d\xcc\xf4\x07\x9b\xd4\xdf\xfd\x53\xe7\xa0\x90\x82\x2c\xbf\x73\xfb\x7c\xf2\x26\x0d\xce\x01\x29\x74\x26\xf5\x56\xdd\x19\xdd\xa6\x1e\x9c\xa9\x60\x28\xcf\x91\x09\xa1\xb8\x17\x52\x4b\xfb\x07\x29\x8b\x54\xa4\xb7\xac\x1a\x2a\xe4\xe3\x78\xb5\xbe\xa0\x94\x50\xfb\x32\x59\xca\x25\x2a\x2b\x48\xdc\x59\x76\x0b\xbf\xcf\x68\xb6\x63\x9f\x5d\x7a\x2d\xb6\x89\xba\x68\x65\xd1\x9f\xda\xd8\x87\xd5\xde\xc9\xbb\xed\x6e\xd8\xe8\x3a\x29\xb7\x79\x86\x6d\xb3\x93\xb1\x2b\xa9\x52\x3e\xd2\xdf\xac\xc0\xa7\xa3\x66\x4a\xd7\x3e\x3d\xd1\xa4\x64\xa9\x65\xfb\x52\x48\x47\x06\x96\x79\x08\x43\x34\xc2\x25\xde\x34\x77\x57\x62\x19\x79\x1c\x09\x6d\x79\xda\x14\x18\xfe\xbd\x00\x71\xa7\x63\x19\x71\x75\x46\x31\x5b\x23\x7d\x28\x73\x9c\xe7\x0d\xa5\x58\x8b\x96\xfb\x6e\x7d\x7d\xb9\xba\xba\x3e\x3f\xaf\x4a\xf1\x3c\x2f\xeb\x07\xf2\x17\xfe\x86\xfc\x9e\x14\xb1\xb0\x9c\xc2\x3f\x25\xbf\xbf\x92\x26\x91\x23\x15\x58\x68\xf3\x0e\x27\x07\x88\x89\x68\x3b\xea\x34\x9e\x8a\x28\xd8\x9e\xd4\x0c\xbb\xbe\x63\xd8\x80\xb2\xb6\x71\x78\x41\x05\xb3\x51\x2e\x23\x65\x45\x4d\x26\x30\xf7\xbd\xa7\x67\x42\x70\x40\xf5\x5a\x67\xa0\xfb\x4d\xf2\xe9\x08\x5a\x07\x6c\x7f\x3e\xe8\x21\x1e\x4e\x43\x5e\x46\x71\x19\x6b\x51\xfb\xbe\x4d\xc2\xb6\xe5\x14\xfc\xb5\xea\x19\xba\x12\xe2\xbe\xe4\xf4\xc0\x17\x96\xdb\x80\xab\x67\x96\xdc\x08\x82\x6e\x78\x4f\xe3\x67\x13\xf6\xb5\xc2\xe7\x0b\x39\x3b\x9a\x3c\xfd\x46\xd1\xf4\xa0\x13\xdb\x83\x43\x9c\x78\x9a\x0a\x6e\x22\x94\x23\x6b\x2f\x65\xad\xfb\x2c\x9c\x1e\x67\xf2\x8c\xe2\x80\x26\x1f\xa7\x42\x9a\x8c\xe5\x88\xda\x4b\x51\xf3\xfe\x4a\x98\x1e\x6f\xfa\x9c\xeb\x10\xa7\x07\xa6\xc2\x9c\x8e\xe6\x48\xdd\x4b\xa9\xf3\xfe\x85\x9a\x1e\x75\xea\x63\xcb\xdf\x4f\x4f\xde\x4c\x85\x39\x15\xcc\x11\xb9\x2f\xd8\x53\xcd\x9f\x97\xd3\xe3\xed\x7d\xc9\xdc\xe3\x9b\xfb\x25\x3f\x15\xea\x64\x88\x47\xe8\x9e\x84\xce\x4d\xdb\x53\xdf\xac\x12\xc2\xcd\xf6\x7f\xa5\xb0\x9d\x1d\x0e\x58\x17\x6d\x3b\xfb\x2f\x00\x00\xff\xff\x7f\xf0\x3a\x54\x91\x1b\x00\x00")

func tmplSwiftStructGoswiftBytes() ([]byte, error) {
	return bindataRead(
		_tmplSwiftStructGoswift,
		"tmpl/swift/struct.goswift",
	)
}

func tmplSwiftStructGoswift() (*asset, error) {
	bytes, err := tmplSwiftStructGoswiftBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/swift/struct.goswift", size: 7057, mode: os.FileMode(420), modTime: time.Unix(1456385346, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tmpl/golang/echo_module.gogo": tmplGolangEcho_moduleGogo,
	"tmpl/golang/include.gogo": tmplGolangIncludeGogo,
	"tmpl/golang/service.gogo": tmplGolangServiceGogo,
	"tmpl/golang/services_file.gogo": tmplGolangServices_fileGogo,
	"tmpl/golang/struct.gogo": tmplGolangStructGogo,
	"tmpl/golang/structs_file.gogo": tmplGolangStructs_fileGogo,
	"tmpl/swift/struct.goswift": tmplSwiftStructGoswift,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"tmpl": &bintree{nil, map[string]*bintree{
		"golang": &bintree{nil, map[string]*bintree{
			"echo_module.gogo": &bintree{tmplGolangEcho_moduleGogo, map[string]*bintree{}},
			"include.gogo": &bintree{tmplGolangIncludeGogo, map[string]*bintree{}},
			"service.gogo": &bintree{tmplGolangServiceGogo, map[string]*bintree{}},
			"services_file.gogo": &bintree{tmplGolangServices_fileGogo, map[string]*bintree{}},
			"struct.gogo": &bintree{tmplGolangStructGogo, map[string]*bintree{}},
			"structs_file.gogo": &bintree{tmplGolangStructs_fileGogo, map[string]*bintree{}},
		}},
		"swift": &bintree{nil, map[string]*bintree{
			"struct.goswift": &bintree{tmplSwiftStructGoswift, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

