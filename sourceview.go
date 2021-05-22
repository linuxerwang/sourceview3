package sourceview

// #include <stdlib.h>
// #cgo pkg-config: gtksourceview-3.0
// #include <gtksourceview/gtksourcebuffer.h>
// #include <gtksourceview/gtksourcegutter.h>
// #include <gtksourceview/gtksourcelanguage.h>
// #include <gtksourceview/gtksourcelanguagemanager.h>
// #include <gtksourceview/gtksourcestyle.h>
// #include <gtksourceview/gtksourcestylescheme.h>
// #include <gtksourceview/gtksourcestyleschemechooser.h>
// #include <gtksourceview/gtksourcestyleschemechooserbutton.h>
// #include <gtksourceview/gtksourcestyleschemechooserwidget.h>
// #include <gtksourceview/gtksourcestyleschememanager.h>
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
		{glib.Type(C.gtk_source_style_get_type()), marshalSourceStyle},
		{glib.Type(C.gtk_source_style_scheme_get_type()), marshalSourceStyleScheme},
		{glib.Type(C.gtk_source_style_scheme_chooser_get_type()), marshalSourceStyleSchemeChooser},
		{glib.Type(C.gtk_source_style_scheme_chooser_button_get_type()), marshalSourceStyleSchemeChooserButton},
		{glib.Type(C.gtk_source_style_scheme_chooser_widget_get_type()), marshalSourceStyleSchemeChooserWidget},
		{glib.Type(C.gtk_source_style_scheme_manager_get_type()), marshalSourceStyleSchemeManager},
		{glib.Type(C.gtk_source_view_get_type()), marshalSourceView},
	}
	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceView"] = wrapSourceView
	gtk.WrapMap["GtkSourceBuffer"] = wrapSourceBuffer
	gtk.WrapMap["GtkSourceGutter"] = wrapSourceGutter
	gtk.WrapMap["GtkSourceLanguage"] = wrapSourceLanguage
	gtk.WrapMap["GtkSourceLanguageManager"] = wrapSourceLanguageManager
	gtk.WrapMap["GtkSourceStyle"] = wrapSourceStyle
	gtk.WrapMap["GtkSourceStyleScheme"] = wrapSourceStyleScheme
	gtk.WrapMap["GtkSourceStyleSchemeChooser"] = wrapSourceStyleSchemeChooser
	gtk.WrapMap["GtkSourceStyleSchemeChooserButton"] = wrapSourceStyleSchemeChooserButton
	gtk.WrapMap["GtkSourceStyleSchemeChooserWidget"] = wrapSourceStyleSchemeChooserWidget
	gtk.WrapMap["GtkSourceStyleSchemeManager"] = wrapSourceStyleSchemeManager
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
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
	gtk.TextView
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
	return &SourceView{gtk.TextView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}}
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

// BeginNotUndoableAction is a wrapper around gtk_source_buffer_begin_not_undoable_action().
func (v *SourceBuffer) BeginNotUndoableAction() {
	C.gtk_source_buffer_begin_not_undoable_action(v.native())
}

// EndNotUndoableAction is a wrapper around gtk_source_buffer_end_not_undoable_action().
func (v *SourceBuffer) EndNotUndoableAction() {
	C.gtk_source_buffer_end_not_undoable_action(v.native())
}

// GetMaxUndoLevels is a wrapper around gtk_source_buffer_get_max_undo_levels().
func (v *SourceBuffer) GetMaxUndoLevels() {
	C.gtk_source_buffer_get_max_undo_levels(v.native())
}

// SetMaxUndoLevels is a wrapper around gtk_source_buffer_set_max_undo_levels().
func (v *SourceBuffer) SetMaxUndoLevels(levels int) {
	C.gtk_source_buffer_set_max_undo_levels(v.native(), C.gint(levels))
}

// SetStyleScheme is a wrapper around gtk_source_buffer_set_style_scheme().
func (v *SourceBuffer) SetStyleScheme(scheme *SourceStyleScheme) {
	C.gtk_source_buffer_set_style_scheme(v.native(), scheme.native())
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

/*
 * GtkSourceStyle
 */

// SourceStyle is a representation of GtkSourceStyle.
type SourceStyle struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyle.
func (v *SourceStyle) native() *C.GtkSourceStyle {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyle(p)
}

func marshalSourceStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyle(obj), nil
}

func wrapSourceStyle(obj *glib.Object) *SourceStyle {
	return &SourceStyle{obj}
}

// Copy is a wrapper around gtk_source_style_copy().
func (v *SourceStyle) Copy() (*SourceStyle, error) {
	c := C.gtk_source_style_copy(v.native())
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceStyle(glib.Take(unsafe.Pointer(c))), nil
}

// Apply is a wrapper around gtk_source_style_apply().
func (v *SourceStyle) Apply(tag *gtk.TextTag) {
	ctag := C.toGtkTextTag(unsafe.Pointer(tag.GObject))
	C.gtk_source_style_apply(v.native(), ctag)
}

/*
 * GtkSourceStyleScheme
 */

// SourceStyleScheme is a representation of GtkSourceStyleScheme.
type SourceStyleScheme struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleScheme.
func (v *SourceStyleScheme) native() *C.GtkSourceStyleScheme {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleScheme(p)
}

func marshalSourceStyleScheme(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleScheme(obj), nil
}

func wrapSourceStyleScheme(obj *glib.Object) *SourceStyleScheme {
	return &SourceStyleScheme{obj}
}

// GetID is a wrapper around gtk_source_style_scheme_get_id().
func (v *SourceStyleScheme) GetID() (string, error) {
	c := C.gtk_source_style_scheme_get_id(v.native())
	if c == nil {
		return "", errNilPtr
	}
	gostr := goString(c)
	C.g_free(C.gpointer(c))
	return gostr, nil
}

// GetName is a wrapper around gtk_source_style_scheme_get_name().
func (v *SourceStyleScheme) GetName() (string, error) {
	c := C.gtk_source_style_scheme_get_name(v.native())
	if c == nil {
		return "", errNilPtr
	}
	gostr := goString(c)
	C.g_free(C.gpointer(c))
	return gostr, nil
}

// GetDescription is a wrapper around gtk_source_style_scheme_get_description().
func (v *SourceStyleScheme) GetDescription() (string, error) {
	c := C.gtk_source_style_scheme_get_description(v.native())
	if c == nil {
		return "", errNilPtr
	}
	gostr := goString(c)
	C.g_free(C.gpointer(c))
	return gostr, nil
}

// GetAuthors is a wrapper around gtk_source_style_scheme_get_authors().
func (v *SourceStyleScheme) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_source_style_scheme_get_authors(v.native())
	if cauthors == nil {
		return nil
	}
	for {
		if *cauthors == nil {
			break
		}
		authors = append(authors, C.GoString((*C.char)(*cauthors)))
		cauthors = C.next_gcharptr(cauthors)
	}
	return authors
}

// GetFileName is a wrapper around gtk_source_style_scheme_get_filename().
func (v *SourceStyleScheme) GetFileName() (string, error) {
	c := C.gtk_source_style_scheme_get_filename(v.native())
	if c == nil {
		return "", errNilPtr
	}
	gostr := goString(c)
	C.g_free(C.gpointer(c))
	return gostr, nil
}

// GetStyle is a wrapper around gtk_source_style_scheme_get_style().
func (v *SourceStyleScheme) GetStyle(id string) (*SourceStyle, error) {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	c := C.gtk_source_style_scheme_get_style(v.native(), cstr1)
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceStyle(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceStyleSchemeManager
 */

// SourceStyleSchemeManager is a representation of GtkSourceStyleSchemeManager.
type SourceStyleSchemeManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceStyleSchemeManager) native() *C.GtkSourceStyleSchemeManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleSchemeManager(p)
}

func marshalSourceStyleSchemeManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeManager(obj), nil
}

func wrapSourceStyleSchemeManager(obj *glib.Object) *SourceStyleSchemeManager {
	return &SourceStyleSchemeManager{obj}
}

// SourceStyleSchemeManagerNew is a wrapper around gtk_source_style_scheme_manager_new().
func SourceStyleSchemeManagerNew() (*SourceStyleSchemeManager, error) {
	c := C.gtk_source_style_scheme_manager_new()
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceStyleSchemeManager(glib.Take(unsafe.Pointer(c))), nil
}

// SourceStyleSchemeManagerGetDefault is a wrapper around gtk_source_style_scheme_manager_get_default().
func SourceStyleSchemeManagerGetDefault() (*SourceStyleSchemeManager, error) {
	c := C.gtk_source_style_scheme_manager_get_default()
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceStyleSchemeManager(glib.Take(unsafe.Pointer(c))), nil
}

// SetSearchPath is a wrapper around gtk_source_style_scheme_manager_set_search_path().
func (v *SourceStyleSchemeManager) SetSearchPath(paths []string) {
	cpaths := C.make_strings(C.int(len(paths) + 1))
	for i, path := range paths {
		cstr := C.CString(path)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cpaths, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cpaths, C.int(len(paths)), nil)
	C.gtk_source_style_scheme_manager_set_search_path(v.native(), cpaths)
	C.destroy_strings(cpaths)
}

// AppendSearchPath is a wrapper around gtk_source_style_scheme_manager_append_search_path().
func (v *SourceStyleSchemeManager) AppendSearchPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_source_style_scheme_manager_append_search_path(v.native(), (*C.gchar)(cstr))
}

// PrependSearchPath is a wrapper around gtk_source_style_scheme_manager_prepend_search_path().
func (v *SourceStyleSchemeManager) PrependSearchPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_source_style_scheme_manager_prepend_search_path(v.native(), (*C.gchar)(cstr))
}

// GetSearchPath is a wrapper around gtk_source_style_scheme_manager_get_search_path().
func (v *SourceStyleSchemeManager) GetSearchPath() []string {
	var paths []string
	cpaths := C.gtk_source_style_scheme_manager_get_search_path(v.native())
	if cpaths == nil {
		return nil
	}
	for {
		if *cpaths == nil {
			break
		}
		paths = append(paths, C.GoString((*C.char)(*cpaths)))
		cpaths = C.next_gcharptr(cpaths)
	}
	return paths
}

// GetSchemeIDs is a wrapper around gtk_source_style_scheme_manager_get_scheme_ids().
func (v *SourceStyleSchemeManager) GetSchemeIDs() []string {
	var ids []string
	cids := C.gtk_source_style_scheme_manager_get_scheme_ids(v.native())
	if cids == nil {
		return nil
	}
	for {
		if *cids == nil {
			break
		}
		ids = append(ids, C.GoString((*C.char)(*cids)))
		cids = C.next_gcharptr(cids)
	}
	return ids
}

// GetScheme is a wrapper around gtk_source_style_scheme_manager_get_scheme().
func (v *SourceStyleSchemeManager) GetScheme(id string) *SourceStyleScheme {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	c := C.gtk_source_style_scheme_manager_get_scheme(v.native(), cstr1)
	if c == nil {
		return nil
	}
	return wrapSourceStyleScheme(glib.Take(unsafe.Pointer(c)))
}

/*
 * GtkSourceStyleSchemeChooser
 */

// ISourceStyleSchemeChooser is an interface type implemented by all structs
// embedding a GtkSourceStyleSchemeChooser.  It is meant to be used as an
// argument type for wrapper functions that wrap around a C GTK function taking
// a GtkSourceStyleSchemeChooser.
type ISourceStyleSchemeChooser interface {
	toSourceStyleSchemeChooser() *C.GtkSourceStyleSchemeChooser
}

/*
 * GtkSourceStyleSchemeChooser
 */

// SourceStyleSchemeChooser is a representation of GtkSourceView's
// GtkSourceStyleSchemeChooser GInterface.
type SourceStyleSchemeChooser struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkSourceStyleSchemeChooser.
func (v *SourceStyleSchemeChooser) native() *C.GtkSourceStyleSchemeChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleSchemeChooser(p)
}

func (v *SourceStyleSchemeChooser) toSourceStyleSchemeChooser() *C.GtkSourceStyleSchemeChooser {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalSourceStyleSchemeChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeChooser(obj), nil
}

func wrapSourceStyleSchemeChooser(obj *glib.Object) *SourceStyleSchemeChooser {
	return &SourceStyleSchemeChooser{obj}
}

// GetScheme is a wrapper around gtk_source_style_scheme_chooser_get_style_scheme().
func (v *SourceStyleSchemeChooser) GetScheme() *SourceStyleScheme {
	c := C.gtk_source_style_scheme_chooser_get_style_scheme(v.native())
	if c == nil {
		return nil
	}
	return wrapSourceStyleScheme(glib.Take(unsafe.Pointer(c)))
}

// SetScheme is a wrapper around gtk_source_style_scheme_chooser_set_style_scheme().
func (v *SourceStyleSchemeChooser) SetScheme(scheme *SourceStyleScheme) {
	C.gtk_source_style_scheme_chooser_set_style_scheme(v.native(), scheme.native())
}

/*
 * GtkSourceStyleSchemeChooserButton
 */

// SourceStyleSchemeChooserButton is a representation of GtkSourceStyleSchemeChooserButton.
type SourceStyleSchemeChooserButton struct {
	gtk.Button

	SourceStyleSchemeChooser
}

// native returns a pointer to the underlying GtkSourceStyleSchemeChooserButton.
func (v *SourceStyleSchemeChooserButton) native() *C.GtkSourceStyleSchemeChooserButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toSourceStyleSchemeChooserButton(p)
}

func (v *SourceStyleSchemeChooserButton) toSourceStyleSchemeChooser() *C.GtkSourceStyleSchemeChooser {
	if v == nil {
		return nil
	}
	return C.toGtkSourceStyleSchemeChooser(unsafe.Pointer(v.GObject))
}

func marshalSourceStyleSchemeChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeChooserButton(obj), nil
}

func wrapSourceStyleSchemeChooserButton(obj *glib.Object) *SourceStyleSchemeChooserButton {
	actionable := &gtk.Actionable{obj}
	chooser := wrapSourceStyleSchemeChooser(obj)
	return &SourceStyleSchemeChooserButton{
		gtk.Button{
			gtk.Bin{
				gtk.Container{
					gtk.Widget{
						glib.InitiallyUnowned{obj},
					},
				},
			},
			actionable,
		},
		*chooser,
	}
}

/*
 * GtkSourceStyleSchemeChooserWidget
 */

// SourceStyleSchemeChooserWidget is a representation of GtkSourceStyleSchemeChooserWidget.
type SourceStyleSchemeChooserWidget struct {
	gtk.Bin

	SourceStyleSchemeChooser
}

// native returns a pointer to the underlying GtkSourceStyleSchemeChooserWidget.
func (v *SourceStyleSchemeChooserWidget) native() *C.GtkSourceStyleSchemeChooserWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toSourceStyleSchemeChooserWidget(p)
}

func (v *SourceStyleSchemeChooserWidget) toSourceStyleSchemeChooser() *C.GtkSourceStyleSchemeChooser {
	if v == nil {
		return nil
	}
	return C.toGtkSourceStyleSchemeChooser(unsafe.Pointer(v.GObject))
}

func marshalSourceStyleSchemeChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeChooserWidget(obj), nil
}

func wrapSourceStyleSchemeChooserWidget(obj *glib.Object) *SourceStyleSchemeChooserWidget {
	chooser := wrapSourceStyleSchemeChooser(obj)
	return &SourceStyleSchemeChooserWidget{
		gtk.Bin{
			gtk.Container{
				gtk.Widget{
					glib.InitiallyUnowned{obj},
				},
			},
		},
		*chooser,
	}
}

// SourceStyleSchemeChooserWidgetNew is a wrapper around gtk_source_style_scheme_chooser_widget_new().
func SourceStyleSchemeChooserWidgetNew() (*SourceStyleSchemeChooserWidget, error) {
	c := C.gtk_source_style_scheme_chooser_widget_new()
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSourceStyleSchemeChooserWidget(glib.Take(unsafe.Pointer(c))), nil
}
