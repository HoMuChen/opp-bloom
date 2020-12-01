opp-bloom
========

This is a opposite to bloom filter implemented by hash table without dealing with hash collision.

Opposite to a bloom filter.False negative matched are possible, but false positives are not.

In other words, a query returns either "definitely in set" or "possibly not in set". 

Usage
-----

```
go get github.com/HoMuChen/opp-bloom
```

```go
import "github.com/HoMuChen/opp-bloom"

//create a set with fixed size array
set := oppbloom.New(1000000)
```

`Add` function add a key into the set
```go
set.Add([]byte(`some key`))
set.Add([]byte(`another key`))
```

`Contain` function return if a key is in the set.

if true, the key must be in the set.

if false, the key might not in the set.
```go
setContain([]byte(`some key`)) //true
setContain([]byte(`another key`)) //false

```
