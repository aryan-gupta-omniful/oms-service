# Order Management Service
Flow of this service is as follows:
1. Post Request on ```http://localhost:8081/api/v1/orders/bulkorder```; Handler function will be triggered and validate the filePath.
2. Validated path will be published to an AWS SQS Queue - ```{"sellerID":1,"filePath":"storage/orders.csv"}```
3. An SQS Listener will be constantly listening to the specified Queue and extract the message and subsequently call the ParseCSV function.
4. After getting parsed, the hub_id and sku_id of every Order Item will be validated via a POST request to ```wms-service```.
5. Once validated, all orders will be pushed inside the a MongoDB database inside 'Orders' Collection with order status - ```on_hold```, and also be published to a Kafka Partition.
6. A Kafka Listener will listem to these orders, extract them and again make an inter-service API call to ```wms-service```.
7. If the quantity ordered is less the amount present in the inventory, Order status will change to ```new_order``` and appropriate quanity will be deducted from the inventory of that hub from which order was placed.

## Project Setup

1. Setup Dependencies
```go
  go mod tidy
```
2. Run Kafka + Zookeeper Docker Image
```go
  docker compose up
```
3. Set-Up SQS Queue in AWS Account and Configure AWS CLI

4. Run Main Program
```go
  go run main.go
```
5. Make Database: ``` oms_service_db ``` in MongoDB and use Atlas/Compass URI

## Previews
- **Initialize SQS, Kafka, Connect to MongoDB and Parse CSV File :**
  
![image](https://github.com/user-attachments/assets/e37f6d20-7a8a-4e05-890d-3c5af94d7c4e)

- **Message Published inside SQS Queue**
  
![image](https://github.com/user-attachments/assets/e643a1bf-f819-4442-b7a5-6d573781b009)

- **Validate Hub Id and SKU Id :**
  
![image](https://github.com/user-attachments/assets/07103848-6a35-430c-a40d-2c1a04cdd6a0)

- **Publish and Consume Message from Kafka :**
  
![image](https://github.com/user-attachments/assets/aa099939-7643-49cb-bdc5-5fa8abb93a0e)

- **Consume Message from Kafka and Validate Inventory Count :**
  
![image](https://github.com/user-attachments/assets/89ddda3b-196e-4046-8c2d-4d4364bd3a34)

- **Update Inventory for Valid Orders :**
  
![image](https://github.com/user-attachments/assets/644b486c-2273-46fd-b004-96ac1ed12497)


## Directory Structure 
```
oms-service/
├── configs/
│   └── config.yaml
├── controllers/
│   └── controller.go
├── init/
│   ├── init.go
│   ├── sqs-consumer.go
│   └── sqs-producer.go
├── intersvc/
│   └── interserviceCall.go
├── kafka/
│   ├── kafka_consumer.go
│   ├── kafka_producer.go
│   └── validate_inventory.go
├── models/
│   ├── bulk_order.go
│   └── customer.go
├── parse_csv/
│   └── parseCSV.go
├── repository/
│   └── mongodb.go
├── routes/
│   └── routes.go
├── storage/
│   └── orders.csv
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```


