# Custom Library for using Templ
```

var IndexPage = fnet.NewComponent("Index").
	View(view_index.Show()).
	Error(view_error_response.Page()).
	Build()

func Handler(w http.ResponseWriter, req *http.Request) {
  IndexPage.Render(w, req)
}

func main() {
  fnet.HandleComponent(fnet.GET, "/", Handler)
  
  fnet.Start("3000", "0.0.0.0")
}

```
