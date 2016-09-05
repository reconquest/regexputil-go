# regexputil-go

This package allows to extract subexp names after regexp.FindStringSubmatch.

## Example:

```go
re := regexp.MustCompile(`(?P<head>[a-z]+)(?P<body>[0-9]+)`)
matches := re.FindStringSubmatch("aaa999")
if len(matches) > 0 {
    head := regexputil.Subexp(re, matches, "head")
    body := regexputil.Subexp(re, matches, "body")
    fmt.Println("<head>:", head)
    fmt.Println("<body>:", body)
}
```

```
<head>: aaa
<body>: 999
```
