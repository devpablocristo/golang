Recursion:
Since recursion quickly scales horizontally, one needs a consistent way to fill the finite list of workers with work but also ensure when workers are freed up, that they quickly pick up work from other (over-worked) workers.

Rather than create a manager layer, employ a cooperative peer system of workers:

1. each worker shares a single inputs channel
2. before recursing on inputs (subIinputs) check if any other workers are idle
3. if so, delegate to that worker
4. if not, current worker continues recursing that branch
5. With this algorithm, the finite count of workers quickly become saturated with work. 
   1. Any workers which finish early with their branch - will quickly be delegated a sub-branch from another worker. 
   2. Eventually all workers will run out of sub-branches, at which point all workers will be idled (blocked) and the recursion task can finish up.

Some careful coordination is needed to achieve this. Allowing the workers to write to the input channel helps with this peer coordination via delegation. A "recursion depth" WaitGroup is used to track when all branches have been exhausted across all workers.

(To include context support and error chaining - I updated your getSubInputs function to take a ctx and return an optional error):