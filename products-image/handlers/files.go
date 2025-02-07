package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/abhiraj-ku/micro_serGO/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Files struct
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// New file creates a new file hanlder
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// ServeHTTP
func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r)
}

// saveFile saves the content of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("unable to save the file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
