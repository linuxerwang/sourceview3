#include <gtk/gtk.h>

static GtkTextView *
toGtkTextView(void *p)
{
	return (GTK_TEXT_VIEW(p));
}

static GtkTextBuffer *
toGtkTextBuffer(void *p)
{
	return (GTK_TEXT_BUFFER(p));
}

static GtkSourceView *
toGtkSourceView(void *p)
{
	return (GTK_SOURCE_VIEW(p));
}

static GtkSourceBuffer *
toGtkSourceBuffer(void *p)
{
	return (GTK_SOURCE_BUFFER(p));
}

static GtkSourceLanguageManager *
toGtkSourceLanguageManager(void *p)
{
	return (GTK_SOURCE_LANGUAGE_MANAGER(p));
}

static GtkSourceLanguage *
toGtkSourceLanguage(void *p)
{
	return (GTK_SOURCE_LANGUAGE(p));
}