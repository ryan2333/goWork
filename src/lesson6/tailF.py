#/usr/bin/env python

import time

def readFile(filename):
	with open(filename) as f:
		f.seek(0,2)
		while True:
			line = f.readline()
			if not line:
				time.sleep(3)
				continue
			yield line

if __name__ == "__main__":
	loglines = readFile('/usr/local/data/logs/catalina.out')
	for line in loglines:
		print line

