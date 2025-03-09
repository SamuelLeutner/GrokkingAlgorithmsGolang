# Exercises

## Exercise 1.1
- Suppose you have a sorted list of 128 names, and you’re searching through it using binary search. What’s the maximum number of steps it would take?
    - **Answer:** It takes 8 steps. Binary search halves the list each step: 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1. That’s 8 divisions.

## Exercise 1.2
- Suppose you double the size of the list. What’s the maximum number of steps now?
    - **Answer:** Now it’s 9 steps. Doubling to 256: 256 -> 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1. One more step than before.

## Big O Run Time Scenarios

### Exercise 1.3
- You have a name, and you want to find the person’s phone number in the phone book.
    - **Answer:** O(log n). A phone book is sorted by name, so binary search works, cutting the search space in half each step.

### Exercise 1.4
- You have a phone number, and you want to find the person’s name in the phone book. (Hint: You’ll have to search through the whole book!)
    - **Answer:** O(n). Phone books aren’t sorted by number, so you’d scan every entry linearly to find a match.

### Exercise 1.5
- You want to read the numbers of every person in the phone book.
    - **Answer:** O(n). Reading all entries requires a full pass through the book, touching each of the n items once.

### Exercise 1.6
- You want to read the numbers of just the As. (This is a tricky one! It involves concepts that are covered more in chapter 4. Read the answer—you may be surprised!)
    - **Answer:** O(1). If the phone book has an index (like a hash table or precomputed offset for “A”), you jump straight to the “A” section in constant time. Surprising, right?