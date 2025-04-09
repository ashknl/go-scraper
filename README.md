# Go-Scraper

Go scraper is an api for scraping google serp results.

## Usage
- Build the project: 
  ```
  go build .
  ```

- start the executable
- Send a `GET` request to `/v1/api` with the following request parameters in request body:
```
{
    "term":<search term> <string>,
    "cc_code": <ISO country code> <string>,
    "lang_code": <ISO language code> <string>,
    "pages":"no of pages to scrape" <int>,
    "count":"no of results per page" <int>,
    "backoff": "Time for expontential backoff" <int>
}


## Todo
- [] : Core scraping feature


```