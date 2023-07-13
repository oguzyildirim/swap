## Infinite Resources, Oversubscription, and Linux Swap Memory

### Introduction

Virtual memory is a key feature of modern operating systems. It allows a computer to use more memory than physically
available in the main memory (RAM). This abstraction creates an illusion of seemingly infinite resources for running
processes. Each process is given its own virtual address space, oblivious to the physical memory constraints. Virtual
memory achieves this by mapping virtual addresses to physical addresses as needed and efficiently moving data between
the main memory and secondary storage.

### Oversubscription: Making the Most of Limited Resources

Oversubscription refers to the practice of allocating more virtual memory to processes than the physical memory capacity
of the system. While this may seem counterintuitive, not all processes require their entire address space to be present
in RAM simultaneously. By oversubscribing memory, the operating system can accommodate a larger number of processes or
handle memory-intensive tasks that would otherwise be limited by the physical memory capacity.

The operating system uses a page replacement algorithm to decide which data to move to swap space. These algorithms,
such as the popular Least Recently Used (LRU) algorithm, prioritize keeping frequently accessed data in RAM to minimize
the performance impact of swapping.

### The Role of Swap Memory in Linux

In Linux systems, swap space plays a critical role in managing virtual memory and oversubscription. Swap space is a
designated portion of the hard disk used as an extension of the physical memory. When the RAM becomes full, the
operating system can transfer less frequently used data from RAM to the swap space, thereby freeing up memory for other
processes.

### Infinite resources

Through the abstraction of virtual memory and oversubscription, the illusion of infinite resources is created, allowing
systems to run more processes or handle larger datasets than physical memory alone would allow.
Swap memory serves as an extension of physical memory, enabling efficient memory management by moving less frequently
used data to secondary storage.
