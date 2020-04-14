# cc-validation



## Installation

1. install `Go` (version 1.13+ is required)

2. Extract zip file to Go directory

3. install sqlite3

3. Install Postman ( for API testing )

## Quick Start

1. open your CLI ( command line interface )

2. navigate to your extracted zip and locate `main.go` in the root folder

3. Run this command to initialize a Local connection for the API

```
    $ go run main.go
    time="2020-04-14T20:29:20+08:00" level=info msg="serving api at http://127.0.0.1:8006"
```
4. Open any API testing tools ( e.g. Postman )

5. Using "POST" method key-in this url in your api testing tool address bar `http://127.0.0.1:8006/piglatin`.
6. If successfull response should be like this

```
  {
    "data": {
        "alt_translation": "",
        "original_text": "",
        "translated_text": ""
    }
  }
```
