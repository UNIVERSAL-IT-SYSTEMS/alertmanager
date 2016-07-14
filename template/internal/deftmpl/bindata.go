// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3b\x7b\x73\xda\xc6\xf6\xff\xeb\x53\x9c\xaa\xf3\x9b\xc6\x33\x08\x6c\xe7\x31\xf5\x03\xff\x86\x60\x39\x66\x2e\x06\x0f\xe0\xa4\x99\x4e\x27\x23\xa4\x05\x36\x91\xb4\xaa\x76\x65\xec\xf6\xf6\xbb\xdf\x73\x56\x02\x24\x10\x36\xf6\xa4\x36\xf7\xd6\x49\x1f\xec\xd1\x9e\xf7\x63\xcf\x6a\x57\x7f\xfe\x09\x1e\x1b\xf1\x90\x81\xf9\xe5\x8b\xe3\xb3\x58\x05\x4e\xe8\x8c\x59\x6c\xc2\x5f\x7f\x35\x68\x7c\x91\x8e\x71\x22\x0b\x3d\x04\x1a\x6b\x51\xae\x7a\x6d\xc2\xc2\xe7\x55\xfb\x46\xb1\x38\x74\x7c\x04\x21\xa4\xf6\x63\x4d\xcf\x93\xff\x1f\x33\x97\xf1\x6b\x16\xd7\x69\x52\x2f\x1b\xa4\x38\x19\xf5\x22\x79\x99\x0c\xbf\x32\x57\x11\xd9\x5f\x09\xa5\xaf\x1c\x95\x48\xf8\x37\x28\x71\x15\x45\x33\x54\x3e\x02\xf6\xfb\xfc\xa1\x39\xe2\x31\x0f\xc7\x84\x73\x48\x38\x5a\x0b\x59\x3d\xd3\x50\x44\xf5\x59\x98\xe7\xf8\x1b\xd0\xa4\x0f\xb1\x48\xa2\xb6\x33\x64\xbe\xac\xf6\x45\xac\x98\x77\xe9\xf0\x58\x56\x3f\x3a\x7e\xc2\x88\xe1\x57\xc1\x43\x30\x81\xa8\x42\xca\x72\xac\xe0\x15\xd1\xaa\x36\x45\x10\x88\x30\x45\xde\xc9\x60\x39\x7a\x3b\x88\xf2\x0a\x51\xa6\x5c\x4d\x8a\x93\xd1\x02\x81\xb8\x66\x45\xee\x1d\x27\x40\x86\xa9\x19\xcb\xb8\xcf\x05\xdf\x99\xff\x5a\xe3\x1b\x8f\x49\x37\xe6\x91\xe2\x22\x34\xef\xb0\xb1\x62\x37\x2a\xf5\xe3\x17\x9f\x4b\x95\x4d\x8d\x9d\x70\x8c\x92\xe1\x20\x95\xeb\xd0\x58\x00\x57\xed\x44\x56\xb1\xb4\x21\x49\x7c\x1a\xd5\x61\xae\x40\x26\x58\xca\xbc\x11\x86\x02\xfd\x84\x32\x15\x48\xe6\xc0\x8f\xa3\xdb\x17\x49\xec\xb2\xc3\xd4\x99\x2c\x64\xb1\xa3\x44\x9c\x86\x9f\x51\x62\xa8\x82\x0d\xa4\xef\xb8\xdf\xaa\x38\x72\x12\x5f\x55\x15\x57\x3e\xcb\xac\xa0\x58\x10\xf9\x8e\x2a\xc6\x62\x75\x9d\xc9\x8b\x74\x12\x49\x29\x10\x94\x91\x2a\x26\xda\x86\xf4\x46\x8e\xef\x0f\x11\xb0\x42\xaf\x54\x7c\x22\x8a\x81\x73\xdf\x44\x9f\x87\xdf\x36\x96\x20\x8a\x19\x05\x8b\xb9\xd9\xec\x1c\xfd\x3b\x0d\xa0\xcb\xc6\x86\x12\x70\x57\x84\x98\x33\x5f\xf9\xa6\x32\xac\x88\x5b\x70\xfc\x84\x47\xee\xc4\x51\x0b\x13\xc7\x22\x78\xbc\xbb\x96\xa9\x61\x1e\x4b\x44\xd9\x3c\x94\x0a\xb2\x45\xc4\xcd\x4b\xd4\xed\x9c\xde\x6a\x3e\x3f\x2c\x3c\x57\x29\xba\x3e\x67\xa1\x7a\xbc\xc6\xeb\x28\x2e\x56\x82\xc7\x39\x7d\x95\x2e\x0f\xa5\x72\x42\x97\xc9\x12\xba\x2b\x05\xec\x0e\xab\x8a\x48\x8e\x59\xc8\xd9\xe3\x9d\x74\x17\xb1\x55\x0f\x65\xf5\x7e\x4d\x79\x2b\x2d\xf0\xc6\xd2\xf2\x52\x58\xbf\x76\x60\x17\x2c\x9c\x93\x02\x21\x05\xea\x42\x7a\xb7\x45\x8a\x8b\xa0\x66\x62\xe5\x34\x2a\xe1\xd7\x63\x52\xf8\xd7\xcc\x5b\xe2\x38\x03\x6f\xce\x73\x86\xb1\xc2\xd5\xda\xc4\xa4\x52\xd7\xf5\x87\x47\x53\xc1\xeb\x2c\x70\xb8\xbf\x20\xb9\x68\x2a\x1e\xec\xf2\x22\xa5\x89\x0a\x7c\xed\xb2\xe3\x1f\x4e\xbb\xcd\xc1\xe7\x4b\x1b\x08\x04\x97\x57\xef\xdb\xad\x26\x98\x56\xad\xf6\xe9\x75\xb3\x56\x3b\x1d\x9c\xc2\x2f\xe7\x83\x8b\x36\xec\x55\x77\x61\x80\x6b\x9e\xe4\x14\x0c\x8e\x5f\xab\xd9\x1d\x74\xfb\x44\xa9\xe8\xb0\x56\x9b\x4e\xa7\xd5\xe9\xeb\xaa\x88\xc7\xb5\x41\xaf\x76\x43\xb4\xf6\x08\x39\xfb\x69\xa9\x1c\x66\xd5\x53\x9e\x79\x82\x9c\x2d\xcb\xe8\xab\x5b\x9f\x81\x83\xd2\x6a\x26\x1e\x8b\x39\x19\x9c\xca\x19\x10\x69\x89\xb4\xc7\xd8\x7e\x24\xc3\xaa\x2b\x82\x1a\xe9\x30\x4e\xc2\x9a\x26\xe7\xb8\x29\x3d\x4b\xab\x66\xcd\xcc\x21\xd1\x82\x83\x09\x83\x8b\xd6\x00\xda\xdc\x65\xa1\x64\xf0\x0a\x07\x3b\x86\xd1\x14\xd1\x6d\xcc\xc7\x13\x0c\x18\x77\x07\xf6\x77\xf7\xde\xc0\x45\x4a\xd1\x30\x2e\x59\x1c\x70\x29\x91\x22\x70\x09\x13\x16\xb3\xe1\x2d\x8c\x91\x0f\x86\x7c\x05\x05\x62\x0c\xc4\x08\xb0\x4c\xc6\x63\x56\xc1\x2e\x0e\x85\xbe\x05\x6c\xe4\x24\x22\x88\xa1\x72\x78\x48\xf1\xe9\x80\x8b\x3c\x0c\x9c\xa9\x26\x48\x46\x8a\x91\x9a\x3a\x71\xaa\xa1\x23\xa5\x70\x39\x4a\xe8\x81\x27\xdc\x24\xc0\x72\xa3\x13\x0b\x46\xdc\xc7\x54\x7a\xa5\x50\x68\xb3\x9f\x61\x98\x3b\x9a\x89\xc7\x1c\xdf\xc0\x04\xa3\x67\xb3\x47\xba\x1f\x13\x89\x82\x98\x49\x15\x73\x6d\x85\x0a\xf0\xd0\xf5\x13\x8f\x64\x98\x3d\xf6\x79\xc0\x33\x0e\x84\xae\x15\x97\x06\x12\xc5\xf5\xbd\xa2\xe5\xac\x40\x20\x3c\x3e\xa2\xff\x33\xad\x56\x94\x0c\x31\x05\x26\x15\xf0\x38\x91\x1e\x26\x0a\x81\x92\x80\xda\x8e\x15\xd2\xa3\x26\x62\x90\xcc\xf7\x0d\xa4\xc0\x51\x6e\xad\xeb\x42\x3a\x3d\x87\x44\x8f\xc8\xa0\x2a\x33\x91\x24\xc8\x74\x82\x5e\x2d\x68\xc2\xa5\x31\x4a\xe2\x10\x59\x32\x8d\xe3\x09\x34\x99\xe6\x48\xd1\x4c\x10\x9a\x3e\x12\xbe\x2f\xa6\xa4\x1a\x2e\xa2\x1e\xcf\x5a\x30\xed\x64\x67\x48\x6d\xa8\x3b\xf7\x2b\x16\x2b\x14\x35\x15\x81\x1c\x10\x2d\xbc\x9a\x3d\x92\x13\xec\x46\x60\xc8\x32\x83\x21\x5f\x34\xaf\x93\x53\x27\x26\xf6\x54\xb3\x15\x77\x7c\x88\xb0\xe6\x11\xbf\x65\x35\xab\xc8\xff\xdc\x86\x7e\xf7\x6c\xf0\xa9\xd1\xb3\xa1\xd5\x87\xcb\x5e\xf7\x63\xeb\xd4\x3e\x05\xb3\xd1\xc7\xb1\x59\x81\x4f\xad\xc1\x79\xf7\x6a\x00\x38\xa3\xd7\xe8\x0c\x3e\x43\xf7\x0c\x1a\x9d\xcf\xf0\xaf\x56\xe7\xb4\x02\xf6\x2f\x97\x3d\xbb\xdf\x87\x6e\xcf\x68\x5d\x5c\xb6\x5b\x36\xc2\x5a\x9d\x66\xfb\xea\xb4\xd5\xf9\x00\xef\x11\xaf\xd3\xc5\x10\x6e\x61\xec\x22\xd1\x41\x17\x88\x61\x46\xaa\x65\xf7\x89\xd8\x85\xdd\x6b\x9e\xe3\xb0\xf1\xbe\xd5\x6e\x0d\x3e\x57\x8c\xb3\xd6\xa0\x43\x34\xcf\xba\x3d\x68\xc0\x65\xa3\x37\x68\x35\xaf\xda\x8d\x1e\x26\x76\xef\xb2\xdb\xb7\x91\xfd\x29\x92\xed\xb4\x3a\x67\x3d\xe4\x62\x5f\xd8\x9d\x41\x15\xb9\x22\x0c\xec\x8f\x38\x80\xfe\x79\xa3\xdd\x26\x56\x46\xe3\x0a\xa5\xef\x91\x7c\xd0\xec\x5e\x7e\xee\xb5\x3e\x9c\x0f\xe0\xbc\xdb\x3e\xb5\x11\xf8\xde\x46\xc9\x1a\xef\xdb\x76\xca\x0a\x95\x6a\xb6\x1b\xad\x8b\x0a\x9c\x36\x2e\x1a\x1f\x6c\x8d\xd5\x45\x2a\x3d\x83\xa6\xa5\xd2\xc1\xa7\x73\x9b\x40\xc4\xaf\x81\xff\x34\x07\xad\x6e\x87\xd4\x68\x76\x3b\x83\x1e\x0e\x2b\xa8\x65\x6f\x30\x47\xfd\xd4\xea\xdb\x15\x68\xf4\x5a\x7d\x32\xc8\x59\xaf\x7b\x51\x31\xc8\x9c\x88\xd1\xd5\x44\x10\xaf\x63\xa7\x54\xc8\xd4\x50\xf0\x08\x4e\xa1\xf1\x55\xdf\x9e\x13\x84\x53\xbb\xd1\x46\x5a\x7d\x42\x26\x15\x67\x93\xab\x86\x65\x61\x45\xd2\x25\xf0\x26\xf0\x43\x59\x2f\x29\x6c\x7b\x07\x07\x07\x69\x3d\x33\x37\x9b\x24\xa9\xb8\xd5\xcd\x91\x08\x95\x35\x72\x02\xee\xdf\x1e\xc2\x4f\xe7\x0c\x97\x14\x8c\x44\x07\x3a\x2c\x61\x3f\x55\x60\x0e\x40\x55\x63\x0c\x39\x0c\x7f\x2c\x6e\x16\xf6\xe0\x7c\x74\x04\x43\x71\x63\x49\xfe\x07\xad\x95\xf8\x3b\xc6\x02\x69\x21\xe8\x08\x34\x51\x7c\x80\x1b\x87\xbd\x37\x11\x02\x02\x2c\x4c\x3c\x3c\x84\xdd\x23\xaa\xad\x13\xe6\x78\xcf\xc9\x3f\x60\xca\x01\xda\x43\xd4\xcd\x6b\xce\xa6\x94\x45\x26\x65\xaf\xc2\xa2\x57\x37\xa7\xdc\x53\x93\xba\xc7\xae\x31\x21\x2d\x3d\x78\x3e\x63\x41\x6d\x26\x2e\x39\xd3\x62\xbf\x27\xfc\xba\x6e\x36\x53\x51\xad\xc1\x6d\xc4\x72\x82\x53\xab\x50\x23\xe7\x1e\xe9\x95\x40\x32\x55\xbf\x1a\x9c\x59\x3f\x3f\xb3\xf8\x7a\xc3\xf2\x7c\xee\xbe\xab\x17\x39\xae\x69\xe1\x4e\x0c\xe3\xb8\x46\x41\x49\x3f\x86\xc2\xbb\x05\x8e\x28\x12\x6b\x2e\x4a\x6c\xea\x81\xba\xa5\xdf\x59\x46\x49\x77\x82\xab\xba\xce\x28\x9b\x56\xf7\x8b\x59\xb3\xfb\xa4\x4a\x5a\x53\x36\xfc\xc6\x91\x91\x7e\x10\x08\x81\x6b\x0a\x21\xa5\x6b\x03\x77\x24\xf3\x16\x93\x28\x36\x34\xb6\xe5\x78\x5f\x13\xa9\x0e\x71\xc5\x09\xd9\x11\xb6\x12\xb4\x32\x21\xc9\xdd\xdd\xff\x3b\xc2\x45\x39\x64\xd6\x1c\x54\x7d\xc7\x82\x23\xd0\x19\x90\x4e\x80\x1f\x78\x40\xc9\x82\x1c\x50\x4e\xdc\x10\x8e\x63\x91\x84\x9e\xe5\x0a\x5f\xc4\x87\xf0\xe3\xe8\x1d\xfd\xcd\x9b\x1f\x22\xc7\xf3\xb4\x54\x14\x0d\xc3\xb1\x9e\x59\x37\xb3\x99\x26\xd9\x5b\x39\xc3\xa7\x0e\x8f\x9c\x4a\x1b\xea\x51\x2a\x3b\xc0\xb1\x8a\x9f\xb1\x8e\x01\x90\x04\x4f\x5c\x49\xaf\x71\x7b\x80\x44\x7c\x0b\x43\x6c\x8c\x92\x28\x11\x15\x0d\x75\xad\x1f\x60\x35\x12\x91\x79\x82\x09\xe6\x2d\x04\x4d\x2b\xab\xf9\x6e\x77\xf7\x89\x53\xa5\x54\x68\xec\x22\xb1\x2a\x20\xdb\xa1\x2f\xdc\x6f\x85\xd8\x0e\x9c\x1b\x2b\x0b\x12\x14\x36\xba\x29\x3c\x74\x7d\xe6\xc4\xc4\x50\x4d\x0a\xf0\x75\x89\x32\x37\x0e\x38\x89\x12\x4b\x29\x51\xb0\x96\x36\x14\x9a\xca\xe3\xd7\x4f\x1d\x56\x45\x7d\x97\x8d\x73\xb7\x12\x33\xb9\xc9\xc9\x3a\x99\x33\x3f\x93\x25\x70\x79\xc2\x6e\x3c\x9b\x5d\x37\x77\xd3\xb1\x8c\x1c\x77\x36\x7e\x52\x45\xb3\x87\xb1\xe3\xf1\x44\x1e\xc2\x6b\x0d\x2b\x29\x00\xa3\x51\xa1\x8a\xa5\x68\x48\x04\x43\x01\x77\xdd\xdc\x83\x1f\xd9\x01\xfd\x2d\x16\x86\xd1\x28\x67\x8b\x6d\xa8\x0e\x0b\x49\x9e\xae\x4a\xbc\x5b\x9b\x70\x05\xeb\x6a\x94\x69\xb6\xd4\xbc\xdd\x45\x23\xeb\x25\x2a\x9b\x8f\x1b\x3a\xc5\xe2\x32\x7f\xe9\x7f\x77\xb5\x53\x56\xfd\x66\xbf\x7b\xbb\xbf\xdf\x2c\x5f\x80\xf6\x29\xae\x4d\xc8\xf2\x2d\x65\x90\xf7\x5e\x8a\x5b\x9e\x91\xb3\x3f\x8b\x73\x8f\xf9\x81\x07\xe8\x17\x26\xa5\xef\x7a\x76\x60\x0f\x27\xc8\xf9\x0b\x0f\xd4\x39\x86\xc5\xbb\xf9\x35\x67\x23\xf4\xde\x03\x60\x95\x6f\xf6\xa6\xbe\x9e\x7f\x4f\x0f\xab\xf2\x65\xef\x56\x0a\xde\x9f\x17\xe1\xf9\x38\x7e\x89\xd3\x4d\x56\xb3\x45\xf4\xec\xa5\xd1\x73\x57\x70\x6c\x7d\xf1\x5b\x6b\xf6\xed\x0a\x82\x6d\x0f\x05\x2c\x3e\xb3\x62\x72\x57\x38\x64\x6a\xe0\xce\x2d\x66\xa3\xba\xb9\xc9\x6b\xd6\x27\x8e\x87\x59\xd5\x3c\x3b\x3b\xcb\xaa\xaf\xc7\x5c\x11\xeb\x97\x72\xb3\xfd\x41\x61\x47\xb0\x4f\xfb\x81\x42\xe1\x1e\x0a\xdf\x2b\xaf\xdc\x6e\x12\x4b\xa2\x1e\x09\x9e\x02\xe6\x1d\x05\x0f\x35\xd1\xac\xb1\x58\xaa\xf0\x6f\x49\x30\x4d\x4f\xbf\x45\xc5\x8a\x19\x20\x4d\x27\xe2\x0a\xe9\xff\xc1\x4a\xab\xfe\xeb\x37\x3f\x33\xcf\x29\x59\xb0\x57\x66\x64\x60\x6d\xe5\xc3\x74\x25\x9f\x03\xe7\xed\x1b\xae\x2f\xa9\x7b\x4f\x3e\x72\x36\xa5\x17\x70\xf7\x9e\xe2\x1c\xd7\x9c\xd2\x18\x5e\x2a\xbc\xe5\xe5\x37\xfd\x73\xdf\xe9\x44\xc9\xaa\xf0\x92\xb2\x7f\x4f\xca\x4a\x15\x8b\x70\xfc\x7c\xa6\xfd\x75\xfd\xf5\x8a\xdf\xb2\xa3\xa9\xe3\x5a\x2a\xe4\x77\x88\xba\x92\x86\x21\x7b\x32\xbb\x43\xb0\x7c\xc6\xf5\x12\x87\xff\x8c\x38\x4c\x7b\xd3\x79\xa8\x1d\x0f\x9f\xcf\xcd\xf4\x22\xb1\xcc\x46\xf7\x5c\x9e\x59\x7f\xc3\xe5\x99\x95\x59\x9f\x77\x99\x56\x85\xb5\x60\x71\xca\x9d\xae\x04\xcf\x1e\x19\x39\x89\xb6\x25\x3c\xee\xb5\xe8\xbd\x37\xa2\xfe\x4b\x83\x25\xdf\x61\x2e\x5f\xd1\x7a\xa6\x86\x72\xd6\x6e\xad\xf4\x94\xd8\xb5\xb1\x98\xba\xbf\x62\x38\xa5\x97\xcc\xa8\x89\xda\xbe\x1a\xf3\xb8\xd5\x74\xc3\xf6\x2e\x7f\x19\xa4\xd4\xbd\x2f\x5d\xe1\xd6\xac\xc6\x5b\x17\x99\x28\xd3\x64\x0b\x65\xda\x3a\x3b\x3d\x24\x83\xef\xea\x88\x5f\x12\xeb\x7f\xb3\xcd\xcd\x6f\xb7\xe6\x97\xea\x16\x1b\xae\x19\xe8\x19\xb6\x5c\xf9\x2b\x7e\x2f\xd1\xf8\xcf\x88\xc6\x97\x4d\xd7\xcb\xa6\xeb\x65\xd3\xb5\xed\xc1\xf2\xb2\xe9\xda\x9a\x96\x6d\x9d\xa3\x70\x36\x9d\xc7\x9d\x3c\xe0\x28\x74\x8e\xb2\x80\x3c\xf9\x55\x8c\xc2\xdd\xa4\xdc\x55\x93\x85\xa3\x0f\x0e\x0e\xee\x3a\xe1\x2e\x9e\xec\xae\x1e\x49\x6e\x47\xd3\xb0\x4d\xed\xcb\x53\xb6\x2e\xfb\x6b\x5b\x97\xd2\x43\xb4\xfb\x5c\x9e\xeb\x6d\x96\x2e\x36\x14\xaf\x61\xe5\xcb\x55\xf1\x23\xd2\xa7\x0b\x88\xfd\x7c\xb5\xd2\x1a\x6d\x5c\xaa\x50\x27\x18\xde\x6e\x76\x0e\xb7\x5a\x3b\x56\xee\x3b\x2c\x57\x86\xe3\x1a\xa6\xf9\x49\xfa\x5f\xa3\x58\x26\xb6\xad\xad\x5d\x73\xbf\x2e\x55\x71\x51\xbf\x8e\x6b\x74\x8d\x95\x20\x74\x1f\xf8\xc4\x30\xca\xbf\x52\x8d\x12\x39\x11\xc8\xf1\x3b\x7c\xa4\xb9\x42\xaa\xf8\x05\xd8\xdf\xf1\xc1\xd6\xf7\xf9\x5e\x6b\xf3\xcf\xb5\xbe\xdf\xd7\x5a\x39\x9e\x1b\x58\x32\x89\xfd\x87\x7f\xa8\xf5\x9f\x00\x00\x00\xff\xff\x65\x23\x7f\x36\x82\x3e\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 16002, mode: os.FileMode(420), modTime: time.Unix(1468335459, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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

