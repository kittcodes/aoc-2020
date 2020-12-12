# Performance Results - Day 01

## Non Solution-Based Tests (Forcing execution through the entire data set)
### Background
These benchmarks are performed by running a built binary with the count parameters spanning the range of 2 - 5. The target value is set to be unreasonably large such that no values will be eliminated and no logic checks are short-circuited. This forces the program to iterate through every possible combination of values.

### Test System
* CentOS 7 VM running on proxmox
* 8GB RAM
* 4 vCPUs (AMD FX-8150 - 3.6GHz)
* 7.2k RPM HDD

### Results
```
$ for i in {2..5}; do echo "========== RUNNING COUNT OF ${i} ==========="; ./day01 -count ${i} -target 999999999999 -input day1.txt; echo ""; done

========== RUNNING COUNT OF 2 ===========
No solution found in the dataset
Execution Time Elapsed (ms): 0.185787

========== RUNNING COUNT OF 3 ===========
No solution found in the dataset
Execution Time Elapsed (ms): 2.31463

========== RUNNING COUNT OF 4 ===========
No solution found in the dataset
Execution Time Elapsed (ms): 105.430074

========== RUNNING COUNT OF 5 ===========
No solution found in the dataset
Execution Time Elapsed (ms): 4248.394477
```


## Solution-Based Tests

### Background 
My specific data set was such that the smallest number in the data set (112) was a part of the 3-component solution, which makes the 3-component solution unrealistically fast. For funsies, I kept the 2020 value and went up to a 5-component check to demonstrate the early breaks when the target is no longer viable

### Results
```
$ for i in {2..5}; do echo "========== RUNNING COUNT OF ${i} ==========="; ./day01 -count ${i} -target 2020 -input day1.txt; echo ""; done

========== RUNNING COUNT OF 2 ===========
Solution found: [889 1131]
Execution Time Elapsed (ms): 0.169798

========== RUNNING COUNT OF 3 ===========
Solution found: [112 666 1242]
Execution Time Elapsed (ms): 0.159868

========== RUNNING COUNT OF 4 ===========
No solution found in the dataset
Execution Time Elapsed (ms): 0.416975

========== RUNNING COUNT OF 5 ===========
No solution found in the dataset
Execution Time Elapsed (ms): 18.492804
```