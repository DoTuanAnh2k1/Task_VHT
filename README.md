# Problem
There is a .txt file with n lines, each line contains an integer. Sort these n integers and return them in another file.
# Solution
When n is less than `10,000,000`, the merge sort algorithm demonstrates high efficiency. The runtime of the algorithm is fast and further improves when employing multiple goroutines concurrently.

Nevertheless, when dealing with the value `4,000,000,000`, a memory overflow issue arises. This situation prompts the adoption of the External Merge Sort algorithm, which involves creating temporary storage files and processing them.

When creating a temporary file, I read from the input file in chunks of `common.CHUNK_SIZE`, sort each chunk, and then write it to the temporary file. Subsequently, I establish a priority queue to merge all temporary files.

# How to run
You must create a folder named "data" with three subfolders: "input," "output," and "temp". 

Place the input file named "number.txt" in the "data/input" folder. If you want to create a file with `4,000,000,000` numbers then change the flag in `main function` to `true`. 

Afterward, run the command `go run main.go`, and the output file will be generated in the "data/output" folder.