import socket
import sys
import subprocess as sp
import os

for i in range (8,9):
	if i != 2:
		folder = "/home/jsubira/DGBP/myn"

		if i==10:
			folder = folder + str(i) + "new"

		elif i!=1:
			folder = folder + str(i)

		os.chdir(folder)
		print "currentDir", os.getcwd()
		#./byfn.sh up -f docker-compose-e2e.yaml
		proc0 = sp.Popen(['./byfn.sh','down','-f','docker-compose-e2e.yaml'],stdout=sp.PIPE, stderr=sp.PIPE,stdin=sp.PIPE)
		proc0.stdin.write("\n")
		proc0.stdin.flush()
		streamdata = proc0.communicate()[0]
		print streamdata
		file0 = open("verbosebyfndown.txt","w")
		file0.write(streamdata)
		print "return code", proc0.returncode

		os.chdir(folder)
		print "currentDir", os.getcwd()
		#./byfn.sh up -f docker-compose-e2e.yaml
		proc0 = sp.Popen(['./byfn.sh','up','-f','docker-compose-e2e.yaml'],stdout=sp.PIPE, stderr=sp.PIPE,stdin=sp.PIPE)
		proc0.stdin.write("\n")
		proc0.stdin.flush()
		streamdata = proc0.communicate()[0]
		print streamdata
		file0 = open("verbosebyfnup.txt","w")
		file0.write(streamdata)
		print "return code", proc0.returncode

		os.chdir("../nodeclient")
		print "currentDir", os.getcwd()
		proc0 = sp.Popen(['rm','-rf','hfc*'],stdout=sp.PIPE, stderr=sp.PIPE)
		streamdata, error = proc0.communicate()
		print streamdata
		print error
		proc0 = sp.Popen(['node','enrollAdmin.js'],stdout=sp.PIPE, stderr=sp.PIPE)
		streamdata, error = proc0.communicate()
		print streamdata
		print error

		file = open("lecturas."+str(i),"w")
		for j in range(0,50):
			proc = sp.Popen( ['node', 'createUser.js' , 'PKI', '0.4.'+ str(j),'Dep1'], stdout=sp.PIPE, stderr=sp.PIPE)
			output, error = proc.communicate()
			file.write(output)
			print "Output: ",output
			print "Error:", error