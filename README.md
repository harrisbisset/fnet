# Custom Library for using Templ

## Basic
```go
var IndexPage = fnet.NewComponent("Index").
	View(view_index.Show()).
	Error(0, fnet.RespError("build error", view_error.DefaultBuildError())).
	Build()

func Handler(w http.ResponseWriter, req *http.Request) {
  	IndexPage.Render(w, req)
}

func main() {
  	fnet.HandleComponent(fnet.GET, "/", Handler)

  	fnet.Start("3000", "0.0.0.0", http.New)
}

```

## Handling Databases/Other Reqs (when rendering)

```go
type indexWrapper struct {
	fnet.Component
	db database.DB
}

var IndexPage = fnet.NewComponent("Index").
	View(view_index.Show()).
	Error(0, fnet.RespError("build error", view_error.DefaultBuildError())).
	Build()

var IndexWrapper indexWrapper = industryWrapper{
	Component: IndexPage,
}

func (i *indexWrapper) DB(d database.DB) *industryWrapper {
	i.db = d
	return i
}

func (i *indexWrapper) Handle(w http.ResponseWriter, req *http.Request) {
	...
	Database functions here
	...
	i.Component.Render(w, req)
}

func main() {
	db = ...
	fnet.HandleComponent(fnet.GET, "/", IndexWrapper.DB(db).Handle)
	fnet.Start("3000", "0.0.0.0")
}
```

## How to treat render errors
```go
type indexWrapper struct {
	fnet.Component
	db database.DB
}

var IndexPage = fnet.NewComponent("Index").
	View(view_index.Show()).
	Error(0, fnet.RespError("build error", view_error.DefaultBuildError())).
	Error(1, fnet.RespError("wrong user input", view_error.IndexUserError())).
	Build()

func Handler(w http.ResponseWriter, req *http.Request) {
	if ... {
		IndexPage.RenderError(0, w, req)
		return
	}

	IndexPage.Render(w, req)
}
```
