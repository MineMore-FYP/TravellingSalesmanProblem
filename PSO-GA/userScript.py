import os
#from collections import OrderedDict
import sys

####----PARAMETER SETTING FOR SOLVING THE TRAVELLING SALESMAN PROBLEM USING THE PARTICLE SWARM OPTIMIZATION ALGORITHM, WITH HYPER PARAMETER OPTIMIZATION PROVIDED BY A GENETIC ALGORITHM----####


# PARTICLE SWARM OPTIMIZATION ALGORITHM - Parameters :
# Lower Bound
lb_iterations=99
lb_size_population=9
lb_beta=0.9
lb_alfa=0.8

# Upper Bound
ub_iterations=101
ub_size_population=11
ub_beta=1.1
ub_alfa=1

# Calculate Cost - Parameters(cost.py):
tsp_file_path = 'assets/qa194.tsp'

#output = "/home/mpiuser/Documents/FYP/TravellingSalesmanProblem/PSO-GA/savedFiles/"
output = "/home/amanda/Desktop/SciFlow/TravellingSalesmanProblem/PSO-GA/savedFiles/"


'''
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
'''
