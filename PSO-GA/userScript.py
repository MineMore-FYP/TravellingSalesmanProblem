import os
#from collections import OrderedDict
import sys

####----PARAMETER SETTING FOR SOLVING THE TRAVELLING SALESMAN PROBLEM USING THE PARTICLE SWARM OPTIMIZATION ALGORITHM, WITH HYPER PARAMETER OPTIMIZATION PROVIDED BY A GENETIC ALGORITHM----####

# Graph instance
amount_vertices=5

# Graph nodes

# Graph structure defined
#tspGraph = OrderedDict()

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
	

# PARTICLE SWARM OPTIMIZATION ALGORITHM - Parameters :
# Lower Bound
lb_iterations=99 
lb_size_population=9
lb_beta=0.9
lb_alfa=0.8

# Upper Bound
ub_iterations=101 
ub_size_population=11
ub_beta=1.5
ub_alfa=1.4


# GENETIC ALGORITHM - Parameters :

# Mating pool size
num_parents_mating = 4
#Population size
sol_per_pop = 8
# Number of Generations
num_generations = 1000
# Random Initial Population Lower Bound
initial_low = -4.0
# Random Initial Population Upper Bound
initial_high = 4.0

		


