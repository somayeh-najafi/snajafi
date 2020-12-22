package eu.fasten.analyzer.complianceanalyzer;


import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;

import java.time.Duration;
import java.util.Arrays;
import java.util.Collections;
import java.util.Properties;



public class ConsumerTXT {
	
	public static void main(String[] args) {
	
              String bootstrapServers="localhost:9092";
              String grp_id="test-group";
              String topic="topic1";
              
              // creating consumer properties
              
              Properties properties=new Properties();
              properties.put("bootstrap.servers",bootstrapServers);
              properties.put("group.id",grp_id);
              properties.put("enableauto.commit","true");
              properties.put("key.deserializer","org.apache.kafka.common.serialization.IntegerDeserializer");
              properties.put("value.deserializer","org.apache.kafka.common.serialization.StringDeserializer");
             
              //creating consumer
              
              KafkaConsumer<String,String> consumer= new KafkaConsumer<String,String>(properties);
              
              //Subscribing
              
              consumer.subscribe(Arrays.asList(topic));
              
              //polling
              while(true) {
            	  ConsumerRecords<String,String> records=consumer.poll(Duration.ofMillis(100));
            	  for(ConsumerRecord<String,String> record: records) {
            		  
            		  
            		  System.out.printf("offset=%d, key= %s,value=%s\n",record.offset(),record.key(),record.value());
            		  
            	  }
              }
              
              
}
	}
              