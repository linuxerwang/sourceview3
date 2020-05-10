package sourceview

// #cgo pkg-config: gtksourceview-3.0
// #include <gtksourceview/gtksourcebuffer.h>
// #include <gtksourceview/gtksourcegutter.h>
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
		{glib.Type(C.gtk_source_buffer_get_type()), marshalSourceBuffer},
		{glib.Type(C.gtk_source_gutter_get_type()), marshalSourceGutter},
		{glib.Type(C.gtk_source_language_get_type()), marshalSourceLanguage},
		{glib.Type(C.gtk_source_language_manager_get_type()), marshalSourceLanguageManager},
		{glib.Type(C.gtk_source_view_get_type()), marshalSourceView},
	}
	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceView"] = wrapSourceView
	gtk.WrapMap["GtkSourceBuffer"] = wrapSourceBuffer
	gtk.WrapMap["GtkSourceGutter"] = wrapSourceGutter
	gtk.WrapMap["GtkSourceLanguage"] = wrapSourceLanguage
	gtk.WrapMap["GtkSourceLanguageManager"] = wrapSourceLanguageManager
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

/*
 * GtkSourceGutter
 */

// SourceGutter is a representation of GtkSourceGutter.
type SourceGutter struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceGutter.
func (v *SourceGutter) native() *C.GtkSourceGutter {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceGutter(p)
}

func marshalSourceGutter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceGutter(obj), nil
}

func wrapSourceGutter(obj *glib.Object) *SourceGutter {
	return &SourceGutter{obj}
}

/*
 * GtkSourceView
 */

// SourceView is a representation of GtkSourceView.
type SourceView struct {
	gtk.Container
}

// SetHighlightCurrentLine is a wrapper around gtk_source_view_set_highlight_current_line().
func (v *SourceView) SetHighlightCurrentLine(highlight bool) {
	C.gtk_source_view_set_highlight_current_line(v.native(), gbool(highlight))
}

// SetShowLineNumbers is a wrapper around gtk_source_view_set_show_line_numbers().
func (v *SourceView) SetShowLineNumbers(show bool) {
	C.gtk_source_view_set_show_line_numbers(v.native(), gbool(show))
}

// SetShowRightMargin is a wrapper around gtk_source_view_get_show_right_margin().
func (v *SourceView) SetShowRightMargin() {
	C.gtk_source_view_get_show_right_margin(v.native())
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

// GetGutter is a wrapper around gtk_source_view_get_gutter().
func (v *SourceView) GetGutter(wt gtk.TextWindowType) (*SourceGutter, error) {
	c := C.gtk_source_view_get_gutter(v.native(), C.GtkTextWindowType(wt))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceGutter(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceBuffer
 */

// SourceBuffer is a representation of GtkSourceBuffer.
type SourceBuffer struct {
	gtk.TextBuffer
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
	return &SourceBuffer{gtk.TextBuffer{obj}}
}

// SourceBufferNew is a wrapper around gtk_text_buffer_new().
func SourceBufferNew() (*SourceBuffer, error) {
	c := C.gtk_text_buffer_new(nil)
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SourceBufferNewWithLanguage is a wrapper around gtk_source_buffer_new_with_language().
func SourceBufferNewWithLanguage(l *SourceLanguage) (*SourceBuffer, error) {
	c := C.gtk_source_buffer_new_with_language(l.native())
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SetText is a wrapper around gtk_text_buffer_set_text().
func (v *SourceBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.asTextBuffer(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// SetLanguage is a wrapper around gtk_source_buffer_set_language().
func (v *SourceBuffer) SetLanguage(l *SourceLanguage) {
	C.gtk_source_buffer_set_language(v.native(), l.native())
}

/*
 * GtkSourceLanguageManager
 */

// SourceLanguageManager is a representation of GtkSourceLanguageManager.
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

// SourceLanguageManagerNew is a wrapper around gtk_text_buffer_new().
func SourceLanguageManagerNew() (*SourceLanguageManager, error) {
	c := C.gtk_source_language_manager_new()
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceLanguageManager(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SourceLanguageManagerGetDefault is a wrapper around gtk_source_language_manager_get_default().
func SourceLanguageManagerGetDefault() (*SourceLanguageManager, error) {
	c := C.gtk_source_language_manager_get_default()
	if c == nil {
		return nil, errNilPtr
	}

	e := wrapSourceLanguageManager(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetLanguage is a wrapper around gtk_source_language_manager_get_language().
func (v *SourceLanguageManager) GetLanguage(id string) (*SourceLanguage, error) {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_source_language_manager_get_language(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceLanguage(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceLanguage
 */

// SourceLanguage is a representation of GtkSourceLanguage.
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
