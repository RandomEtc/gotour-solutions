I ran http://tour.golang.org/ on a plane and this is where I got to.

Most things pass the letter of the exercises but I need to read more Go before I'm confident that my solutions are idiomatic.

I'm not satisfied with 69 or 70, I think I need more structure to stop goroutines gracefully and/or close channels. In 69 I rely on the Tree function always returning 10 nodes, so the channels are buffered with capacity 10 and everything works OK. This feels like I'm cheating somehow! For 70 I ran out of time and the program technically does the right thing but exits in that awkward "ur doin it wrong" kind of way.

Comments/issues/fixes welcome :)

