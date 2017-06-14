FROM openwhisk/dockerskeleton
ENV FLASK_PROXY_PORT 8080
COPY bin/helloslack /action/exec
CMD ["/bin/bash", "-c", "cd actionProxy && python -u actionproxy.py"]
