# Custom Library when using Templ

## Basic
```go
var IndexPage = fnet.NewComponent("Index").
	View(view_index.Show()).
	Error(0, fnet.NewError("build error", view_error.DefaultBuildError()).Build()).
	Build()

func Handler(w http.ResponseWriter, req *http.Request) fnet.Handler {
  	return IndexPage.Render(w, req)
}

func main() {
  	fnet.HandleComponent(fnet.GET, "/", Handler)

  	fnet.Start("80", "0.0.0.0", nil)
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
	Error(0, fnet.NewError("build error", view_error.DefaultBuildError()).Build()).
	Build()

var IndexWrapper indexWrapper = industryWrapper{
	Component: IndexPage,
}

func (i *indexWrapper) DB(d database.DB) *industryWrapper {
	i.db = d
	return i
}

func (i *indexWrapper) Handle(w http.ResponseWriter, req *http.Request) fnet.Handler {
	...	Database functions here ...
	return i.Component.Render(w, req)
}

func main() {
	db = ...
	fnet.HandleComponent(fnet.GET, "/", IndexWrapper.DB(db).Handle)
	fnet.Start("80", "0.0.0.0", nil)
}
```

## How to Render Errors
```go
var IndexPage = fnet.NewComponent("Index").
	View(view_index.Show()).
	Error(0, fnet.NewError("build error", view_error.DefaultBuildError()).Build()).
	Error(1, fnet.NewError("wrong user input", view_error.IndexUserError()).Build()).
	Build()

func Handler(w http.ResponseWriter, req *http.Request) {
	if ... {
		return IndexPage.RenderError(0, w, req)
	}

	if ... {
		return IndexPage.RenderError(1, w, req)
	}

	return IndexPage.Render(w, req)
}
```
