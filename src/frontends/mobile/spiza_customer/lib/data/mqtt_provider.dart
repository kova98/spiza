import 'dart:convert';

import 'package:mqtt_client/mqtt_client.dart';
import 'package:mqtt_client/mqtt_server_client.dart';
import 'package:rxdart/src/subjects/publish_subject.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:spiza_customer/models/order_update.dart';

class MqttProvider {
  MqttServerClient? client;

  subscribe(String topic, PublishSubject<OrderUpdate> orderUpdateSubject) {
    client!.subscribe(topic, MqttQos.atLeastOnce);

    client!.updates!.listen((List<MqttReceivedMessage<MqttMessage>> c) {
      final MqttPublishMessage message = c[0].payload as MqttPublishMessage;
      final payload =
          MqttPublishPayload.bytesToStringAsString(message.payload.message);
      print('Received message: $payload');
      final json = jsonDecode(payload);
      final order = OrderUpdate.fromJson(json);
      orderUpdateSubject.sink.add(order);
    });
  }

  Future connectToMQTT() async {
    client = MqttServerClient.withPort('10.0.2.2', 'flutter_client', 1883);
    client!.logging(on: true);
    client!.onConnected = onConnected;
    client!.onDisconnected = onDisconnected;
    await client!.connect();
  }

  void onConnected() {
    print('Connected to MQTT');
  }

  void onDisconnected() {
    print('Disconnected from MQTT');
  }
}
