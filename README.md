# gin-handler

This is a generic Gin Handler, what it means?

When you are working with gin you need to declare in a route a function to handle your route and most of the time it is not necessary because we don't use logic in this handlers so what we need is just to declare in a handler a unmarshall to our json and send to our function that will really makes something.

For me it is somthing that just takes time and most of type is boilerplate code so for these and some other reasons i've decided to create it.

Usage:

First of all when declaring a new handler we need to declare the input and output types we want and our function wich will be executed in our handler so it will be:

handlers.NewHandler[input type, output type](function to execute)

handlers.NewHandler[string, int](func)

After it we need to say if it have a query or a param values to retrieve in our request, we can do it using one of these two functions:

Query(query word)
Param(param word)

Then it is optional but we can declare a response function (if you don't want to declare we already have one that will return error 500 if it happens any error and 200 if it goes well):

Response(respondMessageString)

And then you just need to declare the method "Handle" without the "()" and it's working.
Handle

When we put it all together it looks like:

With Param:
handlers.NewHandler[int, string](messageIntToString).Param("newString").Handle

With Query:
handlers.NewHandler[int, string](messageIntToString).Query("newString").Handle

With JSON:

handlers.NewHandler[int, string](messageIntToString).Handle

With custome response:

handlers.NewHandler[string](messageString).Query("str").Response(respondMessageString).Handle)

(It doesn't matter if have a query param or nothing, the response will work with any of the options)

WE HAVE AN EXAMPLE FOLDER WITH EXAMPLES TO MAKES IT CLEAR.
