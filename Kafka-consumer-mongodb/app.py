# -*- coding: utf-8 -*-
import time,sys
from pymongo import MongoClient
from kafka import KafkaConsumer

client = MongoClient('mongodb:27017')
db = client.messages
#collection = db.messagecollection
#producer = KafkaProducer(bootstrap_servers='10.106.120.12:30996',api_version=(0,10))
consumer = KafkaConsumer('mytopic', bootstrap_servers='kafka:9092')
for message in consumer:
   print >> sys.stderr, message.value
   db.usercollection.insert({"message": message.value})