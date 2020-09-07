// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// auditor.kubeshield.to_dashboards.v1.yaml
// auditor.kubeshield.to_dashboards.yaml
// auditor.kubeshield.to_dashboardtemplates.v1.yaml
// auditor.kubeshield.to_dashboardtemplates.yaml
// auditor.kubeshield.to_datasources.v1.yaml
// auditor.kubeshield.to_datasources.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _auditorKubeshieldTo_dashboardsV1Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4b\x8f\xdb\x36\x10\xbe\xfb\x57\x0c\xdc\xc3\x5e\x62\x1b\x41\x8b\xa2\xf0\x6d\xb1\x6d\x83\x45\xdb\x34\xd8\x4d\x73\xa7\xc5\xb1\xc4\x2e\x45\xb2\x9c\x91\x37\x6e\xd1\xff\x5e\x0c\x29\x59\xb2\x2d\xd9\x46\x80\x00\x9d\x1b\x5f\xf3\xf8\xe6\x49\x15\xcc\x27\x8c\x64\xbc\x5b\x83\x0a\x06\x3f\x33\x3a\x59\xd1\xf2\xe5\x07\x5a\x1a\xbf\xda\xbd\x9d\xbd\x18\xa7\xd7\xf0\xd0\x10\xfb\xfa\x09\xc9\x37\xb1\xc0\x1f\x71\x6b\x9c\x61\xe3\xdd\xac\x46\x56\x5a\xb1\x5a\xcf\x00\x8a\x88\x4a\x36\x3f\x9a\x1a\x89\x55\x1d\xd6\xe0\x1a\x6b\x67\x00\x56\x6d\xd0\x92\xdc\x01\x50\x21\x2c\x5f\x9a\x0d\x46\x87\x8c\x49\x8a\x53\x35\xae\x41\xf6\xa8\x32\x68\xf5\x0c\x20\x6f\x69\x45\xd5\xc6\xab\xa8\x69\xa9\x1a\x6d\xd8\xc7\x65\x7f\x6b\xc9\x7e\x46\x01\x0b\xe1\x5a\x46\xdf\x84\x35\x8c\x5f\xca\xdc\x5a\xe9\x85\x62\x2c\x7d\x34\xdd\x7a\xd1\x3d\x6a\x57\x47\x5a\xa4\xe3\x10\xa8\xf0\x1a\xd3\x32\x83\x71\x2f\x2f\x9e\xb0\x34\xc4\x31\x19\x9c\xce\xac\x21\xfe\x65\xfc\xfc\x57\x43\x9c\xee\x04\xdb\x44\x65\x87\x86\xa5\x6d\x32\xae\x6c\xac\x8a\x83\x83\x19\x00\x15\x3e\xe0\x1a\xde\x8b\xf2\x41\x15\x28\x7b\xbb\xec\xae\xa4\xfc\xa2\x45\x69\xf7\x56\xd9\x50\xa9\xb7\x99\x55\x51\x61\xad\xb2\x6d\x00\x3e\xa0\xbb\xff\xf0\xf8\xe9\xdb\xe7\xa3\x6d\x80\x10\x7d\xc0\xc8\x07\x18\x32\x0d\xe2\x61\xb0\x0b\xa0\x91\x8a\x68\x02\xa7\x40\xb9\x13\x86\xf9\x16\x68\x09\x04\x24\xe0\x0a\x3b\xd5\x50\xb7\x3a\x80\xdf\x02\x57\x86\x20\x62\x88\x48\xe8\xb8\xc7\xaa\x27\xbf\x05\xe5\xc0\x6f\xfe\xc4\x82\x97\xf0\x8c\x51\xd8\x00\x55\xbe\xb1\x1a\x0a\xef\x76\x18\x19\x22\x16\xbe\x74\xe6\xef\x03\x6f\x02\xf6\x49\xa8\x55\x8c\x2d\xb6\x3d\x19\xc7\x18\x9d\xb2\xb0\x53\xb6\xc1\x37\xa0\x9c\x86\x5a\xed\x21\xa2\x48\x81\xc6\x0d\xf8\xa5\x2b\xb4\x84\xdf\x7c\x44\x30\x6e\xeb\xd7\x50\x31\x07\x5a\xaf\x56\xa5\xe1\x2e\x0f\x0a\x5f\xd7\x8d\x33\xbc\x5f\x15\xde\x71\x34\x9b\x86\x7d\xa4\x95\xc6\x1d\xda\x15\x99\x72\xa1\x62\x51\x19\xc6\x82\x9b\x88\x2b\x15\xcc\x22\xa9\xee\x38\x25\x53\xad\xbf\x89\x6d\xe6\xd0\xdd\x91\xae\xbc\x17\x0f\x13\x47\xe3\xca\xc1\x41\x0a\xb3\x0b\x1e\x90\x30\x03\x43\xa0\xda\xa7\xd9\x8a\x1e\x68\xd9\x12\x74\x9e\x7e\x7a\xfe\x08\x9d\xe8\xe4\x8c\x53\xf4\x13\xee\xfd\x43\xea\x5d\x20\x80\x19\xb7\xc5\x98\x9d\xb8\x8d\xbe\x4e\x3c\xd1\xe9\xe0\x8d\xe3\xb4\x28\xac\x41\x77\x0a\x3f\x35\x9b\xda\xb0\xf8\xfd\xaf\x06\x89\xc5\x57\x4b\x78\x50\xce\x79\x86\x0d\x42\x13\xb4\x62\xd4\x4b\x78\x74\xf0\xa0\x6a\xb4\x0f\x8a\xf0\xab\x3b\x40\x90\xa6\x85\x00\x7b\x9b\x0b\x86\x75\xed\xf4\x72\x46\x6d\x70\xd0\x95\xa1\x9e\xc6\xf3\x4b\x28\x36\xf6\x7c\x13\xc0\x30\xd6\x23\xdb\x97\x38\x65\x72\x5d\x79\x98\x38\x3f\x89\x9d\x43\x35\x91\xa4\x55\xdc\x66\x68\x63\x11\x6a\xc5\x45\x25\xa9\xf0\x51\xbc\x5c\x07\xde\x9f\xa3\x72\x4c\xf3\x79\xf7\x08\x9c\x77\x8b\x83\x26\xfa\x10\x73\xb4\x84\x7b\xd7\x32\x93\xfa\x08\xa6\x0e\xd6\xe0\x69\x18\xf6\x84\x3b\x8c\xfb\xde\xa6\xe5\xc4\xc5\x49\xb4\x32\x4d\xf8\xf4\xfc\x8a\x8a\x51\xed\x47\x6f\x08\xe8\xa9\x5e\xdd\x06\xeb\xef\xdd\x75\x49\x4b\xc9\x8c\xc3\x7b\xd8\xa0\xa4\x63\xc0\xb8\xf5\xb1\x46\xfd\x65\x16\x5d\x16\x26\xc6\x48\x25\x3d\xa4\xfa\x89\xf4\x49\xb6\x00\x45\x85\xc5\x0b\x6a\xd8\xfa\x08\x4a\xd7\x86\x52\x5d\x4f\x49\xe6\xed\xd7\x04\xf8\x10\x22\x37\xe1\xdb\x4d\x1e\xd3\x51\x7b\x14\x67\x93\x8a\xb7\xf1\x07\xca\xda\x54\x65\x09\x8c\x4b\x8b\xfb\x0f\x8f\x79\x8a\xa0\x2f\x8c\xb9\x23\x75\xdf\x09\xa7\x5e\xe7\xbe\xbe\xf6\x2e\xea\xa5\x9f\xb6\xc4\x21\x1d\xd4\x9a\xd2\xea\x7a\x81\xc8\x94\x27\xa4\x0b\x17\xc6\x2c\xe8\xe2\x4b\xf2\x31\xb7\x73\xec\x35\xca\x9e\x90\x48\x51\xc6\x4d\x67\x74\x26\x79\x39\x28\x0a\xa7\x25\x66\x88\x50\xea\x2b\x3e\xe2\x15\x8e\x37\x20\x03\xb7\x45\x2a\x0c\xa2\xf1\x7d\x3f\x29\xde\x04\xd3\xd3\xf0\x5d\x6e\xc9\xa9\xce\x0d\x93\xd1\x38\x62\xe5\x8a\x6b\x06\xa5\xa2\xd7\x85\x37\x42\xf0\xd6\x14\xfb\x3e\xbe\xff\xa0\xdc\xd5\xa5\x17\xcb\x78\x9a\xda\xab\x89\x17\x6a\x69\xa6\x41\xe6\x78\x69\xbf\xd2\xab\xcc\xd6\x48\x03\x1e\x2d\xcd\x49\x83\x2b\x3c\x73\x95\xee\xec\xea\x22\xa3\x37\x98\x5a\xb5\xf5\x65\xdf\x5c\xc9\xa8\x4c\x37\x3a\xf0\x7a\xc1\xc9\x74\xb5\xec\x64\x9a\x28\x3e\xa3\x0e\x26\x78\x35\x5c\x19\x27\x20\xdc\x1c\xb4\xf0\xdc\x6c\x7a\x06\x2a\x76\x95\x4c\x43\x93\x1c\xad\x60\xbe\x9a\x8b\xcb\x8c\xd3\x46\x3e\x2d\x37\x24\x18\xf5\x2c\x97\xf0\xb3\x8f\x80\x9f\x55\x1d\x2c\xbe\x81\x79\xf0\x9a\x56\xd6\x97\x73\x78\x4d\x93\x75\x12\x76\x85\x65\x3b\xbf\x1d\x46\x6d\x5f\x0e\x25\x08\x02\xc2\x35\x27\x33\xfb\x00\x56\x86\xb1\xab\x3c\xdb\xd7\xda\xa7\xa9\x81\xb3\x26\x43\xc6\xd4\xaa\x3b\x4f\x77\xdc\xdd\xb5\x60\xcc\x0c\x7a\x0b\xff\x6f\x31\x37\x31\x33\xde\xce\xa6\x4d\x74\x3d\xa6\xf5\x62\x30\xa8\x8c\x1c\x5f\x94\x3d\x25\x75\x5c\xde\x22\x4f\xae\x37\x8d\xc3\xac\xb8\xa1\x5b\x07\x62\xbf\x21\xf9\x72\xe8\x77\xe8\x5a\x4b\xce\x0d\x3d\x1e\x7f\xce\x1e\x74\x7d\xaa\xf6\x94\x3e\x8b\xe8\x18\xca\xfe\xb4\x93\x30\x02\x90\x8c\x3d\xed\x27\xb5\x4b\x9c\x47\x69\x6a\x31\x22\x05\x2f\x1d\xba\x8d\xfe\xee\xfc\x8e\x06\x9c\xdf\x8c\x70\x7c\xad\x4c\x51\x89\x42\xed\x3f\x07\xbc\x83\xba\xe1\x76\x1a\xdb\x1f\x7a\x68\xfe\x67\x9d\x07\xab\x4c\x89\x8a\xd7\xf2\x89\xfd\xfe\xbb\x09\x9f\xc9\x07\xb7\xc4\x78\xdd\x17\x67\x9b\x19\x88\x35\x70\x6c\x72\x45\x21\xf6\x51\x95\x38\xdc\x19\xa4\x62\xe7\x89\xd6\xa3\xf0\xcf\xbf\xb3\xff\x02\x00\x00\xff\xff\x9d\xd8\x27\xd2\x34\x12\x00\x00")

func auditorKubeshieldTo_dashboardsV1YamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_dashboardsV1Yaml,
		"auditor.kubeshield.to_dashboards.v1.yaml",
	)
}

func auditorKubeshieldTo_dashboardsV1Yaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_dashboardsV1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_dashboards.v1.yaml", size: 4660, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _auditorKubeshieldTo_dashboardsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x6f\x1c\x35\x10\x7f\xbf\xbf\x62\x14\x90\x02\xa8\x77\x47\x54\x54\xc1\xbe\xa0\x28\xd0\x2a\xa2\x94\x2a\x69\xfb\x52\x15\xc9\xb7\x9e\xbb\x35\xf1\xda\xc6\x33\x7b\x69\x8a\xf8\xdf\xd1\xd8\xfb\x95\xcb\xde\x5d\x40\xaa\x84\x9f\x6e\xc7\xf6\x7c\xfc\xe6\xd3\xa7\x82\x79\x87\x91\x8c\x77\x05\xa8\x60\xf0\x23\xa3\x93\x2f\x5a\xdc\x7c\x4f\x0b\xe3\x97\xdb\xb3\x15\xb2\x3a\x9b\xdd\x18\xa7\x0b\xb8\x68\x88\x7d\x7d\x85\xe4\x9b\x58\xe2\x4f\xb8\x36\xce\xb0\xf1\x6e\x56\x23\x2b\xad\x58\x15\x33\x80\x32\xa2\x12\xe2\x1b\x53\x23\xb1\xaa\x43\x01\xae\xb1\x76\x06\x60\xd5\x0a\x2d\xc9\x19\x00\x15\xc2\xe2\xa6\x59\x61\x74\xc8\x98\x44\x39\x55\x63\x01\x42\xa3\xca\xa0\xd5\x33\x80\x4c\xd2\x8a\xaa\x95\x57\x51\xd3\x42\x35\xda\xb0\x8f\x8b\xe1\xd4\x82\xfd\x8c\x02\x96\xc2\x75\x13\x7d\x13\x0a\x98\x3e\x94\xb9\xb5\xd2\x4b\xc5\xb8\xf1\xd1\x74\xdf\xf3\xee\x52\xfb\x75\x4f\x8b\xb4\x1d\x02\x95\x5e\x63\xfa\xcc\x60\x9c\xcb\x8d\x2b\xdc\x18\xe2\x98\x0c\x4e\x7b\xd6\x10\xff\x32\xbd\xff\xd2\x10\xa7\x33\xc1\x36\x51\xd9\xb1\x61\x89\x4c\xc6\x6d\x1a\xab\xe2\x68\x63\x06\x10\x22\x12\xc6\x2d\xbe\x75\x37\xce\xdf\xba\xe7\xa2\x14\x15\xb0\x56\x96\x44\x1b\x2a\x7d\xc0\x02\x5e\x89\x6d\x41\x95\x28\x57\xa8\x59\xc5\xd6\x47\xad\x7d\xc4\x8a\x1b\x2a\xe0\xaf\xbf\x67\x00\x5b\x65\x8d\x4e\x1a\xe5\x4d\x1f\xd0\x9d\xbf\xbe\x7c\xf7\xf4\xba\xac\xb0\x56\x99\x28\x82\x7d\xc0\xc8\x3d\x46\xd9\x6b\x7d\xbc\xf4\x34\x00\x8d\x54\x46\x13\x12\x47\x38\x15\x56\xf9\x0c\x68\x89\x10\x24\xe0\x0a\x61\x9b\x69\xa8\x81\x92\x18\xf0\x6b\xe0\xca\x10\x44\x4c\x26\x3a\x1e\x40\xec\x96\x5f\x83\x72\xe0\x57\x7f\x60\xc9\x0b\xb8\x16\x18\x22\x01\x55\xbe\xb1\x1a\x4a\xef\xb6\x18\x19\x22\x96\x7e\xe3\xcc\xa7\x9e\x33\x01\xfb\x24\xd2\x2a\xc6\x16\xf2\x6e\x19\xc7\x18\x9d\xb2\x02\x42\x83\x4f\x40\x39\x0d\xb5\xba\x83\x88\x22\x03\x1a\x37\xe2\x96\x8e\xd0\x02\x7e\xf5\x11\xc1\xb8\xb5\x2f\xa0\x62\x0e\x54\x2c\x97\x1b\xc3\x5d\x86\x94\xbe\xae\x1b\x67\xf8\x6e\x59\x7a\xc7\xd1\xac\x1a\xf6\x91\x96\x1a\xb7\x68\x97\x64\x36\x73\x15\xcb\xca\x30\x96\xdc\x44\x5c\xaa\x60\xe6\x49\x71\xc7\x29\xcd\x6a\xfd\x45\xef\xaa\xd3\x91\xa6\x7c\x27\x5e\x25\x8e\xc6\x6d\x7a\x72\x8a\xbb\xbd\xb8\x4b\xd4\x81\x21\x50\xed\xb5\xac\xff\x00\xaf\x90\x04\x95\xab\x9f\xaf\xdf\x40\x27\x34\xb9\xe0\x3e\xe6\x09\xed\xe1\x1a\x0d\xc0\x0b\x50\xc6\xad\x31\x66\xc7\xad\xa3\xaf\x13\x47\x74\x3a\x78\xe3\x38\x7d\x94\xd6\xa0\xbb\x0f\x3a\x35\xab\xda\xb0\x78\xfa\xcf\x06\x89\xc5\x3f\x0b\xb8\x50\xce\x79\x86\x15\x42\x13\xb4\x62\xd4\x0b\xb8\x74\x70\xa1\x6a\xb4\x17\x8a\xf0\xb3\xc3\x2e\x08\xd3\x5c\x20\x3d\x0e\xfc\xb8\xbc\x75\x6b\x2a\x3d\x64\xa5\x9a\x75\x8f\x02\x50\xab\x8f\x2f\xd1\x6d\xb8\x2a\xe0\xd9\xd3\x9d\xbd\xa0\x58\x42\xb2\x80\xdf\xdf\xab\xf9\xa7\x0f\x5f\xbd\x9f\xab\xf9\xa7\x6f\xe7\x3f\x7c\xf8\xe6\x7d\xfb\xe3\xeb\x1f\xbf\xdc\xb9\x33\xa9\x64\x47\xce\x0e\xec\xc9\x5d\x71\x3c\xa6\x77\x6c\xec\x2e\x09\xc0\x30\xd6\x0f\x88\xfb\x79\x0c\x08\xa4\x5a\x34\xb9\xbb\x13\xb6\x7d\xe1\x92\x2a\xa1\xb8\x2d\x09\x8d\x45\xa8\x15\x97\x95\xe4\xdf\x1b\x09\xb1\x3a\xf0\xdd\x43\x8b\xc7\xeb\xe4\xa4\xbb\x02\xce\xbb\x79\xaf\x85\xee\x83\x9d\x16\x70\xee\x5a\x56\x52\xa7\xc1\xd4\xc1\x1a\x24\xc0\x2d\xc6\xbb\x3d\x6c\x7b\x3e\x8b\xc9\x03\x7b\x10\xca\x6b\x8f\x9f\x76\x0f\xa8\x18\xd5\x94\x78\x01\x39\x15\xc5\xc7\x00\xf9\x5b\x77\x58\x6a\x80\x24\x62\x7f\x1b\x56\x28\xb9\x1f\x30\xae\x7d\xac\x51\xff\x7b\x3b\x0e\x0b\x12\x23\xa4\x54\xf7\x35\x65\x47\xf2\x1e\xa6\x00\x65\x85\xe5\x0d\x6a\x58\xfb\x08\x4a\xd7\x86\x52\xd3\x48\xf9\xec\xed\xe7\x81\x74\xa7\x2f\x1e\x34\xb4\x9b\x73\xf6\x47\xe6\x64\x34\xed\x51\x5c\x59\x9b\x0a\x38\x81\x71\xe9\xe3\xfc\xf5\x65\x9e\x58\xe8\x3f\x44\xd6\x3d\x45\x5f\x08\x97\x41\xdb\xa1\x74\x0f\x2e\x19\x24\xbb\xbd\xee\xe8\x15\x9a\xd6\xe7\x58\xda\xe7\x95\x67\xb0\xbd\xdb\x53\x9a\x77\x71\x24\x99\x96\xa7\x02\x1c\x74\xc9\xd8\x4b\x4c\x28\xe3\xf6\x61\x9b\x97\xdc\x1b\xa5\xfa\x6e\xd9\x18\xe3\x92\x1a\x95\xf4\x97\xf3\xd7\x97\x07\x79\x1e\xc4\x03\x1e\x13\x8f\x30\x8a\xba\x57\xc3\x04\xfa\x08\x70\xae\xc6\xb7\x72\x67\x4f\x71\x36\x4e\x35\xe3\x88\x95\x2b\xf1\xa0\x11\xa9\x84\x75\x41\x8c\x10\xbc\x35\xe5\xdd\x10\xc5\x6f\x29\x8f\x06\xd2\xd2\x65\xba\x4c\x7d\xda\xc4\xbd\x91\x9c\xd7\x28\x3b\xbc\x74\x71\xe9\x33\x66\x6d\xa4\x8f\x4f\x17\x59\x91\x7f\x90\x63\xaa\xc2\xbd\x45\x5d\x24\x0c\xa6\x52\xab\xb2\x3e\xe4\x8f\x83\x59\x93\xd7\xa3\x5c\x76\xac\x90\xe4\x75\xa4\x9c\xe4\xb5\xa7\xa8\x4c\x3a\x94\xe0\xd6\x70\x65\x9c\x98\x7e\xd0\x88\x21\x59\xe1\x7a\x34\xee\x83\x8a\x5d\x7d\xd2\xd0\x24\xc7\x2a\x38\x59\x9e\x88\x93\x8c\xd3\x46\x1e\x3e\x47\x93\x68\xf4\x7e\x58\xc0\x73\x1f\x01\x3f\xaa\x3a\x58\x7c\x02\x27\xc1\x6b\x5a\x5a\xbf\x39\x81\xdb\x34\x84\x27\x51\xdd\x60\x77\x98\x71\x3b\x93\xfb\xcd\x98\xbf\xd8\x2e\x3c\x73\xba\xb2\x0f\x60\x65\x82\xeb\xf1\x38\xc8\x52\xfb\xd4\xeb\xb9\xd5\x62\xfc\xec\x69\x55\x3d\x49\x67\xdc\x69\x77\x64\xd0\xff\xff\x12\x43\x93\x33\xdb\x63\x59\xb4\x89\xaa\x1f\xea\x3a\x1f\x0d\x0f\x0f\x36\x0f\xc8\x9c\x96\x36\x25\x67\x9e\x27\xc6\xa3\xe3\x67\x7e\x79\x3e\x62\x00\xf5\xab\xf4\xce\xd5\x2f\xd0\xb5\x9a\xef\x9a\x75\x7f\x0c\x79\x70\xbc\xeb\x23\xb5\xa7\xf4\x2a\x44\xc7\xb0\x19\x76\x3b\xfe\x0f\xe0\x90\xf1\xa3\x7d\x89\x76\x41\x7f\x29\x2d\x27\x46\xa4\xe0\xa5\x6f\xb6\xb1\xdb\xed\x9f\xd2\x88\xef\x13\xb8\xad\x4c\x59\x3d\xe0\x6a\xa8\x7b\xd8\x80\x77\x50\x37\xdc\x4e\x44\x77\x7d\x7f\xcb\xcf\xaa\xdd\x40\x94\x19\x4d\x71\x21\x2f\xd5\x67\xdf\x4d\x7a\x47\xde\xb0\x1b\x8c\x87\x91\xdf\x21\x6d\xbb\x7f\x79\xb6\x67\xca\x86\x4a\x9d\x0d\xb4\xe4\x86\x79\xfb\x5f\xcb\x68\x1b\x20\xe3\x55\x00\xc7\x06\xdb\xbf\x11\x7c\x54\x1b\x6c\x29\xff\x04\x00\x00\xff\xff\xff\x71\x1e\x0c\x3d\x12\x00\x00")

func auditorKubeshieldTo_dashboardsYamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_dashboardsYaml,
		"auditor.kubeshield.to_dashboards.yaml",
	)
}

func auditorKubeshieldTo_dashboardsYaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_dashboardsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_dashboards.yaml", size: 4669, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _auditorKubeshieldTo_dashboardtemplatesV1Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x8f\x23\x35\x10\xbd\xe7\x57\x94\xc4\x61\x2f\x24\xd1\x08\x84\x50\xdf\x50\x96\xc3\x08\x90\xd0\xce\x68\xef\x95\x76\xa5\x53\x8c\xdb\x36\x55\xe5\xc0\x2c\xe2\xbf\x23\xbb\xd3\x93\x0f\x9c\xd9\x15\x62\xfa\xe6\xfa\x7c\x7e\xaf\xca\x8d\x89\x3f\x92\x28\xc7\xd0\x01\x26\xa6\x3f\x8d\x42\x39\xe9\xea\xe9\x7b\x5d\x71\x5c\x1f\xee\x16\x4f\x1c\x5c\x07\x9b\xac\x16\xc7\x0f\xa4\x31\x4b\x4f\xef\x69\xc7\x81\x8d\x63\x58\x8c\x64\xe8\xd0\xb0\x5b\x00\xf4\x42\x58\x8c\x8f\x3c\x92\x1a\x8e\xa9\x83\x90\xbd\x5f\x00\x78\xdc\x92\xd7\x12\x03\x80\x29\xad\x9e\xf2\x96\x24\x90\x51\xed\x12\x70\xa4\x0e\x8a\x4d\xf7\x4c\xde\x2d\x00\x26\x93\x43\xdd\x6f\x23\x8a\x33\x1a\x93\xc7\x12\x8e\xd9\xb1\x45\x59\x9d\xa2\x57\x16\x17\x9a\xa8\x2f\xd5\x07\x89\x39\x75\xd0\x0e\x9a\xaa\x1e\x51\xf4\x68\x34\x44\xe1\xf9\xbc\x9c\x93\x8e\xa7\x0b\x34\xd5\x9d\x92\xf6\xd1\x51\x3d\x4e\xa4\xbc\x9f\xe1\x3d\x1e\xe1\x55\x9f\x67\xb5\x9f\xda\xfe\x9f\x59\xad\xc6\x24\x9f\x05\x7d\xeb\x82\xd5\xad\x1c\x86\xec\x51\x1a\x01\x0b\x00\xed\x63\xa2\x0e\x36\x3e\xab\x51\xc1\x7b\x98\x34\xac\x37\x59\x1e\xa9\x3b\xdc\xa1\x4f\x7b\xbc\x9b\xea\xf5\x7b\x1a\x71\xba\x28\x40\x4c\x14\x7e\xf8\xf5\xfe\xe3\x37\x0f\x17\x66\x80\x24\x31\x91\xd8\x0b\x27\xd3\x77\x36\x24\x67\x56\x00\x47\xda\x0b\x27\xab\xd3\xf3\xae\x14\x9c\xa2\xc0\x95\xe9\x20\x05\xdb\xd3\x0c\x8d\xdc\x11\x03\xc4\x1d\xd8\x9e\x15\x84\x92\x90\x52\xb0\x3a\x31\x17\x85\xa1\x04\x61\x80\xb8\xfd\x8d\x7a\x5b\xc1\x03\x49\x29\x03\xba\x8f\xd9\x3b\xe8\x63\x38\x90\x18\x08\xf5\x71\x08\xfc\xe9\xa5\xb6\x82\xc5\xda\xb4\x12\x69\x57\x35\x39\x18\x49\x40\x0f\x07\xf4\x99\xbe\x06\x0c\x0e\x46\x7c\x06\xa1\xd2\x05\x72\x38\xab\x57\x43\x74\x05\xbf\x44\x21\xe0\xb0\x8b\x1d\xec\xcd\x92\x76\xeb\xf5\xc0\x36\x2f\x47\x1f\xc7\x31\x07\xb6\xe7\x75\x1f\x83\x09\x6f\xb3\x45\xd1\xb5\xa3\x03\xf9\xb5\xf2\xb0\x44\xe9\xf7\x6c\xd4\x5b\x16\x5a\x63\xe2\x65\x85\x1e\xac\x6e\xd8\xe8\xbe\x92\xe3\x3a\xe9\xbb\x0b\xac\xf6\x5c\xe4\x55\x13\x0e\xc3\x99\xa3\xce\xdc\x2b\x0a\x94\x99\x03\x56\xc0\x63\xea\x74\x8b\x13\xd1\xc5\x54\xd8\xf9\xf0\xe3\xc3\x23\xcc\xad\xab\x18\xd7\xec\x57\xde\x4f\x89\x7a\x92\xa0\x10\xc6\x61\x47\x32\x89\xb8\x93\x38\xd6\x9a\x14\x5c\x8a\x1c\xac\x1e\x7a\xcf\x14\xae\xe9\xd7\xbc\x1d\xd9\x8a\xee\xbf\x67\x52\x2b\x5a\xad\x60\x83\x21\x44\x83\x2d\x41\x4e\x0e\x8d\xdc\x0a\xee\x03\x6c\x70\x24\xbf\x41\xa5\x37\x17\xa0\x30\xad\xcb\x42\xec\x97\x49\x70\xfe\xd8\x5d\x07\x4f\xac\x9d\x39\xe6\x37\xe9\xf4\xb5\xf7\xab\x2a\x79\xbd\xe6\xd7\x01\xaf\x25\xd7\xe9\x76\x2d\x2b\xc0\x2e\xca\x88\xd6\x95\xe9\xff\xee\xdb\x66\xc4\x84\xbe\x6c\xc7\x40\xd2\x88\x98\x76\xab\xf9\x02\xfc\x7f\x4d\x0c\x87\xe6\xb5\x00\xd8\x68\xbc\xe1\xba\x29\xd3\x75\x00\x8a\xe0\x73\xab\x2b\x8f\xf4\x29\x86\x06\xd9\x5f\x50\xdc\xd8\xfc\x7f\xcb\xcc\xb7\xc4\xfa\x4c\xde\xe1\x2d\x35\x28\x7b\xc9\x42\x0d\x64\x4b\x60\xd7\x30\x5e\xcc\x45\xc3\x5f\x24\x6d\x99\x8f\x9c\x37\x5d\xe6\x5b\xf6\xdc\xec\x7f\xb8\xd1\xf9\xc6\x36\x42\xe5\xc8\x3b\x92\xfb\xc6\x1d\x5f\x67\xef\x35\xe6\xe2\x81\xe4\x0f\xe1\xd6\xc6\x4e\x69\xdb\x18\x3d\xe1\x25\xce\x36\xd9\xcb\xe6\xdf\xfe\xdc\x3f\x5f\xe0\xca\xfc\x82\xe1\xf3\xcf\xd2\xbf\x8c\x5a\x5e\x76\xd7\x81\x49\x9e\xd2\xd5\xa2\xe0\x40\xe7\x96\xbc\x7d\xf9\x53\xcd\x88\xd5\xd0\xb2\x76\xf0\xd7\xdf\x8b\x7f\x02\x00\x00\xff\xff\x0d\xb9\x0d\xda\x45\x0a\x00\x00")

func auditorKubeshieldTo_dashboardtemplatesV1YamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_dashboardtemplatesV1Yaml,
		"auditor.kubeshield.to_dashboardtemplates.v1.yaml",
	)
}

func auditorKubeshieldTo_dashboardtemplatesV1Yaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_dashboardtemplatesV1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_dashboardtemplates.v1.yaml", size: 2629, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _auditorKubeshieldTo_dashboardtemplatesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x23\x35\x10\xbd\xe7\x57\x94\xc4\x61\x2f\x24\xd1\x08\x84\x50\xdf\x50\x16\xa4\x15\x20\xa1\x9d\x61\xef\x95\x76\xa5\x53\x8c\xdb\x36\x55\xe5\x5e\x66\x11\xff\x1d\xd9\xdd\xf9\x9a\xee\x99\x15\x12\x9b\x9b\xdf\xab\x2f\xbf\xaa\x72\x1a\x13\x7f\x20\x51\x8e\xa1\x01\x4c\x4c\x7f\x19\x85\x72\xd2\xcd\xe3\xf7\xba\xe1\xb8\x1d\xee\xf6\x64\x78\xb7\x7a\xe4\xe0\x1a\xd8\x65\xb5\xd8\xbf\x27\x8d\x59\x5a\x7a\x4b\x07\x0e\x6c\x1c\xc3\xaa\x27\x43\x87\x86\xcd\x0a\xa0\x15\xc2\x02\x3e\x70\x4f\x6a\xd8\xa7\x06\x42\xf6\x7e\x05\xe0\x71\x4f\x5e\x8b\x0d\x00\xa6\xb4\x79\xcc\x7b\x92\x40\x46\x35\x55\xc0\x9e\x1a\x28\x98\x1e\x99\xbc\x5b\x01\x8c\x90\x43\x3d\xee\x23\x8a\x33\xea\x93\xc7\x62\x8e\xd9\xb1\x45\xd9\x5c\xac\x37\x16\x57\x9a\xa8\x2d\xd1\x3b\x89\x39\x35\xb0\x6c\x34\x46\x9d\xaa\x68\xd1\xa8\x8b\xc2\xa7\xf3\xfa\xe4\x34\x9d\x6e\xaa\xa9\x74\x4a\xda\x46\x47\xf5\x38\x8a\xf2\xf6\x54\xde\xc3\x54\x5e\xe5\x3c\xab\xfd\xbc\xcc\xff\xc2\x6a\xd5\x26\xf9\x2c\xe8\x97\x2e\x58\x69\xe5\xd0\x65\x8f\xb2\x60\xb0\x02\x48\x42\x4a\x32\xd0\xef\xe1\x31\xc4\x8f\xe1\xa7\x52\xa4\x36\x70\x40\xaf\x85\xd6\x36\x26\x6a\x60\xe7\xb3\x1a\x95\xeb\x68\xde\xcb\xd4\xb8\xe9\xb2\x6a\x68\x59\x1b\xf8\xfb\x9f\x15\xc0\x80\x9e\x5d\xed\xdb\x48\xc6\x44\xe1\x87\xdf\xde\x7d\xf8\xe6\xbe\x3d\x52\x8f\x23\x58\xb2\xc6\x44\x62\x67\xc1\xc6\x56\x9e\x87\xe8\x8c\x01\x38\xd2\x56\x38\xd5\x88\xf0\xa6\x84\x1a\x6d\xc0\x95\xb1\x21\x05\x3b\x12\x0c\x23\x46\x0e\xb4\xa6\x81\x78\x00\x3b\xb2\x82\x50\xbd\x5f\xb0\x5a\xd2\x55\x58\x28\x26\x18\x20\xee\xff\xa0\xd6\x36\x70\x5f\x34\x10\x05\x3d\xc6\xec\x1d\xb4\x31\x0c\x24\x06\x42\x6d\xec\x02\x7f\x3a\x47\x56\xb0\x58\x53\x56\x7d\xed\x26\x22\x07\x23\x09\xe8\x8b\x08\x99\xbe\x06\x0c\x0e\x7a\x7c\x02\xa1\x92\x03\x72\xb8\x8a\x56\x4d\x74\x03\xbf\x46\x21\xe0\x70\x88\x0d\x1c\xcd\x92\x36\xdb\x6d\xc7\x76\x5a\x9b\x36\xf6\x7d\x0e\x6c\x4f\xdb\x36\x06\x13\xde\x67\x8b\xa2\x5b\x47\x03\xf9\xad\x72\xb7\x46\x69\x8f\x6c\xd4\x5a\x16\xda\x62\xe2\x75\x2d\x3c\x58\xdd\xbd\xde\x7d\x75\x6e\xd5\x9b\xab\x4a\xed\xa9\xb4\x54\x4d\x38\x74\x67\xb8\x0e\xe1\x8b\xba\x97\x11\x04\x56\xc0\xc9\x6d\xac\xff\x22\x6f\x81\x8a\x2a\xef\x7f\xbc\x7f\x80\x53\xd2\xda\x82\x5b\xcd\xab\xda\x17\x37\xbd\x08\x5f\x84\xe2\x70\x20\x19\x1b\x77\x90\xd8\xd7\x88\x14\x5c\x8a\x1c\xac\x1e\x5a\xcf\x14\x6e\x45\xd7\xbc\xef\xd9\x4a\xa7\xff\xcc\xa4\x56\xfa\xb3\x81\x1d\x86\x10\x0d\xf6\x04\x39\x39\x34\x72\x1b\x78\x17\x60\x87\x3d\xf9\x1d\x2a\x7d\x71\xd9\x8b\xc2\xba\x2e\x92\x7e\x5e\xf8\xeb\x37\xef\xd6\x70\x54\xeb\x0c\x9f\x9e\xa5\xd3\x6f\x69\x87\x6a\xe7\x9e\x6f\xf9\x2d\xfd\xb2\x63\x9d\x61\x37\xc7\x00\x0e\x51\x7a\xb4\xa6\x4c\xf8\x77\xdf\x2e\xf0\x63\xbd\x65\xfe\x3b\x92\x19\x3f\x6e\xce\xc2\x6e\xff\x3f\xe1\x0d\xbb\x85\x8b\x00\xb0\x51\xbf\x48\xbc\xd0\x88\xe7\x34\x8a\xe0\xd3\x3c\x1b\xf7\xf4\x29\x86\x99\xa8\x9f\x0d\x6b\x6c\xfe\xbf\x7b\xe5\xe5\x86\xbc\xea\x33\x7c\x19\xa5\xcb\x86\xb1\xd0\xac\x9e\x35\xb0\x9b\x41\x37\x3d\x9f\xb1\xa5\x65\x73\x70\x52\x76\x81\x30\x3f\x47\xf3\x42\xd6\x61\x31\xdf\xe2\x36\x8d\x6a\x78\x47\xf2\x6e\x76\xa3\xd7\x54\x7a\x59\xa1\x38\x90\x7c\x14\x9e\xef\xdb\xe8\xb2\x8f\xd1\x13\x5e\xd7\xb6\x24\xe8\x7a\xf1\x4f\xfa\xc2\x9e\x4a\xbe\x01\xcf\x99\x5f\x7f\x42\x9e\x41\xa7\x31\x81\xe1\x0e\x7d\x3a\xe2\xdd\x05\xab\x6b\xb3\x9e\xbe\x9b\xae\x68\x80\xfa\xad\xe0\x1a\x30\xc9\x34\xfd\xfb\x47\xc1\x8e\x26\xe4\xdf\x00\x00\x00\xff\xff\x38\xe2\x23\xf2\x09\x0a\x00\x00")

func auditorKubeshieldTo_dashboardtemplatesYamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_dashboardtemplatesYaml,
		"auditor.kubeshield.to_dashboardtemplates.yaml",
	)
}

func auditorKubeshieldTo_dashboardtemplatesYaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_dashboardtemplatesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_dashboardtemplates.yaml", size: 2569, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _auditorKubeshieldTo_datasourcesV1Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x8f\xe3\x36\x0c\xbd\xe7\x57\x10\xe8\x61\x2e\x1b\x07\x83\x16\x45\xe1\xdb\x62\xb6\x28\x06\xfd\x5a\xec\x2c\xf6\xae\xc8\x8c\xad\x8e\x4d\xaa\x24\x95\x76\x5a\xf4\xbf\x17\x92\xec\xc4\x49\x06\xdd\x5e\xd6\x37\x91\xf4\x23\xf5\xde\xa3\x5c\x0c\x9f\x50\x34\x30\xb5\xe0\x62\xc0\x3f\x0d\x29\x9f\xb4\x79\xfe\x4e\x9b\xc0\xbb\xe3\xfd\xe6\x39\x50\xd7\xc2\x43\x52\xe3\xe9\x03\x2a\x27\xf1\xf8\x0e\x0f\x81\x82\x05\xa6\xcd\x84\xe6\x3a\x67\xae\xdd\x00\x78\x41\x97\x83\x1f\xc3\x84\x6a\x6e\x8a\x2d\x50\x1a\xc7\x0d\xc0\xe8\xf6\x38\x6a\xae\x01\x70\x31\x36\xcf\x69\x8f\x42\x68\x58\xba\x90\x9b\xb0\x85\x1c\xd3\x21\xe0\xd8\x6d\x00\x6a\x28\x03\xd7\x8e\xda\xb8\xd4\x05\x63\x69\xce\x65\x8d\xf1\x46\x23\xfa\x0c\xdb\x0b\xa7\xd8\xc2\xeb\x45\x15\x6e\x6e\xef\x9d\x61\xcf\x12\x96\xf3\x76\xf9\x69\x3e\x5d\x8c\x51\xd2\x31\xaa\xe7\x0e\xcb\xb1\xb2\xf1\xee\x34\x57\x09\x8e\x41\xed\xc7\xab\xc4\x4f\x41\xad\x24\xe3\x98\xc4\x8d\x17\x77\x29\x71\x0d\xd4\xa7\xd1\xc9\x3a\xb3\x01\x50\xcf\x11\x5b\xf8\x25\x0f\x1c\x9d\xc7\x3c\xc5\xb1\x6a\x54\x06\xde\xce\xd4\x1c\xef\xdd\x18\x07\x77\x5f\xb1\xfc\x80\x93\xab\xf7\x01\xe0\x88\xf4\xf6\xfd\xe3\xa7\xaf\x9f\x2e\xc2\x00\x51\x38\xa2\xd8\xe9\xea\xf5\x5b\x99\x60\x15\x05\xe8\x50\xbd\x84\x68\xc5\x1d\x77\x19\xb0\x56\x41\x97\xd5\x47\x05\x1b\x70\x19\x0d\xbb\x79\x06\xe0\x03\xd8\x10\x14\x04\xa3\xa0\x22\x59\x71\xc4\x05\x30\xe4\x22\x47\xc0\xfb\xdf\xd0\x5b\x03\x4f\x28\x19\x06\x74\xe0\x34\x76\xe0\x99\x8e\x28\x06\x82\x9e\x7b\x0a\x7f\x9d\xb0\x15\x8c\x4b\xd3\xd1\x19\xce\xec\x9e\xbf\x40\x86\x42\x6e\x84\xa3\x1b\x13\xbe\x01\x47\x1d\x4c\xee\x05\x04\x73\x17\x48\xb4\xc2\x2b\x25\xda\xc0\xcf\x2c\x08\x81\x0e\xdc\xc2\x60\x16\xb5\xdd\xed\xfa\x60\x8b\xf9\x3d\x4f\x53\xa2\x60\x2f\x3b\xcf\x64\x12\xf6\xc9\x58\x74\xd7\xe1\x11\xc7\x9d\x86\x7e\xeb\xc4\x0f\xc1\xd0\x5b\x12\xdc\xb9\x18\xb6\x65\x74\xb2\xb2\x41\x53\xf7\x95\xcc\xeb\xa2\x77\x17\xb3\xda\x4b\x56\x58\x4d\x02\xf5\xab\x44\xb1\xd6\x7f\x28\x90\x1d\x06\x41\xc1\xcd\xbf\xd6\x5b\x9c\x89\xce\xa1\xcc\xce\x87\xef\x9f\x3e\xc2\xd2\xba\x88\x71\xcd\x7e\xe1\xfd\xfc\xa3\x9e\x25\xc8\x84\x05\x3a\xa0\x54\x11\x0f\xc2\x53\xc1\x44\xea\x22\x07\xb2\x72\xf0\x63\x40\xba\xa6\x5f\xd3\x7e\x0a\x96\x75\xff\x3d\xa1\x5a\xd6\xaa\x81\x07\x47\xc4\x06\x7b\x84\x14\x3b\x67\xd8\x35\xf0\x48\xf0\xe0\x26\x1c\x1f\x9c\xe2\x17\x17\x20\x33\xad\xdb\x4c\xec\xff\x93\x60\xfd\x98\x5d\x17\x57\xd6\x56\x89\xe5\xe9\xf9\x7c\xa1\x39\x4b\x7a\x59\xfa\xfa\x2a\x56\x75\x34\x8b\xd1\xfd\x80\x84\x52\x96\xe7\xba\xe2\xca\x18\xbf\xde\xfc\x90\x5d\x92\x85\x9a\x58\xcb\x1a\x21\x19\xf4\xe7\xec\xd2\xe1\x06\x16\xe0\xc0\xb2\xac\x6f\xf5\x4f\x03\x8f\x06\x9e\x45\x50\x23\x53\x77\x5a\xc1\x25\x7f\xa7\x2b\xe4\x37\xaf\x20\xfe\x31\x04\x3f\xe4\x81\x66\x07\x00\x13\x4c\xa9\xbe\x0a\xb0\x7f\x29\x60\x6f\xdf\x3f\xce\x0e\x6c\x6e\x10\x0e\x2c\x93\xb3\x36\xaf\xf7\xb7\xdf\xdc\x64\x2b\xe3\x79\xf5\x7b\x94\xcf\x6b\x71\x13\xac\x44\xb4\x60\x92\xea\x53\xae\xc6\xe2\x7a\x5c\x47\xd2\xfe\xb4\xc7\x8b\x12\xb3\xa2\xf0\xf7\x3f\x9b\x7f\x03\x00\x00\xff\xff\x35\x6d\xe8\xe5\x43\x07\x00\x00")

func auditorKubeshieldTo_datasourcesV1YamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_datasourcesV1Yaml,
		"auditor.kubeshield.to_datasources.v1.yaml",
	)
}

func auditorKubeshieldTo_datasourcesV1Yaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_datasourcesV1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_datasources.v1.yaml", size: 1859, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _auditorKubeshieldTo_datasourcesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4f\x8f\x1b\xb7\x0f\xbd\xfb\x53\x10\xf8\x1d\xf6\x12\x8f\xb1\xf8\x15\x45\x31\xb7\x60\xd3\x16\x8b\xfe\x0b\xb2\x69\xee\xb4\x44\xcf\xa8\xab\xa1\x54\x91\x72\xba\x2d\xfa\xdd\x0b\x49\x33\xe3\xf1\x6e\x92\x9e\xea\x9b\x1e\x39\x8f\xd4\x7b\xa4\x8c\xd1\x7d\xa0\x24\x2e\x70\x0f\x18\x1d\xfd\xa1\xc4\xe5\x24\xdd\xe3\x37\xd2\xb9\x70\x38\xdf\x1e\x49\xf1\x76\xf7\xe8\xd8\xf6\x70\x97\x45\xc3\xf4\x8e\x24\xe4\x64\xe8\x0d\x9d\x1c\x3b\x75\x81\x77\x13\x29\x5a\x54\xec\x77\x00\x26\x11\x16\xf0\xbd\x9b\x48\x14\xa7\xd8\x03\x67\xef\x77\x00\x1e\x8f\xe4\xa5\xe4\x00\x60\x8c\xdd\x63\x3e\x52\x62\x52\xaa\xa5\x18\x27\xea\xa1\x60\x32\x3a\xf2\x76\x07\xd0\xa0\x42\xdc\x2a\x4a\x87\xd9\x3a\x0d\xa9\xbb\xa4\x75\x1a\x76\x12\xc9\x14\xda\x21\x85\x1c\x7b\xf8\x74\x52\xa3\x9b\xcb\x1b\x54\x1a\x42\x72\xcb\x79\xbf\x7c\x34\x9f\xae\xda\xa8\xe1\x18\xc5\x04\x4b\xf5\xd8\xd4\x78\xb3\xf6\x55\x41\xef\x44\x7f\x78\x16\xf8\xd1\x89\xd6\x60\xf4\x39\xa1\xbf\xba\x4b\xc5\xc5\xf1\x90\x3d\xa6\x6d\x64\x07\x10\x13\x09\xa5\x33\xfd\xca\x8f\x1c\x3e\xf2\x77\xa5\x11\xe9\xe1\x84\x5e\x4a\x58\x4c\x88\xd4\xc3\xcf\xe5\x3e\x11\x0d\x95\x26\x25\x1f\xd3\x6c\xcc\x7c\x27\x51\xd4\x2c\x3d\xfc\xf5\xf7\x0e\xe0\x8c\xde\xd9\xea\x4b\x0b\x86\x48\xfc\xfa\xed\xfd\x87\xff\x3f\x98\x91\x26\x6c\x60\x29\x1c\x22\x25\x5d\x75\x69\x56\xad\x43\xb2\x62\x00\x96\xc4\x24\x17\x2b\x23\xdc\x14\xaa\x96\x03\xb6\x8c\x05\x09\xe8\x48\x70\x6e\x18\x59\x90\x5a\x06\xc2\x09\x74\x74\x02\x89\xea\x15\x59\x6b\x4b\x1b\x5a\x28\x29\xc8\x10\x8e\xbf\x91\xd1\x0e\x1e\x8a\x0c\x49\x40\xc6\x90\xbd\x05\x13\xf8\x4c\x49\x21\x91\x09\x03\xbb\x3f\x57\x66\x01\x0d\xb5\xa4\x47\xa5\x59\xf4\xe5\xe7\x58\x29\x31\xfa\x22\x42\xa6\x57\x80\x6c\x61\xc2\x27\x48\x54\x6a\x40\xe6\x0d\x5b\x4d\x91\x0e\x7e\x0a\x89\xc0\xf1\x29\xf4\x30\xaa\x46\xe9\x0f\x87\xc1\xe9\xb2\x16\x26\x4c\x53\x66\xa7\x4f\x07\x13\x58\x93\x3b\x66\x0d\x49\x0e\x96\xce\xe4\x0f\xe2\x86\x3d\x26\x33\x3a\x25\xa3\x39\xd1\x01\xa3\xdb\xd7\xc6\x59\xeb\x6e\x4d\xf6\x7f\xab\x55\x37\x9b\x4e\xf5\xa9\xb8\x2a\x9a\x1c\x0f\x2b\x5c\x67\xed\xb3\xba\x97\x81\x03\x27\x80\xf3\x67\xad\xff\x8b\xbc\x05\x2a\xaa\xbc\xfb\xf6\xe1\x3d\x2c\x45\xab\x05\xd7\x9a\x57\xb5\x2f\x9f\xc9\x45\xf8\x22\x94\xe3\x13\xa5\x66\xdc\x29\x85\xa9\x32\x12\xdb\x18\x1c\x6b\x3d\x18\xef\x88\xaf\x45\x97\x7c\x9c\x9c\x16\xa7\x7f\xcf\x24\x5a\xfc\xe9\xe0\x0e\x99\x83\xc2\x91\x20\x47\x8b\x4a\xb6\x83\x7b\x86\x3b\x9c\xc8\xdf\xa1\xd0\x7f\x2e\x7b\x51\x58\xf6\x45\xd2\x7f\x17\x7e\xfb\xa6\x5d\x27\x36\xb5\x56\x78\x79\x7d\xbe\x9c\xd4\x96\x71\x93\xf6\xa9\x55\x6b\x5e\xd4\xd5\xb7\xdf\x13\x53\xda\x6c\xec\x67\x46\xe0\x97\x17\xe9\x65\x1e\x8a\x29\x53\x90\xba\x28\xc4\x0a\xc3\x25\xba\xf0\x3f\x23\x05\x38\x85\xb4\x2c\x67\x9b\x93\x0e\xee\x15\x4c\x48\x89\x24\x06\xb6\xeb\x8a\x2d\xf1\x1b\xd9\xf0\xbe\x82\x8f\xa3\x33\xe3\x0b\x56\x27\x8b\xd7\x10\x18\xa6\xdc\x36\x1e\x8e\x4f\x95\xea\xf5\xdb\xfb\x79\xd2\xba\xdd\x8b\x6e\x26\xd4\xbe\x2c\xef\xd7\x5f\x3d\x8b\x35\x85\xcb\x5a\x0f\x94\xbe\xac\xfc\x33\xe8\xbc\xfc\xdb\x9d\x6f\xd1\xc7\x11\x6f\x2f\x58\xb5\x61\x3f\xff\xe7\x6c\xc2\x00\x4d\xaf\x1e\x34\x65\x9a\x5f\xd6\x90\x70\xa0\x19\xf9\x27\x00\x00\xff\xff\x91\x26\xd4\x9c\x45\x07\x00\x00")

func auditorKubeshieldTo_datasourcesYamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_datasourcesYaml,
		"auditor.kubeshield.to_datasources.yaml",
	)
}

func auditorKubeshieldTo_datasourcesYaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_datasourcesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_datasources.yaml", size: 1861, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"auditor.kubeshield.to_dashboards.v1.yaml":         auditorKubeshieldTo_dashboardsV1Yaml,
	"auditor.kubeshield.to_dashboards.yaml":            auditorKubeshieldTo_dashboardsYaml,
	"auditor.kubeshield.to_dashboardtemplates.v1.yaml": auditorKubeshieldTo_dashboardtemplatesV1Yaml,
	"auditor.kubeshield.to_dashboardtemplates.yaml":    auditorKubeshieldTo_dashboardtemplatesYaml,
	"auditor.kubeshield.to_datasources.v1.yaml":        auditorKubeshieldTo_datasourcesV1Yaml,
	"auditor.kubeshield.to_datasources.yaml":           auditorKubeshieldTo_datasourcesYaml,
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
	"auditor.kubeshield.to_dashboards.v1.yaml":         {auditorKubeshieldTo_dashboardsV1Yaml, map[string]*bintree{}},
	"auditor.kubeshield.to_dashboards.yaml":            {auditorKubeshieldTo_dashboardsYaml, map[string]*bintree{}},
	"auditor.kubeshield.to_dashboardtemplates.v1.yaml": {auditorKubeshieldTo_dashboardtemplatesV1Yaml, map[string]*bintree{}},
	"auditor.kubeshield.to_dashboardtemplates.yaml":    {auditorKubeshieldTo_dashboardtemplatesYaml, map[string]*bintree{}},
	"auditor.kubeshield.to_datasources.v1.yaml":        {auditorKubeshieldTo_datasourcesV1Yaml, map[string]*bintree{}},
	"auditor.kubeshield.to_datasources.yaml":           {auditorKubeshieldTo_datasourcesYaml, map[string]*bintree{}},
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
