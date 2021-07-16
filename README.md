# magicTGArchive
##Creators:
API | Database | Backend | Frontend :   https://github.com/Hertucktor & Android App: https://github.com/HimeUrikohime

##Overview - What is this project about?
This Project will make storing digital copies of Magic: The Gathering cards as easy as opening a new booster pack.
You will be able to store your own collection of mtg cards and build decks with these cards and share your collection information with your friends.

##How does it work (an abbreviation)
Using http requests to the internal REST APIs will make you able to manipulate your own mtg collection (for further API Usage Info read docs/api)
The Frontend will allow you to manage your collection in a user-friendly and accessible way.
The Android App helps you by inserting your cards to your collection. You will be able to take a picture, this image will be sent
to an internal API and compared to an image dataset, furthermore identifying your card and inserting it into your collection.

##How to run the current project

run docker localy (you will need docker desktop : https://www.docker.com/products/docker-desktop )
To start the database:
```shell
make db 
```
after your db is up and running execute the card importer (this will take up to an hour)
```shell
make importer
```

finally your database is running and filled with all needed information to start the main event
```shell
make api
```

###Configure Credentials for MongoDB in Docker
If you don't want to run your app with the default credentials change them like so:
```shell
nano init-mongo.js
```

If you want to change the database config
```shell
nano config.yml
```
Test
