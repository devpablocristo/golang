https://go101.org/article/channel.html

don't (let computations) communicate by sharing memory, (let them) share memory by communicating (through channels)


https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none

Channels should usually have a size of one or be unbuffered. By default, channels are unbuffered and have a size of zero. Any other size must be subject to a high level of scrutiny. Consider how the size is determined, what prevents the channel from filling up under load and blocking writers, and what happens when this occurs.

Bad	Good
// Ought to be enough for anybody!
c := make(chan int, 64)

// Size of one
c := make(chan int, 1) // or
// Unbuffered channel, size of zero
c := make(chan int)