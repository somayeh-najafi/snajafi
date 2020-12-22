package eu.fasten.analyzer.complianceanalyzer;

import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.KafkaProducer;

import java.util.Properties;


public class ProducerTXT {
	
	public static void main(String[] args) {
		
        String bootstrapServers="localhost:9092";
        String topic="topic1";
             
        // creating producer properties
        
        Properties prop=new Properties();
        prop.put("bootstrap.servers",bootstrapServers);
        prop.put("key.serializer","org.apache.kafka.common.serialization.IntegerSerializer");
        prop.put("value.serializer","org.apache.kafka.common.serialization.StringSerializer");

        //creating producer

        Producer<Integer,String> producer= new KafkaProducer<Integer,String>(prop);
        
        //sending message to topic
        
        for (int i=0; i< 2;i++) {
        	ProducerRecord<Integer,String> producerRecord= new ProducerRecord<Integer,String>("topic1",i,"test  message#"+Integer.toString(i));
        	producer.send(producerRecord);
        }
        producer.close();
	}

}
