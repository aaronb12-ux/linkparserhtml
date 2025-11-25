# HTML Link Parser

A Go package that parses HTML files and extracts all links (`<a>` tags) with their href attributes and text content.

## What It Does

Extracts links from HTML and returns:
- The `href` attribute value
- All text content inside the link (including text nested in other tags)

## Example

Given this HTML:
```html

  Something in a span
  Text not in a span
  Bold text!

```

Returns:
```go
Link{
  Href: "/dog",
  Text: ["Something in a span", "Text not in a span", "Bold text!"]
}
```

## How to Run
```bash
go get golang.org/x/net/html
go run main.go
```

## Implementation Notes

- Uses `golang.org/x/net/html` package for HTML parsing
- Traverses the HTML document tree using `Descendants()`
- Filters for text nodes (`html.TextNode`) to extract actual content
- Handles nested HTML elements inside links
