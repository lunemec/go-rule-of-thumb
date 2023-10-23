# Rules of thumb for Go
> In English, the phrase "rule of thumb" refers to an approximate method for doing something, based on practical experience rather than theory.

As a software engineer, you likely have a good understanding of data structures and the `Big O` complexities associated with different usage patterns. However, determining the most suitable data structure for your specific use case can be a challenging decision.

Often, you'll find yourself in a situation where you need to weigh the benefits of creating a new data structure optimized for your access pattern. The question then arises: when is it worthwhile to invest the effort in building a custom data structure?

The decision isn't as simple as solely relying on `Big O` notation, which primarily reflects time complexity. Real-world performance depends on various factors, such as memory locality, the number of allocations, pointer chasing, and more.


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
Is it more efficient to `append([]T, elems...)` or in a `for` loop one-by-one?

> **TL;DR**: ALWAYS use `append([]T, elems...)` because `for` looping may trigger multiple array re-sizings, whereas `append` will always allocate only once.

| Type          | len(A) | len(B) | ns/op       | B/op       | allocs/op   |     |
| ------------- | ------ | ------ | ----------- | ---------- | ----------- | --- |
| append_expand | 10     | 1000   | 994.1 ns/op | 8192 B/op  | 1 allocs/op | ✅   |
| append_for    | 10     | 1000   | 2533 ns/op  | 19936 B/op | 7 allocs/op |

## Strings concatenation
Is it more efficient to `"str1" + var`, `fmt.Sprintf()`, `strings.Join()` or `strings.Builder`. When does it make sense to add `sync.Pool`?

> **TL;DR**: use `strings.Builder` when `len(str) < 100 & N ops < 1000`, use `sync.Pool + strings.Builder` when doing this for every request. For `len(str) > 100` use `+` or `strings.Join`.

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
| plus_sign            | 1000     | 10    | 3220 ns/op    | ✅
| sprintf              | 1000     | 10    | 4613 ns/op    |
| strings_join         | 1000     | 10    | 3342 ns/op    |
| strings_builder      | 1000     | 10    | 13307 ns/op   |
| strings_builder_pool | 1000     | 10    | 13747 ns/op   |
| plus_sign            | 1000     | 100   | 40945 ns/op   |
| sprintf              | 1000     | 100   | 49810 ns/op   |
| strings_join         | 1000     | 100   | 36396 ns/op   | ✅
| strings_builder      | 1000     | 100   | 142254 ns/op  |
| strings_builder_pool | 1000     | 100   | 149184 ns/op  |
| plus_sign            | 1000     | 500   | 224290 ns/op  | ✅
| sprintf              | 1000     | 500   | 296963 ns/op  |
| strings_join         | 1000     | 500   | 233688 ns/op  |
| strings_builder      | 1000     | 500   | 783015 ns/op  |
| strings_builder_pool | 1000     | 500   | 657683 ns/op  |
| plus_sign            | 1000     | 1000  | 557672 ns/op  | ✅
| sprintf              | 1000     | 1000  | 715151 ns/op  |
| strings_join         | 1000     | 1000  | 571112 ns/op  |
| strings_builder      | 1000     | 1000  | 1326209 ns/op |
| strings_builder_pool | 1000     | 1000  | 1106394 ns/op |

### Notes
- More "Rules of thumb" will be added over time.
- All benchmarks were conducted on a **Macbook Pro M1 (2020) 16GB RAM**, using **Go 1.21.3**. 
- As always, benchmark your data with your usecase, and on hardware where it will run (AWS/gcloud).
