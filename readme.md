# VideoCrawler

A fast webcrawler for finding persistant video links embedded on a webpage written in Go.

VideoCrawler provides a clean method for finding video links that are embedded on a webpage.

Just provide a webpage containing a video (ex: https://archive.org/details/wayback-machine-intro), and the webcrawler will parse the webapge HTML to find a '.mp4' link. 
**NOTE:** This means that library will only work correctly if there is an 'mp4' file referenced directly in the HTML (for now).

**This library is able to parse dynamic webpages that are rendered by javascript using the power of chromedp** (https://github.com/chromedp/chromedp)

## Example
```
package main

import (
	"fmt"

	"github.com/mastrofran/videocrawler"
)

func main() {

	url := "https://archive.org/details/wayback-machine-intro"
	link := videocrawler.GetVideo(url)
	fmt.Println(link)
}
```

## Installation

Add VideoCrawler to your project:

```
go get github.com/mastrofran/videocrawler@latest
```

and import it to your project:
```
import (

	.
    .
    .
	"github.com/mastrofran/videocrawler"
    
)
```

## Notes

- VideoCrawler is a lightweight library that performs one function and one function only: to scrape mp4 files embedded directly in a webpage HTML. It takes a url (string) as the only argument and returns an array (with elements of type string) of video links found on the given url.

- This library currently does not work on websites that deliver videos using the blob (Binary Large OBject) protocol, nor will it parse videos that are rendered via JW Player (for now).

- This is a work in progress. Please open an issue if something seems to be behaving incorrectly.
