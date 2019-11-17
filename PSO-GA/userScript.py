import os
from collections import OrderedDict
import sys

####----PARAMETER SETTING FOR SOLVING THE TRAVELLING SALESMAN PROBLEM USING THE PARTICLE SWARM OPTIMIZATION ALGORITHM, WITH HYPER PARAMETER OPTIMIZATION PROVIDED BY A GENETIC ALGORITHM----####

# Graph instance
amount_vertices=5

# Graph nodes

# Graph structure defined
tspGraph = OrderedDict()

# Define graph vertices in the following format
# (source node, destination node, cost of edge) 
tspGraph= {(0, 1, 1), 
	(1, 0, 1), 
	(0, 2, 3), 
	(2, 0, 3),
	(0, 3, 4),
	(3, 0, 4),
	(0, 4, 5),
	(4, 0, 5),
	(1, 2, 1),
	(2, 1, 1),
	(1, 3, 4),
	(3, 1, 4),
	(1, 4, 8),
	(4, 1, 8),
	(2, 3, 5),
	(3, 2, 5),
	(2, 4, 1),
	(4, 2, 1),
	(3, 4, 2),
	(4, 3, 2)}
	

# Particle Swarm Optimization Algorithm - Parameters
iterations=100 
size_population=10
beta=1
alfa=0.9

#print(tspGraph)

#for key, value1, value2 in tspGraph:
#	print("key",key)
#	print("value1",value1)
#	print("value2",value2)
		


