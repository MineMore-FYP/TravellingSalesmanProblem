from operator import attrgetter
import random, sys, time, copy

import numpy as np
import pandas as pd


import json

import userScript



with open(userScript.output + "1_costJson.json", 'r') as myfile:
	data=myfile.read()
	# parse file
	obj = json.loads(data)
	path = obj['path']
	path = path.replace(" ", '')
	path = path[1:-1].split(',')
	newPath =[]
	print("#####################################################")
	for i in path:
		#print(i)
		x = i.replace("'", "")
		newPath.append(x)
	print(newPath)

	cost = obj['cost']
