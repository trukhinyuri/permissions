// Code generated by fileb0x at "2018-06-22 12:08:58.527305118 +0300 MSK m=+0.022796994" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(871ec00059aa54727bbfbf7e689d5341.f7d60b70267ef579d9aad25918a6b087)

package static

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path"

	"context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileSwaggerJSON is "swagger.json"
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x7d\x5b\x8f\xdb\x38\xd2\xe8\x7b\x7e\x05\xe1\x73\x80\x24\xd8\xb8\x7b\xce\xee\x62\x81\x33\x6f\x3d\xe9\x4c\xc6\x98\xf4\x8c\xd1\x97\xf9\x16\xd8\x1e\x18\xb4\x54\xb6\xb9\x91\x48\x85\xa4\xfa\x92\x41\xfe\xfb\x07\x92\xba\x5b\x37\x4b\xa2\x5b\xdd\xd1\x53\x3a\x32\xef\x55\xac\x2a\xd6\xf5\xaf\x57\x08\xcd\x1c\x46\x45\xe8\x83\x98\xfd\x88\xfe\xf3\x0a\x21\x84\x66\x38\x08\x3c\xe2\x60\x49\x18\x3d\xfd\xaf\x60\x74\xf6\x0a\xa1\x3f\xdf\xa9\xb6\x01\x67\x6e\xe8\xb4\x6b\x2b\xee\xf1\x76\x0b\x7c\xf6\x23\x9a\xfd\xfd\xe4\x87\x99\xfe\x46\xe8\x86\xcd\x7e\x44\x7f\x99\xbe\x2e\x08\x87\x93\x40\xf5\x55\xad\x96\xc0\x7d\x22\x04\x61\x54\x20\x01\xfc\x8e\x38\x80\x84\x64\x1c\x04\x0a\x05\x70\xc4\x41\xb0\x90\x3b\x20\x10\xa6\x2e\xc2\x9e\xc7\xee\x05\x92\x0c\xf9\x98\xe2\x2d\x20\xec\x38\x20\xf4\x87\xa4\xa1\x9e\x14\xa1\x99\x24\xd2\x83\xe2\x14\x67\xcb\x85\x5a\xec\x37\xb3\x31\x2c\x77\x22\x5d\xd9\xa9\x19\x0c\xd2\x4f\x08\xcd\xb6\x20\x33\xff\x55\xc3\xe2\x6d\x7a\x12\xd1\xb7\xcc\x0c\xb3\xe4\xfb\x9f\xef\xd2\x4e\x22\xf4\x7d\xcc\x1f\xd5\x6a\x2e\x41\x86\x9c\x46\xbb\x8b\x67\x3c\x99\x65\x1a\xb3\x00\xb8\x3e\xdb\x85\xab\x3a\x7c\x04\x79\x23\x80\x9f\xc5\x8b\xcb\xb4\x0c\x30\xc7\x3e\x48\xe0\xc5\x15\xfd\x95\xf9\x5b\xad\xf9\x31\xd0\x27\x21\x24\x27\x74\x9b\x19\x41\xff\xba\x61\xdc\xc7\x6a\x97\xb3\x30\x24\x6e\xf1\x57\x8a\x7d\xdd\xf7\xdf\x73\xb5\x8a\xf9\xe2\xbc\xd8\x80\x68\x38\xee\x00\xbb\xc0\x8b\xbf\x71\xf8\x12\x12\x0e\x6a\x1f\x92\x87\x90\xf9\xf1\xdb\xbb\xea\xe5\x02\x0d\xfd\xc2\x86\xf4\x77\x75\x64\x85\x19\x14\x32\xba\x3e\xa1\xb3\xdc\xd7\x3f\xdf\x1d\xb2\xff\xc2\x0e\x2f\x99\x07\xf6\xf7\xd8\x07\x24\x85\x2b\x74\xbd\x03\xb4\x38\x47\x6c\x83\xe4\x0e\xe2\x7f\x34\x7a\x31\x8a\xd6\xb0\xc3\xde\x46\x7d\xbd\xdf\x11\x67\xa7\x7f\xd3\x27\x46\x84\xe4\x58\x32\x8e\x1c\x4c\x51\x00\x5c\x4d\x69\x7e\x75\xf4\xc0\x15\x87\xa4\xc6\x9d\xef\x2f\xc9\x1c\xd0\x97\x10\xf8\x63\x16\x14\xdf\x4a\x6f\x03\x07\x11\x30\x9a\xbf\x68\xfa\x87\xbf\xff\xf0\x43\xe1\xd3\xfe\x6e\x73\xf7\x26\x77\xf1\xd1\x9b\xeb\xdf\xcf\x7f\x47\xc2\xd9\x81\x8f\xdf\xce\xaa\x00\x31\x73\x61\x83\x43\x4f\x36\xce\xe4\xec\x80\xf3\x47\x04\x9c\xb3\x3d\xa0\x9b\x49\xf6\x86\x40\x68\xf6\x7f\x39\x6c\x54\xef\xff\x73\xea\xc2\x86\x50\xa2\x46\x13\xa7\xc0\xf9\xef\x67\x8b\x8f\x50\x40\xd4\x6f\xa5\x67\x15\xff\x65\xfe\x8d\x16\x3f\x3b\xd5\x60\x2b\x25\x52\x41\x38\x30\x91\x3a\x13\x82\x6c\x69\x4c\x5d\x3d\xb8\x03\x0f\x6d\x18\x57\xf4\xb7\x40\x96\x4f\xd0\x8d\x00\x57\xff\xb8\x26\x9e\x47\xe8\x16\x05\x21\x0f\x58\x03\x51\xbb\x02\x79\x19\x0f\x31\x51\xb6\x89\xb2\x3d\x3f\xca\x56\x73\x02\xf1\x88\x6b\xe6\x3e\x96\x0f\x57\xf6\x4b\xe1\xb4\x7b\x12\x9c\xb2\x0b\x76\x09\x5f\x42\x10\xb2\x0d\x09\x1a\x8c\x5c\x47\x24\x44\x80\x7c\xa9\x04\x59\x41\x5b\x04\xd8\xe9\x20\x37\xfe\x96\x76\x6d\xa2\xc8\x1f\x41\x6a\xea\x9b\xce\x86\xde\xe8\xf9\x11\xa3\xde\xe3\xdb\x06\x01\xf2\xcc\xf3\x32\x73\x4d\x74\x76\xa2\xb3\x2f\x81\xce\xd6\x9f\x40\x61\x8f\x67\x8a\x06\xa9\x5d\x6c\x88\xa7\xf0\x1d\x09\x50\xc8\x2f\xc1\x45\xf7\x44\xee\x90\xc3\x7c\x1f\x9f\xa0\x2b\x00\x74\x3b\x73\x31\xbb\x9d\xa1\x00\x3b\x9f\xd5\x23\x53\x09\x37\x3e\xe3\x80\xd4\x4b\x56\x1d\x2b\x61\xf4\xa4\x6a\x7b\x66\xf8\x01\x76\xa7\x4e\xd8\xd7\x68\xfc\x43\x05\x32\x12\x2a\x61\xbb\x3f\x57\xbc\x92\x00\x6f\x61\x14\xeb\x00\xbe\x3a\x64\x2d\xc3\x33\xa1\x0c\xd1\x8c\xc7\xe9\xcb\x3f\xb6\x05\xa2\xfa\xfb\xaf\x3f\x29\x9e\x5e\xcd\x4c\x9e\x29\xc7\x7b\x97\xbc\x31\x98\x18\x96\xa3\xbd\xe7\x80\x25\xa4\x0c\x4d\x5f\x43\x16\xca\xe4\x1d\xd1\x96\xbf\x9d\xa9\x66\x66\xb4\x64\xfa\x89\xc7\x4d\x3c\x6e\x7a\x4b\x0c\xf4\x96\x48\x6e\x55\xe6\xa6\x0d\xff\x96\xf8\x7f\xed\xc9\x38\x72\xf4\x1a\xdc\x97\xff\xa4\x38\xfd\x8b\xb8\xdf\x0e\x56\xf5\x1c\x40\x85\x2f\x41\x90\xaf\x83\x51\x61\x33\xda\x44\x85\x27\x2a\x3c\x51\x61\xbb\x54\xd8\xdc\xb4\x66\x2a\x7c\x1c\xb0\xc6\xaa\x26\xb4\x7f\xe5\xe2\xa3\xab\x82\x43\x80\xe5\xee\x00\x34\xb5\xf8\x38\x50\x6f\x03\xf2\xf5\x25\x72\x95\xe3\xa9\xa8\xf4\x55\x4e\xa7\x6b\x61\xd7\x9c\xf4\x52\x13\xb7\x98\xf4\x52\x63\xd3\x4b\x3d\x17\x1d\x4c\x9e\x82\x4c\x4a\x98\xde\x4a\x98\x50\x28\xa1\x3f\x12\xfe\xeb\xc8\xf7\xa4\x70\x99\x88\xf7\x24\xea\x5b\x16\xf5\x27\x5d\x8b\x1d\x3a\xea\x82\x07\x12\x06\xa5\xa4\xe7\x7a\xc8\xd4\x43\xa6\x9d\x10\x6c\x7a\x9d\x79\xde\x24\x0a\x4f\xd4\x74\x72\xf2\x2b\x11\x16\xcd\x55\x7d\xd1\x8f\xf2\x3d\x25\xaf\x8d\x97\x79\x32\x5d\xc3\x9b\x7c\x92\xe8\x26\x1a\xf4\x3d\x3d\xc7\xbf\x4f\x3d\xe7\x20\xef\xef\x94\x54\xbc\xe0\x17\xb7\x65\x7b\x5b\x1d\x35\x9e\xac\x69\x13\x41\x9e\x9e\xd8\x96\x9f\xd8\x93\x21\x6d\x32\xa4\x3d\xad\xca\xa0\x15\x2f\x30\x6d\x27\x5e\x30\xf1\x82\x49\x38\x7f\xd1\xb4\xf3\x7b\xd1\x77\x1c\x27\xda\x3a\xa7\xf9\x30\x06\x55\x7d\x27\x82\x74\x8c\xb6\xfa\x90\xff\x21\x72\x97\x9d\x7a\xa2\xc0\x13\x05\x9e\x28\xf0\xa4\x1e\x99\xd4\x23\x83\x12\x6c\x8e\x69\x96\x64\xa7\x84\x1a\x49\xa6\x6f\x53\x43\x58\x79\xea\x26\xa9\xf9\xcb\x44\xa6\x27\x32\x3d\x29\x4d\x86\x0f\x2a\x37\xb7\x6b\x52\x9b\xa0\x29\x42\xbe\x9f\xbe\xe4\x10\xe6\x50\x54\x98\x1c\xc8\x1d\x0a\x3a\x94\x89\x41\x4c\x0c\x62\x62\x10\x83\x32\x08\x73\xc3\x26\x1e\x61\x89\x47\x7c\x37\xba\x21\x0e\x11\x48\x2c\xc6\x3d\xaa\x19\xda\xda\x61\x55\xb3\x49\xf7\x3e\x71\x8c\x89\x63\x58\xb4\xc3\xaa\xe5\x4c\xcc\x62\x58\x4d\x96\xfa\xfb\x05\xf2\x8b\x80\xb3\xff\x82\x23\xf3\x49\x10\x5b\xc5\xc6\x2c\xe3\x9e\x2d\x23\x63\xa2\x99\x9a\x23\x61\xa2\x81\x27\xe6\x30\x31\x87\x89\x39\x0c\xc4\x1c\xa2\x3b\xf5\x74\x51\x30\xd1\xed\x7f\xb9\xf9\x46\x62\x42\x7a\xfa\x57\xf4\xd7\xb7\xd3\x2d\x67\x61\xd0\xc1\x2a\xdb\x96\xb0\x7e\x04\x19\x53\x55\x64\xe6\x6a\xb0\xc0\x46\x23\x7f\x34\xeb\x9a\xe8\xeb\x44\x5f\x27\xb3\x6b\x3e\x87\x7e\x74\x9b\xaa\x65\xd5\x60\x4f\x38\x19\x93\xc0\x9a\xa7\x06\x03\x44\x85\xe7\x28\xc6\x14\x14\x7e\x18\x7d\x3e\x73\x5d\x03\x09\x24\x59\x42\xa8\xdb\x67\x81\x72\xf5\xa9\x5f\xb3\x49\x1e\x9e\xe8\xf5\x44\xaf\x47\x40\xaf\x9f\x87\x8c\x1f\x13\x8e\xe1\xa5\xfc\xbf\x37\xd2\x72\x43\xed\xb0\xeb\x82\x9b\xa1\x79\xdf\x91\xb0\x7f\xfa\x97\xfe\xf7\xf0\x44\x83\xed\xb5\x29\x3b\x4c\xd3\x72\x3d\x6c\x13\x71\x18\x1f\xfc\x35\xf0\x2c\xa3\x69\x17\x36\x7f\x05\x86\xb9\x5f\xe8\xfe\x93\x21\x77\xe2\x34\x13\xa7\x19\x35\xa7\x19\x70\x4f\xfa\xe2\xd7\xec\x48\x53\x96\xef\x81\x73\x96\x11\xc1\xa7\x60\x9f\xdf\xbb\x03\x53\x5b\x1e\x18\x79\x2f\x19\xce\x97\x09\x40\x40\x1b\xce\xfc\x03\x19\xa0\x19\x4b\x83\xff\x67\xce\xfc\xe9\xb1\x35\xb1\xc0\x89\x05\x4e\x2c\x70\x70\x16\x68\xeb\xb1\xf5\x62\x3d\x99\x4a\xde\x58\xe6\x91\x63\xd3\x56\x7d\xe6\xba\x79\x86\xc2\xb8\xb9\xe4\x07\x3f\xac\xce\x5c\xd7\x88\x13\x93\x02\x6f\xe2\x29\x13\x4f\x99\x14\x78\xcd\x54\x73\x9f\x64\x3c\xc5\x23\x24\xd2\x24\x69\x25\xde\x77\xc1\x55\xa6\x70\xea\x89\xcd\x4c\x6c\x66\x62\x33\xcf\xec\xe9\xf2\x32\x1c\x6b\x23\x95\x57\x3c\xda\x4b\xad\x3c\xfe\x2a\xda\xc0\x2c\x33\x60\x32\xf3\xcc\x68\x1c\x3f\xc1\x1d\x78\x59\xa6\x53\x81\x2c\xfb\xe4\x6f\x46\x19\xcd\x05\x53\x70\xc0\x6e\xf1\xff\x91\x16\x2e\xf3\xf5\x9e\x93\xfc\x07\x76\x4f\x81\xc7\xbb\xfb\x33\xc7\x38\xcf\x1c\x87\x85\x34\xef\x27\x5c\xac\xbf\x10\x35\x41\xf3\x39\xf2\x88\xd0\xa5\x18\xd6\x2c\xa4\xae\x02\xb2\xe9\xfc\xae\xb8\x35\xb6\xce\xdd\x3a\x75\x0d\x03\xe0\x92\x14\xd0\x6a\x86\xf7\xa7\xaf\x1b\x27\xa2\xfb\xae\x3e\x68\xec\x2d\xcb\x47\x2d\x3b\xe4\x4a\x8c\x7b\x98\x6f\xd9\x3c\xbe\x58\xc9\x69\xd4\x68\x3d\x75\x87\xa8\xc0\x84\x7e\xae\x13\x79\xe2\x30\x2a\x31\xa1\xc0\x43\xff\x84\x82\x3c\x75\x76\xa7\x19\x86\x7f\x7a\x07\xd4\x65\xfc\xb4\xa2\xa5\xa6\xac\x3e\xa6\x78\x0b\xfc\x34\xf8\xbc\x3d\xf5\x99\x0b\x5e\xb4\x86\x04\x4e\x95\x32\x64\x0d\xe0\xaa\xfa\xa0\x68\x15\x02\xa5\x82\x84\xa9\x5e\xef\xba\x84\x6e\x8b\x2f\xe3\x5e\xe0\x05\xb1\x07\xdc\xf2\x6b\xa8\x78\xb0\x56\xdb\x44\xc6\xca\x57\x25\xf0\xd2\x5c\xa8\x10\xf8\xd4\x44\x7f\x33\xd4\x17\x7c\x4c\xbc\x59\x35\xf8\x6f\xe2\xd1\x87\x05\x7f\x02\xd3\x22\x48\x7d\x42\xff\x60\x5e\xe8\x17\x52\xec\xd5\x41\xb4\xbc\x4b\x25\x40\x9d\x1d\xa6\x5b\x05\xd2\x3b\xdd\x09\xe9\xd4\x9a\x58\x18\x49\xa0\x3b\x5c\x1d\x1c\x60\x87\xc8\xc7\x2a\x30\x94\xd5\xbb\xcd\xc0\x81\x50\xf9\xaf\x7f\xd6\xc0\xe1\x7d\x3c\xfc\x31\xaf\xa1\x39\xa1\xe6\x8b\x78\x87\x89\x87\xd7\x1e\x5c\x31\x2f\x8c\xe0\x52\x09\xae\x62\x5b\x45\x42\x45\xfc\xb7\x91\xcb\xf4\x5d\x53\xa2\x18\x0f\x69\x9f\x8b\x26\xc9\x5d\xe5\xad\x58\x33\xe6\x01\xa6\xb3\x3a\xc2\xa7\xfb\x97\xde\x39\xe2\x1e\x72\xdb\x72\xc3\x2e\xce\x2b\x86\xf4\xf1\x16\x2a\x89\x3e\xe6\x1c\xe7\x9f\xe4\x33\x22\xc1\x1f\x8a\xc8\x2f\xcc\xe4\xa5\x0b\xf3\x88\x4f\x64\x4b\x82\x15\x03\xf5\x93\xe9\x53\x3a\xde\xa1\xb4\x2a\xb7\xd0\xdf\xf2\xa4\x28\x4b\x06\xb9\xd7\x79\xd4\x9b\xcb\x4f\xd6\xee\xd5\x2e\x5c\x9f\x38\xcc\x3f\xcd\x74\x38\xfd\x1c\xae\x61\xee\x78\x04\xa8\xac\xa4\x84\xc5\x7b\x22\x3e\x91\x7a\x3a\x58\xda\x21\x2b\xa0\xe0\xb8\x45\x72\xdf\x7a\x48\x29\xe9\x10\xbd\x31\xb6\x42\x43\xb4\x47\x54\xda\xa1\x72\xb2\xfd\xb1\x41\xf4\x27\x25\x1f\x46\xd2\x54\x1c\xf2\xde\xc4\xdc\x2a\xfb\x28\xb8\xf2\xe8\x4f\xc9\x10\x07\x9f\xdd\x41\x5e\x04\x6d\x86\x6d\xe6\x1d\xf4\x9f\xdc\x6b\x47\x3f\xb6\x12\xe1\xb8\x11\x17\x92\x1e\x5d\xef\xdf\x65\x61\xca\xf1\xc8\x98\x59\x00\x88\x96\x80\xd2\xaf\x82\xfc\x6b\xc0\xdc\xc1\xd8\xc2\xd2\x0c\x99\x46\x79\xbe\x82\xcc\x3f\xf5\x71\x19\x8f\xbd\x6b\xcc\xc9\x66\xd3\x8c\xdd\x25\xad\x4b\xa5\x36\x0a\xe0\x82\xab\x8f\x4f\xea\xc6\x89\x0c\xd7\x9d\x7e\x99\x81\x56\x3d\x18\xb9\x59\x77\x8e\x9d\xdb\xa7\x33\x6b\xe2\x79\x73\x78\x90\x4a\x28\xf7\xca\x41\xc0\xe8\x86\x6c\x7f\x26\x1e\x94\xbd\xaf\x5b\x4b\xb3\x7a\x94\xd5\x26\x3f\xcc\x50\xbc\x3a\x1a\xa2\xe3\xb9\xab\x86\x23\xa3\xed\xe6\xd0\x2f\x70\x50\x87\xed\x71\x1b\x45\x1e\xf4\x18\xe6\x39\xa2\x3f\x23\x1f\x07\x5d\x09\xb6\x3e\x98\xcc\xe9\xba\x58\xe2\xf6\xc4\x3b\x8a\x52\x5c\xe1\xa2\x86\x69\x4f\xbf\xa4\x1a\x2a\x21\xdd\xc5\x12\x10\xa1\xe8\xf2\xe7\xf7\xff\xf8\xc7\x3f\xfe\x3f\x8a\x1e\x31\xef\x3a\x41\xd3\x84\x69\xba\x67\xb2\x1c\x53\xf4\x5e\x5a\x89\x9e\xc9\xf1\x9e\x67\xb6\x5f\x18\xcc\xf8\x0f\x34\x6f\xd5\x34\x1c\x76\xa3\x86\x7d\x57\x6e\xd4\x92\x54\x6c\xd4\x5c\x5d\xc7\xfd\x3d\xab\x24\x1b\xdd\x5d\x3b\xcf\xe3\x46\xe5\x7d\x53\xed\xaa\xee\x1c\xd2\x08\x36\x76\x7e\xdc\xf7\xa4\x9a\xde\x0e\xf9\x86\x55\x67\x65\xa4\x98\x1e\x1a\x12\x3d\x92\x1a\xc8\xd6\x9b\x21\x25\xc3\xed\xde\x0a\xe9\xbe\x47\x88\xe4\xa6\x79\x3d\xd4\x4c\x9b\x3d\x80\x45\x9f\x09\x45\x2e\x04\x1e\x7b\xf4\xa1\xfb\x73\x40\xeb\x25\x66\xef\xaa\xd9\x8d\x97\x7b\xeb\xb7\x60\x38\xcc\xf7\x31\x75\x9f\x4a\xd5\xf1\x3e\x9e\xbe\x94\x5e\x46\x72\x8f\x65\x24\x35\xf0\x31\x4a\xcb\x3e\xa8\x9a\x59\x38\xd0\x3b\x5b\x0b\xfe\x40\xef\x5a\x2e\x32\xd7\x72\x4f\xb9\xd5\x5d\x65\xa6\x7b\xf7\x56\x4e\x95\xbc\x32\xed\x33\xe0\x80\x71\x69\x1f\x95\x96\x8c\xcb\x96\x30\x5a\xea\x05\x95\xae\xd5\xa8\x7b\x57\x7e\xad\xf9\xe9\x69\xd0\xdf\x34\xbe\xb0\x6b\x89\xea\x4b\xab\x35\x14\x5a\xd0\x6b\xd5\x2e\x4f\xb3\x15\x96\x28\x72\x9d\xcc\x3c\xd4\x5b\x40\x0d\x9c\xfb\x3f\x67\x92\x39\xcc\x6b\x4f\xae\x2d\xde\x0b\x4b\xa6\x92\xfc\x65\xf8\x56\xb6\xfb\x56\xf4\x62\x59\x38\xab\xd1\xe1\x5b\x74\x83\x5a\x60\x9c\x69\xa9\x70\x2e\x32\x7a\xbd\xc9\x49\x76\x6f\x91\xbe\xf3\xe0\x0e\x82\x83\x7a\xac\x95\x76\xbc\x68\x8d\x65\x6a\x9b\x9d\xb1\xec\x42\x75\x2e\x05\x78\x66\x29\xdd\x07\x0f\xa9\x5c\x66\x36\x73\x14\xb6\x21\xc2\x75\xbf\x65\x5f\x85\xeb\xc2\xa2\xc7\x81\xbb\xfa\xb9\xdf\x8c\xb8\x99\x66\x68\x3e\xef\xfe\xe6\xf0\xf0\x1a\xba\xdb\x85\x3e\xe9\xde\xe5\x2a\xac\xe7\xa9\x3f\x6c\x01\xa2\xba\x74\xbc\xd5\x10\x33\xbd\x6e\x44\x21\xa6\x10\x11\x81\x70\x62\xa9\x30\x50\xd3\x1c\x4f\xab\x58\x08\xdd\xc6\x5e\x51\xda\x88\x11\x79\x79\xb5\xd7\x93\x57\x81\xdd\xbe\x4f\xc4\x6f\x6d\x7c\x22\xf2\xba\x9f\xbd\xc3\x99\x59\xf0\xa0\x38\x4f\x1f\x7c\x75\xc0\x8a\x1b\xe5\xc5\x90\xf4\xb5\x28\xba\x12\xff\x64\xd5\xa2\xee\xcd\xc8\x21\xf0\x88\x83\x0f\x78\x35\xda\xb4\xe6\x67\xd6\x6c\x5b\xe8\x6d\xff\xda\x8b\x57\x54\xbe\xe2\x71\x2b\x6d\x27\x3d\x6b\x3f\x3d\x6b\xd9\x4d\xb1\x23\x2b\x5f\x16\x2e\x62\x41\x0a\x89\x6c\xe9\x25\x6c\xae\x00\xc9\xc4\x9b\x67\x71\x8e\xde\x30\xea\x3d\x22\xb2\xc9\x90\x13\xc5\x06\x02\xcc\xb5\x1b\x42\x3c\xe8\xdb\x8e\x00\x8e\xa7\xaa\x72\xa4\x11\x12\xcb\xb0\xa5\x4a\x20\x25\x84\x57\xa6\x57\x39\xaf\x67\x12\x7b\x2b\x27\x08\x1b\x0e\x41\xb7\x43\xef\x97\x37\x28\x14\x78\x0b\x68\xfd\x88\xb0\xe7\xa5\x02\xb5\x50\x48\x2e\x77\x44\x94\xa9\xe5\x0e\x81\x6a\xd8\x04\xd6\x6b\xb5\x92\xf7\xcb\x9b\xba\xfd\xf8\xe0\x33\xfe\xd8\x6a\x4b\x97\x67\x17\xe3\xd8\xd2\x85\x59\x73\xb9\xf6\x02\xb8\xc8\xfb\x9e\xd5\x00\xfe\x8f\xa8\xf1\xe8\x64\xaf\x02\x3e\xb6\xe1\xe1\xa6\xa9\xe2\xe4\x6a\x06\x4e\x41\x82\x40\xe6\x16\xa8\xfb\x76\x88\x0a\xb8\x92\xfb\xc6\xce\x38\x2b\xcb\xf4\x28\xf1\xfa\xa9\x27\x4c\x1c\xb0\xfb\xb8\xb2\x4e\x1b\xb1\xfb\xd8\xb4\x8e\x27\xa4\xce\x21\x3d\x1a\x58\x6e\xd2\xa9\x1a\xd6\x14\xb8\x5a\x30\xb1\xbd\x1e\x33\x4d\xc9\x5a\x46\x76\x8d\xff\xd8\x23\x4a\xd5\xf7\x38\x6a\x5b\x25\x92\xa3\x88\xc0\x21\x73\xc6\xdd\x2f\x73\x05\xa1\x6c\xcd\x80\x47\x4f\x3b\x9b\xcc\x9d\x85\x96\x95\x6f\xa0\x9e\xf6\xce\xec\x6b\xca\xd2\xbb\x22\xf3\xda\x6b\xf7\xb0\xc8\x6c\x7d\x74\xf0\x23\x1c\x1c\x69\x74\x3e\xf9\xe4\xdd\x35\x90\xac\xea\x53\xa9\x77\x30\x0f\xa2\xd4\xfb\x3f\x76\xfc\x47\x6f\xee\x89\xdc\xb1\x50\xa2\x35\xf1\x3c\x42\xb7\x6f\x9f\x53\x28\xc0\x3b\x5b\x1a\xaf\xd1\xc4\x18\x9c\x33\x1f\x93\x5a\x3a\xaa\x1b\xa8\xfb\x7c\x19\x07\xd8\xcd\xe7\xc8\xd5\x5f\xdb\xde\xe4\x72\x95\x86\x19\xe3\x00\x55\x85\xeb\x82\xbb\x5a\x3f\x76\x06\xc2\x99\x1a\xe0\xa7\xc7\x43\x9f\xfe\xed\x94\x5b\x8a\x77\xcc\x25\xf1\xa1\xc7\xfb\xbe\x08\x8b\x03\x5f\xe8\xb9\xf3\x1c\x8f\xaf\xef\x3e\x06\x59\x46\xb7\x4a\xc6\xa1\x47\x59\x79\x44\x48\x6b\x8c\xa3\x00\x83\x7a\xa6\x91\xec\x75\x7c\x40\xfb\x90\xf3\x7b\x28\x42\xe9\x03\xbd\xd3\xef\x22\x78\x9c\xdf\x61\x2f\x04\x14\x60\xc2\xd5\xa3\x08\xe8\x1d\xe1\x8c\x1a\xb9\x0a\x73\xa2\x24\xdb\xce\x3a\x4f\x3d\xf4\x9e\xba\xf3\xa9\x2d\xac\x66\x55\x9d\x85\x3c\xdd\x7b\x64\x22\xc2\x07\x5e\xe7\x0f\xf5\x81\x6b\x4f\x28\x21\x31\x75\x31\x77\x91\x00\x4e\xb0\x47\xbe\xea\x90\x97\xb3\xe5\xc2\x84\x46\xdf\xd2\x0b\x10\x5a\x97\x31\x9f\x23\x87\x51\xd5\x5c\x9a\x9f\x90\x6f\x7e\xf9\xf1\x96\xfe\x0d\xdd\xce\x08\xbd\xc3\x1e\x71\x51\x6c\x4c\xb8\x9d\x99\xef\x5f\x42\x26\x31\x82\x07\x47\xbb\xa5\xc7\x5f\x75\x5b\xa3\x6b\x35\xf3\xcc\x6e\xe9\xc9\xc9\x09\x48\xe7\xe4\xe4\xe4\x96\x2e\xce\xd5\x7c\x21\x25\x5f\x42\x88\x66\x23\x2e\x50\x49\x36\xc4\x31\xbd\x1c\xe6\xc2\x2d\x3d\x07\x89\x89\xa7\x5f\xf3\x2c\x30\x2e\x87\x5a\xdb\x02\x0f\x85\x45\x0a\xf4\x99\x50\x17\x9b\xc9\x37\x04\x3c\x17\xbd\x8e\xdf\x43\xaf\x91\x1f\x0a\x89\xd6\x80\x28\xa3\xf3\xaf\xc0\x19\xd2\xd8\x10\xaf\x95\x32\x89\x80\xb2\x70\xbb\x43\x92\x6c\x77\x52\x9b\x5f\x36\x00\x2e\xda\xb2\x60\x07\x3c\x6e\x97\x58\x64\x5e\x7f\x64\xee\x6b\xe4\x32\x10\xaf\x25\x82\x07\x22\xa4\x6a\xf2\xb3\x9a\x35\xbf\x54\x01\x5a\xc7\x97\xbf\x70\xa2\x8f\x0c\xad\x8f\xe3\x89\x7c\xc5\x22\x60\x94\xdf\x2f\x7d\xe6\x2d\xf5\x8c\xe6\xa4\xda\xc6\x12\x56\xb8\x61\x71\x5e\xa5\xf2\x8c\x50\xa2\xbb\x8d\x3b\xea\x5f\xa3\x4f\x5d\xed\xa4\x0c\x2c\x09\xb4\x46\x81\xf5\xcb\xf5\xf5\xf2\xa8\x04\xc7\x24\x4c\xd8\xa3\x30\x8b\xf3\x7a\x1a\x63\x6e\x32\x87\x80\x83\xd0\xaf\xc4\xdc\xa5\xce\x64\xa8\x38\x18\xd9\xd5\x85\x6e\x8d\x0b\xbf\xaa\xc6\xe5\x00\x3b\x00\xa3\xae\x8e\x6c\x7f\xae\x38\xf3\x5f\xf3\x3b\x2f\x39\x75\xd5\xa2\x70\xee\xea\xb8\x34\x2f\xcf\x65\xbc\xa8\x44\xca\x6a\x65\xf3\xd1\xb6\x79\xd5\x84\x5b\x57\x7b\xc8\x25\x80\xdf\x11\x27\x4e\x6a\x53\xb1\xd5\x62\xd6\x8c\xa3\xec\xe7\xe7\x22\xfd\x2b\xee\x27\x65\x0f\x99\xfd\xa4\x9e\xf4\x19\x26\x61\x48\x69\xcd\x06\x9f\x81\x3f\x7e\xd9\x11\x2d\xce\x1b\x34\x51\xa6\x81\x3a\x22\xe2\x56\xbd\x1a\xf2\x3c\xae\x84\xbf\x8d\x34\x16\x70\x51\xf0\xf9\x3d\x98\x1c\xfe\xd6\x2c\x17\x97\xd2\xbf\x6b\xbc\x6d\xdd\x6d\x14\x62\xed\x82\x6e\x79\x3e\x27\xc8\x1e\x9e\x98\x16\x79\x4d\x25\x89\xba\x0d\xe4\x2a\xca\x43\x0f\xc4\x8b\x89\x1b\x9b\x5c\x10\x86\x73\x41\xd0\x88\x61\x49\x11\x71\x19\x7a\x6d\x7d\xc0\x2f\xb3\x08\x3a\xb2\xbb\x0b\x4d\x36\x87\x5c\xbb\xd2\x7b\x0c\x7d\xd5\x46\xc9\x38\xb6\x60\x15\xd3\xa9\x96\x39\x45\x8a\xad\xc7\x01\xb0\x4f\x84\x7e\xae\x81\x93\xfa\xd9\x24\xcd\xa0\x9f\xd1\x1b\x05\x1e\x0e\x5b\x93\x1f\x90\x30\x7a\xaa\xdd\xd0\xcc\x9f\x20\x9d\xb7\x5d\x69\xaf\x1a\x7d\x08\x52\x7b\x34\x95\x2f\x3c\x04\x6a\x17\x56\x97\xf0\xc1\xcc\x51\xb5\x04\x22\x56\x7d\x9d\x00\x17\xa2\xce\x0d\xd0\xcb\x63\xc6\x81\x86\x93\x0c\x44\x0b\xcf\x31\xa0\xd2\xea\xb9\x5d\x01\x95\x55\x87\x56\x16\x69\x5f\x7e\xb9\xd5\x06\xac\x86\xd6\x77\x16\x24\x93\x95\xd5\xdf\x5a\xd5\x24\xb9\xb9\x7a\xdf\x76\xdf\x49\xbd\xf6\x23\x1a\x36\x23\xe2\x9d\xf4\x65\x0a\x5e\x61\xb2\x41\x19\x42\x1e\xe9\x6b\xb9\x81\xd9\xf4\xf8\x70\x8b\x6d\x09\xad\x7d\x95\x47\x2d\x34\x3c\xd4\x9f\x02\x61\x6a\x74\xc2\xad\x74\x3d\xcf\x35\x93\x8b\xde\x76\xb3\x01\x3e\xdb\x2c\x39\xa2\xc4\xee\xae\xf9\xe7\x1a\x0b\xe2\x98\x1f\xba\xf3\x4b\xd5\x3b\x9f\xed\x59\x88\x7b\xc6\xdd\xf6\x4c\xd4\x0c\xd1\x99\xbc\xeb\xde\xe5\x41\x65\xf1\x52\xba\x8e\xbd\x2c\xec\x65\x3c\xb7\xe3\x02\x0b\x41\xee\xe0\x86\x8a\x70\xad\xc0\xbe\x6e\x9d\xda\xa7\xa1\x67\x53\x9a\x1f\xec\x79\x51\xaa\x1f\x81\xc2\x64\x8c\x5e\x09\x7f\x62\x5b\xc6\x53\xd9\x12\x2e\x93\xf9\xc7\x95\x2f\x28\xc9\xad\x9e\x85\x66\x39\xb9\x4f\x52\x19\x67\x73\x03\xef\x8d\xa3\x13\x72\xb6\x75\xe0\xa9\xe9\x55\x9d\xc0\x33\x76\xe1\xc9\x27\x95\xcf\xf8\xee\xf4\x70\xdd\xd9\x77\xf2\x1e\xca\x6b\xa7\xca\x1b\xdb\x5a\x88\x9a\x8f\x1f\x56\xf0\x20\x57\x91\x0a\xd9\x96\x2f\xe6\x05\x7e\xf8\xf0\x20\xaf\xe2\x49\x2a\x97\x42\xe8\x31\x96\xb2\xa0\x2d\x96\x22\x39\xde\x6c\x88\x63\x6f\x15\xd7\xd1\x04\x15\xd6\xba\x32\xc7\xfb\xa1\x26\x2f\x7a\xc8\x5b\xcc\xd6\x9b\xbf\xbc\x6d\x13\xf6\xd6\xf4\x2a\xb9\xf2\xc8\x3c\xc1\x05\xf9\xfa\x72\x6e\xfc\x74\x31\xa7\x8b\x79\x9c\x8b\x79\x30\x1b\xee\xca\x81\x5f\x4a\x84\x76\x4d\x59\x8a\xa3\x44\x6f\x37\x01\xf4\x12\xd4\x02\x0e\x00\x68\xae\x43\x25\x40\xb9\x6a\x55\x01\xd0\xfa\x64\x3a\x26\x12\x23\xb5\x4a\x94\xac\xf8\x40\xb6\x70\x60\x16\xf7\x94\x25\x68\xbf\xac\x23\x67\x04\x1d\x3b\xbe\x98\x69\xdb\x1c\xbb\x69\x99\x35\xd5\x47\xac\x35\xce\xb9\xaa\x4e\x3d\x39\xec\x5b\xba\xd8\x20\xe3\x77\x7f\x45\xbe\x02\x0a\x38\xbb\x23\xea\xf5\x46\x19\x9d\x07\xc0\x05\x11\x52\xfb\x56\x1a\x2f\xfb\x7b\xe2\x79\x68\x0d\x28\xd2\x64\x9f\xf4\x62\xd7\x2b\x9d\xf4\xc9\x1e\xd3\xd6\xc9\xcb\xc7\xeb\x7c\x9d\x83\x5f\x47\x0f\xec\xcc\x18\x15\xea\x7e\xf3\x62\xb4\x2d\x17\x7c\x88\xe6\xa9\x97\x0c\x6c\x64\xdb\xa7\xc7\xd9\xe0\x82\xb6\xda\x60\x7f\xd3\x46\x6d\xb5\x02\xb1\x0a\xc2\xb5\x57\x2d\x57\x35\x0f\xbf\x34\xfd\x8f\xfc\x80\xd5\x22\x93\xd5\xeb\x6e\xa4\xb2\x9a\x1b\x1f\x70\x52\x9d\xd8\x9c\x86\xfe\xba\x7a\x6e\x97\x85\x6b\xaf\xee\xa6\x2f\xf5\xd8\xe5\xc2\x8b\x55\x41\xb8\x56\x0a\x8e\xa6\x5e\xd9\xdc\x79\xb4\x80\x9a\x03\x88\x72\xd1\x29\x51\xc0\xd2\x21\xa4\xdc\x6b\xac\x1a\xb9\x62\xb5\xc3\x16\x5c\xbc\xd0\x25\x0e\x83\x8b\x22\x54\x32\x71\x70\x5b\x90\xb1\x9a\xbe\x87\x7b\x76\x3a\x55\xcb\x94\x6a\x69\xfb\x72\xfe\x7a\xfc\xc7\xb8\x61\xe9\x2b\xcd\x94\x2d\xf3\xf4\x6b\x52\xe5\xbb\x63\xbc\x90\x2c\x2f\xc2\xb8\x25\x35\x2d\xc2\xed\xce\x24\xce\x8b\xf5\xb9\xbb\x33\xf1\xd6\x72\x75\x15\x83\xb7\xad\x55\x3d\x8a\xf4\x70\x80\x06\xe7\x58\xeb\x79\x06\x6a\x1c\xed\x08\xb7\x0a\x05\x70\x5b\x8f\x39\xed\x2c\x77\x23\xa0\x32\x20\x23\xb3\x84\x83\xed\x8e\xed\xf3\x90\x25\xcb\xa8\xb3\x4e\x9a\x7a\x75\xb6\x4e\x22\x2a\xa2\x57\x75\x0e\x1c\xfb\xb6\x52\x63\x9c\x5d\x8c\x44\xed\x53\x28\xc5\x67\xcd\xd9\xa2\x8c\x77\xd6\x67\xfa\x2d\xab\xd1\x7c\x14\xad\x43\x93\x33\x64\xbe\x61\xde\x1b\x32\x2d\x51\xdd\xd3\xf3\x25\x1d\xc8\x16\x44\x52\xd3\x69\x3b\x80\xa4\xdb\x1e\x9b\x53\xe4\xef\x67\xa1\xdc\x0d\xe7\xf2\xc1\x70\x28\x77\x3d\x5d\x3e\x92\xa2\x52\xef\x8a\xb5\x33\x57\x92\x7d\x86\x43\xe2\xf1\xb3\xbd\x3a\xc7\xe4\xeb\x41\xae\xb3\x33\x17\x53\x01\x95\x17\xc1\x2a\x47\x1d\x7d\xe0\xe3\x2d\x7b\x95\x5f\x5e\x8b\x3a\xbd\x4f\xb4\x4c\xaa\xc5\x59\x0d\x95\xe1\xb0\x57\x23\x4a\x4f\xec\x3d\x10\x45\xfb\xe1\x66\x11\x2b\x47\x82\x44\xcb\xcb\xe6\xec\x47\x49\x9b\x5c\x8c\x18\x5a\x72\xb8\x04\x0f\xb0\x00\x14\x8f\xd1\x99\x0f\x2c\xc4\x6f\xa1\xdf\xf4\xa8\x29\xbd\xd2\xd1\xd4\x35\xdd\xdb\x25\xb5\xab\x1b\xfc\x4a\xf2\x27\x0e\x56\x5a\x7b\x98\x6e\x4f\x05\xf8\x77\x71\x00\x48\x02\xc0\xc8\x31\xcd\xd4\x69\x6b\xbe\x5b\xa5\xed\xd5\x25\x8b\x5d\xe4\x8c\xd5\x06\xe2\xeb\xd6\x39\xc7\x6c\xc8\x39\xe8\xbc\xde\x91\xe3\x5c\x36\xb8\x06\xee\x57\x87\x3b\x07\xee\x8d\xd8\xf5\x2a\xbe\x37\x03\x95\xb8\xf4\x55\xad\xb1\x73\x20\x10\xdc\x8f\xd7\x71\x30\x5e\xd9\x25\x08\xc9\xf8\x01\xa8\x93\xef\x90\xc3\x1d\x6e\x7e\xea\x8b\x3c\xda\xcb\xbf\x37\xc2\x58\x8a\x15\xf8\x1e\x30\x23\x97\xd4\x7e\x1f\x0f\xe4\x4e\x87\xc3\x46\x01\x8f\x3a\x07\x7e\x47\x48\xe7\xfa\xea\x40\x0c\xad\xb7\x58\x15\x03\x1f\xe3\xef\x41\xa6\x6e\x44\x33\x06\xf4\xca\xce\x5f\x5d\x4f\x20\xb7\xc8\xce\xc9\xff\xcd\x20\x35\xe5\x05\xb2\x5b\xb6\x94\xd4\xc1\x4c\x51\x28\xc6\x31\x8a\xd7\x4e\xa9\xd6\x7a\x0f\x13\x93\x46\x59\xd1\x84\x50\x73\x02\xea\x33\x5e\xb3\x50\x9a\x48\x80\x74\x79\xd9\x44\xf6\x7d\x6a\x97\x3b\xf9\x80\xe0\x9a\xa7\xc4\x8d\x00\xfe\x91\xb3\x30\x30\x4f\x94\x72\x80\x47\x6f\x20\x0f\xee\xc0\x5b\x19\x2e\x6c\x59\xdb\x6c\x56\xf3\x49\x4d\x18\xd5\x79\xad\x54\x3c\x8f\x42\x05\xbf\x55\x47\xd8\xa7\x9e\x83\x86\x41\x95\x66\xa8\x7d\x6a\x8f\xf8\xd5\x55\x88\xc5\x2a\x30\x89\x2c\x3c\x07\x44\x13\x85\xc7\xab\x23\x6b\xea\xe3\xdb\x62\x6b\xde\xf8\x40\xeb\x94\x76\xb6\xe6\xae\xd3\x11\xdb\xd6\x0e\x97\x29\x86\x2d\x2a\xff\x96\x89\xd3\x8f\x31\xac\x1a\x34\xbb\xc8\xd7\xf5\x39\x5e\xc6\x92\x36\x6c\x80\xd5\xe5\x9c\x59\x32\xb7\x58\xa3\xcb\xed\x55\xfb\x72\x2a\x31\x31\x68\x7e\x87\xc0\x63\x8f\x3d\x1c\xa0\x74\xf7\xea\x5a\x85\xab\x20\xf4\xbc\x95\x00\x87\x43\xa5\x80\x54\x80\x7d\xbb\x70\xbf\xee\x01\x44\x3a\x9f\xca\x32\xf4\xbc\x2b\xb3\xaa\x17\x91\x5b\xe2\x90\x62\x0d\x4b\xe6\x1e\xb7\x4a\x43\xf6\xc6\xa3\x67\x5e\x9e\xc1\xee\x5e\xac\x38\xb7\x0f\x41\xe0\x1b\xcb\x28\x24\x6d\x2a\xeb\x27\xf4\xa2\xfb\xc1\x0e\x8b\xee\x17\x72\xa9\x7b\x57\x49\x4d\x12\x73\xb9\x72\x58\x48\xa5\xb5\xf2\x03\x7a\x8e\xf7\x7a\x8a\xaa\xeb\xcb\x65\x33\x8b\x08\x98\x8b\x74\xd3\x61\x79\xc4\x95\x1a\xf2\x4c\x8e\x10\xed\x9a\x0c\xa0\x71\x93\x3d\x09\xa3\xaf\xd1\x53\x0d\x61\xcd\x00\xcd\xdc\xd6\x35\x66\xdd\xd1\x99\x38\x97\x9c\x6d\x88\x57\x97\x44\x22\x6a\xa1\x60\x92\xc9\x52\xa7\xdf\xf7\x99\x67\xff\xd0\x8f\xfa\x03\x2d\x90\xe5\xf7\x70\xed\x61\xe7\xb3\xc2\x9c\x83\x7d\xd3\x73\x53\xfc\x94\x0e\x73\x36\xa8\x07\xfc\x41\xa2\x5d\xbe\xac\x7f\xc3\x93\xf6\x5c\xb5\x3e\x30\x05\xd8\x50\x89\xbc\x3c\x2c\xe4\xaa\x67\x5e\x01\x2c\x64\x8d\xf7\x0e\x87\x0d\x70\x8e\xbb\x3b\xb2\x5d\xc6\x03\x8c\x4f\x11\x6c\x3c\x87\xfa\x24\x08\x9c\x5c\x37\x27\xd7\xcd\x43\x5c\x37\x9f\x89\x47\xd0\xe4\xca\x78\x24\x8d\x95\xa1\x40\x67\xae\xab\xf5\xa2\x2d\x8c\x94\xa5\x1d\x2a\x43\x14\x95\x10\x41\xb7\x19\xdb\x80\xf9\xac\x15\xcd\xdd\xa5\x08\xd3\x7d\x38\x25\xb5\xfd\x03\x6e\x1b\x84\x5c\xd6\xbc\x39\x04\x39\x72\x70\x1d\x4b\x00\xb2\xdd\xe3\x2c\x56\xaf\x2f\x1e\x61\xc6\x52\x15\x30\x2e\x51\x5c\xf0\x5e\xe7\x43\x43\x6f\xae\xdf\x2f\x11\xe3\xe8\xe6\x7c\xf9\xf6\x09\x92\x48\xb7\x10\xd2\x2f\x75\xae\x45\xe0\xcd\xc8\x52\x68\x69\xb2\x4c\x9b\x3f\x25\x8b\xe2\x4c\x11\x85\xfb\x96\x85\xa4\x0f\x4d\x3b\x95\x93\xd1\x1c\x1c\x48\x67\x87\x5f\x68\x2e\xaa\xd2\xad\x76\x97\x47\xdf\xe7\x0f\xeb\xbb\x91\x77\xdb\x46\xee\x37\x07\xec\x67\xf2\x64\x25\x05\x3a\xe2\x00\xfe\xfe\xf9\xb1\x56\xfd\xe8\x61\x6c\x6d\xb3\x5b\x48\xac\x5b\x60\x5d\x99\x43\xeb\xde\xe9\x47\x07\x9a\xcf\x5a\x9f\x26\x19\x78\xbf\xbc\xd1\x09\x00\x2f\xcf\x2e\x3a\xbb\xaf\x05\xe1\x6c\x3f\xeb\x48\x7b\x5f\xb5\x46\x0d\xb7\x5a\x23\xa1\xc8\xb7\xa3\xf0\xad\x4e\x6b\xd3\x46\x63\x7d\x79\x76\xa1\xd6\x76\x41\xec\x2c\x6e\x9c\x8a\xe8\x9c\x4d\xdf\x8e\x35\xb4\x61\xe6\x4c\x9a\x8e\xc6\x34\x6f\x7c\xaf\x4f\x59\xbe\xb7\xfc\xc8\x4a\x74\x3f\x2b\xea\xb7\x5a\x8d\x9f\xf6\xac\x9b\x45\xb4\xd1\x13\x48\x22\xbd\x7c\xae\xbd\xca\x6b\xac\x63\xca\x33\xa9\x3c\x93\xec\x80\x27\x4d\xb7\x77\x87\x0f\xf1\x13\xd4\xad\x0f\xf2\xfb\xa8\x74\x16\xe8\x3e\xce\x38\xee\x40\x58\xab\x7c\x55\x3f\x67\x1d\xff\x78\xe8\x41\x57\x02\xbb\x63\x42\xe6\xa5\xb6\xd4\xdb\xae\x05\xc8\x98\xe8\xae\x31\xfc\x85\x89\xaa\x44\x02\x35\xfe\x82\xbd\x35\xf3\x39\x77\xc2\x7a\xd5\x7c\xa5\xe7\xa1\xf4\x44\x83\xc1\xbb\x39\x10\xe1\xd3\xd5\x9e\x71\x7a\x24\xc8\x47\xaf\x98\x17\xaa\xd3\x6a\x51\xdb\xb1\xa4\xb5\x21\x24\xd1\xdf\x92\x21\x1e\x52\x24\xa2\x36\xed\x3d\x92\x1b\x12\xbc\xdb\x0a\x7d\x8f\x46\x2f\xcf\xf4\xc2\x39\xe3\x4f\x95\x13\xf5\x83\x99\xbc\x5c\x59\xc7\xe4\xca\xee\xb9\xfc\xc6\xe4\xfe\xd1\x8c\x02\x5d\xaf\x8a\xd7\xb0\x88\xa1\xa9\xcd\x30\xba\xb1\x03\xd5\x80\x71\x33\xd6\x94\x67\x5f\x02\xa6\xc4\x8e\xf4\x44\xae\x3b\x5d\x8c\x54\x53\x9d\x9a\x56\xbe\x44\x23\xba\xaf\x4d\x46\xff\x4c\xab\xbc\xdd\xdf\xdc\xe1\xbe\xa6\xff\x68\x14\x5b\x32\x46\x91\xb1\xd7\x22\x7c\xb4\xd5\xf1\xc1\xc9\x03\x47\x82\xfb\x9e\xd1\x0d\xd9\x5e\xe0\xa0\x19\x64\x65\x1d\xf2\xd0\x73\xf4\x6f\xc8\xc7\x81\x81\x20\xda\x70\xe6\x6b\x37\xac\x8c\xbd\x69\x80\x3a\x01\x95\x6e\x9f\xd9\xad\x8c\xea\x98\xdb\xd6\x46\x2a\x6d\x5f\x57\x23\xe9\xb8\x47\x9c\xdf\xc6\x48\x4e\x58\x07\xd7\xb4\xb3\x04\x44\xc1\x3e\x43\xc9\x08\x01\xe3\x72\xaa\x13\x77\x04\xfe\x6b\xd1\xc7\x79\xe8\x0a\xfe\x59\x67\x81\xe0\xa9\x5e\x13\x8b\xa5\x78\x19\x4e\xd1\xe6\x86\x59\xe3\xe3\x25\x71\x79\x0d\xde\x7c\x5c\x56\x9c\x6c\xfc\x04\x2e\xf1\x4e\x28\xd2\xf8\xf8\xb1\xbc\x38\x47\x6f\x18\xf5\x1e\x11\xd9\x24\xc5\x6d\x89\xb6\x6f\xe8\xca\xe1\xf1\x88\x6f\xbb\xfa\x84\x46\xfd\x8f\x5c\xd7\xb8\x3d\xcd\x5e\xe6\x43\x2e\x2b\xe9\x76\x72\x38\x3a\x46\x73\x20\xda\x2d\x31\xdf\x82\x5c\xe5\x86\x34\xb4\xdb\x58\x96\x5b\x53\x75\x4b\x57\xca\x62\x3c\x6a\x21\x10\xb5\x6c\xf7\xed\xa2\x10\x0a\x67\xb5\x97\x7d\x2a\x39\x60\x4b\xb9\x3b\xf5\x0c\x63\x0c\xab\x8d\x10\xbc\xa1\x54\x5b\x49\x6d\xeb\x23\x14\x6b\xeb\xb7\xa5\x66\x01\x36\x6d\x56\x7c\xda\x99\x5f\x7a\xbf\xed\xea\xf3\xfa\x0d\xc4\x14\x5a\xbf\xee\xf6\x12\x00\x8e\x04\x01\xa5\xf6\x6f\xba\x00\x7f\x1d\x1b\x95\x9a\x0d\xee\x35\xbd\x2a\xbd\x8f\x04\x48\xed\x7c\x64\x5c\xba\x91\x64\xc8\xd7\xbd\x15\x07\xeb\xe9\xdf\x35\x78\xe8\x77\x28\x80\x1f\x4a\xae\x0f\x8b\x32\xa5\x79\x7a\x6e\xd1\x05\xea\x0a\x64\x6c\xdf\x3a\x00\xbc\xa9\x89\x31\x86\x6b\x94\x16\xd7\xfc\x27\x93\x15\x77\x1f\xae\x89\x87\x85\xfa\x39\xef\x4b\xd4\xa6\x30\x42\xc6\x2e\x5a\xb5\x11\x61\x7e\x07\x5b\x7b\xd1\xf6\x4d\x1c\xcd\xf1\xd4\x78\x59\x8e\x1b\x29\x61\xc9\xec\x2f\x3d\x10\x1b\x88\x14\x89\x89\x1f\xe8\x5d\xdd\x79\xa7\xad\x14\x59\x4f\x0c\x3e\x40\xef\x08\x67\xd4\xd7\x15\x0e\x30\x27\x78\xed\xf5\x39\x5a\xc8\x2d\xe2\x29\x15\xd5\xea\x38\xc6\x46\xd4\xa3\x43\xd7\x39\xda\x45\x0b\x60\x99\x86\x39\x78\x25\xd6\xfd\xc8\xee\x6f\xa5\x5a\x54\xb3\x2a\xa2\xca\x71\xa6\x26\x4b\x6a\xb3\xbf\x55\x2e\x0f\xea\xa8\x20\x56\xea\xbb\x51\x01\xb4\x9c\xc7\x86\x56\xf1\x65\xde\x82\x29\xfc\x2c\x96\x7f\xec\x7f\xd3\xca\x04\xfb\x12\xd1\x7e\xcc\xe5\x22\xdb\x40\x57\x32\x6e\xa6\xab\x84\xa9\x69\x81\xcc\xe7\x35\x88\xa8\x0a\x8c\x40\x22\xea\xdb\x2b\xa7\xab\x85\x67\xa7\xc5\x2a\x03\x85\xfa\x02\x0d\x2e\x3d\x03\x4d\x7a\x23\xaa\xcc\xfd\x11\x24\x6c\xbd\x23\x4c\xea\x8f\x96\x28\xfe\x47\xb4\x96\x63\xba\xea\x9a\xfd\x37\x3a\xeb\x5e\x1d\x5a\xef\xf6\xaa\x53\x99\xdb\xa8\xb2\xd2\xa0\xf5\x6d\xfb\xa4\x30\x6a\x4a\x95\x63\xd9\x47\xb8\x6c\xa6\xf6\x05\xec\x9b\x13\x27\x75\x4a\x0d\x7e\x8c\x2a\x5e\x7d\x1d\x9b\xcd\x5b\x43\x27\x08\xa9\xc1\xd1\x4c\xab\xbc\x92\x22\xd4\x3f\xa4\x59\x23\x90\x4e\x81\xd2\x3b\x4b\x67\x3c\xdc\x5e\xc2\x3d\x3d\xfc\x01\x16\xa4\xfc\x40\x9d\xc5\xaf\xfd\x04\x39\x7b\x69\x5f\xba\xd7\x99\xca\x6e\x69\x2c\xac\x3a\x75\xe0\xd5\x26\xd2\xa2\xf7\x6f\x29\x7a\xe4\x5a\x2b\x34\x49\x48\x18\x85\xfb\xac\x03\xad\xbe\xbb\x5d\xe3\x6a\xb2\x17\xff\xcf\x67\x18\xa1\x35\x10\x68\x2e\x21\xf0\x88\x83\x45\x23\x54\xe2\x86\xfb\x00\x09\x63\xb5\x13\x8f\xc7\xea\x9c\x6f\x3d\xea\xdf\x1a\x2c\x7c\x7f\xf5\xc3\x66\x1f\xc9\x2f\x68\x5c\xa0\x8b\x84\xdc\x66\xc9\xa0\xac\x79\xd6\x27\xbd\x44\x73\xa3\xe9\x31\xa1\xdb\xef\x51\x5e\x2e\x88\xae\xa3\x91\x08\x6f\x44\xce\xd6\xbb\x07\x65\x01\x5c\x5d\x4e\xad\x6c\x33\x48\x33\xc6\x3c\x21\xd8\xd1\xf9\x82\xc4\x48\xd4\x4d\x67\xf1\x72\xa6\xa4\x26\xa3\x49\x6a\x62\xa3\xd8\x66\xff\xf2\x96\x0b\xd1\x50\xe0\xb2\x77\x06\x8c\x85\xa8\xcf\x81\x21\x56\x84\xae\x12\x94\xec\x33\xcf\x82\x26\x28\xf9\x54\x79\x65\x9e\x6f\x00\xf2\xf0\xf1\xc1\xd9\xe1\x99\xd7\x9d\x5d\x5e\xaa\xce\xa3\x0b\x3b\x2e\x0f\xcc\xeb\xc6\x92\x06\x4f\x37\xdc\xc5\x34\x79\x0c\xeb\x63\x7f\x19\x31\x26\xe2\xf5\x12\x83\x6a\x92\x48\x0d\x41\x94\x7e\x4c\xb3\x8a\x01\x9c\x59\x0b\x3d\x47\x84\x8f\x1f\x0b\x29\x4c\x2a\x50\x28\x8c\x9b\x96\x05\x86\x16\xf1\xab\x04\xb9\x8b\x92\x68\x75\x1c\xf1\x93\xe1\x48\xc6\xf1\xa0\x01\x55\x32\x2d\x15\xc6\x68\x27\x83\xc8\xe9\x60\x44\xbe\x06\x16\x84\x87\x17\x4d\x23\x32\x50\x15\xed\x11\x20\x67\x24\xd3\xb4\x23\x8b\x0e\x3d\x8c\x64\xfe\xde\x4a\x06\xb5\x0e\x14\x31\xbe\xdd\x23\x21\x3e\x9e\xb1\x82\xaf\x15\xe0\x2a\x40\x26\x7a\xe6\x81\xb2\x0f\xaa\x96\x40\x8a\x4e\x62\x84\x30\xfa\x05\xb0\xdb\xc8\x8c\x33\xfa\x98\x9d\x6e\xaf\xb9\xb0\x56\xc6\xfc\x7b\xae\x46\x99\x27\x7a\x51\x9d\x33\x20\xfa\x68\x2c\x58\x51\x17\x81\xde\x00\x75\x98\x0b\x2e\x22\x14\xad\xb1\x80\x7f\xfd\xb3\x73\xfd\xb8\x5c\xce\xb6\x59\x5e\xd9\x9a\x12\xee\x03\x4b\x20\x5a\x25\xf4\x85\x13\xdd\x31\x21\x09\xdd\xce\xe3\xd2\xc4\xa8\x60\x04\x18\x82\x35\x94\xea\x85\x0b\xeb\xd0\x42\xc9\x1d\x11\x64\xed\x45\x4a\x6b\x63\x79\xdb\xc5\xd5\xd0\x3b\xae\x69\x9c\x5a\x65\x01\xbc\xc1\x55\x35\x6e\x12\x8b\x9d\x7d\x3d\x53\xad\x96\xbc\xd5\x4a\xb7\x76\x04\xe8\x46\xd8\xe4\x11\xbd\xc4\xdd\x4f\x85\xc7\x76\x1d\x48\x32\x46\x39\xc5\x24\xf4\x3b\xfd\x5d\x5a\x03\x4d\xd1\x1e\xd2\x23\x1d\xb9\x05\xf9\xec\x59\xa9\x12\x46\x84\x15\xb1\xe7\x53\x03\x62\x24\x71\x33\xf3\x39\xe2\x21\xa5\xda\x2c\x10\x77\xed\x5a\x9c\x14\xfc\xc0\xc3\x32\x67\x93\x2d\xda\x68\x69\x31\x93\x69\x33\x93\x59\x73\x4c\x9d\xee\x35\xd1\x7e\x32\xdd\xcb\xb3\x54\x8c\xd4\x43\xd2\xee\xd3\xc7\x92\xf5\x26\x05\x6d\x9f\xb1\x0b\x69\x6e\x73\x7e\x16\x31\x7e\x75\x76\xb3\x88\x07\x28\x7f\x11\xf2\xee\xfa\xbf\x9b\xcb\x4f\x63\x64\xda\xf1\x2d\x17\x2d\xb8\x77\xae\x6d\xf6\x59\x71\x38\x79\xa8\x0c\x37\x89\xa7\xb0\xc9\xd8\x13\x02\xd8\x32\xe8\x24\x59\xd3\xc8\xe0\xd7\x5c\xe4\xb8\xb4\xc4\xb1\x29\xb8\x8b\x1c\xe6\x07\x58\x6a\xe9\xf4\xae\x6f\xa5\xe3\x9f\x42\xe2\xb9\x56\x83\x6f\x4b\xef\xe3\x05\xfe\x2f\xe3\x16\x2a\x24\x5f\x10\x6a\x65\xdc\x25\x96\xd5\x6c\xaa\xcf\xb8\x1c\xac\x25\xf8\x4a\x4a\x69\x97\x43\xe3\xd8\x95\xa1\x23\x3f\xcf\x1a\x94\x37\x2f\xe3\xa2\x0b\x70\x5f\x55\xe5\xca\xcf\xd7\x93\xab\x3b\xb2\xea\x7a\x74\xe5\x36\x60\x1c\x60\x87\xc8\x47\x5b\x49\xb8\xe2\xe1\xc7\x5b\x01\x73\xaa\x64\xf0\xfc\x2a\x19\xd8\x2a\x0b\x90\x48\x78\x2d\xca\x02\xf4\x90\x7c\x1b\xab\x0f\x44\x1e\x50\x3d\xcb\x21\x9b\x41\xaa\x25\xe2\x4e\x8e\xc1\xad\xcf\xd2\xba\xd3\x70\x0f\xb7\x26\x43\x1b\xdb\x66\xe8\x2f\x69\x5d\x19\x7e\x98\x24\xe8\x2f\x52\xfe\x8a\x3c\x46\x7a\xd0\x6c\xfc\x40\x7e\x89\x6d\xd3\x67\x97\xb4\xae\x5c\x62\x9c\x34\xbb\xe5\x12\x4b\x72\xe9\x96\x2e\x54\x90\xaf\x07\x2c\x34\xd3\xba\xba\xd8\xc1\x0e\xd3\x6d\xba\x50\xa4\xfd\xef\x3a\xb3\xd2\xef\x1d\xdb\xcd\x0a\x1b\x41\x63\x9a\x65\x85\xf6\x35\xf1\x3c\x05\x85\x28\x6c\x43\x41\xe6\x4d\x90\x88\x19\x6f\x7b\xcb\x38\xdd\xdc\xd0\xba\x08\x01\x35\x39\x87\xb2\x27\xd1\xd5\xc1\x2c\x1d\xe3\x19\xb9\x98\x35\x38\x98\x05\xe1\xda\x23\x4e\xf7\xe1\x97\xa6\xff\x71\xe5\x84\x80\x93\x6a\x35\x8f\xf1\x1d\xaf\x44\x26\x16\xae\xbd\x3a\x4c\x5a\x72\xe2\x54\x16\xa7\x34\xae\xdb\x2b\x1d\x6e\x6b\xd9\x3f\x5c\xc7\xfb\xd6\x4b\x0f\x36\x97\x11\x09\x17\xc5\x55\x8c\x21\x2c\x28\x0a\xb0\x6b\xd0\x29\x65\x5a\xe5\x2d\x10\x71\xd4\x66\x3f\x03\xd1\xf7\x14\x71\xd8\x51\x83\x04\x9c\x47\x2e\x4e\xd5\x50\xfa\xc0\xb5\xdf\x8f\x90\x98\xba\x98\xbb\x48\x00\x27\xd8\x23\x5f\xf1\xda\x03\x74\xb6\x5c\x20\x9d\xc8\xf9\x96\x5e\x80\x10\x51\x80\x97\xc3\xa8\x6a\x2e\xcd\x4f\xc8\x37\xbf\xfc\x78\x4b\xff\x86\x6e\x67\x84\xde\x61\x8f\x98\xec\xf7\xea\x78\x6e\x67\xe6\xfb\x97\x90\x49\x8c\xe0\xc1\xd1\x51\x8a\xf1\x57\xdd\xd6\x64\xe7\x33\xf3\xcc\x6e\xe9\xc9\xc9\x09\x48\xe7\xe4\xe4\xe4\x96\x2e\xce\xb5\x39\x91\x92\x2f\x21\x44\xb3\x11\x17\xa8\x24\x1b\xe2\x98\x5e\x0e\x73\xe1\x96\x9e\x83\xc4\xc4\xd3\xbe\x10\x2c\x88\xea\x6b\xaa\xb3\x82\x87\xc2\x22\x05\xfa\x4c\xa8\x8b\xcd\xe4\x1b\x02\x9e\x8b\x5e\xc7\xf7\xfd\x35\xf2\x43\x21\xd1\x1a\x10\x65\x74\xfe\x15\x38\x43\x77\xd8\x0b\x93\x1d\x50\x26\x11\x50\x16\x6e\x77\x48\x92\xed\x4e\xea\x34\x1e\x1b\x00\x17\x6d\x59\xb0\x03\x1e\xb7\x4b\x32\x7b\xbc\xfe\xc8\xdc\xd7\xc8\x65\x20\x5e\x4b\x04\x0f\x44\x48\xd5\xe4\x67\x35\x6b\x7e\xa9\x02\xb4\x9e\xf5\x33\x3c\xce\xf5\x8c\x28\xc0\xa4\x8f\xbb\x8d\x6b\x8e\xe3\x89\xd2\xe5\x45\xc0\x28\xa7\x9c\xfa\xcc\x45\x83\x6d\x3d\x3d\xa2\xac\x66\x33\x2d\x9d\x9a\x1e\x94\x19\x4e\x1d\x9e\x06\x73\xb9\xcd\xfd\x48\x96\x9c\xfc\xe5\x2e\xbb\xb2\xce\x0e\x38\x7f\x6c\x2b\xbb\x94\xd3\x25\x75\xa7\x31\xf9\x08\x74\xe1\x56\x55\x4b\xd1\x98\xde\x99\xeb\x47\x17\xbd\xae\xac\xfa\x6a\x27\x65\x60\x8d\xed\xa9\x19\x7e\xb9\xbe\x5e\x1e\x48\x4d\x9b\x0e\x3c\x4b\x14\xe3\x03\xac\x25\x8b\x86\xf8\x64\x30\x30\x47\x87\x16\xe7\xdd\xef\xa7\xa2\x41\x0d\x57\xe0\x03\xe7\xbf\x12\xea\x16\x56\xa0\x3a\x36\x20\x7b\xd7\x42\x37\xfd\xb0\x57\x90\x16\x1b\xba\xda\x3b\xd1\x38\x51\xdb\xe2\xbc\x61\x57\x95\x58\xdb\x61\xd9\xe5\xc8\xb4\x05\x3a\xf7\x98\xe1\x2a\x49\x26\x7c\x31\x1b\x02\xe3\xb6\x20\xcf\x3c\x2f\x2d\xc9\xf9\xfb\xaf\x3f\x31\xf7\xb1\x8f\xdb\xfc\x53\x15\x1c\x3d\xf4\x10\xd5\xea\x71\xc6\xb8\x96\x39\x91\xa8\x28\xa3\xf1\xef\xeb\x7f\x22\x4f\xe1\xb5\x38\xe0\x69\xa8\x09\x26\x04\x49\x8f\x24\x59\x7f\xa3\xfc\x9a\xfa\x4f\xce\xe7\x99\x30\xf2\x84\xc8\xe0\x3e\x0e\x2d\xbc\x98\x2c\xe7\x4f\x5b\x41\x96\xcd\x94\xa5\x46\xde\x6f\x1f\x76\x38\x86\xd4\xd8\x36\x7c\xc5\xfa\x39\x6c\xa6\x8e\x2b\xef\x86\x56\x94\xf8\xf8\x61\x05\x0f\x72\x25\xf6\x12\xab\xa3\x41\xcb\xd3\xe1\x87\x0f\x0f\x72\x2f\xdb\x66\x71\x25\x84\x1e\x63\x25\x0b\xda\xbc\x12\xc9\xf1\x66\x53\xad\xf6\x1a\x60\x15\xd7\xd1\x0c\xc7\x4d\xb1\x6d\xcc\x57\xfd\x7c\x17\xf5\xf0\xb5\xc5\xfb\x2b\x52\x8e\x95\x53\xea\x94\x96\x6a\xb1\xb7\x2c\x01\x98\xed\x5c\x36\xef\x8e\xe4\xd7\x5b\xd8\xaa\xf6\xdd\x5d\xe8\x01\xec\x39\xfb\x76\xd4\xcd\x54\x42\xe5\xa5\xd7\x23\xac\xda\xf8\x2f\xb8\x2a\x38\xb7\x7d\x8d\xc2\xaa\xb1\x5b\x25\x7f\x18\xe8\x05\xd0\x03\xf6\xbf\xe4\x0f\x71\xaa\x21\x7b\x30\xdd\x7f\x09\x35\x64\x9f\x1e\x0f\x6f\xf2\x17\x6e\xc2\xc3\x09\x0f\x8f\x89\x87\x19\xb6\x6d\x29\xb5\xc0\x71\x9f\x63\xc7\x8b\x34\x3e\x1e\xc8\x2a\x6b\x39\x37\x93\x8d\xb4\xcb\x5e\xd6\xb1\xc4\x74\x92\x8b\x1a\x39\x18\xf4\xc7\xcf\xeb\x36\xc8\x29\xee\x55\xac\x6e\x7b\x96\x69\xc7\xdc\x89\x66\x12\xb4\xb7\x4c\xfe\x30\x66\xa5\xc6\x53\xc5\xeb\x77\x04\x6d\x31\xc9\x44\x8b\x58\xee\x34\xff\x42\x14\xcb\x2d\x9e\x37\xc4\x26\x35\xd4\x60\x5a\x24\xbb\xf9\x13\x52\x5c\xd5\xfc\x37\x4a\x8c\x70\xc8\xc3\x79\x3f\x97\xc2\xfe\xea\x57\x3a\x27\x9b\x35\x8d\x8f\x99\xe4\xbd\x9e\xe3\xc9\x54\x33\x03\x7a\x2f\x1f\xec\xe9\xd8\x9b\x50\xed\x03\x7f\x34\xa9\x63\x8e\x4b\xb8\x9e\x75\xb6\x99\x23\xc8\x80\xaf\xa2\x49\x67\xa9\x3b\x71\xb2\xaf\xd9\xcf\xc4\x93\x85\xac\x33\x55\x95\xb7\x0a\x38\x75\x16\x3b\xc2\x6c\xcc\x10\x48\x80\x9a\x40\x82\x8b\xee\x89\xdc\x21\x87\xf9\x3e\x3e\x41\x57\x00\xe8\x76\xe6\x62\x76\x3b\x43\xd1\x3e\xb4\x26\xdf\x67\x5c\x31\x0f\x43\x33\x08\xa3\x19\x0d\x57\x7c\x88\x66\xe0\xf4\xbb\xa6\x04\xb3\x2f\x21\x14\x4d\xb6\x1a\xb5\x17\xe7\x6d\x36\x51\xe1\xc0\x5c\xdc\x9c\xe1\xf0\x59\xaf\x81\x78\x51\x85\xba\x4e\x66\x4d\x01\x96\xbb\xd2\x37\xb2\xe4\x21\xe4\x96\xba\xc4\x5b\xf8\x2d\xcc\x16\xfa\x98\xf9\x84\x12\x5f\x7f\xfa\xe1\x5d\x63\xf2\xaa\x78\x19\x41\x2e\x21\x6c\xe5\xc9\x2c\x81\x2f\x13\xb7\xc4\xde\x73\x02\x5f\xb5\x9d\xd7\xd8\x8b\x07\x85\x49\x34\x66\x29\x54\x02\xf3\x5b\x77\xb8\x64\x72\xde\x0f\xb7\xe2\x44\xf5\x52\xb6\xe4\x6c\xbf\x43\x57\x7b\x15\xae\x85\x24\x32\x94\xe0\x46\x61\x33\xc3\x2d\xfa\x7a\x17\x7b\x78\xe8\xec\x28\xe6\x1f\xfd\x38\x61\x14\xad\x61\x87\xbd\x8d\xfa\x7a\xbf\x23\xce\x4e\xff\x86\x5d\x85\x4d\x42\x72\x2c\x19\x47\x0e\xa6\x28\x00\xae\x26\x33\xbf\x3a\x79\x93\x6e\xbc\x7d\x6d\xe0\xdb\x3f\x83\x12\x5c\x32\x3b\x34\x09\x7c\xfa\xec\x33\x9e\x39\xca\xda\x93\x85\x89\x99\xdb\x24\xf0\x69\x07\x01\x35\xc4\x25\xf3\x60\x7f\x59\x40\xf5\xbd\xfa\x4f\x9e\x9b\xe4\xd2\xf6\xa8\x03\xdb\xd7\x8e\x55\xed\xa7\xb0\x6e\x9d\x60\xb2\xcb\xca\x13\x3e\xc0\x41\x04\x8c\x8a\x0c\x2f\x9f\x19\x67\x9e\x6a\x01\xc2\x38\xca\x14\x7c\x7e\x66\xc2\xd9\x81\x8f\xdb\x14\xcf\x4e\x7d\x5b\x5f\x65\xd9\xa1\x5e\xd3\xab\x6f\xff\x1b\x00\x00\xff\xff\xbb\xf8\xe1\x4b\x57\xb2\x01\x00")

func init() {
	if CTX.Err() != nil {
		panic(CTX.Err())
	}

	var err error

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
