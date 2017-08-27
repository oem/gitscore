FROM scratch
MAINTAINER Ömür Özkir <oemuer.oezkir@gmail.com>
ADD ca-certificates.crt /etc/ssl/certs/
ADD list /
ADD dashboard /
ENV PATH="/:${PATH}"
