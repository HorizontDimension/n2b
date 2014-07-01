package models

import (
	"github.com/HorizontDimension/twiit"
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"mime/multipart"
)

func FilesFs(s *mgo.Session) *mgo.GridFS {
	s.SetMode(mgo.Monotonic, true)
	return s.DB("n2b").GridFS("fs")
}

//Session must be closed outside this func
func AddFile(s *mgo.Session, fhs *multipart.FileHeader, fileId bson.ObjectId) error {
	//iterate over multiapartFileHeader
	f, err := fhs.Open()
	if err != nil {
		twiit.Log.Error("Failed to open multipart", "error", err)
		return err
	}

	gridfile, err := FilesFs(s).Create("")
	if err != nil {
		twiit.Log.Error("Failed to create gridfile", "error", err)
		return err
	}

	gridfile.SetId(fileId)
	gridfile.SetContentType(fhs.Header.Get("Content-Type"))

	_, err = io.Copy(gridfile, f)

	//please handle error ws notification
	if err != nil {
		twiit.Log.Error("Failed to copy image do gridfile", "error", err)

		return err
	}

	err = gridfile.Close()
	if err != nil {
		twiit.Log.Error("Failed to close gridfile", "error", err)
		return err
	}

	err = f.Close()
	if err != nil {
		twiit.Log.Error("Failed to close multipart file", "error", err)
		return err
	}

	return nil
}
