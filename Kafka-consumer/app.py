# -*- coding: utf-8 -*-
import time,sys
from kafka import KafkaConsumer
count = 1
#producer = KafkaProducer(bootstrap_servers='10.106.120.12:30996',api_version=(0,10))
consumer = KafkaConsumer('mytopic', bootstrap_servers='kafka:9092')
for message in consumer:
   print >> sys.stderr, message