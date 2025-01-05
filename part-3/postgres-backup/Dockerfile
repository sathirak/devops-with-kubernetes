FROM google/cloud-sdk:alpine

RUN apk add --no-cache bash postgresql-client

WORKDIR /usr/src/app

COPY job.sh .

RUN chmod +x job.sh

CMD ["./job.sh"]
