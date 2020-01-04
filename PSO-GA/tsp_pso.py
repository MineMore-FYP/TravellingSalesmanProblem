# encoding:utf-8

'''
	Solution for Travelling Salesman Problem using PSO (Particle Swarm Optimization)
	Discrete PSO for TSP

'''

from operator import attrgetter
import random, sys, time, copy

import numpy as np
import pandas as pd

import parsl
from parsl import load, python_app

import json
import tsp_graph
import userScript
import sys
# insert at 1, 0 is the script path
#sys.path.insert(1, '/home/clusteruser/TravellingSalesmanProblem/PSO-GA/configs')


from configs.local_threads import local_threads
#from configs.local_htex import local_htex
#from remote_htex import remote_htex

parsl.load(local_threads)
#parsl.load(local_htex)
#parsl.load(remote_htex)

'''
# define PSO input parameter : number of iterations
#iterations=sys.argv[1]
iterations=99
INTiterations=int(iterations)

# define PSO input parameter : size of population
#size_population=sys.argv[2]
size_population=9
INTsize_population=int(size_population)

# define PSO input parameter : beta
#beta=sys.argv[3]
beta=0.9
FLOATbeta=float(beta)

# define PSO input parameter : alpha
#alfa=sys.argv[4]
alfa=0.8
FLOATalfa=float(alfa)
'''

# Lower Bound
lb_iterations = userScript.lb_iterations
lb_size_population= userScript.lb_size_population
lb_beta= userScript.lb_beta
lb_alfa= userScript.lb_alfa

# Upper Bound
ub_iterations=  userScript.ub_iterations
ub_size_population= userScript.ub_size_population
ub_beta= userScript.ub_beta
ub_alfa= userScript.ub_alfa

step = sys.argv[1]

# class that represents a particle
class Particle:

	def __init__(self, solution, cost):

		# current solution
		self.solution = solution

		# best solution (fitness) it has achieved so far
		self.pbest = solution

		# set costs
		self.cost_current_solution = cost
		self.cost_pbest_solution = cost

		# velocity of a particle is a sequence of 4-tuple
		# (1, 2, 1, 'beta') means SO(1,2), prabability 1 and compares with "beta"
		self.velocity = []

	# set pbest
	def setPBest(self, new_pbest):
		self.pbest = new_pbest

	# returns the pbest
	def getPBest(self):
		return self.pbest

	# set the new velocity (sequence of swap operators)
	def setVelocity(self, new_velocity):
		self.velocity = new_velocity

	# returns the velocity (sequence of swap operators)
	def getVelocity(self):
		return self.velocity

	# set solution
	def setCurrentSolution(self, solution):
		self.solution = solution

	# gets solution
	def getCurrentSolution(self):
		return self.solution

	# set cost pbest solution
	def setCostPBest(self, cost):
		self.cost_pbest_solution = cost

	# gets cost pbest solution
	def getCostPBest(self):
		return self.cost_pbest_solution

	# set cost current solution
	def setCostCurrentSolution(self, cost):
		self.cost_current_solution = cost

	# gets cost current solution
	def getCostCurrentSolution(self):
		return self.cost_current_solution

	# removes all elements of the list velocity
	def clearVelocity(self):
		del self.velocity[:]


# PSO algorithm
class PSO:

	def __init__(self, graph, iterations, size_population, beta=1, alfa=1):
		self.graph = graph # the graph
		self.iterations = iterations # max of iterations
		self.size_population = size_population # size population
		self.particles = [] # list of particles
		self.beta = beta # the probability that all swap operators in swap sequence (gbest - x(t-1))
		self.alfa = alfa # the probability that all swap operators in swap sequence (pbest - x(t-1))

		# initialized with a group of random particles (solutions)
		solutions = self.graph.getRandomPaths(self.size_population)

		# checks if exists any solution
		if not solutions:
			print('Initial population empty! Try run the algorithm again...')
			sys.exit(1)

		#print("#############################################")
		#print(self.size_population)

		count = 0


		#previous gbest
		#print("#############################################")
		#print(step)
		if step!= "1":
			i = str(int(step)-1)
			with open(userScript.output + i+"_costJson.json", 'r') as myfile:
					data=myfile.read()

			# parse file
			obj = json.loads(data)
			path = obj['path']
			path = path.replace(" ", '')
			path = path[1:-1].split(',')
			newPath =[]
			for i in path:
				#print(i)
				x = i.replace("'", "")
				newPath.append(x)
			#print(newPath)
			costOfPrevious = obj['cost']

		# creates the particles and initialization of swap sequences in all the particles
		for solution in solutions:
			if count == 0 and step != "1":
				#extra new particle
				particle1 = Particle(solution=newPath, cost=costOfPrevious)
				self.particles.append(particle1)
			# creates a new particle
			#print(solution)
			particle = Particle(solution=solution, cost=graph.getCostPath(solution))
			count = count + 1
			# add the particle
			self.particles.append(particle)

		# updates "size_population"
		self.size_population = len(self.particles)

		#print("#############################################")
		#print(self.size_population)

	# set gbest (best particle of the population)
	def setGBest(self, new_gbest):
		self.gbest = new_gbest

	# returns gbest (best particle of the population)
	def getGBest(self):
		return self.gbest


	# shows the info of the particles
	def showsParticles(self):

		print('Showing particles...\n')
		for particle in self.particles:
			print('pbest: %s\t|\tcost pbest: %d\t|\tcurrent solution: %s\t|\tcost current solution: %d' \
				% (str(particle.getPBest()), particle.getCostPBest(), str(particle.getCurrentSolution()),
							particle.getCostCurrentSolution()))
		print('')


	def run(self):

		# for each time step (iteration)
		for t in range(1,self.iterations):
			self.gbest = min(self.particles, key=attrgetter('cost_pbest_solution'))

			# for each particle in the swarm
			for particle in self.particles:
				#print("particle : " + str(particle))
				particle.clearVelocity() # cleans the speed of the particle
				temp_velocity = []

				'''
				if step!="1" and t == 1:
					solution_gbest = newPath
					#solution_pbest = newPath
					#solution_particle = newPath
				else:
					#check what these are in t==1
					solution_gbest = copy.copy(self.gbest.getPBest()) # gets solution of the gbest
				'''
				solution_gbest = copy.copy(self.gbest.getPBest()) # gets solution of the gbest
				solution_pbest = particle.getPBest()[:] # copy of the pbest solution
				solution_particle = particle.getCurrentSolution()[:] # gets copy of the current solution of the particle

				# generates all swap operators to calculate (pbest - x(t-1))
				for i in range(self.graph.amount_vertices):
					if solution_particle[i] != solution_pbest[i]:
						# generates swap operator
						swap_operator = (i, solution_pbest.index(solution_particle[i]), self.alfa)

						# append swap operator in the list of velocity
						temp_velocity.append(swap_operator)

						# makes the swap
						aux = solution_pbest[swap_operator[0]]
						solution_pbest[swap_operator[0]] = solution_pbest[swap_operator[1]]
						solution_pbest[swap_operator[1]] = aux

				# generates all swap operators to calculate (gbest - x(t-1))
				for i in range(self.graph.amount_vertices):
					if solution_particle[i] != solution_gbest[i]:
						# generates swap operator
						swap_operator = (i, solution_gbest.index(solution_particle[i]), self.beta)

						# append swap operator in the list of velocity
						temp_velocity.append(swap_operator)

						# makes the swap
						aux = solution_gbest[swap_operator[0]]
						solution_gbest[swap_operator[0]] = solution_gbest[swap_operator[1]]
						solution_gbest[swap_operator[1]] = aux


				# updates velocity
				particle.setVelocity(temp_velocity)

				# generates new solution for particle
				for swap_operator in temp_velocity:
					if random.random() <= swap_operator[2]:
						# makes the swap
						aux = solution_particle[swap_operator[0]]
						solution_particle[swap_operator[0]] = solution_particle[swap_operator[1]]
						solution_particle[swap_operator[1]] = aux

				# updates the current solution
				particle.setCurrentSolution(solution_particle)
				# gets cost of the current solution
				cost_current_solution = self.graph.getCostPath(solution_particle)
				# updates the cost of the current solution
				particle.setCostCurrentSolution(cost_current_solution)

				# checks if current solution is pbest solution
				if cost_current_solution < particle.getCostPBest():
					particle.setPBest(solution_particle)
					particle.setCostPBest(cost_current_solution)

			#filename = 'savedFiles/iterations_' + str(t) + '.txt'
			#pbestcost = particle.getCostPBest()
			#np.savetxt(rfilename, t, fmt = '%i')
			#print(particle.getCostPBest())

#gbest_path_with_cost_at_tail = []

@python_app
def createPsoInstance(a,b,c,d):
	#print("New PSO instance started")
	#import tsp_graph
	# creates a PSO instance
	pso = PSO(tsp_graph.tsp_graph, a, b, c, d)
	pso.run()

	#pso.showsParticles() # shows the particles


	gbest_path = pso.getGBest().getPBest()
	gbest_path_cost = pso.getGBest().getCostPBest()

	#gbest_path.append(gbest_path_cost)
	# shows the global best particle
	#print('gbest: %s\n' % (gbest_path))

	#print("PSO COMPLETED FOR THE ITERATION " + str(i) + " POPULATION " + str(j) + " BETA " + str(k) + " ALFA " + str(l))
	return [gbest_path,gbest_path_cost]


def stepf():
	columns = ['ITERATION','POPULATION','BETA','ALFA']
	df_new = pd.DataFrame(columns=columns)

	gbest_paths_of_all_psos = []
	for i in range(0,10):
		#print("parsl iteration" + str(i))
		gbest_path1 = createPsoInstance(100,20,0.9,0.8)
		gbest_paths_of_all_psos.append(gbest_path1)
		#costs_of_all_psoInstances.append(gbest_path_cost1)
		df_new = df_new.append({'ITERATION' : 10 , 'POPULATION' : 10 , 'BETA' : 0.9 , 'ALFA' : 0.8},  ignore_index=True)


	#print(df_new)

	gbest_path1_values = []

	for i in gbest_paths_of_all_psos:
		gbest_path1_values.append(i.result())
		#gbest_path1_values.append(i)

	#print(gbest_path1_values)

	path = []
	cost = []

	for i in gbest_path1_values:
		path.append(i[0])
		cost.append(i[1])

	df_new['Path'] = path
	df_new['Cost'] = cost

	print(df_new)


	df_new.to_csv(userScript.output + step + "_pso_instances.csv", index = None, header=True)
	print("tsp pso")

if __name__ == "__main__":

	stepf()
	print("TSP-PSO Step ", step, " completed." )
