package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	sourceview "github.com/linuxerwang/sourceview3"
)

const txt = `# GtkSourceView3

- GtkSourceBuffer
- GtkSourceView
- GtkSourceGutter
- GtkSourceLanguage
- ... and **many** others
`

func main() {
	gtk.Init(nil)

	testBuilder()
	// testManual()
	gtk.Main()
}

func testBuilder() {
	builder, _ := gtk.BuilderNew()
	builder.AddFromString(glade)

	win := extractWindow(builder, "window")
	win.ShowAll()

	sv := extractSourceView(builder, "sv")
	sv.ShowAll()

	lm, _ := sourceview.SourceLanguageManagerGetDefault()
	l, _ := lm.GetLanguage("markdown")
	buf, _ := sv.GetBuffer()
	buf.SetLanguage(l)
	buf.SetText(txt)
}

func extractWindow(builder *gtk.Builder, id string) *gtk.Window {
	obj, err := builder.GetObject(id)
	if err != nil {
		fmt.Printf("failed to extract object with id %s\n", id)
		return nil
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		fmt.Printf("object %s is not a *gtk.Window\n", id)
		return nil
	}
	return window
}

func extractPaned(builder *gtk.Builder, id string) *gtk.Paned {
	obj, err := builder.GetObject(id)
	if err != nil {
		fmt.Printf("failed to extract object with id %s\n", id)
		return nil
	}

	p, ok := obj.(*gtk.Paned)
	if !ok {
		fmt.Printf("object %s is not a *gtk.Paned\n", id)
		return nil
	}
	return p
}

func extractSourceView(builder *gtk.Builder, id string) *sourceview.SourceView {
	obj, err := builder.GetObject(id)
	if err != nil {
		fmt.Printf("failed to extract object with id %s\n", id)
		return nil
	}

	sv, ok := obj.(*sourceview.SourceView)
	if !ok {
		fmt.Printf("object %s is not a *sourceview.SourceView\n", id)
		return nil
	}
	return sv
}

func testManual() {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	sv, _ := sourceview.SourceViewNew()

	win.Add(sv)
	win.SetDefaultSize(800, 600)
	win.ShowAll()
}

const glade = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.22.1 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <requires lib="gtksourceview" version="3.0"/>
  <object class="GtkWindow" id="window">
    <property name="can_focus">False</property>
    <property name="default_width">800</property>
    <property name="default_height">600</property>
    <child>
      <placeholder/>
    </child>
    <child>
      <object class="GtkScrolledWindow">
        <property name="visible">True</property>
        <property name="can_focus">True</property>
        <property name="shadow_type">in</property>
        <child>
          <object class="GtkSourceView" id="sv">
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="pixels_above_lines">5</property>
            <property name="pixels_below_lines">5</property>
            <property name="wrap_mode">word-char</property>
            <property name="left_margin">2</property>
            <property name="right_margin">2</property>
            <property name="monospace">True</property>
            <property name="show_line_numbers">True</property>
            <property name="show_line_marks">True</property>
            <property name="tab_width">4</property>
            <property name="auto_indent">True</property>
            <property name="show_right_margin">True</property>
            <property name="smart_home_end">always</property>
            <property name="highlight_current_line">True</property>
            <property name="draw_spaces">GTK_SOURCE_DRAW_SPACES_TAB | GTK_SOURCE_DRAW_SPACES_NEWLINE | GTK_SOURCE_DRAW_SPACES_NBSP | GTK_SOURCE_DRAW_SPACES_LEADING | GTK_SOURCE_DRAW_SPACES_TRAILING</property>
            <property name="smart_backspace">True</property>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
