# [Sync](https://pkg.go.dev/sync)

## [WaitGroup](https://pkg.go.dev/sync#WaitGroup)

- [src/features/sync/waitgroup/waitgroup.go](../src/features/sync/waitgroup/waitgroup.go)
- [src/features/sync/waitgroup/loop.go](../src/features/sync/waitgroup/loop.go)

```go
var wg sync.WaitGroup
wg.Add(1)
wg.Done()
wg.Wait()
```

## [Mutex](https://pkg.go.dev/sync#Mutex) & [RWMutex](https://pkg.go.dev/sync#RWMutex)

Wikipedia: [Mutual exclusion](https://en.wikipedia.org/wiki/Mutual_exclusion)

- [src/features/sync/mutex/mutex.go](../src/features/sync/mutex/mutex.go)

```go
var lock sync.Mutex
var wg sync.WaitGroup

wg.Add(1)
lock.Lock()
count++ // critical section, bottleneck
lock.Unlock()
wg.Wait()
```

- [src/features/sync/mutex/rwmutex.go](../src/features/sync/mutex/rwmutex.go)
- [src/features/sync/mutex/simple_rwmutex.go](../src/features/sync/mutex/simple_rwmutex.go)

```go
var m sync.RWMutex
var wg sync.WaitGroup

wg.Add(2)
m.Lock()
m.Unlock()
m.RLocker().Lock()
m.RLocker().Unlock()
wg.Wait()
```

```bash
Readers  RWMutext      Mutex
1        40.907µs      21.196µs
2        25.871µs      4.505µs
4        5.412µs       2.92µs
8        11.241µs      5.123µs
16       22.959µs      37.643µs
32       54.154µs      58.339µs
64       39.508µs      44.931µs
128      53.686µs      63.769µs
256      213.584µs     120.785µs
512      230.939µs     159.073µs
1024     266.688µs     247.973µs
2048     513.899µs     500.5µs
4096     1.044435ms    972.898µs
8192     2.101028ms    1.912819ms
16384    4.433551ms    3.921538ms
32768    8.347544ms    7.30968ms
65536    17.62768ms    15.897422ms
131072   33.773806ms   28.726714ms
262144   66.001342ms   55.079544ms
524288   135.539631ms  113.173349ms
```

## [Cond](https://pkg.go.dev/sync#Cond)

Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.

- [src/features/sync/cond/cond.go](../src/features/sync/cond/cond.go): signal
- [src/features/sync/cond/broadcast.go](../src/features/sync/cond/broadcast.go)
- [src/features/sync/cond/signal.go](../src/features/sync/cond/signal.go)

```go
c := sync.NewCond(&sync.Mutex{})
var isTrue = false

go func() {
  isTrue = true
  fmt.Println("Yes")
  c.Signal()
}()

c.L.Lock()
for isTrue == false {
  c.Wait()
}
c.L.Unlock()
```

## [Once](https://pkg.go.dev/sync#Once)

How often Go itself uses `Once`:

```bash
grep -ir sync.Once $(go env GOROOT)/src | wc -l

128
```

- [src/features/sync/once/once.go](../src/features/sync/once/once.go)

```go
var once sync.Once

once.Do(func() {
  fmt.Println("One")
})
once.Do(func() {
  fmt.Println("Two")
})
```

`One`

## [Pool](https://pkg.go.dev/sync#Pool)

- [src/features/sync/pool/pool.go](../src/features/sync/pool/pool.go)
- [src/features/sync/pool/memory.go](../src/features/sync/pool/memory.go)
- [src/features/sync/pool/server_test.go](../src/features/sync/pool/server_test.go)
- [src/features/sync/pool/pool-server_test.go](../src/features/sync/pool/pool-server_test.go)

```go
myPool := &sync.Pool{
  New: func() interface{} {
    return struct{}{}
  },
}

instance := myPool.Get()
myPool.Put(instance)
```

### Benchmark

```bash
go test -bench=. -benchtime=10s server_test.go

BenchmarkNetworkRequest-8             10        1005006833 ns/op
PASS
ok      command-line-arguments  11.066s
```

```bash
go test -bench=. -benchtime=10s pool-server_test.go

BenchmarkNetworkRequest-8           5800           6536593 ns/op
PASS
ok      command-line-arguments  51.041s
```
