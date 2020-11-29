package main

import (
   "context"
   "github.com/kubemq-io/kubemq-go"
   "log"
   "time"
)

func main() {
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   client, err := kubemq.NewClient(ctx,
      kubemq.WithAddress("kubemq-cluster-grpc.kubemq", 50000),
      kubemq.WithClientId("java-builder-mvn"),
      kubemq.WithTransportType(kubemq.TransportTypeGRPC))
   if err != nil {
      log.Fatal(err)
   }
   defer client.Close()
   channel := "somayeh.reply.build.mvn"

   sendResult, err := client.NewQueueMessage().
      SetChannel(channel).
      SetBody([]byte("some-simple_queue-queue-message")).
      Send(ctx)
   if err != nil {
      log.Fatal(err)
   }
   log.Printf("Send to Queue Result: MessageID:%s,Sent At: %s\n", sendResult.MessageID, time.Unix(0, sendResult.SentAt).String())
}
