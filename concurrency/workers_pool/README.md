# Workers Pool

We cannot let an app create an unlimited amount of Goroutines. Goroutines are light, but the work they perform could be very heavy. A workers pool helps us to solve this problem.

With a pool of workers, we want to bound the amount of Goroutines available so that we have a deeper control of the pool of resources. This is easy to achieve by creating a channel for each worker and having workers with either an idle or busy status.

Creating a Worker pool is all about resource control: CPU, RAM, time, connections, and so on. The workers pool design pattern helps us to do the following: 

- Control access to shared resources using quotas
- Create a limited amount of Goroutines per app
- Provide more parallelism capabilities to other concurrent structures