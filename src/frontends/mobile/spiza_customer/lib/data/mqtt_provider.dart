import 'dart:convert';

import 'package:mqtt_client/mqtt_client.dart';
import 'package:mqtt_client/mqtt_server_client.dart';
import 'package:rxdart/rxdart.dart';

class Subscription<T> {
  final String topic;
  final PublishSubject<T> subject;
  final T Function(Map<String, dynamic>) parser;

  Subscription(this.topic, this.subject, this.parser);
}

class MqttProvider {
  MqttServerClient? client;
  final _subscriptions = <String, Subscription>{};

  void subscribe<T>(String topic, PublishSubject<T> subject,
      T Function(Map<String, dynamic>) parser) {
    final sub = Subscription<T>(topic, subject, parser);
    client!.subscribe(sub.topic, MqttQos.atLeastOnce);
    _subscriptions[sub.topic] = sub;
    print('Subscribed to ${sub.topic}');
  }

  Future connectToMQTT() async {
    client = MqttServerClient.withPort('10.0.2.2', 'flutter_client', 1883);
    client!.logging(on: true);
    client!.onConnected = onConnected;
    client!.onDisconnected = onDisconnected;
    await client!.connect();
    client!.updates!.listen((List<MqttReceivedMessage<MqttMessage>> c) {
      final MqttPublishMessage message = c[0].payload as MqttPublishMessage;
      final payload =
          MqttPublishPayload.bytesToStringAsString(message.payload.message);
      print('Received message: $payload');
      final json = jsonDecode(payload) as Map<String, dynamic>;
      final topic = message.payload.variableHeader!.topicName;
      final subscription = _subscriptions[topic];
      if (subscription == null) {
        print('No subscription for topic $topic');
        return;
      }
      final parsedMessage = subscription.parser(json);
      subscription.subject.sink.add(parsedMessage);
    });
  }

  void onConnected() {
    print('Connected to MQTT');
  }

  void onDisconnected() {
    print('Disconnected from MQTT');
  }
}
