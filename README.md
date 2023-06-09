# Mike-Go-Rest

This is an application for me to test playing with Golang and a REST API.

## Prerequisites

Create an app within [Spotify](https://developer.spotify.com/dashboard) and get the ```CLIENT_ID``` and ```CLIENT_SECRET```

## Usage

Set ```CLIENT_ID``` and ```CLIENT_SECRET``` as environment variables.

```go run main.go <path to text file>```

The text file can be in these formats, one per line.

``` 
song - artist
song
song-artist
```