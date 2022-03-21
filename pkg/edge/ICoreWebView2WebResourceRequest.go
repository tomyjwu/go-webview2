package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2WebResourceRequestVtbl struct {
	_IUnknownVtbl
	GetUri     ComProc
	PutUri     ComProc
	GetMethod  ComProc
	PutMethod  ComProc
	GetContent ComProc
	PutContent ComProc
	GetHeaders ComProc
}

type ICoreWebView2WebResourceRequest struct {
	vtbl *_ICoreWebView2WebResourceRequestVtbl
}

func (i *ICoreWebView2WebResourceRequest) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2WebResourceRequest) GetMethod() (string, error) {
	// Create *uint16 to hold result
	var _method *uint16
	res, _, err := i.vtbl.GetMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_method)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if windows.Handle(res) != windows.S_OK {
		return "", syscall.Errno(res)
	}
	// Get result and cleanup
	uri := windows.UTF16PtrToString(_method)
	windows.CoTaskMemFree(unsafe.Pointer(_method))
	return uri, nil
}

func (i *ICoreWebView2WebResourceRequest) GetUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _uri *uint16
	_, _, err = i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

func (i *ICoreWebView2WebResourceRequest) Release() error {
	return i.vtbl.CallRelease(unsafe.Pointer(i))
}
