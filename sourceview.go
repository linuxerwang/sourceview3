package sourceview

// #cgo pkg-config: gtksourceview-3.0
// #include <gtksourceview/gtksourcebuffer.h>
// #include <gtksourceview/gtksourcelanguage.h>
// #include <gtksourceview/gtksourcelanguagemanager.h>
// #include <gtksourceview/gtksourceview.h>
// #include "sourceview.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var errNilPtr = errors.New("cgo returned unexpected nil pointer")

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_view_get_type()), marshalSourceView},
		{glib.Type(C.gtk_source_buffer_get_type()), marshalSourceBuffer},
		{glib.Type(C.gtk_source_language_get_type()), marshalSourceLanguage},
		{glib.Type(C.gtk_source_language_manager_get_type()), marshalSourceLanguageManager},
	}
	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceView"] = marshalSourceView
	gtk.WrapMap["GtkSourceBuffer"] = marshalSourceBuffer
	gtk.WrapMap["GtkSourceLanguage"] = marshalSourceLanguage
	gtk.WrapMap["GtkSourceLanguageManager"] = marshalSourceLanguageManager
}

type SourceView struct {
	gtk.Container
}

// native returns a pointer to the underlying GtkSourceView.
func (v *SourceView) native() *C.GtkSourceView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceView(p)
}

// native returns a pointer to the underlying GtkSourceView.
func (v *SourceView) asTextView() *C.GtkTextView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextView(p)
}

func marshalSourceView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceView(obj), nil
}

func wrapSourceView(obj *glib.Object) *SourceView {
	return &SourceView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}
}

// SourceViewNew is a wrapper around gtk_source_view_new().
func SourceViewNew() (*SourceView, error) {
	c := C.gtk_source_view_new()
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

func SourceViewNewWithBuffer(buffer *SourceBuffer) (*SourceView, error) {
	c := C.gtk_source_view_new_with_buffer(buffer.native())
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

// GetBuffer is a wrapper around gtk_source_view_get_buffer().
func (v *SourceView) GetBuffer() (*SourceBuffer, error) {
	c := C.gtk_text_view_get_buffer(v.asTextView())
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceBuffer(glib.Take(unsafe.Pointer(c))), nil
}

type SourceBuffer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceBuffer.
func (v *SourceBuffer) native() *C.GtkSourceBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceBuffer(p)
}

// native returns a pointer to the underlying GtkSourceBuffer.
func (v *SourceBuffer) asTextBuffer() *C.GtkTextBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextBuffer(p)
}

func marshalSourceBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceBuffer(obj), nil
}

func wrapSourceBuffer(obj *glib.Object) *SourceBuffer {
	return &SourceBuffer{obj}
}

// SourceBufferNew() is a wrapper around gtk_text_buffer_new().
func SourceBufferNew() (*SourceBuffer, error) {
	c := C.gtk_text_buffer_new(nil)
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

func SourceBufferNewWithLanguage(l *SourceLanguage) (*SourceBuffer, error) {
	c := C.gtk_source_buffer_new_with_language(l.native())
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

func (v *SourceBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.asTextBuffer(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

func (v *SourceBuffer) SetLanguage(l *SourceLanguage) {
	C.gtk_source_buffer_set_language(v.native(), l.native())
}

type SourceLanguageManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceLanguageManager.
func (v *SourceLanguageManager) native() *C.GtkSourceLanguageManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceLanguageManager(p)
}

func marshalSourceLanguageManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceLanguageManager(obj), nil
}

func wrapSourceLanguageManager(obj *glib.Object) *SourceLanguageManager {
	return &SourceLanguageManager{obj}
}

// SourceLanguageManagerNew() is a wrapper around gtk_text_buffer_new().
func SourceLanguageManagerNew() (*SourceLanguageManager, error) {
	c := C.gtk_source_language_manager_new()
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceLanguageManager(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

func SourceLanguageManagerGetDefault() (*SourceLanguageManager, error) {
	c := C.gtk_source_language_manager_get_default()
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceLanguageManager(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

func (v *SourceLanguageManager) GetLanguage(id string) (*SourceLanguage, error) {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_source_language_manager_get_language(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceLanguage(glib.Take(unsafe.Pointer(c))), nil
}

type SourceLanguage struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceLanguageManager.
func (v *SourceLanguage) native() *C.GtkSourceLanguage {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceLanguage(p)
}

func marshalSourceLanguage(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceLanguage(obj), nil
}

func wrapSourceLanguage(obj *glib.Object) *SourceLanguage {
	return &SourceLanguage{obj}
}
