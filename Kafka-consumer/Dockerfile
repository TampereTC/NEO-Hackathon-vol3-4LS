FROM centos:7

# Update
ENV http_proxy http://10.144.1.10:8080
RUN yum -y install epel-release && yum clean all

# Install python
ENV https_proxy http://10.144.1.10:8080
RUN yum -y install python-pip

# Install app dependencies
RUN pip --proxy http://10.144.1.10:8080 install kafka

# Bundle app source
COPY app.py .

CMD ["python", "app.py"]
