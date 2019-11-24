# encoding:utf-8

'''
	Solution for Travelling Salesman Problem using PSO (Particle Swarm Optimization)
	Discrete PSO for TSP

	References: 
		http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.258.7026&rep=rep1&type=pdf
		http://www.cs.mun.ca/~tinayu/Teaching_files/cs4752/Lecture19_new.pdf
		http://www.swarmintelligence.org/tutorials.php

	References are in the folder "references" of the repository.
'''

from operator import attrgetter
import random, sys, time, copy

from sys import argv

import numpy as np
import pandas as pd
import math
from datetime import datetime

from io_helper import read_tsp

from userScript import tsp_file_path

# class that represents a graph
class Graph:

	def __init__(self, amount_vertices):
		self.edges = {} # dictionary of edges
		self.vertices = set() # set of vertices
		self.amount_vertices = amount_vertices # amount of vertices


	# adds a edge linking "src" in "dest" with a "cost"
	def addEdge(self, src, dest, cost = 0):
		# checks if the edge already exists
		if not self.existsEdge(src, dest):
			self.edges[(src, dest)] = cost
			self.vertices.add(src)
			self.vertices.add(dest)


	# checks if exists a edge linking "src" in "dest"
	def existsEdge(self, src, dest):
		return (True if (src, dest) in self.edges else False)


	# shows all the links of the graph
	def showGraph(self):
		print('Showing the graph:\n')
		for edge in self.edges:
			print('%d linked in %d with cost %d' % (edge[0], edge[1], self.edges[edge]))

	# returns total cost of the path
	def getCostPath(self, path):
		
		total_cost = 0
		for i in range(self.amount_vertices - 1):
			total_cost += self.edges[(path[i], path[i+1])]

		# add cost of the last edge
		total_cost += self.edges[(path[self.amount_vertices - 1], path[0])]
		return total_cost


	# gets random unique paths - returns a list of lists of paths
	def getRandomPaths(self, max_size):

		random_paths, list_vertices = [], list(self.vertices)

		initial_vertice = random.choice(list_vertices)
		if initial_vertice not in list_vertices:
			print('Error: initial vertice %d not exists!' % initial_vertice)
			sys.exit(1)

		list_vertices.remove(initial_vertice)
		list_vertices.insert(0, initial_vertice)

		for i in range(max_size):
			list_temp = list_vertices[1:]
			random.shuffle(list_temp)
			list_temp.insert(0, initial_vertice)

			if list_temp not in random_paths:
				random_paths.append(list_temp)

		return random_paths


# class that represents a complete graph
class CompleteGraph(Graph):

	# generates a complete graph
	def generates(self):
		for i in range(self.amount_vertices):
			for j in range(self.amount_vertices):
				if i != j:
					weight = random.randint(1, 10)
					self.addEdge(i, j, weight)


def getTime():
	now = datetime.now()
	c_time = now.strftime("%H:%M:%S")
	return c_time

def createGraph():

	problem = read_tsp(tsp_file_path)
	
	points = problem[['city', 'x', 'y']]
	#print(points)

	vertices = len(points.index)

	# creates the Graph instance
	graph = Graph(amount_vertices=vertices)
	
	columns = ['City1','City2', 'Cost']
	df_new = pd.DataFrame(columns=columns)

	print('Calculating costs between all the edges .....')
	items = range(1,vertices)
	for i in items:
		points2 = pd.DataFrame(np.roll(points, i, axis=0))
		points2.columns = ['city2','x2', 'y2']
		for j in range(0,vertices):
			xcost = points['x'].iloc[j] - points2['x2'].iloc[j]
			ycost = points['y'].iloc[j] - points2['y2'].iloc[j]
			
			distance = math.sqrt(xcost**2 + ycost**2)

			x = round(distance)
			df_new = df_new.append({'City1' : points['city'].iloc[j] , 'City2' : points2['city2'].iloc[j], 'Cost' : x} , ignore_index=True)
			
	
	print(df_new)
	np.savetxt(r'savedFiles/cities_with_costs.txt', df_new.values, fmt='%s %s %i')
	print("Caluculation Done!")
	listnew = df_new.values.tolist()
	print("Created new list")
	
	tspGraph = set(tuple(x) for x in listnew)
	print("Created new set")
	
	# This graph is in the folder "images" of the repository.
	for value1, value2, key in tspGraph:
		graph.addEdge(value1, value2, key)
	print("Added all the edges!")
	
	return graph
	
