# <a href="https://leetcode.com/problems/two-sum/description/">1. Two Sum</a>

## 📝 Description

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## 🧠 How I solved the problem 

I implemented a brute-force solution using nested loops. To avoid the "double-counting" bug (using the same element twice), I implemted a condition that continue the loop if the "indexes" are equal with eacother. This ensures that we only compare distinct pairs and find the target sum without using any element more than once.

## ➗ Complexity

* **Time complexity**: *O(n^2)* – Each element is compared with every other element.
* **Space complexity**: *O(1)* – No extra data structures are used.


## 📊 Benchmark

I made it in release mode for more accurate results:
```bash
go run .
```

Hardware: *Apple Mac Mini M4*

### 🤏 Small Input Test

* **Execution Time**: *52.083µs*
* **Memory Delta**:   *0 bytes*
* **Current Memory**: *12667144 bytes*

### 😖 Stress Test (Large Input)

* **Execution Time**: *7.833µs*
* **Memory Delta**:   *16 bytes*
* **Current Memory**: *12994824 bytes*