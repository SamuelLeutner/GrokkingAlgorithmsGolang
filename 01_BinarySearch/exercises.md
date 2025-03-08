# Exercises

  1.1 Suppose you have a sorted list of 128 names, and you’re searching through it using binary search. What’s the maximum number of steps it would take?
  > 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1 = **8 steps**

  1.2 Suppose you double the size of the list. What’s the maximum number of steps now?
  > 256 -> 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1 = **9 steps**

## Give the run time for each of these scenarios in terms of big O

  1.3 You have a name, and you want to find the person’s phone number in the phone book.
  > O(log n)

  1.4 You have a phone number, and you want to find the person’s name in the phone book. (Hint: You’ll have to search through the whole book!)
  > O(n)

  1.5 You want to read the numbers of every person in the phone book.
  > O(n)

  1.6 You want to read the numbers of just the As. (This is a tricky one! It involves concepts that are covered more in chapter 4. Read the answer you may be surprised!)
  > O(1)
