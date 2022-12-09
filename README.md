# adventOfCode2022

I've decided to use this one to learn Golang

# 01

The problem so far is trivial, most of my time here is spent learning how to get the data into the program, set up a map of ints to int array/slices, and get the data in that.

For part 2, the natural solution is a heap.  I found a min-heap and couldn't figure out how to modify it, so I just stored the negative of each value.

I also learned that `go run main.go` won't load other files, even if they're in the same package.

# 02

Part 1
Another easy problem, just needed to learn to do more with maps and 2d arrays.

Part 2
Pretty easy, just use maps to select what to play and get the points.

# Day 3

Part 1
Annoying that there isn't a set and related functions.  Need to ask someone how they normally handle things like this.

Part 2
Also relatively simple, just making and comparing sets/dicts to find commonality.  O(n) solution via adding triplets to sets, then comparing sets 1&2, then that union to 3.

Got tripped up by using range with strings giving back index, then value, and not just value like I expected.

# Day 4

At first glance, this is basically frequency queries. -- Yup, it is.

Part 2 is just a variant, 5 mins and 2 lines changed.

# Day 5

It's funny when it's literal.  Using a stack data structure to simulate stacks of boxes.

This block annoys me.  There must be a better way to do everything here!

```go
for i := 0; i < number; i++ {
    // Is there not a better way?
    moved := crates[from][len(crates[from])-1]
    crates[from] = crates[from][:len(crates[from])-1]
    crates[to] = append(crates[to], moved)
}
```

Part 2 victory again.  < 5 minutes, < 5 lines of code changed
