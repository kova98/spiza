# Spiza

Spiza is a food delivery system.

My goal is to build a usable, fully simulated, observable food delivery system with fictional users, restaurants and couriers. 

https://github.com/kova98/spiza/assets/28999034/3b959387-9a90-4aa8-b6b6-14968a4d7267

## Components
### Simulator
Simulates couriers, restaurants and users

### Monitor
Serves a dashboard displaying all current domain information including all active orders and a real-time map with all couriers and restaurants

### Api
A web API that serves domain entities, as well as websocket endpoints for real time information

### Restaurant Web Client
A web client used by a restaurant to manage its menu, items and orders in real time

### User Mobile App
A mobile app used by users to order and track deliveries

### Broker
A lightweight MQTT-based broker for pub/sub communication between the components

## Stack
**Backend:** go  
**Web:** Svelte  
**Mobile:** Flutter  

