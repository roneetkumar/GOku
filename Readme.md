#__Routers__:

a company wants to increase the reliability of its network
and is piloting hardware enhancements in one of the data
centers prior to a full-scale roll out. To facilitate
the routing of the incoming packets, there is a network
of N routers in the data center. Every router is connected
to every other router of the network either through a direct
link or via some other routers.

to increase the fault tolerance of the network, the company
wants to identify routers which would result in a disconnected
network if they went down and add replicas of these routers
to the network.

Write an algorithm to identify all such routers that need to
be connected to the network all the time.

__Input:__
the input to the function/method consists of three
arguments:
numRouters, an integer representing the number of routers
in the data center (>=3)
numLinks, an integer representing the number of links
between the pair of routers.
Links, a list of pairs of integers - A ,B reprensenting
a link between the routers A and B. the network will be
connected.
```
numRouters = 7
numLinks = 7
links = [[1,2],[1,3],[2,4],[3,4],[3,6],[6,7],[4,5]]
```
__Output:__

Return a list of integers representing the routers which
need to be connected to the network all the time.
```
[3 , 4 , 6]
```
