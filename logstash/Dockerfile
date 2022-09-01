FROM logstash:8.4.0

RUN /opt/logstash/bin/logstash-plugin install logstash-integration-jdbc
RUN mkdir /opt/logstash/vendor/jdbc
RUN curl -o /opt/logstash/vendor/jdbc/postgresql-42.4.2.jar https://jdbc.postgresql.org/download/postgresql-42.4.2.jar

ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]
