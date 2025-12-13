Here are some notes:

* day 1: not so hard but some corner cases.
  I used Python because that what I'm the most confortable with.
  I rewrote the part 1 in go for fun.

* day 2: trying to do it in go (and not python).
  I had the intuition about the regex for the p1, but the golang regex engine
  does not accept `\1`, so I did it another way.
  It turns out I was on the correct path and the p2 is way easier with a regex.
  I used Perl for that.
  I really struggle to write Perl now (compared to 12 years ago), but the regex
  engine is just so powerful.
  Implicit variables and implicit type conversion also made the code very compact,
  which is nice when writing but terrible when reading (and the implicite type
  conversion is a potential source of error).

* day 9: the _classic_ "is this point in this shape?" problem. I still struggle to solve it.

* day 10:
  * writing a function that returns all (non repeated) combinations of elements
    in a list is suprisingly very hard. Luckily I remembered about the binary _hack_, after
    45 minutes of struggling.
  * thanks to the AOC author this is no use at all for the part 2 of the day


* day 11: classic path search in a tree as there is every year. Some notes:
  * using a FIFO means a breadth first search, and fill all the memory very quickly
  * using a LIFO means a depth first search, and use a constante (little) amount of memory
  * using a cache (a map) to record how many paths are valid from a point avoid going
    multiple time through the same branches
  * my algorithm probably does not work if path are looping: first it will loop endlessly,
    secondly the cache will probably be wrong. But the input does not contains such looping
    path fortunately.
  * I cheated a little bit, because while the problem statement says that the path must go
    through "dac" and "ftt" in any order, we can check that there is no path from "dac" to
    "ftt", which means the order is always "ftt" and then "dac".
