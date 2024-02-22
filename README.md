# Problem
There is a .txt file with n lines, each line contains an integer. Sort these n integers and return them in another file.
# Solution
When n is less than `10,000,000`, the merge sort algorithm demonstrates high efficiency. The runtime of the algorithm is fast and further improves when employing multiple goroutines concurrently.

Nevertheless, when dealing with the value `4,000,000,000`, a memory overflow issue arises. This situation prompts the adoption of the External Merge Sort algorithm, which involves creating temporary storage files and processing them.