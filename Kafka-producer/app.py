# -*- coding: utf-8 -*-
import time
from kafka import KafkaProducer
count = 1
#producer = KafkaProducer(bootstrap_servers='10.106.120.12:30996',api_version=(0,10))
producer = KafkaProducer(bootstrap_servers='kafka:9092',api_version=(0,10))
while True:
   message = 'message %i' %count
   producer.send('mytopic', message)
   time.sleep( 10 )
   count += 1
