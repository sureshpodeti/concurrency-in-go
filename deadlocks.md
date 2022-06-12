Occurs when all processes are blocked while waiting for each other and program cannot proceed further

### Conditions for deadlock to occur
There are four conditions, known as the coffman conditions, that must be present simultaneously for a deadlock to occur
#### Mutual exclusion
A concurrent process holds atleast one resource at any one time making it non-sharable 
![Tux, the Linux mascot](/2.png)

#### Hold and wait
A concurrent process holds a resource and is waiting for an additional resource
![Tux, the Linux mascot](/3.png)

#### No Preemption
A resource held by a concurrent process cannot be taken away by the system. It can only be freed by the process holding it
![Tux, the Linux mascot](/4.png)

#### Circular wait
A concurrent process must be waiting on a chain of other concurrent processes such that P1 is waiting on P2, P2 on P3, and so on,
and there exists a Pn which is waiting on P1. This forms a ciruclar loop.
![Tux, the Linux mascot](/5.png)
