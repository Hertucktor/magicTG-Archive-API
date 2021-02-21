package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"magicTGArchive/internal/pkg/imgHandler"
	"net/http"
)

var dbCollName = "setImages"

func returnSingleImg(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: returnSingleImg")
	var img imgHandler.Img
	var readResult bson.M
	var err error
	var imgBaseURL = "https://media.magic.wizards.com/images/featured/"

	vars := mux.Vars(r)
	setName := vars["setName"]

	readResult, err = imgHandler.SingleSetImg(setName, dbCollName)

	marshalled, err := json.Marshal(readResult)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't marshal")
	}

	if err = json.Unmarshal(marshalled, &img); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't unmarshal")
	}

	requestURL := imgBaseURL + img.PicName + "." + img.Extension

	marshalled, err = json.Marshal(requestURL)
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't marshal")
	}

	if _, err = w.Write(marshalled); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write url to user")
	}
}

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
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	//tempFileLocation := "/var/folders/mq/sb338gd91hd0psz4dx8mq2t40000gp/T/\n"
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