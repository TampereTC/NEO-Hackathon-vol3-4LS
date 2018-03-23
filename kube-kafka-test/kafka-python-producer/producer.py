#!/usr/bin/python3

from kafka import KafkaProducer
import time

# Create a Kafka producer
producer = KafkaProducer(bootstrap_servers='kafka:9092')

# Before sending any messages, write a time stamp
print(time.time(), file=open("timestamp.txt","a"))

# Send 60000 messages to a new topic
for i in range(0,60000):
    producer.send('muntopic123',b'mun message'+bytes(str(i),'utf-8'))

# Close producer to flush the messages
producer.close()

# Sleep 100 seconds to keep the container alive
time.sleep(100)
