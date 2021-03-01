package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func uploadImg(w http.ResponseWriter, r *http.Request){
	log.Info().Msgf("Endpoint Hit: uploadImg")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	if err := r.ParseMultipartForm(128 << 20); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't parse a request body as multipart/form-data")
	}

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't retrieve file")
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: couldn't close file")
		}
	}()
	log.Info().Msgf("Uploaded File: %+v", handler.Filename)
	log.Info().Msgf("File Size: %+v", handler.Size)
	log.Info().Msgf("MIME Header: %+v", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp/images/upload", "upload-*.png")
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't create temp file")
	}
	defer func() {
		if err = tempFile.Close(); err != nil {
			log.Error().Timestamp().Err(err).Msg("Error: couldn't close temp file")
		}
	}()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't read from r")
	}
	// write this byte array to our temporary file
	if _, err = tempFile.Write(fileBytes); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't write bytes to temp file")
	}
	// return that we have successfully uploaded our file!
	if _, err = fmt.Fprintf(w, "Successfully Uploaded File\n"); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't write to w")
	}

}
