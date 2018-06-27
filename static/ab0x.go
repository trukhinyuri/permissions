// Code generated by fileb0x at "2018-06-27 13:07:00.144911195 +0300 MSK m=+0.014651406" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(b6e8a429627329ed9e112c816563026c.f7d60b70267ef579d9aad25918a6b087)

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
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x3d\x6b\x6f\x1b\xb7\x96\xdf\xf3\x2b\x08\xed\x02\x4d\xb0\xb1\xdd\xbd\xf7\xe2\x02\xdb\x6f\x6e\x9c\xa6\xc6\x8d\x53\xc1\x8e\xbb\x0b\xdc\x14\x02\x35\x43\x49\x6c\x66\xc8\x29\xc9\x71\xe2\x16\xf9\xef\x0b\x92\xf3\x20\x47\xf3\xe0\x3c\x28\x8d\x13\x7d\x8a\x33\xe2\xfb\x3c\x79\xce\xe1\x39\x7f\x3d\x03\x60\x11\x50\xc2\xd3\x18\xf1\xc5\x0f\xe0\xdf\xcf\x00\x00\x60\x01\x93\x24\xc2\x01\x14\x98\x92\x8b\xdf\x39\x25\x8b\x67\x00\xfc\xf6\x52\xb6\x4d\x18\x0d\xd3\xc0\xad\x2d\xff\x04\xb7\x5b\xc4\x16\x3f\x80\xc5\xdf\xce\xbf\x5f\xa8\x6f\x98\x6c\xe8\xe2\x07\xf0\x97\xee\x1b\x22\x1e\x30\x9c\xc8\xbe\xb2\xd5\x12\xb1\x18\x73\x8e\x29\xe1\x80\x23\xf6\x80\x03\x04\xb8\xa0\x0c\x71\x90\x72\xc4\x00\x43\x9c\xa6\x2c\x40\x1c\x40\x12\x02\x18\x45\xf4\x13\x07\x82\x82\x18\x12\xb8\x45\x00\x06\x01\xe2\xea\x43\xd1\x50\x4d\x0a\xc0\x42\x60\x11\xa1\xea\x14\x97\xcb\x6b\xb9\xd8\x2f\x7a\x63\x50\xec\x78\xb9\xb2\x0b\x3d\x18\x2a\x3f\x01\xb0\xd8\x22\x61\xfc\x57\x0e\x0b\xb7\xe5\x49\x64\xdf\x8c\x19\x16\xc5\xf7\xdf\x5e\x96\x9d\x78\x1a\xc7\x90\x3d\xca\xd5\xdc\x22\x91\x32\x92\xed\x2e\x9f\xf1\x7c\x61\x34\xa6\x09\x62\xea\x6c\xaf\x43\xd9\xe1\x0d\x12\xf7\x1c\xb1\xcb\x7c\x71\x46\xcb\x04\x32\x18\x23\x81\x58\x75\x45\x7f\x19\x7f\x03\xb0\xf8\x4f\x86\x36\x72\xa8\xff\xb8\x28\x7b\x5c\xc8\x41\xaf\xaf\x7e\x46\x30\x44\x6c\x61\xb4\xff\xf2\x72\xc0\x40\xb7\x34\x42\x13\x0c\x75\x97\xae\xb9\xc0\x22\x15\x28\xd4\xcb\xb3\x46\xab\x3d\x5a\x86\x78\x42\x89\x0d\x35\xf5\xc3\xdf\xbe\xff\xbe\xf2\x69\x1f\xfb\x2c\x20\x58\x58\x04\x9e\xbf\xff\xe5\xea\x17\xc0\x83\x1d\x8a\xe1\x8b\xc6\x4d\x2d\x42\xb4\x81\x69\x24\xf6\x67\x2a\xf7\x57\xac\xf0\x02\x31\x46\x59\xfd\x96\xf2\xbf\xf4\xbf\xd9\x1c\x8b\x0b\x18\xc6\x98\xd4\x22\x66\x92\x4e\x8c\x98\x97\x9c\xe3\x2d\xc9\x29\x2a\x42\x0f\x28\x02\x1b\xca\x24\xcd\x55\x48\xf1\x1c\xdc\x73\x14\xaa\x1f\xd7\x38\x8a\x30\xd9\x82\x24\x65\x09\xed\x40\xe4\x3b\x24\x6e\xf3\x21\x4e\xd8\xdc\x36\x1a\x81\xb1\xe2\x5c\x6b\x1a\x3e\x2e\x5e\xda\xbf\x61\xd2\xf4\x0b\x43\x7f\xa4\x98\x21\x79\xd6\x82\xa5\xa8\xf2\xab\xc6\xe4\x3d\x3c\xb5\xd6\x1e\xa2\x0d\x26\x58\x42\x8c\x5f\xd4\x41\xeb\x16\xfd\x91\x22\x2e\x16\xd6\x08\x5f\xfc\x92\x68\x86\x8f\x1c\x89\x23\x13\xa1\x04\x0a\x4f\x60\x30\x40\x3e\xbc\x2b\xbb\x76\x51\xe1\x1b\x24\x14\xc5\x95\xb3\x81\xe7\x6a\x7e\x40\x49\xf4\xf8\xa2\x43\x50\x5c\x46\x91\x31\xd7\x89\xb6\x5c\x47\xfb\x09\x47\xea\x6c\xc6\x8c\xb1\x84\x5b\xf4\x2e\x8d\xc7\x8d\x81\x98\x1c\xe6\x2d\x8e\xb1\xf0\x2c\xfa\x0c\x04\xcb\xc7\x59\x38\xb3\x0c\xf1\x98\x28\x06\x45\xd7\xbf\xa3\x40\x54\xfa\x69\x95\x31\x41\x4c\xe0\xbd\xa5\x19\x0c\x6e\x8f\x96\x6a\x67\x80\x8c\xc1\xc7\xbd\x09\x34\x27\x14\x28\x6e\xea\xdf\xc8\xd6\x0a\xf2\x58\xd4\xf4\xfa\xf2\xac\xeb\xcb\x97\x66\xd6\xe7\x97\x23\xbd\x2c\xe4\x3e\xe5\xd3\x72\x9c\x57\x0c\x41\x81\x4a\x86\x03\x3e\x61\xb1\xa3\xa9\x28\x64\xbb\x2b\xff\xb9\x94\xcd\xf4\x68\xe5\x29\x9f\x78\xd0\x4c\xe5\x7b\x01\x22\x03\x6c\xd3\xcb\xf7\xff\x76\xe7\x43\x20\x50\x6b\x08\x67\x23\xe6\x2f\xfe\xc2\xe1\x97\xde\x2a\x77\x0f\xca\xbb\x45\x1c\xff\x39\x19\xe5\xe9\xd1\x4e\x94\xf7\xc4\x28\x4f\x83\xad\x9b\xf2\xfa\x1e\x51\x31\x8b\xf7\x4b\x74\x89\xbf\x4c\xed\xe5\x78\x14\x7c\x38\x15\x5d\x5d\x88\xcb\xe9\x1c\xec\x37\x27\xbd\x7c\x22\xbd\xfc\xa4\x0a\x9f\x54\xe1\x7c\xe4\x08\x09\x34\x29\x6d\x5f\xa9\x21\x4b\x9b\x97\x1b\x89\xeb\x5e\x97\x51\x74\x22\xf4\x43\x10\xa9\x86\xfb\x1c\xc4\x8c\x52\x11\x0f\x63\x9e\x7d\xc3\x20\x11\x86\xae\x98\x14\x7d\x81\xa0\x0a\x59\x3b\x6c\xaf\xa5\xda\xa1\x96\x7b\x92\x42\x4f\xc0\xf2\xaa\x41\x75\x38\xe5\xb0\xb7\xbd\x8a\x51\x29\xf5\xbc\x13\xff\x0c\x8c\xc0\x89\xde\xaa\x4d\xe4\x4e\xb6\x98\x65\xde\xd3\xd1\x12\x93\xcd\xd4\x46\xce\xba\x65\x36\xf0\x89\x92\xe7\x4a\xc9\x19\x80\x8e\x67\x5c\xc9\x50\xe9\xe8\xa6\x95\x9c\x78\x2e\x7e\xc8\xfe\x6a\xb8\xaa\xf9\x37\x6e\xa6\x1c\x93\x6d\x6e\x60\xe9\x26\xb1\x93\x39\x65\x2a\xa1\xf0\x54\x8c\x32\x27\x4b\xa8\x41\xae\x7f\x65\x7f\x7d\xb9\xd8\x32\x9a\x26\x03\x22\x63\x5c\x65\xdf\x1b\x24\x72\xc1\x07\xf4\x5c\x1d\x36\x95\x6c\xe4\x37\x7a\x5d\x27\xea\x9c\x9b\xca\x66\x03\xf3\x70\xf6\x94\x3d\x44\x3d\x88\x2d\x45\x02\x46\xa1\xe2\xd7\xed\x56\x74\x26\xe7\xcb\x30\xd4\x90\x97\x97\xd3\x1c\x15\xdc\xdd\x19\xa1\x3a\xcb\xf7\xf4\xa4\xe1\x7e\x13\xc2\x37\x5b\x74\x0e\xf7\xe9\x85\xef\xdf\x3a\xd9\x95\x46\x56\x18\x86\x28\x34\x50\x76\x7e\x12\xf8\xe2\x2f\xf5\x6f\x7f\xbf\xa4\xfb\x2d\x74\x07\x49\x19\x59\x4b\x37\x19\x1d\xc7\x28\x5e\x23\x66\x92\xb3\x9b\x6d\xf4\x0e\x69\x19\x7d\xa3\xfa\x9f\x6c\x4f\xbe\xe9\xb9\x76\x14\x05\x81\x27\x67\x05\xdb\xc3\x9b\x63\x30\x86\xa3\x98\x9e\x7a\xbb\x38\x9c\xa9\x3b\x73\x70\x68\x9a\x4e\x8c\xc8\xf8\x0d\xa3\x71\x4f\xd2\xd6\x63\x29\x28\xfd\xc4\x68\x7c\x12\xd6\xf3\x21\x6e\x5f\x02\xf2\xd8\xae\x97\x1a\xb9\xa8\x05\x93\x4f\xbb\xac\x54\x66\x2d\x52\xa1\x4c\x7b\x07\x7b\x0b\xc3\xcb\x30\xd4\xfc\xec\xa4\xda\x7e\x23\xaa\xed\x3e\xc4\x8f\x21\xc4\x32\xe5\x4d\xa9\xb7\x73\xa2\x5d\xd3\x93\x5a\xfc\xfd\xc5\x77\xf0\x4e\x31\x53\x87\x89\xe9\x64\xfc\xfd\xfa\x3d\x8b\x56\xf0\x5a\xcf\xd0\x9f\x5e\x61\x35\x47\x31\xea\x78\x0e\x58\x6d\xa3\x9f\x53\x38\xea\x53\xf3\x7c\x1c\x3a\x12\xf5\x49\xb0\x84\x03\xc7\xb3\x1e\x20\xb6\xcd\x89\x7a\x75\xdb\x13\xf5\x7e\x4b\xd8\x3e\xc3\xbb\x5d\xbd\x7e\x78\x98\x87\xfa\x96\xa6\xa8\x9e\x66\xe8\x4b\x9f\x71\x13\x74\xd5\x1f\xff\x17\x8b\x9d\x39\xf5\x89\x9a\x0e\x6b\x21\xc9\xc3\xe9\x0e\x48\x4c\x99\xd9\x30\x1f\xed\x68\xe9\x04\xfa\x5b\x12\x7b\x90\x48\x55\xa2\xf4\x0c\x4a\xad\x08\x99\x93\x6f\x60\xde\x8a\xa2\x06\xd7\x29\x34\xb5\xd6\x3f\xf0\x74\x84\x27\x43\x19\x32\x79\x7c\xd7\x28\x67\x70\xbd\x26\xca\x66\x27\x45\xf3\x29\x5d\x13\xe5\x72\x4e\x84\x6f\x5e\x13\xe5\xdf\x87\xa7\xfd\x22\x91\x95\x01\xa7\x32\x9d\x95\x66\xd0\x6f\xd1\x03\x8a\x4c\x62\xcf\x23\xa0\xb8\x60\x98\x6c\x0b\x9c\x59\x20\x92\xc6\x16\x9d\x2d\x08\x25\x16\x3d\x32\x04\xc3\xea\xff\x33\xad\xc2\xf8\xfa\x89\x61\xfb\x03\xfd\x44\x4a\xba\xf9\xcd\xe2\x5a\x97\x41\x40\x53\x62\xc7\xd4\x57\x4e\x39\x6f\x02\xce\xce\x40\x84\xb9\x00\x74\x03\xd6\x34\x25\xa1\xd4\xb0\x74\xe7\x97\xcf\xda\x43\xc7\x9a\x42\xc6\x16\x70\x7f\xfa\xae\x10\xb4\x05\x0c\x43\x75\xd0\x30\x5a\x36\x07\xa2\x55\x0f\xb9\x11\x31\x3e\x9f\x6d\xe9\x59\x4e\xdd\xc5\x69\xb4\x68\x71\xaa\x43\x02\x83\x8f\x70\xab\xfa\x6c\xb1\x38\x0f\x28\x11\x10\x13\xc4\xd2\xf8\x9c\x20\x71\x11\xec\x2e\x8c\x3b\xca\xc5\x03\x22\x21\x65\x17\x0d\x2d\xa5\x8e\x76\xa6\x93\xa9\xb1\x8b\xe4\xe3\xf6\x22\xa6\x21\x8a\xb2\x35\x14\x70\x6a\x74\x62\xb4\x00\xae\xa9\x0f\xc8\x56\xc1\x41\x49\xb4\x3a\xd1\x54\x18\x62\xb2\xad\x7a\xd6\x46\x81\x57\xea\x93\x15\xe0\x76\xc4\xef\x65\x4a\xe8\xb3\x1a\x78\xa9\x6c\x61\x15\xd9\xd9\x46\x52\xea\xb7\x0d\x65\x31\x94\xe7\xb4\x40\x31\xc4\xd1\xa2\x19\xfc\xf7\xf9\xe8\xd3\x82\xbf\x80\x69\x15\xa4\x31\x26\xbf\xd2\x28\x8d\x2b\x76\xbf\x36\x88\xd6\x77\x69\x04\x68\xb0\x83\x64\x2b\x41\xfa\xa0\x3a\x01\x65\x3c\x86\x1c\xa8\x30\xc0\xe1\x70\x0d\x60\x02\x03\x2c\x1e\x9b\xc0\x80\x89\x40\x5b\xc4\x9a\xe0\x80\x89\xf8\xe7\x3f\x5a\xe0\xf0\x2a\x1f\xfe\x90\x64\xa8\x4f\xa8\x9b\x10\x1f\x20\x8e\xe0\x3a\x42\x77\x34\x4a\x33\xb8\x34\x82\xab\xda\x56\xb2\x50\x9e\xff\xfd\x69\x87\x83\xcc\xa0\x11\x40\x02\x58\x4a\xc6\x10\x9a\xc0\x0f\x8d\x54\xb1\xa6\x34\x42\x90\x2c\xda\x18\x9f\xea\x5f\x4b\x73\x38\xec\x43\x6d\xd6\xb0\xa6\x06\x60\x0d\x19\xc3\x2d\x6a\x64\xfa\xfb\x91\xc1\x0d\x11\xc1\xc3\x98\xfc\xb5\x9e\xbc\x76\x61\x11\x8e\xb1\x70\x64\x58\x39\x50\xdf\xea\x3e\xb5\xe3\xf5\xe5\x55\xd6\x42\xdf\xd9\xac\xc8\x64\x83\x2c\x1a\x3c\xea\xfd\xed\x5b\x6f\x74\xb5\x4b\xd7\xe7\x01\x8d\x2f\x8c\x0e\x17\x1f\xd3\x35\x3a\x0b\x22\x8c\x88\x68\xe4\x84\x55\x3a\xe1\x6f\x71\x3b\x1f\xac\xed\x60\x2a\x28\x30\x6f\x51\xd0\xdb\x08\x2d\xa5\x1c\x62\x34\xc6\x36\x84\x28\xec\x31\x15\x37\x54\x2e\xb6\x3f\x37\x88\xfe\x28\xf5\xc3\x4c\x9b\xca\x8d\x16\x5d\xc2\xad\xb1\x8f\x84\x2b\xcb\xfe\x54\x09\x42\x63\xfa\x80\x6c\x15\xb4\x1b\xb6\xc6\x55\xec\xdf\xd6\xf5\x43\x99\x52\x0a\xe5\xb8\x13\x17\x8a\x1e\x43\xe9\xef\xb6\x32\xe5\x7c\x74\x4c\x13\x00\xdc\x11\x50\xea\x56\x60\xdf\x06\x34\x0d\xe6\x11\x5a\xdd\x90\xe9\xd4\xe7\x1b\xd8\xfc\xb1\x8f\x4b\x47\x69\xbf\x87\x0c\x6f\x36\xdd\xd8\x5d\xd3\xba\x56\x6b\x23\x08\x85\x59\x66\x57\xa1\x1a\x17\x3a\xdc\x70\xfe\xa5\x07\x5a\x8d\x10\xe4\x7a\xdd\x96\x38\xf7\xcf\x67\xd6\x38\x8a\xce\xd0\x67\x21\x95\xf2\xa8\x1e\x04\x94\x6c\xf0\xf6\x27\x1c\xa1\xba\xfb\xb5\xb3\x36\xab\x46\x59\x6d\xec\x61\xa6\x92\xd5\xd9\x10\x03\xcf\x5d\x36\x9c\x19\x6f\xd7\x87\x7e\x03\x93\x36\x6c\xcf\xdb\x48\xf6\xa0\xc6\xd0\xd7\x11\xf5\x19\xc4\x30\x19\xca\xb0\xd5\xc1\x18\xa7\x1b\x42\x01\xdd\x99\x77\xf6\xaa\x73\x05\xab\x86\xa0\xea\xfa\x55\x43\xa9\xa4\x87\x50\x20\x80\x09\xb8\xfd\xe9\xd5\xdf\xff\xfe\xf7\xff\x01\xd9\x25\xe6\xe5\x20\x68\xea\x87\xad\xe1\xa5\xa8\xc7\x14\xb5\x17\x27\xd5\xb3\x38\xde\x2b\x63\xfb\x95\xc1\xb4\x35\xbc\x7b\xab\xba\xe1\xb4\x1b\xd5\xe2\xbb\x71\xa3\x9e\xb4\x62\x6d\xe6\x1a\x3a\xee\x2f\xa6\x91\x6c\x76\xb4\x76\x65\xe3\x46\x23\xbd\xc9\x76\x4d\x34\x07\x14\x82\xcd\x5d\x1e\x8f\x3d\xa9\xae\xbb\x83\xdd\xb0\xe9\xac\xb4\x16\x33\xc2\x42\xa2\x46\x92\x03\xf9\xba\x33\x94\x6c\xd8\xed\xae\x50\xee\x7b\x86\x48\xae\x9b\xb7\x43\x4d\xb7\xd9\x03\x58\xf6\x19\x13\x10\xa2\x24\xa2\x8f\x31\x1a\x7e\x1d\x50\x76\x89\xc5\xcb\x66\x71\x13\x59\x77\x7d\x07\x81\x43\xe3\x18\x92\xf0\x58\xa6\x8e\x57\xf9\xf4\xb5\xfc\x32\xd3\x7b\x3c\x23\xa9\x86\x8f\x36\x5a\x8e\x41\x55\x63\xe1\x88\x3c\xf8\x5a\xf0\x6b\xf2\xe0\xb8\x48\xab\xe5\x9e\x71\x6b\xb8\xc9\x4c\xf5\x1e\x6d\x9c\xaa\xb9\x65\xfa\x17\xc0\x09\x65\xc2\x3f\x2a\x2d\x29\x13\x8e\x30\x5a\xaa\x05\xd5\xae\x55\x9b\x7b\x57\x71\xab\xfb\xe9\x38\xe8\xaf\x1b\xdf\xf8\xf5\x44\x8d\xe5\xd5\x0a\x0a\x0e\xfc\x5a\xb6\xb3\x79\xb6\xc4\x12\xc9\xae\x8b\x99\xa7\xba\x0b\xc8\x81\xad\xff\x33\x2a\x68\x40\x23\x77\x76\xed\x91\x2e\x3c\xb9\x4a\x6c\x62\xf8\x52\xb7\x7b\x27\x7e\xb1\xac\x9c\xd5\xec\xf0\x2d\xa3\x20\x07\x8c\xd3\x2d\x25\xce\x65\x4e\xaf\xe7\x96\x66\xf7\x02\x28\x9a\x47\xe1\x24\x38\xa8\xc6\x5a\x25\x50\xec\xdc\xb1\x4c\x6e\x73\x30\x96\xdd\xc8\xce\xb5\x00\x37\x96\x32\x7c\xf0\x94\x88\xa5\xb1\x99\x83\x88\x0d\x9e\xae\xc7\x2d\xfb\x2e\x5d\x57\x16\x3d\x0f\xdc\x55\xd7\xfd\x6e\xc4\x35\x9a\x81\xb3\xb3\xe1\x77\x8e\x08\xae\xd1\x70\xbf\xd0\x5b\xd5\xbb\xde\x84\xf5\x34\xed\x87\x0e\x20\x6a\x0b\xa8\x6c\x86\x98\xee\x55\x96\x84\xcb\xed\xb9\x98\x03\x58\x78\x2a\x34\xd4\x94\xc4\x53\x26\x16\x4c\xb6\x35\xb5\xf2\x7a\xd8\xc9\x9b\xc0\xee\x3f\x26\xe2\x9d\x4b\x4c\x84\x6d\xfb\xd9\x3b\x9c\x85\x87\x08\x8a\xab\xf2\xc2\xd7\x06\xac\xbc\x91\xad\x86\x94\xb7\x45\x3e\x94\xf9\x17\xab\xe6\x6d\x77\x46\x86\x54\xa9\xc6\x1e\xb7\x46\x9f\xde\x7c\x63\xcd\xbe\x95\x5e\xf7\xdb\x5e\xbe\xa2\xfa\x15\xcf\xdb\x68\x7b\xb2\xb3\x8e\xb3\xb3\xd6\x51\x8a\x1f\x5d\xf9\xb6\x42\x88\x15\x2d\x24\xf3\xa5\xd7\x88\xb9\x0a\x24\x8b\x68\x9e\xeb\x2b\xf0\x9c\x92\xe8\x11\xe0\x8d\xc1\x4e\xa4\x18\x48\x20\x53\x61\x08\xf9\xa0\x2f\x06\x02\x38\x9f\xaa\x29\x90\x86\x0b\x28\x52\x47\x93\x40\xc9\x08\xef\x74\xaf\x7a\x59\x4f\x05\x8c\x56\x41\x92\x76\x1c\x82\x6a\x07\x5e\x2d\xef\x41\xca\xe1\x16\x81\xf5\xa3\x4a\xd9\x5f\xb2\x17\x89\xe4\x62\x87\x79\x9d\x59\xae\x0f\x54\xd3\x2e\xb0\xbe\x97\x2b\x79\xb5\xbc\x6f\xdb\x4f\x8c\x62\xca\x1e\x9d\xb6\x74\x7b\x79\x33\x8f\x2d\xdd\xe8\x35\xd7\x5b\x2f\x10\xe3\x76\xec\x59\x0b\xe0\x7f\xcd\x1a\xcf\x4e\xf7\xaa\xe0\xa3\x8b\x0c\xd7\x4d\xa5\x24\x97\x33\x30\x82\x04\xe2\x40\x53\x81\xa4\xb7\x3e\x26\xe0\x46\xe9\x9b\x07\xe3\xac\x3c\xf3\xa3\x22\xea\xa7\x9d\x31\x31\x04\xc3\xc7\x95\x77\xde\x08\xc3\xc7\xae\x75\x1c\x91\x3b\xa7\xe4\x60\x60\xb9\x2f\xa7\xea\x58\x53\x12\x2a\xc5\xc4\xf7\x7a\xf4\x34\x35\x6b\x99\x19\x19\xff\xba\xc7\x94\x9a\xe9\x38\x6b\xdb\xa4\x92\x83\x8c\xc1\x01\x7d\xc6\xc3\x89\xb9\x81\x51\x3a\x0b\xe0\xd9\xf3\xce\x2e\x77\x67\xa5\x65\xe3\x1d\x68\xa4\xbf\xd3\xbc\x4d\x79\xba\x57\x18\xb7\x3d\xb7\x8b\x85\xb1\xf5\xd9\xc1\x0f\x33\x14\x08\x6d\xf3\xb1\xd3\x9d\xb7\x40\xb2\xa9\x4f\xa3\xdd\x41\x5f\x88\xca\xe8\xff\x3c\xf0\x1f\x3c\xaf\x94\x39\x7c\xf1\x94\x9e\x02\xbc\xf4\x65\xf1\x9a\xcd\x1b\x83\x2b\x1a\x43\xdc\xca\x47\x55\x03\x49\xcf\xb7\xf9\xeb\xf6\xb3\x33\x10\xaa\xaf\xae\x94\x5c\x6f\xd2\xd0\x63\xf4\x30\x55\x84\x21\x0a\x57\xeb\xc7\xc1\x40\xb8\x94\x03\xfc\xf8\xd8\xf7\xea\xef\x66\xdc\x92\xb2\xe3\x4c\xe0\x18\x8d\xb8\xdf\x57\x61\xd1\xf3\x86\x6e\x9d\xe7\x7c\x62\x7d\xf7\x31\xc8\x33\xba\x35\x0a\x0e\x35\xca\x2a\xc2\x5c\x78\x13\x1c\x15\x18\xb4\x0b\x8d\x62\xaf\xf3\x03\xda\x6b\x2b\xee\xa1\x0a\xa5\xd7\xe4\x41\xdd\x8b\xd0\xe3\xd9\x03\x8c\x52\x04\x12\x88\x99\xbc\x14\x21\xf2\x80\x19\x25\x5a\xaf\x82\x0c\x4b\xcd\x76\xb0\xcd\x53\x0d\xbd\x67\xee\x3c\xb6\x87\x55\xaf\x6a\xb0\x92\xa7\x7a\xcf\x4c\x45\x78\xcd\xda\xe2\xa1\x5e\x33\x15\x09\xc5\x05\x24\x21\x64\x21\xe0\x88\x61\x18\xe1\x3f\xd5\x93\x97\xcb\xe5\x35\x50\x6f\x98\x3f\x90\x1b\xc4\x95\x2d\xe3\xec\x0c\x04\x94\xc8\xe6\x42\xff\x04\x62\xfd\xcb\x0f\x1f\xc8\x7f\x81\x0f\x0b\x4c\x1e\x60\x84\x43\x90\x3b\x13\x3e\x2c\xf4\xf7\x3f\x52\x2a\x20\x40\x9f\x03\x15\x96\x9e\x7f\x55\x6d\xb5\xad\x55\xcf\xb3\xf8\x40\xce\xcf\xcf\x91\x08\xce\xcf\xcf\x3f\x90\xeb\x2b\x39\x5f\x4a\xf0\x1f\x29\xca\x66\xc3\x21\x22\x02\x6f\x70\xa0\x7b\x05\x34\x44\x1f\xc8\x15\x12\x10\x47\xea\x36\x4f\x13\x1d\x72\xa8\xac\x2d\xe8\x73\x65\x91\x1c\x7c\xc4\x24\x84\x7a\xf2\x0d\x46\x51\x08\xbe\xcb\xef\x43\xdf\x81\x38\xe5\x02\xac\x11\x20\x94\x9c\xfd\x89\x18\x05\x0a\x1b\xf2\xb5\x12\x2a\x00\x22\x34\xdd\xee\x80\xc0\xdb\x9d\x50\xee\x97\x0d\x42\x21\xd8\xd2\x64\x87\x58\xde\xae\xf0\xc8\x7c\xf7\x86\x86\xdf\x81\x90\x22\xfe\x9d\x00\xe8\x33\xe6\x42\x36\xf9\x49\xce\x6a\x2f\x95\x23\x65\xe3\xb3\x09\x8e\x8f\xd1\xa1\xd5\x71\x1c\x29\x56\x2c\x03\x46\x3d\x7d\xa9\x33\x77\xb4\x33\xea\x93\x72\x7d\x4b\xd8\x10\x86\xc5\x58\x93\xc9\x33\x43\x89\xe1\x3e\xee\xac\x7f\x8b\x3d\x75\xb5\x13\x22\xf1\xa4\xd0\x6a\x03\xd6\xcf\xef\xdf\x2f\x0f\xca\x70\x82\x1d\x62\xb9\x31\xd1\xe0\x30\xd7\x57\xed\x3c\x46\x53\x32\x43\x09\x43\x5c\xdd\x12\x2d\xa2\xbe\xbe\x1a\x8e\xec\x92\xa0\x9d\x71\xe1\x5f\xb2\x71\x3d\xc0\x7a\x60\xd4\xdd\x81\xfd\xcf\x0d\x67\xfe\x2f\x7b\xe7\x35\xa7\x2e\x5b\x54\xce\x5d\x1e\x97\x92\xe5\x2a\x35\xc5\xde\xa9\x57\x91\xb2\xd9\xd8\x7c\xb0\x6d\xde\x75\xe1\xd6\xdd\x1e\x72\x71\xc4\x1e\x70\x80\xc0\xf5\x55\xcb\x56\xab\x59\x33\x0e\xb2\x9f\x9f\xaa\xfc\xaf\xba\x9f\x52\x3c\x18\xfb\x29\x23\xe9\x0d\x21\xa1\x59\x69\xcb\x06\x9f\x40\x3c\x7e\xdd\x11\x5d\x5f\x75\x58\xa2\x74\x03\x79\x44\x38\x6c\xba\x35\xd8\x32\xae\x46\xbe\xcd\xf4\x2d\xe0\x75\x25\xe6\xb7\x37\x3b\x7c\xd7\xad\x17\xd7\xf2\xbf\xf7\x70\xeb\xdc\x6d\x16\x6a\xed\x35\xd9\x32\x3b\x27\xc8\x1e\x9e\xe8\x16\xb6\xa5\x12\x67\xdd\x26\x0a\x15\x65\x69\x84\xf8\x57\xf3\x6e\xec\x14\x82\x30\x5d\x08\x82\x42\x0c\x4f\x86\x88\xdb\x34\x72\x8d\x01\xbf\x35\x11\x74\x66\xb4\x8b\xba\x7c\x0e\x56\xbb\x5a\x3a\x46\x63\xcd\x46\xc5\x38\xbe\x60\x95\xf3\x29\xc7\x9c\x22\xd5\xd6\xf3\x00\xd8\x5b\x4c\x3e\xb6\xc0\x49\xfe\xac\x93\x66\x90\x8f\xe0\xb9\x04\x0f\x43\x5b\xcc\x85\xce\x0e\x78\xa1\xc2\xd0\xf4\x9f\x48\x04\x2f\x86\xf2\x5e\x39\xfa\x14\xac\xf6\x60\x26\x5f\xf4\x39\x91\xbb\xf0\xba\x84\xd7\x7a\x8e\xa6\x25\x60\xbe\x1a\x1b\x04\x78\xcd\xdb\xc2\x00\x23\x1b\x33\x7a\x3a\x4e\x0c\x88\x56\xae\x63\x88\x08\xaf\xe7\x76\x87\x88\x68\x3a\xb4\xba\x97\xf6\xf5\xc4\x2d\x37\xe0\xf5\x69\xfd\x60\x45\xb2\x58\x59\x3b\xd5\xca\x26\x05\xe5\xaa\x7d\xfb\xbd\x27\x8d\xda\x0f\xef\xd8\x0c\xcf\x77\x32\x56\x28\x44\x95\xc9\x26\x15\x08\x36\xd2\xb7\x4a\x03\xbd\xe9\xf9\xe1\x16\xdd\x62\xd2\x7a\x2b\xcf\x5a\x28\x78\xc8\x3f\x39\x80\x44\xdb\x84\x9d\x6c\x3d\x4f\x35\x93\x8b\xda\x76\xb7\x03\xde\x6c\x56\x1c\x51\xe1\x77\x57\xf2\x73\x0d\x39\x0e\xf4\x0f\xc3\xe5\xa5\xec\x6d\x67\xe1\xe5\xfc\x13\x65\xa1\xbb\x10\xd5\x43\x0c\x66\xef\xaa\x77\xfd\xa3\xb2\x7c\x29\x43\xc7\x5e\x56\xf6\x32\x1f\xea\xb8\x81\x9c\xe3\x07\x74\x4f\x78\xba\x96\x60\x5f\x3b\xa7\xf6\xe9\xe8\xd9\x95\xe6\x07\x46\x51\x96\xea\x87\x83\xb4\x18\x63\x54\xc2\x9f\xdc\x97\x71\x2c\x5f\xc2\x6d\x31\xff\xbc\xf2\x05\x95\x49\xae\x9b\xa1\x59\xb4\x91\x24\x6e\x66\x18\xce\xcc\x7a\x4a\x21\x1e\x9b\x6f\xac\xd7\x63\x90\x89\x73\xa8\xce\xdb\x98\xe1\x21\xdb\x64\x6d\x94\x50\x65\xaf\x8a\x45\x3c\x60\x8e\xd7\x11\x02\xaa\x83\x4e\xc1\xb5\x33\x92\xa8\x0f\xdc\x70\xcb\xc3\xba\x18\x7e\x5e\xa1\xcf\x62\x95\x19\xbe\x07\xba\x9c\x3a\xc3\xc9\x6f\xe0\xe7\xd7\x9f\xc5\x5d\x36\x49\xe3\x4a\x30\x39\xc4\x4a\xae\x49\xf7\x4a\x04\x83\x9b\x0d\x0e\x3c\xae\xe2\x7d\x36\xc3\x61\xad\x4b\x6a\xe0\xd5\x38\xe9\xac\x86\x6f\x11\xd1\x8d\xbc\xbf\x3d\x51\x04\x3f\xe0\xd3\xcf\x4a\x42\x67\x6f\x2a\x7b\xf9\x08\xd1\x51\x6e\xdd\x73\xfb\x15\xda\x9e\xcc\xca\xe7\x71\xb7\xcc\xf0\x4f\x70\xbb\x45\xec\xfc\x77\x4e\x49\x6b\x95\xc2\x3d\x19\xa5\x92\x3d\xbb\x06\x87\xb6\xf4\x6a\x4e\x0e\x9d\x87\x87\xda\x35\x96\x8c\xb8\xd0\x11\x61\xa1\xfb\x0f\x88\xa6\x8a\x08\x6d\x7a\xe9\xe3\xed\xf9\x73\x85\x4b\xfb\x8a\xf3\xb7\xb8\x34\x77\x62\xd3\x1e\x97\x52\xb2\x69\xee\x8b\x4f\x8f\x65\xd3\xf5\x8f\xba\xa6\x9a\xbc\xfa\xfa\xca\x63\x26\x78\x9b\x78\x5d\x93\xc1\xb7\xf4\xaa\x21\x79\xa0\xcd\xbb\x1c\xff\xf9\xf5\x50\xfc\x89\x30\x4f\x84\x79\x18\xc2\xec\x2d\x86\x87\x4a\xe0\xaf\x25\xfb\x87\xa1\x8d\xa7\x38\x3c\x7c\x66\x90\x2e\x80\xda\x25\x94\x1c\x00\x6a\x75\x68\x04\xa8\xaa\x43\xd4\x00\xd0\x76\xfd\x5b\xbf\xf2\x2b\x3d\xde\x35\x2b\xee\x29\x16\x7a\x56\x08\x29\x45\x82\x8a\xf9\x3d\x70\xb6\xe9\xb9\xe3\x8b\x9e\xd6\xe5\xd8\x75\x4b\x33\x0c\x2c\x13\xad\x79\x3e\x6f\x79\xea\xc5\x61\x7f\x20\xd7\x1b\xa0\xdf\x74\xdd\xe1\x3f\x11\x48\x18\x7d\xc0\x21\x0a\x55\x30\x73\x82\x18\xc7\x5c\xa8\xb8\x7d\xfd\x82\xeb\x13\x8e\x22\xb0\x46\x20\xb3\xe1\x9c\x8f\x12\xd7\x2b\x95\x50\xd0\x9f\xd0\x56\x85\x31\xe6\xfb\xb0\xc7\x82\xdf\xc0\xd7\x3d\xc6\x18\x0d\xae\x64\x6d\x8d\xf4\xad\x17\xbc\xce\xe6\x69\xd7\x0c\x7c\x54\x72\x21\x87\xd9\xe0\x35\x71\xda\xe0\x78\xb7\x79\x6b\x25\x1c\xbe\x4a\xd2\x75\xd4\xac\x57\x75\x0f\xbf\xd4\xfd\x0f\x7c\x81\x55\x2a\x93\x57\x72\xd7\x5a\x59\x0b\xc5\x27\xac\xc5\xac\x48\xd2\x78\xdd\x3c\x77\x48\xd3\x75\xd4\x46\xe9\x4b\xd6\x68\x4d\xf4\xab\x08\xb7\x6a\xc1\xd9\xd4\x2b\x9f\x3b\xcf\x16\xd0\x72\x00\x59\x9e\x53\xa9\x0a\x78\x3a\x84\x52\x7a\xcd\xd5\xdb\x53\x2d\xfe\xed\x20\xc5\x2b\x5d\xf2\x27\xd6\xd9\xeb\x47\xe3\x8d\xf5\x16\x89\xdc\x05\x3c\xe2\xe9\x4f\x39\x95\x63\xba\xce\xb2\x7d\xbd\x7c\x3d\xfc\x65\x5c\x8b\xf4\x95\x12\xca\x9e\x65\xfa\x7b\xdc\x14\x17\xaa\x23\x5c\x3d\x2f\x42\x87\xbc\x76\x2d\x22\x1c\x2e\x24\xae\xaa\x95\x8c\x87\x0b\x71\x67\xbd\xba\x97\xf3\x6c\x42\xab\xea\x41\xb4\x87\x1e\x16\x9c\x43\xad\xe7\x09\x98\x71\xb4\xb7\x2a\xe5\x88\xf9\xba\xcc\x29\x5f\x56\xb5\x10\x73\xc3\x12\x7a\x7b\xcd\xdc\x73\x5c\x16\xcb\x68\x8b\x7c\xd1\xb5\x50\x7d\x9d\x44\x4d\xc9\x64\xcb\xa9\x07\x63\x5f\x69\x97\x2e\x6f\x66\x62\xf6\x39\x94\x57\xb0\x4e\x76\xb6\x67\x91\x37\x54\x87\xc3\x5a\x1d\xba\x02\xed\xed\x86\x76\xa4\x7d\x59\x7b\x7e\x64\x54\x65\x39\x90\x2f\x88\x54\x3c\xa1\x9d\x00\xa9\x2b\x84\x3f\x8f\x80\xfb\x5f\x2e\x53\xb1\x9b\x2e\x9c\x90\xc2\x54\xec\x46\x86\x13\x16\x05\x0b\x5f\x56\x63\x8a\x56\x82\x7e\x44\xa4\x6f\x24\x52\xd6\x6b\x70\xbe\x17\x35\xc8\x7b\x73\xe6\xfa\xf8\x05\x37\x45\x58\x1d\xf8\x7c\x4b\x2a\xda\xcb\x73\xa8\x01\x7f\xa4\x65\x12\xa5\xce\x2a\xa8\x4c\x87\xbd\x0a\x51\x46\x62\x6f\x4f\x14\x1d\x87\x9b\x55\xac\x9c\x09\x12\x2d\x6f\xbb\x33\xeb\x15\x6d\xac\xf7\xc7\x60\xc9\xd0\x2d\x8a\x10\xe4\x08\xe4\x63\x0c\x96\x03\xd7\xfc\x5d\x1a\x77\x5d\x6a\x6a\x49\x3a\x9b\xba\xa5\xbb\x5b\x44\x57\xdb\xe0\x77\x82\x1d\xf9\x21\xec\x3a\x82\x64\x7b\xc1\x51\xfc\x90\x87\x7f\x15\x00\xcc\x82\x9e\x75\x0d\xd0\x6e\xda\xaa\x6d\x2f\x89\x2c\x0f\xbf\xd6\x5e\x1b\x94\x93\xdb\xe0\xfc\xe5\x29\x63\x48\xd5\x8c\xc8\x82\xb2\xcd\x87\x9b\xe8\xd3\xaa\x7f\xe0\xf9\xde\x88\x43\x49\xf1\x95\x1e\xa8\x26\x5c\xbc\x69\x8d\x83\x1f\x99\xa2\x4f\xf3\x0d\x4a\xcf\x57\x76\x8b\xb8\xa0\xac\x07\xea\xd8\x1d\x2c\xdc\x61\xfa\xa7\xb1\xc8\xa3\x5e\x90\x8d\x46\x18\x4f\xef\xd0\xbe\x05\xcc\xb0\x0a\xa6\xec\xe3\x81\xd8\xa9\x54\x0b\xd9\x63\x7a\x55\x5f\x65\x20\xa4\xad\xbe\xea\x91\x9f\xb2\x5b\xac\xaa\x8f\xea\xf3\xef\x89\x51\x93\xa8\x1b\x03\x46\x55\x7e\x69\xae\x55\x63\x2d\x72\x70\x61\x19\x3d\x48\x4b\xe9\x1a\x73\xcb\x9e\x12\x06\xe9\x29\x2a\x85\x9e\x66\x71\xdb\xa9\xb5\x5a\xef\x61\x62\xd1\xc8\x54\x4d\x30\xd1\x27\x20\x3f\xc3\x35\x4d\x85\x7e\x65\x56\x2e\xcf\x2c\x92\x32\x22\x97\xf7\xd4\x8f\x27\xb2\x3b\x50\x84\x1e\x50\xb4\xd2\x52\xd8\xb3\xb5\x59\xaf\xe6\xad\x9c\x30\xab\x21\xde\x68\x78\x9e\x85\x09\x7e\x2b\x8f\x70\x4c\xc0\xb8\x82\x41\x93\x65\xc8\x3d\x6d\x54\x7e\xeb\xaa\xbc\xf3\xad\x08\x09\x13\x9e\x13\xa2\x89\xc4\xe3\xd5\x81\x2d\xf5\x39\xb5\xf8\x9a\x37\x3f\xd0\x36\xa3\x9d\xaf\xb9\xdb\x6c\xc4\xbe\xad\xc3\x75\x86\x61\x8f\xc6\xbf\x65\x11\xf4\xa3\x1d\xab\x1a\xcd\x6e\xec\x9a\x71\x87\xcb\x86\xe5\x22\x06\x68\x5b\x3e\xb3\x25\x0d\xab\xf5\x1f\xc3\x51\x75\x95\x4f\xe5\x8b\x26\xcd\x1d\x94\x44\xf4\x71\x44\x00\x94\xea\xde\x5c\x07\x77\x95\xa4\x51\xb4\xe2\x28\x60\xa8\x51\x41\xaa\xc0\xde\xed\x29\xf9\xf0\xc7\xa9\x2a\x57\xd7\x32\x8d\xa2\x3b\xbd\xaa\xaf\x22\x6f\x51\x9f\x42\x40\x4b\x1a\x1e\xb6\x02\x90\x49\xf1\xe0\x89\x97\xfe\xf1\xbb\x17\x2f\xc1\xed\x53\x30\xf8\xce\x12\x3d\x45\x9b\xc6\xda\x3c\xa3\xf8\x7e\xb2\x83\x7c\x38\x41\x2e\x55\xef\x26\xad\x49\x40\x26\x56\x01\x4d\x89\xf0\x56\xda\x46\xcd\xf1\x4a\x4d\xd1\x44\xbe\x4c\x74\x8b\x88\x84\x86\x40\x35\x9d\x56\x46\xdc\xc9\x21\x2f\xc5\x0c\xd1\xae\xcb\x01\x9a\x37\xd9\xd3\x30\xc6\x3a\x3d\xe5\x10\xde\x1c\xd0\x34\x74\xae\x5f\x1e\xce\xce\xc5\xb9\x64\x74\x83\xa3\xb6\x0c\x0a\x59\x0b\x09\x13\x23\x03\xaa\xba\xdf\x1b\xd7\xfe\xa9\x2f\xf5\x3d\x3d\x90\xf5\x74\xb8\x8e\x60\xf0\x51\x62\x4e\xef\xd8\x74\x6b\x8a\x1f\xcb\x61\x2e\x27\x8d\x80\xef\xa5\xda\x41\x01\xdd\xaf\xb4\x57\xb2\x75\xcf\xf4\x92\x53\x25\x89\x8c\x20\x17\x23\x5f\xc5\xbf\x85\x5c\xb4\x3e\x8a\xdf\x20\xc6\xe0\xf0\x40\xb6\xdb\x7c\x80\xf9\x19\x82\x75\xe4\xd0\x98\xe4\xb3\xa7\xd0\xcd\x53\xe8\x66\x9f\xd0\xcd\x27\x12\x11\x74\x0a\x65\x3c\x90\xc5\x4a\x73\xa0\xcb\x30\x54\x76\x51\x07\x27\x65\x6d\x87\xc6\x27\x8a\x52\x89\x20\x5b\xc3\x37\xa0\x3f\x2b\x43\xf3\x70\x2d\x42\x77\x9f\xce\x48\xed\xff\x80\x5d\x1f\x21\xd7\x35\xef\x7e\x82\x9c\x05\xb8\xce\xe5\x01\xb2\xdf\xe3\x14\x34\xa0\x51\xcb\x11\x1a\x9e\xaa\x84\x32\x21\x0f\x47\x75\x51\xb9\x36\xc1\xf3\xf7\xaf\x96\x80\x32\x70\x7f\xb5\x7c\x71\x84\x02\x05\x0e\x4a\xfa\xad\xca\xe3\x8b\x58\x37\xb2\x54\x5a\xea\x0a\x06\xfa\x4f\x41\xb3\x77\xa6\x80\xa0\x4f\x4a\x7f\x9f\x3e\xa5\xa1\xa5\xa3\x05\x30\x11\xc1\x0e\x7e\xa5\x79\x0e\x6b\xb7\x3a\x5c\x1f\x7d\x65\x1f\xd6\x37\xa3\xef\xba\xbe\xdc\xef\x7e\xb0\x6f\xe4\x60\x2c\x8a\x3f\xe5\x0f\xf8\xc7\xe7\x5e\x5c\x8d\xe3\x87\xb9\xb7\xcd\x6f\x91\xca\x61\x0f\xeb\xea\x02\x5a\xf7\x4e\x3f\x3b\x50\xbb\x22\x4a\x99\x64\xe0\xd5\xf2\x5e\x25\x97\xbd\xbd\xbc\x19\x1c\xbe\x96\xa4\x8b\xfd\xac\x23\xee\xb1\x6a\x9d\x16\x6e\xb9\x46\x4c\x40\xec\xc7\xe0\xdb\x9c\xd6\xc6\xc5\x62\x7d\x7b\x79\x23\xd7\x76\x83\xfd\x2c\x6e\x9e\x86\x68\xcb\xa7\xef\xc7\x1b\xda\x31\xb3\x91\xa6\xa3\x1b\xf9\xcb\xc6\x59\xfd\x3d\xcd\x82\xa4\x34\x2d\x18\x8e\x15\x44\x35\x37\x9d\x6b\xca\x74\x77\x5d\x89\x4e\xec\xdf\x8d\xd4\x7d\xae\xe7\x5c\x76\xb1\x4e\x5b\xd9\x1d\xb5\xc9\x10\x28\x93\xd8\x7c\x82\x89\xf2\x62\x8b\x83\x01\x78\x9f\x0f\x70\x14\x18\x56\xd3\x2b\x56\x21\xc9\x5d\x6c\x52\x02\x8b\xc8\xce\x19\xdc\x28\x32\x54\xfe\x02\x23\x25\x79\x91\xe9\xf2\xbc\x4b\x52\xec\x60\x9f\x98\x54\xd5\xba\x57\x8c\x51\x23\x78\x87\x8f\x33\x0f\x7e\x9b\xb6\x1a\xfa\xe5\xcf\x66\x90\x29\x4b\x23\x34\x54\x98\xef\x28\x17\xf6\x0d\xa1\x8c\xec\x74\x00\x19\xe5\xc3\xad\xd3\x3f\x53\xde\x94\xb4\xa2\x25\x36\x75\xb4\x17\xc8\x0a\x5d\x6d\x77\x03\x35\x46\xb9\x8a\x88\x77\x04\x57\x74\x3f\x7a\x79\x7b\xb7\x17\x08\x31\x13\xe4\x23\x77\x34\x4a\xe5\x69\x39\xd4\xa8\xae\x69\xad\x19\x49\xf6\xb7\xa0\x80\xa5\x04\xf0\xac\x8d\x7b\xf4\x7b\x47\xa1\x1a\x5f\x69\x16\xb2\xd1\xeb\xb3\x0a\x31\x46\xd9\xb1\x72\xbb\xbf\xd6\x93\xd7\x1b\x86\xa9\x58\xf9\x3d\x97\x77\x54\xec\x1f\xcd\x2c\xd0\xf5\xae\x4a\x86\x55\x0c\x2d\xfd\xd3\x19\xc5\x4e\x54\xcb\x2e\x34\x3c\x77\x4f\xbe\x94\x5d\x8d\xcf\xf2\x48\x61\x62\x43\x1c\xa2\xa7\x7a\x7b\x4e\x71\x6b\x33\xa2\xd7\xae\x00\x13\xa3\x95\x1d\x63\xa2\x69\x78\x6c\x98\x49\x36\x8a\x2f\x1d\xa3\x2a\xd8\x5b\x11\x3e\xdb\xea\xfc\xe0\x14\xa1\x40\xa0\xf0\x15\x25\x1b\xbc\xbd\x81\x49\x37\xc8\xea\x3a\xd8\xd0\x0b\xd4\x6f\x20\x86\x89\x86\x20\xd8\x30\x1a\xab\x90\x3f\xc3\xb7\x39\x41\xbd\xa3\xc6\x10\x63\x73\x2b\xb3\x3a\x66\xd7\x1a\x8f\xb5\xed\xdb\x6a\x3d\x1e\xf6\x88\xed\x6d\xcc\xe4\x84\xab\xa5\x35\x5a\xbc\x4e\x79\x1d\x8e\x89\x74\x84\x84\x32\x71\xaa\x77\x7b\x00\xf9\xeb\x31\x9e\x3e\xa4\x31\x1c\xe1\xed\xba\xd2\xdd\xeb\x03\x53\x92\x63\xdd\x26\xae\x97\xfc\xeb\x08\xc0\xd7\x14\xe6\x4d\x8e\xd7\xbc\x01\xed\x88\x1c\x65\xa2\xe1\x64\xf3\x2b\x70\x4d\x24\x4c\x95\xc7\xe7\x97\xe5\xeb\x2b\xf0\x9c\x92\xe8\x11\xe0\x4d\x51\xa4\x1f\x2b\x5f\x9a\x00\x74\x53\x5c\xaa\x5f\x0c\x8d\x3f\xce\xfa\x4f\x1e\x4f\x31\x15\xcf\x5e\xda\xcf\x7b\x1b\xf9\x76\x71\x38\xea\x3d\xf0\x44\xbc\x5b\x40\xb6\x45\x62\x65\x0d\xa9\x79\xb7\x8e\x62\x70\xe6\xea\x9e\x48\xca\xe3\xdb\xe7\xca\xa3\xe7\xba\xdd\xbb\xbd\x78\xa9\x9c\xd5\x5e\xa6\xb3\xe2\x80\x3d\xe5\x89\x55\x33\xcc\xf1\x09\x77\x86\xe0\x1d\x25\x67\xf7\x15\x93\x43\x14\x9d\x1d\xb7\xa5\x6e\x05\xb6\x6c\x56\xbd\xda\xe9\x5f\x46\xdf\xed\xda\x73\x48\x4e\x24\x14\x9c\x6f\x77\x7b\xc9\x26\x67\x82\x80\x42\xf9\xc9\x6e\x50\xbc\xce\xdd\x49\xdd\xc1\x1d\x2d\xbd\x1a\x23\xdd\x38\x12\x2a\xd0\x2d\xf3\x05\x0a\x0a\x62\xd5\x5b\x4a\xb0\x91\xb1\x84\xb3\xf0\x0c\xf6\x7b\xd1\xdc\xe5\x36\x9c\xce\x75\x7e\x87\x44\xee\xdf\xea\x01\xde\xd2\xb9\x98\xc3\x35\x4b\xc1\xac\xff\x63\x64\x60\xde\x87\x6b\xe1\x5c\x97\x3f\xdb\x71\x6b\x2e\x45\x38\x9a\xfc\x9a\xc6\x46\xb8\xfe\x1d\xf9\xda\x8b\xe1\xb4\x46\xfc\xd8\x78\x59\x8f\x1b\x25\x63\x31\xf6\x57\x1e\x88\x0f\x44\xca\xd4\xc4\xd7\xe4\xa1\xed\xbc\xcb\x56\x92\xad\x17\x0e\x1f\x44\x1e\x30\xa3\x24\x56\xd5\x34\x20\xc3\x70\x1d\x8d\x39\x5a\x64\x2d\xe2\x98\x86\x6a\x79\x1c\x73\x63\xea\xd9\xa1\xab\x7a\x00\xdc\x01\x58\xba\xa1\x05\xaf\xc2\xbb\x9f\xf9\xfd\xbd\x54\x26\xeb\x36\x45\x34\x05\x69\xb5\x64\xe4\xed\x8e\xed\xb3\x72\xee\xce\x0a\x62\xb5\xb1\x1b\x0d\x40\xb3\x22\x36\x94\x89\xcf\xb8\x0b\x96\xf0\xf3\x58\xc6\x7a\x3c\xa5\xd5\x29\xf6\x35\xaa\xfd\x9c\xcb\x5e\xbb\x40\x57\x50\xa6\xa7\x6b\x84\xa9\x6e\x01\xf4\xe7\x35\xe2\x59\xc5\x21\x0e\x78\xd6\x77\x54\xfe\x60\x0f\xd7\x4e\x8f\x15\x2d\x2a\xb5\x2c\x3a\x42\x7a\x26\x9a\xf4\x9e\x37\xb9\xfb\x33\x48\xf8\xba\x47\xe8\x34\x33\x8e\x28\xfe\x6b\xb6\x96\x43\x86\x85\xeb\xfd\x77\x06\x86\xdf\xf5\xad\xdb\x7f\x37\xa8\x5c\x7f\x56\xc5\x6b\xd2\x3a\xfd\x63\xd2\x65\x75\xa5\x65\xf2\x1c\x8f\x5e\x37\x93\xb0\xad\x09\xa3\x92\x74\x79\x29\x40\x3d\x87\x20\x7a\x7d\xd7\x50\xc9\x68\x5a\x70\xd4\x68\x65\x1b\x29\x52\xf5\x43\x99\xa1\x04\xa8\x74\x3b\xa3\x33\xc2\xe6\xc3\xed\x25\x77\x54\xc3\xf7\xf0\x20\xd9\x03\x0d\x56\xbf\xf6\x93\x31\xed\xa5\x18\x1a\x5e\xd3\xcc\xdc\xd2\x5c\x44\x75\x19\x1e\xad\x5c\xa4\x1d\x91\xe6\x35\xad\xf7\x42\xcd\x8d\x00\x5a\x45\xbb\x43\xdf\x70\x99\x84\xff\xdb\x13\x7c\x0d\x38\x11\x68\x6e\x51\x12\xe1\x00\xf2\x4e\xa8\xe4\x0d\xf7\x01\x92\xe6\x66\x27\x96\x8f\x35\x38\xb7\x7f\xd6\xdf\x19\x2c\x6c\x7f\xf5\xd3\x66\xba\xb1\x17\x34\x2f\xd0\x65\x4a\x6e\xb7\x66\x50\xd7\xdc\x8c\x49\xaf\xb1\xdc\x28\x7e\x8c\xc9\xf6\x5b\xd4\x97\x2b\xaa\xeb\x6c\x34\xc2\x7b\x6e\xf9\x7a\xf7\xa0\xcc\x11\x93\xc4\xa9\x8c\x6d\x1a\x69\xe6\x98\x93\x06\x06\x2a\x37\x15\x9f\x89\xb9\xe9\x32\x5f\xce\x29\x81\xce\x6c\x12\xe8\xf8\x28\xec\x3a\xbe\x94\xea\x35\xef\x28\xa6\x3a\x3a\xdb\xca\x35\x6f\xcf\xb7\xc2\x57\x98\xac\x0a\x94\x1c\x33\xcf\x35\x29\x50\xf2\x58\x39\x8c\x9e\xee\x63\xf7\xe9\xdf\xa2\x9b\xc3\xd3\x68\xb8\xb8\xbc\x95\x9d\x67\xf7\xc4\xbd\xfe\xf1\xe3\x30\x91\x34\x79\x6a\xeb\x83\x3c\x5a\x3c\x8a\x8e\x98\x33\xf1\x76\x8d\x41\x36\x29\xb4\x86\x24\x4b\x75\xe7\xf6\xb0\xb4\x3b\x98\xb5\xd2\x73\x46\xf8\xf8\xa6\x92\x2e\xa7\xee\x68\x54\x1b\x79\x36\xca\x9d\x2e\x2f\x38\xba\x90\xdf\x7c\xfc\xea\xf3\x0e\x6c\xf5\x20\xc6\x7d\xd6\x14\x5f\x7b\xcc\xc2\x5d\x40\x53\x47\x78\x38\xea\xa6\x37\xd9\xa2\xda\x56\x3c\x2a\xd3\xab\x4b\xde\x04\x39\x49\x4b\xaa\x57\x9d\x07\x6c\x9c\x30\x57\x71\xa8\x2d\x12\x7d\x70\xba\xb3\xfd\x59\xda\xed\xa6\x53\x3e\x2e\x2f\xb9\x4c\x3d\xf3\x69\x96\x88\x55\xa0\x35\x03\xec\x68\x82\xc5\xc4\x65\x07\x26\xaa\x5b\x96\xac\x54\xe3\xee\x8c\x18\xa9\x07\x56\xf5\x55\x2b\x16\x06\x54\xb9\x3b\x02\x58\x9e\x75\xa5\x70\x98\xe8\x30\x42\xb0\x3e\x11\xee\x3d\x2f\xf0\x39\x01\xae\x01\x64\x7c\x64\xa2\x42\xff\xa0\x72\x04\x52\x76\x12\x33\x84\xd1\xcf\x08\x86\x9d\x1a\xbc\x61\xc4\xdd\xa9\xf6\x4a\x75\x57\x16\xdc\xff\x3b\x93\xa3\x9c\x15\xce\x14\x95\x68\x24\xfb\xa8\xdd\xde\x59\x17\x0e\x9e\x23\x12\xd0\x10\x85\x52\x31\x5c\x43\x8e\xfe\xf9\x8f\xc1\x05\x4e\xad\xa4\xa2\x0b\xdb\x43\x53\x32\xee\x9e\x35\x7a\xbd\x32\xfa\xca\x89\xee\x28\x17\x98\x6c\xcf\xf2\xda\xf9\xa0\xe2\x39\xf4\xa6\xc5\x56\xd6\xa1\x6e\x32\x0f\x98\xe3\x75\x94\x79\xba\xb4\xbb\x7e\x87\x40\x9d\x85\xf4\x89\xbb\xa2\x38\x62\x1d\xf1\xed\x79\x93\xfc\xae\x3a\x36\x9c\xdd\x6b\x4d\x76\x65\xa9\x77\x63\x40\xf7\xdc\xa7\x8c\x18\x75\x47\x7e\x5b\x51\xea\xdb\x40\x62\x78\xf2\xa5\x90\x50\xf7\x81\x97\x65\x91\x4e\xc9\x7b\xf0\x88\x7a\x19\x3e\xae\x92\x4f\xc9\xfe\x38\x23\xac\xc8\xc3\x25\x3b\x10\xa3\x78\x6c\x77\x76\x06\x58\x4a\x88\xf2\x25\xe6\x5d\x87\x56\xcf\x46\x71\x12\x41\x61\x05\x72\x54\x03\x3b\x48\x35\xd5\x76\xb7\x90\x59\x33\x48\x82\xe1\x45\x3b\x7f\xd4\xdd\xeb\x53\xdb\xcc\x34\xac\xda\xef\xd5\xc7\x93\xcb\xb7\x04\xed\x98\xb1\x2b\x79\xd8\xad\xe0\xac\x1c\xbf\x06\xc7\x66\xe5\x03\xd4\xdf\x08\xd9\x70\xd3\xd5\xfd\xed\xdb\x39\x0a\xed\x9c\xca\xb9\x83\xf4\xb6\xda\x9a\xd7\x8a\xfe\xec\xa1\xf1\x8d\x5a\x3e\x85\x4f\xc1\x5e\x30\x40\xc7\x97\x6a\xc5\x9a\x66\x06\xbf\xee\x2a\xfc\xb5\x35\xf8\x75\x45\x78\x10\xd0\x38\x81\x42\x69\xa7\x0f\x63\x4b\xf1\xff\x98\xe2\x28\xf4\xfa\x62\xbf\x96\x1e\x6f\xe0\xef\x94\x79\x28\xe1\x7f\x83\x89\x97\x71\x97\x50\x34\x8b\xa9\x31\xe3\x32\xe4\x2d\x2b\xe0\x6d\x8e\x65\xf5\xd0\xf0\x43\x08\xeb\x08\x92\xed\x85\x46\xd4\x0a\xce\xeb\xe0\xf0\x16\x94\xd7\x37\xe3\xb3\xb3\xec\xc1\x40\x89\xfb\x5e\xaa\x3c\x8d\x2f\xdd\x1c\xdb\x95\x54\xdb\x60\xd1\x5c\x89\xb5\xde\xb9\x04\x13\x18\x60\xf1\xe8\xcb\xb7\xf0\x2a\x1f\xff\x49\xd6\x19\xf5\x1d\x03\xf3\xd4\x12\x86\x1c\xc2\x0f\x54\x5f\x15\xf4\x00\x59\x72\xda\xea\x8b\x66\x01\x92\x23\x2b\xf3\xeb\x41\x9a\xe1\xe6\xe5\xdd\xc0\xcb\x03\x9a\x43\xaa\xac\xec\x88\x46\x91\xa1\xfa\x92\xe2\x9a\xae\x55\x6b\x6a\x5a\x37\x3e\x93\x2e\x8a\xd6\x68\x99\xd3\xf5\xbc\x5b\x0f\x6a\xbe\x73\xb2\x97\xe8\x5a\x52\xa2\xa6\x75\xe3\x12\xf3\x42\x12\x8e\x4b\xec\xca\x8e\x9e\x4f\xcd\xf1\x9f\x3d\x16\x6a\xb4\x6e\x2e\x00\xb4\x83\x64\x5b\x2e\x14\xa8\x38\xe1\xc1\x52\x7b\x10\xd9\x39\x17\xd9\xf2\xfe\x94\x67\x44\xb0\xb1\x3e\x73\xbd\xc2\x4e\xd0\xe8\x66\xe6\x3d\x61\x8d\xa3\x48\x42\x21\x7b\x5e\x26\x21\xf3\x3c\x29\x14\x90\x17\x7b\x68\x34\x30\x5d\x71\x4f\x01\x3c\xa4\xc4\x5f\x8b\x0e\x60\x9e\xc4\x50\x25\xa0\x1c\xe3\x09\x85\xc2\x76\x04\xc2\x26\xe9\x3a\xc2\xc1\xf0\xe1\x97\xba\xff\x61\x23\x80\x12\x86\x9b\x2d\x4b\xfa\x8d\x4b\x23\x32\xd1\x74\x1d\xb5\x61\xd2\x92\xe1\xa0\xb1\x60\xb3\x7e\x62\xb2\x52\x69\x01\x3c\xbf\x63\x51\x79\x09\xda\xd5\x18\x9f\xcb\xc8\xb4\x9c\xea\x2a\xe6\xf0\x7c\x31\x7b\x08\xdc\x61\xc6\x32\x5a\xd9\x4e\x8f\xfc\x75\xf9\x38\x9f\xd4\xb7\xf4\x32\xba\xaf\x12\xf6\x2c\x9b\x79\x51\xca\xfb\x62\xdf\x8b\x9f\x70\x24\x2a\x91\x28\x4d\x29\xbc\x2a\x10\xbd\x04\x1c\x29\x23\xe4\x46\x0f\x01\x38\x92\x13\x08\x14\x82\x4f\x58\xec\x40\x40\xe3\x18\x9e\x83\x3b\x84\xc0\x87\x45\x08\xe9\x87\x05\xc8\x76\xac\xe0\x1e\x53\x86\xcc\x2a\xd0\x46\xe9\x8e\xfc\x20\xf5\xc0\xe5\x77\x75\x3f\x5a\xfc\x91\xa2\xbc\x1e\x51\x8e\x82\x79\x0d\x46\x87\x4d\x34\x68\x18\xd5\xcd\xe9\x00\xd6\xeb\xab\xfd\x45\x55\x12\x44\xe9\x35\xa9\xf2\x10\x75\x2e\x18\xc1\x52\x64\x2d\xb5\xb0\xa3\x4f\xba\xdc\xa2\xc8\x56\xdd\x8a\x4b\xe3\xff\xe0\x55\x2f\xe5\x0d\x2b\x35\xf3\x9c\x2c\x62\x4c\x70\xac\x3e\x7d\xff\xb2\x33\x0c\x2f\x5f\x4a\x62\xbd\x87\x6b\x84\xe7\x12\xb1\x65\xc1\xed\x46\xcf\x89\xd8\xca\x75\x5e\x5d\xf9\x72\x52\xd0\x64\x63\xd6\x42\x66\xaf\xd2\x66\x5f\xb8\xdc\xa5\x6b\x2e\xb0\x48\x05\x0a\xb3\x8a\xb3\xc6\xc2\xdd\x42\x32\x53\x81\x23\x7e\xb1\x13\x22\x91\x7f\x55\xa2\x32\x4b\x8e\x71\xb1\x3f\xd5\x9e\x4b\xe3\xfa\x4a\x47\xde\xf8\x5b\x84\x35\xcb\xde\xfc\xb7\x34\x42\x87\x58\x81\x31\x4f\x85\xc3\xe6\x45\x46\x0c\x06\xab\x4a\x74\xb4\x08\xc6\x60\x87\x18\x7b\x04\xba\x59\x01\x74\x1e\xec\x50\x0c\xeb\xf3\x5b\xb7\x6f\x44\x8f\xd7\x12\x5d\xfb\x9a\x15\x76\xa8\x2f\xe5\xf2\x9f\x7d\xf9\xff\x00\x00\x00\xff\xff\xc5\x4a\x02\x6c\x8c\x64\x01\x00")

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
