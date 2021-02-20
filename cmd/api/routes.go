package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func handleRequests(){
	var port = ":8080"
	log.Info().Msgf("Starting API on port:", port)

	myRouter := mux.NewRouter().StrictSlash(true)

	//Status calls
	status := myRouter.PathPrefix("/status").Subrouter()
	status.HandleFunc("/alive",statusAlive).Methods(http.MethodGet)
	status.HandleFunc("/check",statusCheck).Methods(http.MethodGet)

	//CRUD Operations for card info
	api := myRouter.PathPrefix("/api").Subrouter()
	api.HandleFunc("/card", createNewCardEntry).Methods(http.MethodPost)
	api.HandleFunc("/card/all", returnAllCardEntries).Methods(http.MethodGet)
	api.HandleFunc("/card/number/{number}/set/name/{setName}", returnSingleCardEntry).Methods(http.MethodGet)
	api.HandleFunc("/card/number/{number}/set/name/{setName}", updateSingleCardEntry).Methods(http.MethodPut)
	api.HandleFunc("/card/number/{number}/set/name/{setName}", deleteSingleCardEntry).Methods(http.MethodDelete)

	//CRUD Operations for img info
	api.HandleFunc("/img/set/name/{setName}", returnSingleImg).Methods(http.MethodGet)
	api.HandleFunc("/img/upload", upload).Methods(http.MethodPost)

	//Open http connection
	if err := http.ListenAndServe(port, myRouter); err != nil {
		log.Panic().Timestamp().Err(err).Msg("Panic: problem with TCP network connection")
	}
}

func upload(w http.ResponseWriter, r *http.Request){
	log.Info().Msgf("Endpoint Hit: upload with method:",r.Method)

	if r.Method == "GET" {
		cruTime := time.Now().Unix()
		hash := md5.New()
		if _, err := io.WriteString(hash, strconv.FormatInt(cruTime, 10)); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't write s to w")
		}
		token := fmt.Sprintf("%x", hash.Sum(nil))

		newTemplate, err := template.ParseFiles("upload.gtpl")
		if err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't create template")
		}
		if err = newTemplate.Execute(w, token); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't execute template")
		}
	}else {
			if err := r.ParseMultipartForm(32 << 20); err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't parse multi part form")
			}

			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't provide file")
			}

			defer func() {
				if err = file.Close(); err != nil {
					log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't close file")
				}
			}()

			if _, err = fmt.Fprintf(w, "%v", handler.Header); err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't format handler.Header")
			}

			f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't open file")
			}

			defer func() {
				if err = f.Close(); err != nil {
					log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't close f")
				}
			}()

			if _, err = io.Copy(f,file); err != nil {
				log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't copy file to f")
			}

		}

}