# Barrier Pattern

Its purpose is simple – put up a barrier so that nobody passes until we have all the results we need, something quite common in concurrent applications.

Imagine the situation where we have a microservices application where one service needs to compose its response by merging the responses of another three microservices. This is where the Barrier pattern can help us.

Our Barrier pattern could be a service that will block its response until it has been composed with the results returned by one or more different Goroutines (or services). And what kind of primitive do we have that has a blocking nature? Well, we can use a lock, but it's more idiomatic in Go to use an unbuffered channel. (Contreras, 2017)

Check out the following link as well to see how the author implemented a reusable barrier, which is implemented very cool:

https://medium.com/golangspec/reusable-barriers-in-golang-156db1f75d0b