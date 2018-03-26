Repository
===

Find all Git repository under the directory.

Usage
---

```go
    repos, err := reposioty.Get(basePath)
    
    if err != nil {
        panic(err)
    }
    
    for _, repo := range repos {
        fmt.Println(repo)
    }
```

Installation
---

```bash
go get -u github.com/sawadashota/reposioty
```

License
---

MIT

Author
---

Shota Sawada