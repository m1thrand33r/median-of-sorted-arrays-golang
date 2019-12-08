# median-of-sorted-arrays-golang

Golang implementation of the [median-of-sorted-arrays project](https://github.com/m1thrand33r/median-of-sorted-arrays), eases running code when it needs to be on the `GOPATH`.

- Please clone this project into your Go workspace under the local `GOPATH` in order to compile and run the test.
- Tested with Golang version 1.13.1.

## Problem description

There are 2 sorted arrays A and B, each of the same size. 

Write an algorithm to find the median of the combined set of all numbers contained in both arrays.

Median: See See https://en.wikipedia.org/wiki/Median.

### Examples

- Odd n example: The median of the sorted array `{ 10, 11, 12, 15, 20 }` is 12.
- Even n example: The median of the arrays `{1, 12, 15, 26, 38}` and `{2, 13, 17, 30, 45}` is 16.
  - Looking at the sorted combined array `{1, 2, 12, 13, 15, 17, 26, 30, 38, 45}` we see `{15, 17}` are the median
  - In this case we use the floor of the average `(15 + 17) ? 2 = 16`
