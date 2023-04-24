# gin-handler

This is a generic Gin Handler, what it means?

When you are working with gin you need to declare in a route a function to handle your route and most of the time it is not necessary because we don't use logic in this handlers so what we need is just to declare in a handler an unmarshall to our json and send to our function that will really make something.

It is something that just takes time and most of the time is boilerplate code so for these and some other reasons I've decided to create it for my use and others witch need something simple to use.

Usage:

First of all when declaring a new handler we need to declare the input and output types we want and our function witch will be executed in our handler, and say if we have error or success custom handlers, and if we don't have we need to pass nil, so it will be like:

- handlers.NewHandler\[input type, output type](nil,nil)

- handlers.NewHandler\[string, int](nil,nil)

Then we need to say if it has a query or a param nor a json value to retrieve in our request, so we will call one of the three pre-defined functions to do it:

- HandleParam\("param to get value", function to handle)

- HandleQuery\("query to get value", function to handle)

- HandleJSON(function to handle)

When we put it all together it looks like:

With Param:
handlers.NewHandler\[int, string](nil,nil).HandleParam("newString", func(int) (string,error))

With Query:
handlers.NewHandler\[int, string](nil,nil).HandleQuery("newString", func(int) (string,error))

With JSON:

handlers.NewHandler\[customType, customType](func(customType) (customType, error))

Of course, you can declare a variable when you have all of the handlers with the same type and just call the variable with respective handler function.

WE HAVE AN EXAMPLE FOLDER WITH EXAMPLES TO MAKE IT CLEARER.
