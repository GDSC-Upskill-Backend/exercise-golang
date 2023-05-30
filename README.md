# Task 1 (Basic Statistics)

## Atention

Before you start, make sure you have read the [general instructions](../README.md), and
Install the dependencies with command:

```bash
go mod tidy
```

## Description

This task is about basic statistics. You will be given a list of integers, and you have to find the following:

- mean
- median
- max

## Input

The first line of the input contains an integer N, the number of integers. The next line contains N integers separated by a space. before the calculation, you have to sort the list in ascending order.

## Output

if N is even, then the median is the average of the two middle numbers. For example, if the list is [1, 2, 3, 4], the median is (2 + 3) / 2 = 2.5. If N is odd, then the median is the middle number. For example, if the list is [1, 2, 3], the median is 2.
