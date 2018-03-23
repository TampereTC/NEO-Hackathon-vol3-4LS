FROM python:3

ENV https_proxy http://10.144.1.10:8080

RUN pip install kafka

ADD producer.py /

CMD [ "python", "./producer.py" ]
