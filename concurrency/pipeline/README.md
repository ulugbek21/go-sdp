# Pipeline Pattern

When using a Pipeline pattern, we are mainly looking for the following benefits:

- We can create a concurrent structure of a multistep algorithm
- We can exploit the parallelism of multicore machines by decomposing an algorithm in different Goroutines

However, just because we decompose an algorithm in different Goroutines doesn't necessarily mean that it will execute the fastest. We are constantly talking about CPUs, so ideally the algorithm must be CPU-intensive to take advantage of a concurrent structure. The overhead of creating Goroutines and channels could make an algorithm smaller.
