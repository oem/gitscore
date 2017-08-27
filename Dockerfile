FROM alpine:3.5
MAINTAINER Ömür Özkir <oemuer.oezkir@gmail.com>
ADD gitscore gitscore
ADD gitscore-dashboard gitscore-dashboard
ENTRYPOINT gitscore-dashboard
