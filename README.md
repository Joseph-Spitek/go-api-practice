# go-api-practice

This repository contains a practice project that I have worked on in order to familiarize myself with GoLang. This program will read an array from the `/configs/places.json` file, then make concurrent HTTP requests to a federal weather API utilizing GoRoutines.

## Usage
Compile the program using the following command.
```bash
go build
```
Then, run the resulting file.
```bash
./go-api-practice
```

Edit the `/config/places.json` file to change how weather data is collected.