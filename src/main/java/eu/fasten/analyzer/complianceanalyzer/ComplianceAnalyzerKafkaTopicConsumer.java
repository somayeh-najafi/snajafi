package eu.fasten.analyzer.complianceanalyzer;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.*;

import eu.fasten.core.plugins.KafkaPlugin;

public class ComplianceAnalyzerKafkaTopicConsumer implements KafkaPlugin {

    /**
     * Consumes a license & compliance Kafka topic and sends it to a PostgreSQL instance.
     *
     * @param kafkaRecord   the input Kafka topic to be sent to a PostgreSQL instance.
     * @throws SQLException 
     * @throws ClassNotFoundException 
     */
	
	
	public static void main(String[] args) throws SQLException, ClassNotFoundException {
		
		Class.forName("org.postgresql.Driver");
        try (Connection connection = DriverManager.getConnection("jdbc:postgresql://35.234.104.111/mddb", "postgres", "Cae2eimea7ae")) {
 
           
        	if (connection != null) {
                System.out.println("Connected to the database!");
            } else {
                System.out.println("Failed to make connection!");
            }
        	
        	
        	//System.out.println("Java JDBC PostgreSQL Example");
            // When this class first attempts to establish a connection, it automatically loads any JDBC 4.0 drivers found within 
            // the class path. Note that your application must manually load any JDBC drivers prior to version 4.0.
//          Class.forName("org.postgresql.Driver"); 
 
            //System.out.println("Connected to PostgreSQL database!");
        }
	
	
	}
	
    public void consume(String kafkaRecord) {
        // TODO Consuming the Kafka topic

        // TODO Sending its content to a PostgreSQL instance
    }
}
