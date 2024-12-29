# Custom Library when using Templ

## Basic
```go
var IndexPage = fnet.NewComponent(fnet.WithTempl(view_index.Show()))

func main() {
  	http.HandleFunc("/", IndexPage.Render)

  	http.ListenAndServe("8080", nil)
}

```

## Handling Databases/Other Reqs (when rendering)

```go
var IndexPage = fnet.NewDataComponent(fnet.SetRender(
	func(v DB) fnet.Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			fnet.RenderTempl(view_index.Show(v), w, r)
		}
	},
))

type DB struct {
	Name string
	Conn ...
}

func main() {
	db := &DB{
		Name: "Database Name"
	}

	http.HandleFunc("/", IndexPage.RenderWithData(db))

  	http.ListenAndServe("8080", nil)
}
```
