#!/usr/bin/python3

from kafka import KafkaConsumer
import time

# Create a Kafka consumer, with a new topic
consumer = KafkaConsumer('muntopic123', bootstrap_servers='kafka:9092')

# Read first 60000 messages from the topic
count=0
for message in consumer:
    count+=1
    if count == 60000:
        break

# When first 100000 messages have been read, write a time stamp
print(time.time(), file=open("timestamp.txt","a"))

# Sleep 100 seconds to keep container alive
time.sleep(100)
