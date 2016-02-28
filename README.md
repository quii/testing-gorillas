# Testing Gorillas

In this example I will illustrate the ease of testing a web server which happens to be using [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux).

I have seen a number of tutorials on the web that seem to overcomplicate this task, even introducing special testing libraries to get around the fact that Gorilla introduces something which is a little tough to unit-test.
````go
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name, exists := mux.Vars(r)["name"]

	if !exists {
		name = "world"
	}

	w.Write([]byte("Hello, " + name))
}
````

The first line in the function is how you can get the path variable from a HTTP request. The problem is it makes it problematic to unit test the handler in isolation. The following test fails because we cant hook into Gorillas path handling (at least, not easily afaik).

````go

func TestMyHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/chris", nil)
	res := httptest.NewRecorder()

	helloHandler(res, req)

	if res.Body.String() == "Hello, world" {
		t.Error("Fail! It should not use the default, it should see Chris!")
	}
}
````

The solution is very easy once you read the documentation for [`httptest.NewServer`](https://golang.org/pkg/net/http/httptest/#NewServer), which says it requires a `http.Handler`. The `mux.NewRouter` returns a `http.Handler` so you can just pass your *router with it's handlers wired up* and then you can test it as if it's a fully running web server.

So create a function which encapsulates both your routing and the handlers that run against them.

````go
func newHelloServer() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	return r
}
````

Now in your test you can call your function and pass it to `httptest.NewServer`

````go
func TestMyRouterAndHandler(t *testing.T) {
	svr := httptest.NewServer(newHelloServer())
	defer svr.Close()

	res, err := http.Get(svr.URL + "/hello/chris")

	if err != nil {
		t.Fatal("Problem calling hello server", err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if string(greeting) != "Hello, chris" {
		t.Error("Expected hello Chris but got ", greeting)
	}
}
````

Granted, this is not *unit* testing the handler's behaviours in isolation but there is an inherent coupling (with Gorilla at least), between the routing and the handling; so you may as well just accept it and test like I do in the example.

In practice, a handler should delegate it's domain logic to another service anyway (which you could then mock and inject here) so testing "handling" and "routing" together seems like a fair enough trade-off.

For me this illustrates the real strength of Gorilla, in that because it builds on the interfaces defined in the standard library you can re-use the great testing tools already available to you.
