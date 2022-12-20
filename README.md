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

# Day 6

Thinking maybe something sort of like an LRU cache.

Ended up just storing N sized blocks in a pseudo-set each loop, probably not the most efficient

Part two super victory, solution in less than 30 seconds

Relatively lazy solution that I don't think would have worked with a really large dataset.  But it worked!  On to the next!

# Day 7

I don't think I actually need any kind of tree here.  Can just do a flat array with the full directory path as the name, kind of like what IIRC Github does.

Then just use a switch to process each line.

Part II
Should have been a super victory, but I don't know how subtraction works apparently ¯\_(ツ)_/¯

# Day 8

I smell a graph problem!

Not yet, more like a holding water variant.

Not my best work in part 2.  I think a two-cursor solution might be more efficient, but on to the next!

# Day 9

Messy collision-like if statements FTW!

Part 2 was pretty easy once I read the rules.  Only the tail counts still!!!

# Day 10

Sort of an emulator/opcodes problem.  Got bit by not DRYing.  Still didn't dry.

# Day 11

3 in the test set and 7 in the final?  Good chance best solution will be exponential!

OOP approach for part 1.  Got confused because sort modifies the original slice whilst append does not.  Also got confused by range, it makes a copy of what it iterates over, so modifying needs a pointer!

I've got a terrifying but for part II.  My solution works fine in the test data for 20 rounds, but is wrong for 1000!?!?!?

Ahhh, because the numbers are getting large enough to overflow

```
{0 [7779672550538772604 2216785406302739480 824468344374183412 4489397885352936922 -6273093796060914110 -6528033743753191667] Operation: new = old * 19 23 2 3 5029}
{1 [8720456354528955014 -3508958547517842620 -3984809899671244867 5712001249291805870] Operation: new = old + 6 19 2 0 4967}
{2 [] Operation: new = old * old 13 1 3 462}
{3 [] Operation: new = old + 3 17 0 1 5245}
```

Now the problem seems to be precision.  Moving on for now

# Day 12
Finally, a graph problem!

Part 2 was almost a super win.  And my alternate path was silly, just 1 different!
