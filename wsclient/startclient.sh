#!/bin/bash

for ((i=1;i<=400;i=i+1))
do 
	nohup ./wsclient > "client_$i.log" 2>&1 &
done
