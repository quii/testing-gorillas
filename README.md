# Testing Gorillas

In this example I will illustrate the ease of testing a web server which happens to be using [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux).

I have seen a number of tutorials on the web that seem to overcomplicate this task, even introducing special testing libraries to get around the fact that gorilla introduces something which is a little tough to unit-test.
````go
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"] // where does this come from and how can i set it in tests?
	w.Write([]byte("Hello, " + name))
}
````

The solution is very easy once you read the documentation for [`httptest.NewServer`](https://golang.org/pkg/net/http/httptest/#NewServer), which says it requires a `http.Handler`. The `mux.NewRouter` returns a `http.Handler` so you can just pass your *router with it's handlers wired up* rather than your handlers and then you can test it as if it's a fully running web server.

Granted, this is not *unit* testing the handler's behaviours in isolation but there is an inherent coupling (in Go at least), between the routing and the handling; so you may as well just accept it and test like I do in the example.

In practice, a handler should delegate it's domain logic to another service anyway so testing "handling" and "routing" together seems like a fair enough trade-off.
