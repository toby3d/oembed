# oEmbed
oEmbed is a format for allowing an embedded representation of a URL on third party sites. The simple API allows a website to display embedded content (such as photos or videos) when a user posts a link to that resource, without having to parse the resource directly.

## Start using telegraph
Download and install it:  
`$ go get -u gitlab.com/toby3d/oembed`

Import it in your code:  
`import "gitlab.com/toby3d/oembed"`

## Example
```go
package main

import "gitlab.com/toby3d/oembed"

var targetUrl = "https://www.youtube.com/watch?v=8jPQjjsBbIc"

func main() {
  // optional: checks what url has YouTube provider
  if !oembed.HasProvider(targetUrl) {
    return
  }

  // extract oEmbed object of source url
  data, err := oembed.Extract(targetUrl)
  if err != nil {
    // provider not found / source not found / bad response...
    panic(err)
  }

  // use data as you want
}
```
