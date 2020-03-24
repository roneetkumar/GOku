## __Routers__

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

**Write an algorithm to identify all such routers that need to
be connected to the network all the time.**

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

----------------------------------------------------------------


## __Stores__

A company plans to open stores downtown in the
city of Techlandia. Downtown Techlandia consists of city
blocks, represented as a grid of M*N
blocks. Each block represents either building denoted by
1 or open land area denoted by 0.
Adjacent blocks with value 1 are considered clusters of
buildings. Diagonal blocks with value 1 are not considered
part of the same cluster. this company plans to have a store
in each cluster of buildings.

**Write an algorithm to find the number of stores that
this company can open in downtown Techlandia.**

__Input:__
the input to the function/method consists of three arguments:
rows, an integer representing the number of rows in the grid.
columns, an integer representing the number of columns in
the grid.
grid, a two-dimensional integer array representing downtown
Techlandia.

```
rows = 5
columns = 4
grid = [
   [ 1, 1, 0, 0 ]
   [ 0, 0, 1, 0 ]
   [ 0, 0, 0, 0 ]
   [ 1, 0, 1, 1 ]
   [ 1, 1, 1, 1 ]
]
```

__Output:__
Returns an integer representing the number of stores that
this company could build in downtown Techlandia.

```
output : 3
```
