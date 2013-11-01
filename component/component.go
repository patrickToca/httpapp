package component

import (
	"net/http"
)

func ServeFiles(root string) *FileServerComponent {
	return &FileServerComponent{root}
}

func URLMapper() *URLMapComponent {
	return &URLMapComponent{http.DefaultServeMux}
}

func Redirect(location string) *RedirectComponent {
	return &RedirectComponent{location}
}
