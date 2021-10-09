FROM mysql
COPY ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf

RUN apt-get install -y tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

EXPOSE 3306
CMD ["mysqld"]