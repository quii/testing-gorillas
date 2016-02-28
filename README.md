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

The solution is very easy once you realise that `mux.NewRouter` just returns a `http.Handler`, which you can then call with a `http.Request` and a `httptest.ResponseRecorder` (credit to [MistakenForYeti](https://www.reddit.com/user/MistakenForYeti) for pointing this out).

So just create a function which creates your router and wires up your handlers:

````go
func newHelloServer() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	return r
}
````

Now in your test you can call your complete server and make assertions.

````go
func TestMyRouterAndHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/chris", nil)
	res := httptest.NewRecorder()
	newHelloServer().ServeHTTP(res, req)

	if res.Body.String() != "Hello, chris" {
		t.Error("Expected hello Chris but got ", res.Body.String())
	}
}
````

Granted, this is not *unit* testing the **handler's** behaviours in isolation but there is an inherent coupling (with Gorilla at least), between the routing and the handling; so you may as well just accept it and test like I do in the example.

In practice, a handler should delegate it's domain logic to another service anyway (which you could then mock and inject here) so testing "handling" and "routing" together seems like a fair enough trade-off.

For me this illustrates the real strength of Gorilla, in that because it builds on the interfaces defined in the standard library you can re-use the great testing tools already available to you.
