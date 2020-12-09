# Kafka topic consumer

The
[ComplianceAnalyzerKafkaTopicConsumer](src/main/java/eu/fasten/analyzer/complianceanalyzer/ComplianceAnalyzerKafkaTopicConsumer.java)
consumes a
[Kafka topic](src/main/resources/inputKafkaTopic.json)
and sends its content to a
[PostgreSQL instance](https://hub.docker.com/_/postgres).

## Run

```bash
mvn clean compile exec:java
```

## Resources

- An example of
[Kafka topic consumption](https://github.com/fasten-project/fasten/blob/ede82dfe7d2154b754a21a75a1eface797b52c82/analyzer/javacg-opal/src/main/java/eu/fasten/analyzer/javacgopal/OPALPlugin.java#L60-L97)
in Java.
- [How to launch a PostgreSQL instance](https://github.com/docker-library/docs/blob/master/postgres/README.md#how-to-use-this-image)
locally.
- [JOOQ library](http://www.jooq.org)
to interact with SQL databases.
