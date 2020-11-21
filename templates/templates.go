// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// html/templates/aggregator.html
// html/templates/home.html
// html/templates/layouts/layout.html
// html/templates/login_form.html
// html/templates/stats.html
package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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

// Mode return file modify time
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _aggregatorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcf\x52\xdb\x30\x10\xc6\xef\x79\x8a\x1d\x91\xc1\x0e\x34\x36\x1c\x7a\x21\xb6\x0f\x85\xe9\xf4\x50\x7a\x68\xe1\x01\x84\xbd\x89\xc5\xd8\x72\x46\xbb\x4e\xcb\x64\xfc\xee\x1d\x4b\xb6\x70\x20\x34\x33\xbd\xad\xf4\x7d\xdf\x4f\x7f\x76\x93\xf2\x3a\x9b\x01\x00\xec\xf7\x46\xea\x0d\xc2\x5c\xe9\x02\xff\x7c\x82\x79\x5e\x4a\xad\xb1\x82\x9b\x14\xa2\x5b\x57\x53\xd7\xc1\x7e\xff\x2a\xd9\x25\xea\xc2\x6d\x47\xf7\x52\xe9\x5b\x2f\xcd\x92\xb8\x67\x27\x2c\x9f\x2a\x04\x55\xa4\xa2\x46\x22\xb9\x41\x12\x50\x48\x96\xcb\xd6\x54\xa9\xe8\x73\x5f\x24\xe1\xe3\xcf\xef\xd0\x75\xf1\x33\x35\x3a\x7e\xe3\x93\x2d\x37\x4b\x83\x6b\x83\x54\xa6\x82\x4d\x8b\x47\x84\xa5\xd2\x8c\x66\x27\xab\x54\x7c\x16\xee\x45\x09\x97\x28\x0b\x57\xbb\xb5\x79\x5d\x0c\x06\x07\x5a\x2b\xac\x8a\x54\x0c\x77\x17\xd9\x50\x24\x31\x97\xff\x4e\x3c\xa8\x7a\xbc\xcc\xba\x31\xb5\x64\x46\x93\x0a\x56\x35\x7e\xb5\x4b\x34\x22\xeb\x3d\xa7\x49\x8f\x84\x26\xba\x53\xb4\xad\xe4\xcb\x0f\x59\xa3\xc8\xfa\x9d\xd3\xb9\x7b\xf7\x57\x22\x1b\x8a\xc3\x44\x12\x8f\x6f\xee\xf7\xed\x6f\x24\xb1\x6d\x48\x36\x4b\x28\x37\x6a\xcb\x40\x26\x7f\xd7\x06\x62\xc9\x2a\x8f\x9f\x29\x7e\x6a\x1a\x26\x36\x72\xbb\xb4\xb1\xa8\x56\x3a\x7a\x26\x91\x25\xb1\x8b\xff\x27\xe7\xa0\x77\x1f\x43\xdd\xdd\xe7\x61\xd1\xe4\x6d\x8d\x9a\x17\x91\x41\x59\xbc\x84\xeb\x56\xe7\xac\x1a\x1d\x2e\x60\xef\x1f\x3b\x0f\x83\xb3\x71\x74\x82\x45\xe4\x4f\x7c\xe8\x0f\x0c\x17\x2b\x6b\xec\x16\xab\x99\x2d\xe2\x8b\x11\x02\xfd\x38\x8f\xfd\x0a\x77\xb2\x6a\x71\x8a\x35\xc8\xad\xd1\x10\x9c\x05\x70\x09\x56\x75\x20\x87\xf1\x90\x96\xd0\x9c\x86\x9c\x57\xbc\xf2\x1c\xb8\x84\xe0\x7c\xc3\xab\xc0\x01\x2f\xe2\x37\xc8\xe9\x1c\xbd\x47\xee\xa4\xb1\x0e\x48\x41\xe3\x6f\xb8\x93\x8c\x83\x69\xe5\x3d\x25\xa4\xd6\x13\x6d\x90\xbf\x35\xad\xa1\x70\x22\xd6\x13\xf1\x5e\xe9\x96\xf1\x40\xa6\x89\xfc\x0b\xf3\x46\x17\x56\xf6\xba\x5a\x43\x58\x42\x02\xd7\x57\xd3\x5b\x8d\xa7\x06\x57\xfd\x33\x4b\xbf\xdf\x1d\xe4\xea\xa3\xb9\xda\xe7\xea\x0f\x72\x74\x34\x47\x3e\x47\x47\x72\xc3\xd7\x97\xfd\x6f\xdf\x58\xb8\xaf\x68\xe8\xa4\x9f\xbb\xbf\x01\x00\x00\xff\xff\xd1\xfa\x12\x79\x0e\x05\x00\x00")

func aggregatorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_aggregatorHtml,
		"aggregator.html",
	)
}

func aggregatorHtml() (*asset, error) {
	bytes, err := aggregatorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aggregator.html", size: 1294, mode: os.FileMode(420), modTime: time.Unix(1605896240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb4\x0b\x4f\xcd\x49\xce\xcf\x4d\x55\x28\x2e\x2d\x52\xc8\x49\x54\x48\x4f\x2d\x2e\xc9\xcc\xcf\x53\x48\x29\x55\x48\xca\x2f\x51\x08\x29\xcf\x2c\x49\xce\x50\x28\xc8\x2f\x2d\x52\xb0\x52\xa8\xae\x56\xd0\xf3\x4d\xcc\xcc\x73\xce\x48\xcc\xcb\x4b\xcd\x51\xa8\xad\xb5\xd1\xcf\x30\xb4\x03\x04\x00\x00\xff\xff\x73\x88\xd5\x8b\x47\x00\x00\x00")

func homeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_homeHtml,
		"home.html",
	)
}

func homeHtml() (*asset, error) {
	bytes, err := homeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.html", size: 71, mode: os.FileMode(420), modTime: time.Unix(1605896240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xcf\x52\xdb\x3c\x10\xbf\xe7\x29\x16\xf1\x1d\x3f\x47\xf7\x8e\xec\x03\xd0\x1e\x3a\x2d\xa5\x40\x0f\x3d\xae\xad\xc5\x56\x2a\x24\x57\x5a\x87\x66\x32\x7e\xf7\x8e\x22\x9b\x98\x00\x87\xe6\x22\x65\xb5\xfb\xfb\xb3\xd2\x5a\x9d\x5d\x7d\xbb\xbc\xff\x79\xf3\x11\x3a\x7e\xb4\xd5\x4a\xe5\x65\xa5\x3a\x42\x5d\xad\x00\x00\x14\x1b\xb6\x54\xdd\x3f\x19\x6e\x3a\xa8\x3d\x43\x01\xfb\x3d\xac\xbf\xa2\x71\x97\x1d\x3a\x47\x16\xc6\x51\xc9\x9c\x96\x4b\xce\x8a\x02\x2e\xef\xee\xa0\x28\xa6\x80\x35\xee\x17\x04\xb2\xa5\x88\xbc\xb3\x14\x3b\x22\x16\xd0\x05\x7a\x28\x45\x02\xbb\xc0\x48\x3f\x6e\xbf\xc0\x38\xca\xc8\xc8\xa6\x91\x4d\x8c\xb2\xf6\x9e\x23\x07\xec\xd7\x8f\xc6\xad\x9b\x18\x45\xb5\x3a\x12\x6c\xbe\x0f\x14\x76\x80\x4e\xc3\xe7\x3b\xa8\x07\xa7\x2d\xc1\x93\x84\x1b\xdf\xf7\x14\xd6\x9b\x78\xa4\x8f\x4d\x30\x3d\x43\x0c\xcd\x7b\x74\x9b\x28\x37\xbf\x13\xde\x81\x6a\x13\x45\xa5\x64\xae\xfa\x17\x88\xa3\xe0\x2c\xe7\x0d\x30\x25\x73\x6b\x57\xaa\xf6\x7a\x37\x81\x3b\xdc\x42\x63\x31\xc6\x52\x38\xdc\xd6\x18\x20\x2f\x05\xfd\xe9\xd1\xe9\xc2\xb6\x73\xc0\x9a\xb6\x63\xa8\xdb\xbc\x11\xb9\xfe\x80\x81\x2f\x11\x8a\x3a\xa0\xd3\x73\x8f\xcf\x45\xf5\xc9\xbb\x86\x8d\x77\x11\xf4\x00\x17\x9e\x95\xc4\x45\x75\x3d\x30\x7b\x77\x02\xc1\xbe\x6d\x2d\x05\x01\xbc\xeb\xa9\x14\x39\x47\x80\x46\xc6\xe9\xac\x14\x8d\xb7\x16\xfb\x48\x73\x18\x43\x4b\x5c\x8a\xf3\x0c\x71\x8d\x5b\x01\x18\x0c\x16\x8d\x77\x1c\xbc\x7d\x06\x3f\x9e\x64\x8f\xa4\x4b\xf1\x80\x36\x01\x1d\xa2\x16\xeb\xf4\x5c\xee\x0f\x34\xc9\xbd\x69\x31\xc9\x5f\x58\xce\xf7\xd2\xe3\x3b\xb2\x0b\xd3\xa4\x74\x25\x53\xca\xc2\xaa\xcc\x3e\x16\x11\x6d\x9e\xdb\x3f\xdb\x99\xfb\x7d\xb4\x67\xf4\x52\xfa\x89\x8a\xc1\x9e\x68\x48\x37\x3a\x6d\xc3\xc9\x4d\xcd\xbf\xfd\x1e\x02\xba\x96\xe0\xbf\xe0\x07\xa6\xff\xa7\xf5\x8a\x18\x8d\x8d\xf0\xa1\x84\xf5\xf5\xb3\x6d\x18\xc7\x57\x00\xca\x9a\x05\x6b\x61\x98\x1e\x45\xf5\xe2\x1d\x14\x69\xee\x16\x73\xf6\x82\x61\x7d\x9b\xfe\xc0\x38\x8a\xea\xd5\xd1\x15\xc5\xe6\x30\xd5\x58\x29\x69\xcd\x9b\xe2\xc9\xe9\x53\x55\x4a\x0e\x76\xd9\x69\x6d\xb6\xd3\x0b\x97\x0e\xa7\xed\x7e\x0f\x3b\x43\xf6\x50\xab\x64\x1e\x82\x34\x15\xe9\xc3\xf3\x37\x00\x00\xff\xff\x1d\xe8\xd8\xc2\x8f\x04\x00\x00")

func layoutsLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsLayoutHtml,
		"layouts/layout.html",
	)
}

func layoutsLayoutHtml() (*asset, error) {
	bytes, err := layoutsLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/layout.html", size: 1167, mode: os.FileMode(420), modTime: time.Unix(1605896240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _login_formHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x4d\x6e\xe3\x30\x0c\x85\xf7\x39\x05\xc1\xcd\xac\x62\x5f\xc0\xf6\x66\x80\xd9\xcd\x60\xd0\xb4\x07\x90\x2d\x26\x11\x2a\x8b\x2e\x45\xa5\x0d\x82\xdc\xbd\xb0\xe5\xfc\x16\x45\x16\x5d\xf2\xe9\xbd\x27\x7e\x00\x2b\xeb\x76\xd0\x79\x13\x63\x8d\x6b\x96\x7e\xd9\xb2\x2a\xf7\xd8\x2c\x00\x00\xaa\x51\x02\x61\x4f\xf9\x15\xc1\x74\xea\x38\xd4\x88\xd0\x93\x6e\xd9\xd6\x38\x70\x54\x3c\x55\x78\xde\xb8\xb0\x9c\xac\xb9\x61\x6a\xb9\xff\x63\x23\x9c\x86\x2b\xc3\x64\xf2\xa6\x25\x0f\x6b\x96\xd9\x94\x22\x49\x30\x3d\x61\xf3\x8f\x7b\xb0\xbf\x92\x3a\xef\xa2\x51\x4a\x52\x95\x93\xf9\xae\xc0\x85\x21\x29\x8c\x91\x1a\xcf\x61\x18\xbc\xe9\x68\xcb\xde\x92\xd4\xf8\x32\xcb\x45\x51\xe0\xcd\x46\x27\x3f\x4c\x53\xc7\x41\x85\x3d\x82\xb3\xf7\xcb\x80\xee\x07\xaa\x51\xe9\x43\x11\x84\xde\x92\x13\xb2\x57\xa8\xa5\x75\xbb\x9f\x92\x0f\x26\xc6\x77\x16\x8b\xcd\x5f\x56\xb0\x04\xa3\x40\x8f\xa1\xcf\xb9\x5b\xe8\xff\xb3\xfc\x05\xfa\xe4\xff\x0e\xfa\xd2\x97\xa1\x2f\xf3\x63\xf0\x36\xa9\x72\x98\x83\x31\xb5\xbd\xbb\x1c\x49\xab\x01\x5a\x0d\xcb\x41\x5c\x6f\x64\x8f\xcd\x8a\xa0\xe3\x10\xa8\x53\x92\xaa\xcc\xd1\xab\xae\xcc\x98\xab\xb6\xce\x5a\x0a\x38\x13\x77\x51\xd6\x85\xf2\xeb\xa8\xec\x8c\x4f\x54\xe3\xe1\x00\xc5\xef\xd5\xd3\x9f\xe7\x51\x85\xe3\x11\xa1\x9c\x6f\xb9\x1c\xa1\x9a\x45\xde\xf4\x33\x00\x00\xff\xff\x3e\xd6\x11\x49\xf8\x02\x00\x00")

func login_formHtmlBytes() ([]byte, error) {
	return bindataRead(
		_login_formHtml,
		"login_form.html",
	)
}

func login_formHtml() (*asset, error) {
	bytes, err := login_formHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "login_form.html", size: 760, mode: os.FileMode(420), modTime: time.Unix(1605896240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _statsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x6a\xc4\x30\x10\x44\xfb\xfb\x8a\xc5\x5c\x99\x58\x5c\x1b\xd6\x6a\x92\x36\xd5\x91\x0f\x58\x5b\x4b\x64\x30\xd2\x65\x77\x5d\x04\xe1\x7f\x0f\x67\x39\x71\x12\x8e\x6b\x84\x66\x79\xcc\x0c\x83\xf1\xe4\xcf\x46\x36\xaa\x8d\x1f\x33\x2b\x5c\xf2\x2c\x50\x0a\xb4\xaf\x34\xa6\xe7\x48\x29\xf1\x04\xcb\x82\x2e\x9e\xfc\x01\x8d\xfa\x89\x61\x98\x48\xb5\x6b\xaa\x58\xdf\xc7\x3e\x4b\x60\xe1\xd0\xf8\x03\x00\x00\x5a\x64\x0a\xf5\x5f\xb5\xec\x62\x03\x40\x87\x7c\xe1\xae\x19\xf2\xd4\xf8\x37\x65\x41\x67\xf1\x3e\x75\x1e\xb2\xf0\x5f\x0c\xdd\xb7\xf5\xf5\xfe\x13\x8a\xd6\xe7\xf0\xb9\x63\xa5\x08\xa5\x77\x86\xe3\xac\x2c\x0f\x70\xd4\xab\x13\x3c\x75\xd0\xbe\x90\x51\xfb\x7b\x82\x65\xb9\x57\x3b\xf8\x52\xaa\xcb\xba\x8a\x85\xdb\x40\xf5\xff\x4f\xec\x65\x6b\x27\x4e\x61\x4b\x43\xb7\x15\x46\xb7\x0e\xea\xbf\x02\x00\x00\xff\xff\xbf\xf4\xba\x07\x99\x01\x00\x00")

func statsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_statsHtml,
		"stats.html",
	)
}

func statsHtml() (*asset, error) {
	bytes, err := statsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stats.html", size: 409, mode: os.FileMode(420), modTime: time.Unix(1605896240, 0)}
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
	"aggregator.html":     aggregatorHtml,
	"home.html":           homeHtml,
	"layouts/layout.html": layoutsLayoutHtml,
	"login_form.html":     login_formHtml,
	"stats.html":          statsHtml,
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
	"aggregator.html": &bintree{aggregatorHtml, map[string]*bintree{}},
	"home.html":       &bintree{homeHtml, map[string]*bintree{}},
	"layouts": &bintree{nil, map[string]*bintree{
		"layout.html": &bintree{layoutsLayoutHtml, map[string]*bintree{}},
	}},
	"login_form.html": &bintree{login_formHtml, map[string]*bintree{}},
	"stats.html":      &bintree{statsHtml, map[string]*bintree{}},
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
