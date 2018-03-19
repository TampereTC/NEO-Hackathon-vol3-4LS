FROM centos:7

# Update
RUN yum -y install epel-release && yum clean all

# Install python
RUN yum -y install python-pip

# Install app dependencies
RUN pip install kafka

# Bundle app source
COPY app.py .

CMD ["python", "app.py"]