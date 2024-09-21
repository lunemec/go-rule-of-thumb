# Rules of thumb for Go
> In English, the phrase "rule of thumb" refers to an approximate method for doing something, based on practical experience rather than theory.

As a software engineer, you likely have a good understanding of data structures and the `Big O` complexities associated with different usage patterns. However, determining the most suitable data structure for your specific use case can be a challenging decision.

Often, you'll find yourself in a situation where you need to weigh the benefits of creating a new data structure optimized for your access pattern. The question then arises: when is it worthwhile to invest the effort in building a custom data structure?

The decision isn't as simple as solely relying on `Big O` notation, which primarily reflects time complexity. Real-world performance depends on various factors, such as memory locality, the number of allocations, pointer chasing, and more.

## General rules of thumb that always apply
Always KISS (keep it simple, stupid).

1) make it work
2) make it right
3) make it fast

## Disclaimer
❗These rules are not a dogma! Please don't link to this document saying "you should use this because rules-of-thumb says so". Always measure and benchmark your own code with your own data.

Examples here are what is called micro-optimization, before diving into these, profile your code, find real bottlenecks, and fix low hanging fruit there first.


## Needle in a haystack
When is it more efficient to convert a *slice* into a *map* for locating an element `x` within the set `A` (x ∈ A)?

> **TL;DR**: use `map` when `len(haystack) > 100 && len(needles) > 100`

Depending on size of the *haystack* and number of *needles*, this will differ:

| Type  | Haystack | Needles | ns/op        |     |
| ----- | -------- | ------- | ------------ | --- |
| slice | 10       | 10      | 46.90 ns/op  | ✅   |
| map   | 10       | 10      | 203.2 ns/op  |
| slice | 10       | 100     | 332.6 ns/op  | ✅   |
| map   | 10       | 100     | 709.6 ns/op  |
| slice | 10       | 500     | 1670 ns/op   | ✅   |
| map   | 10       | 500     | 3119 ns/op   |
| slice | 10       | 1000    | 3130 ns/op   | ✅   |
| map   | 10       | 1000    | 6123 ns/op   |
| slice | 100      | 10      | 244.1 ns/op  | ✅   |
| map   | 100      | 10      | 2028 ns/op   |
| slice | 100      | 100     | 2145 ns/op   | ✅   |
| map   | 100      | 100     | 2550 ns/op   |
| slice | 100      | 500     | 10919 ns/op  |
| map   | 100      | 500     | 4795 ns/op   | ✅   |
| slice | 100      | 1000    | 22762 ns/op  |
| map   | 100      | 1000    | 7793 ns/op   | ✅   |
| slice | 500      | 10      | 1099 ns/op   | ✅   |
| map   | 500      | 10      | 9804 ns/op   |
| slice | 500      | 100     | 10887 ns/op  |
| map   | 500      | 100     | 10303 ns/op  | ✅   |
| slice | 500      | 500     | 54101 ns/op  |
| map   | 500      | 500     | 12983 ns/op  | ✅   |
| slice | 500      | 1000    | 112415 ns/op |
| map   | 500      | 1000    | 15738 ns/op  | ✅   |
| slice | 1000     | 10      | 2187 ns/op   | ✅   |
| map   | 1000     | 10      | 19861 ns/op  |
| slice | 1000     | 100     | 20728 ns/op  |
| map   | 1000     | 100     | 20219 ns/op  | ✅   |
| slice | 1000     | 500     | 103448 ns/op |
| map   | 1000     | 500     | 22292 ns/op  | ✅   |
| slice | 1000     | 1000    | 207279 ns/op |
| map   | 1000     | 1000    | 25299 ns/op  | ✅   |

## Deduplication
When is it more efficient to deduplicate a `slice` as opposed to using a `map[]struct{}` for the same purpose?

> **TL;DR**: use `map` when `len(haystack) > 100`. Use `slice` when the elements are pre-sorted.

| Type  | Haystack | ns/op       |     |
| ----- | -------- | ----------- | --- |
| slice | 10       | 220.6 ns/op | ✅   |
| map   | 10       | 354.9 ns/op |
| slice | 100      | 3284 ns/op  | ✅   |
| map   | 100      | 3984 ns/op  |
| slice | 500      | 21366 ns/op |
| map   | 500      | 18901 ns/op | ✅   |
| slice | 1000     | 58297 ns/op |
| map   | 1000     | 37787 ns/op | ✅   |

## Subsets
When checking if A is subset of B (A ⊆ B), when is it more efficient to iterate both slices in nested loop `A x B` `O(n^2)`, and when does it make sense to use `map`, or `sort` + binary search?

> **TL;DR**: when use `slice` when `len(A) << len(B)`, use `map` when `len(A) > 500 && len(B) > 500`.

| Type                 | len(A) | len(B) | ns/op        |     |
| -------------------- | ------ | ------ | ------------ | --- |
| slice                | 10     | 10     | 25.39 ns/op  | ✅   |
| slice_sort_binsearch | 10     | 10     | 193.3 ns/op  |
| map                  | 10     | 10     | 164.7 ns/op  |
| slice                | 10     | 100    | 28.84 ns/op  | ✅   |
| slice_sort_binsearch | 10     | 100    | 2051 ns/op   |
| map                  | 10     | 100    | 2044 ns/op   |
| slice                | 10     | 500    | 28.82 ns/op  | ✅   |
| slice_sort_binsearch | 10     | 500    | 12733 ns/op  |
| map                  | 10     | 500    | 9694 ns/op   |
| slice                | 10     | 1000   | 28.92 ns/op  | ✅   |
| slice_sort_binsearch | 10     | 1000   | 37443 ns/op  |
| map                  | 10     | 1000   | 19409 ns/op  |
| slice                | 100    | 100    | 1550 ns/op   | ✅   |
| slice_sort_binsearch | 100    | 100    | 2661 ns/op   |
| map                  | 100    | 100    | 3441 ns/op   |
| slice                | 100    | 500    | 2040 ns/op   | ✅   |
| slice_sort_binsearch | 100    | 500    | 13988 ns/op  |
| map                  | 100    | 500    | 10596 ns/op  |
| slice                | 100    | 1000   | 2137 ns/op   | ✅   |
| slice_sort_binsearch | 100    | 1000   | 39404 ns/op  |
| map                  | 100    | 1000   | 20172 ns/op  |
| slice                | 500    | 500    | 34328 ns/op  |
| slice_sort_binsearch | 500    | 500    | 20569 ns/op  |
| map                  | 500    | 500    | 16648 ns/op  | ✅   |
| slice                | 500    | 1000   | 36999 ns/op  |
| slice_sort_binsearch | 500    | 1000   | 52155 ns/op  |
| map                  | 500    | 1000   | 25562 ns/op  | ✅   |
| slice                | 1000   | 1000   | 129238 ns/op |
| slice_sort_binsearch | 1000   | 1000   | 75645 ns/op  |
| map                  | 1000   | 1000   | 33859 ns/op  | ✅   |

## Append
```go
append([]T, elems...) // append_expand
```
vs
```go
for _, e := range elems {
    arr = append(arr, e) // append_for
}
```
vs
```go
for _, e := range elems {
    arr = append(arr, e) // append_for_prealloc
}
```
vs
```go
for i, e := range elems {
    arr[i] = e // append_for_index (pre-allocated)
}
```

> **TL;DR**: ALWAYS use `append([]T, elems...)` because `for` looping may trigger multiple array re-sizings, whereas `append` will always allocate only once. If you must use `for` loop (extra logic), try to pre-allocate the slice.

Even though regular `append()` has time complexity `O(1)` (amortized constant-time), because every time it needs to allocate more space, it grows the underlying data array by 2x (until 512 elements, after 512 it grows less), simply by having to allocate + copy makes it significantly slower than if you are able to calculate the resulting size and pre-allocating.

| Type                | len(A) | len(B) | ns/op       | B/op       | allocs/op   |     |
| ------------------- | ------ | ------ | ----------- | ---------- | ----------- | --- |
| append_expand       | 10     | 1000   | 878.7 ns/op | 8192 B/op  | 1 allocs/op | ✅   |
| append_for_index    | 10     | 1000   | 1049 ns/op  | 8192 B/op  | 1 allocs/op |
| append_for_prealloc | 10     | 1000   | 1148 ns/op  | 8192 B/op  | 1 allocs/op |
| append_for          | 10     | 1000   | 2115 ns/op  | 19936 B/op | 7 allocs/op |

## Strings concatenation
Is it more efficient to `"str1" + var`, `fmt.Sprintf()`, `strings.Join()` or `strings.Builder`? When does it make sense to add `sync.Pool`?

> **TL;DR**: use `strings.Builder` when `len(str) < 100 & N ops < 1000`, use `sync.Pool + strings.Builder` when doing this for every request. For `len(str) > 100` use `+` or `strings.Join`.
>
> Use `fmt.Sprintf` for regular string formatting (not just concatenation).

| Type                 | len(str) | N ops | ns/op         |     |
| -------------------- | -------- | ----- | ------------- | --- |
| plus_sign            | 10       | 10    | 377.2 ns/op   |     |
| sprintf              | 10       | 10    | 1174 ns/op    |
| strings_join         | 10       | 10    | 457.2 ns/op   |     |
| strings_builder      | 10       | 10    | 226.6 ns/op   | ✅   |
| strings_builder_pool | 10       | 10    | 242.4 ns/op   |     |
| plus_sign            | 10       | 100   | 4976 ns/op    |
| sprintf              | 10       | 100   | 13242 ns/op   |
| strings_join         | 10       | 100   | 5832 ns/op    |
| strings_builder      | 10       | 100   | 1275 ns/op    |
| strings_builder_pool | 10       | 100   | 1265 ns/op    | ✅   |
| plus_sign            | 10       | 500   | 62458 ns/op   |
| sprintf              | 10       | 500   | 106616 ns/op  |
| strings_join         | 10       | 500   | 67195 ns/op   |
| strings_builder      | 10       | 500   | 6515 ns/op    | ✅   |
| strings_builder_pool | 10       | 500   | 6670 ns/op    |
| plus_sign            | 10       | 1000  | 209530 ns/op  |
| sprintf              | 10       | 1000  | 308757 ns/op  |
| strings_join         | 10       | 1000  | 219302 ns/op  |
| strings_builder      | 10       | 1000  | 13754 ns/op   |
| strings_builder_pool | 10       | 1000  | 13660 ns/op   | ✅   |
| plus_sign            | 100      | 10    | 533.6 ns/op   |     |
| sprintf              | 100      | 10    | 1342 ns/op    |
| strings_join         | 100      | 10    | 586.6 ns/op   |     |
| strings_builder      | 100      | 10    | 975.8 ns/op   |     |
| strings_builder_pool | 100      | 10    | 1028 ns/op    |
| plus_sign            | 100      | 100   | 6670 ns/op    | ✅   |
| sprintf              | 100      | 100   | 14949 ns/op   |
| strings_join         | 100      | 100   | 7562 ns/op    |
| strings_builder      | 100      | 100   | 9713 ns/op    |
| strings_builder_pool | 100      | 100   | 9918 ns/op    |
| plus_sign            | 100      | 500   | 71459 ns/op   |
| sprintf              | 100      | 500   | 116144 ns/op  |
| strings_join         | 100      | 500   | 75915 ns/op   |
| strings_builder      | 100      | 500   | 41344 ns/op   | ✅   |
| strings_builder_pool | 100      | 500   | 43655 ns/op   |
| plus_sign            | 100      | 1000  | 227323 ns/op  |
| sprintf              | 100      | 1000  | 519672 ns/op  |
| strings_join         | 100      | 1000  | 316674 ns/op  |
| strings_builder      | 100      | 1000  | 93747 ns/op   | ✅   |
| strings_builder_pool | 100      | 1000  | 99357 ns/op   |
| plus_sign            | 500      | 10    | 1900 ns/op    | ✅   |
| sprintf              | 500      | 10    | 2741 ns/op    |
| strings_join         | 500      | 10    | 2066 ns/op    |
| strings_builder      | 500      | 10    | 7532 ns/op    |
| strings_builder_pool | 500      | 10    | 6380 ns/op    |
| plus_sign            | 500      | 100   | 20481 ns/op   | ✅   |
| sprintf              | 500      | 100   | 30771 ns/op   |
| strings_join         | 500      | 100   | 21401 ns/op   |
| strings_builder      | 500      | 100   | 68556 ns/op   |
| strings_builder_pool | 500      | 100   | 69828 ns/op   |
| plus_sign            | 500      | 500   | 139848 ns/op  | ✅   |
| sprintf              | 500      | 500   | 193818 ns/op  |
| strings_join         | 500      | 500   | 143040 ns/op  |
| strings_builder      | 500      | 500   | 313665 ns/op  |
| strings_builder_pool | 500      | 500   | 322586 ns/op  |
| plus_sign            | 500      | 1000  | 370027 ns/op  | ✅   |
| sprintf              | 500      | 1000  | 515227 ns/op  |
| strings_join         | 500      | 1000  | 379155 ns/op  |
| strings_builder      | 500      | 1000  | 647395 ns/op  |
| strings_builder_pool | 500      | 1000  | 573904 ns/op  |
| plus_sign            | 1000     | 10    | 3220 ns/op    | ✅   |
| sprintf              | 1000     | 10    | 4613 ns/op    |
| strings_join         | 1000     | 10    | 3342 ns/op    |
| strings_builder      | 1000     | 10    | 13307 ns/op   |
| strings_builder_pool | 1000     | 10    | 13747 ns/op   |
| plus_sign            | 1000     | 100   | 40945 ns/op   |
| sprintf              | 1000     | 100   | 49810 ns/op   |
| strings_join         | 1000     | 100   | 36396 ns/op   | ✅   |
| strings_builder      | 1000     | 100   | 142254 ns/op  |
| strings_builder_pool | 1000     | 100   | 149184 ns/op  |
| plus_sign            | 1000     | 500   | 224290 ns/op  | ✅   |
| sprintf              | 1000     | 500   | 296963 ns/op  |
| strings_join         | 1000     | 500   | 233688 ns/op  |
| strings_builder      | 1000     | 500   | 783015 ns/op  |
| strings_builder_pool | 1000     | 500   | 657683 ns/op  |
| plus_sign            | 1000     | 1000  | 557672 ns/op  | ✅   |
| sprintf              | 1000     | 1000  | 715151 ns/op  |
| strings_join         | 1000     | 1000  | 571112 ns/op  |
| strings_builder      | 1000     | 1000  | 1326209 ns/op |
| strings_builder_pool | 1000     | 1000  | 1106394 ns/op |

## If vs switch
Is there even any difference? In theory, `switch` should be faster (at least for some types) if the
compiler is able to transform it into a jump table.

> **TL;DR**: Use which ever one is more readable.

| Type   | N statements | ns/op        |     |
| ------ | ------------ | ------------ | --- |
| if     | 1            | 0.9470 ns/op |
| switch | 1            | 0.9486 ns/op |
| if     | 5            | 1.270 ns/op  |
| switch | 5            | 1.578 ns/op  |

It looks like Go doesn't support jump tables yet? The tests I tried compile into same code for both switch/if statements. You can try to hand-roll jump table [similar to the #19791](https://github.com/golang/go/issues/19791).

Read more:
- https://github.com/golang/go/issues/5496
- https://github.com/golang/go/issues/19791
- https://github.com/golang/go/issues/10870
- https://go-review.googlesource.com/c/go/+/357330
- https://go-review.googlesource.com/c/go/+/395714

## Negative space programming (a.k.a. asserts everywhere)

What is the cost of adding `assert`? Does it make any significant impact?

> **TL;DR**: Use asserts whenever possible to improve reliability of your software. The cost is almost non-existent.

| Type         | N statements | ns/op        |     |
| ------------ | ------------ | ------------ | --- |
| no assert    | 1            | 0.3453 ns/op |
| assert       | 1            | 0.4979 ns/op |
| assert       | 5            | 1.791 ns/op  |
| defer assert | 1            | 2.411 ns/op  |

Read more:
- https://spinroot.com/gerard/pdf/P10.pdf
- https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md#safety
- https://alfasin.com/2017/12/21/negative-space-and-how-does-it-apply-to-coding/

## Pass by reference vs copy
When should you pass a reference (pointer), and when should you use pass by value?

> **TL;DR**: Pass by reference if you want to mutate the data, otherwise pass a copy. 

Performance-wise, this one is almost impossible to give general advice for. If your struct (or nested structs)
are very big (it depends on the types of fields too), copying will become slower. 
But if you have many more pointers, you increase GC pressure and your program will
spend more time on waiting on memory pointer lookup.

References (pointers) vs copied values is way more complicated, 
and there is tons of resources on this topic, great one is 
[this article](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go) by Dave Cheney.

## Range over func
With [Go 1.23 came new feature - range over func](https://go.dev/blog/range-functions), lets check when it makes sense to use that over
pre-allocating a slice and putting values in it.

I'm quite suprised to see that range over func adds extra 3x number of allocations somewhere.
Not sure where, that is to be measured later. 

Here are some results:
```
BenchmarkRangeFunc/slice(10)_iterations(10)-8         	  976581	      1175 ns/op	     880 B/op	      11 allocs/op
BenchmarkRangeFunc/iter_func(10)_iterations(10)-8     	  770889	      1556 ns/op	     528 B/op	      33 allocs/op
BenchmarkRangeFunc/slice(10)_iterations(100)-8        	  112356	     10747 ns/op	    8080 B/op	     101 allocs/op
BenchmarkRangeFunc/iter_func(10)_iterations(100)-8    	   80978	     14675 ns/op	    4848 B/op	     303 allocs/op
BenchmarkRangeFunc/slice(10)_iterations(500)-8        	   22530	     54421 ns/op	   40080 B/op	     501 allocs/op
BenchmarkRangeFunc/iter_func(10)_iterations(500)-8    	   16900	     70454 ns/op	   24048 B/op	    1503 allocs/op
BenchmarkRangeFunc/slice(10)_iterations(1000)-8       	   10000	    105975 ns/op	   80080 B/op	    1001 allocs/op
BenchmarkRangeFunc/iter_func(10)_iterations(1000)-8   	    8494	    140141 ns/op	   48048 B/op	    3003 allocs/op
BenchmarkRangeFunc/slice(100)_iterations(10)-8        	  123470	      9899 ns/op	    9856 B/op	      11 allocs/op
BenchmarkRangeFunc/iter_func(100)_iterations(10)-8    	  116560	     10279 ns/op	     528 B/op	      33 allocs/op
BenchmarkRangeFunc/slice(100)_iterations(100)-8       	   13488	     88975 ns/op	   90496 B/op	     101 allocs/op
BenchmarkRangeFunc/iter_func(100)_iterations(100)-8   	   12710	     94388 ns/op	    4848 B/op	     303 allocs/op
BenchmarkRangeFunc/slice(100)_iterations(500)-8       	    2715	    443446 ns/op	  448897 B/op	     501 allocs/op
BenchmarkRangeFunc/iter_func(100)_iterations(500)-8   	    2498	    492394 ns/op	   24048 B/op	    1503 allocs/op
BenchmarkRangeFunc/slice(100)_iterations(1000)-8      	    1345	    934711 ns/op	  896899 B/op	    1001 allocs/op
BenchmarkRangeFunc/iter_func(100)_iterations(1000)-8  	    1252	    975041 ns/op	   48048 B/op	    3003 allocs/op
BenchmarkRangeFunc/slice(500)_iterations(10)-8        	   25249	     47466 ns/op	   45056 B/op	      11 allocs/op
BenchmarkRangeFunc/iter_func(500)_iterations(10)-8    	   24255	     49390 ns/op	     528 B/op	      33 allocs/op
BenchmarkRangeFunc/slice(500)_iterations(100)-8       	    2761	    434421 ns/op	  413697 B/op	     101 allocs/op
BenchmarkRangeFunc/iter_func(500)_iterations(100)-8   	    2652	    455068 ns/op	    4848 B/op	     303 allocs/op
BenchmarkRangeFunc/slice(500)_iterations(500)-8       	     556	   2159299 ns/op	 2052103 B/op	     501 allocs/op
BenchmarkRangeFunc/iter_func(500)_iterations(500)-8   	     532	   2263910 ns/op	   24048 B/op	    1503 allocs/op
BenchmarkRangeFunc/slice(500)_iterations(1000)-8      	     277	   4309678 ns/op	 4100116 B/op	    1001 allocs/op
BenchmarkRangeFunc/iter_func(500)_iterations(1000)-8  	     266	   4501908 ns/op	   48048 B/op	    3003 allocs/op
BenchmarkRangeFunc/slice(1000)_iterations(10)-8       	   12734	     94393 ns/op	   90112 B/op	      11 allocs/op
BenchmarkRangeFunc/iter_func(1000)_iterations(10)-8   	   12175	     98567 ns/op	     528 B/op	      33 allocs/op
BenchmarkRangeFunc/slice(1000)_iterations(100)-8      	    1383	    862572 ns/op	  827395 B/op	     101 allocs/op
BenchmarkRangeFunc/iter_func(1000)_iterations(100)-8  	    1328	    907260 ns/op	    4848 B/op	     303 allocs/op
BenchmarkRangeFunc/slice(1000)_iterations(500)-8      	     277	   4285715 ns/op	 4104203 B/op	     501 allocs/op
BenchmarkRangeFunc/iter_func(1000)_iterations(500)-8  	     267	   4479069 ns/op	   24048 B/op	    1503 allocs/op
BenchmarkRangeFunc/slice(1000)_iterations(1000)-8     	     139	   8939082 ns/op	 8200234 B/op	    1001 allocs/op
BenchmarkRangeFunc/iter_func(1000)_iterations(1000)-8 	     127	   9272349 ns/op	   48048 B/op	    3003 allocs/op
```

## Notes
- More "Rules of thumb" will be added over time.
- All benchmarks were conducted on a **Macbook Pro M1 (2020) 16GB RAM**, using **Go 1.21.3**. 
