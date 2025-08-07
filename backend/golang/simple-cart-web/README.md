
# A simple Checkout Backend Service

I design the service with MVC structure,
This checkout backend service that will support different promotions with the given inventory and this service have 4 routes


| Method  | Route |Description|
| ------------- | ------------- |-------------|
|    GET        | /ping          | check the service is up/down     |
|    GET  | /products       |  get list products      |
|    POST  | /add-to-cart       | add to cart       |
|    GET  | /list-cart       | will display product after add to cart     |

service will running on port :8001, you can check page directly like this
http://localhost:8001/ping

This service need a database
So you need a setup database like Mysql, i will share the DDL for that

```
CREATE DATABASE shop;
CREATE TABLE products (
    skuNo varchar(15) NOT NULL PRIMARY KEY,
    name varchar(150),
    price float,
    inventoryQty int,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP NULL DEFAULT NULL
);
CREATE TABLE cart (
    cartId varchar(32) NOT NULL PRIMARY KEY,
    skuNo varchar(15),
    qty int,
    CONSTRAINT FOREIGN KEY (skuNo) REFERENCES products (skuNo),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- this data for testing

INSERT INTO products SET skuNo='12OP90', name='Google Home', price=49.99, inventoryQty=10;
INSERT INTO products SET skuNo='43N23P', name='MacBook Pro', price=5399.99, inventoryQty=5;
INSERT INTO products SET skuNo='A304SD', name='Alexa Speaker', price=109.5, inventoryQty=10;
INSERT INTO products SET skuNo='234234', name='Raspberry Pi B', price=30, inventoryQty=2;
```

```
curl http://localhost:8001/products

response:
{
	"status": 200,
	"message": "success",
	"data": [{
		"skuNo": "12OP90",
		"name": "Google Home",
		"price": 49.99,
		"inventoryQty": 10
	}, {
		"skuNo": "234234",
		"name": "Raspberry Pi B",
		"price": 30,
		"inventoryQty": 2
	}, {
		"skuNo": "43N23P",
		"name": "MacBook Pro",
		"price": 5399.99,
		"inventoryQty": 5
	}, {
		"skuNo": "A304SD",
		"name": "Alexa Speaker",
		"price": 109.5,
		"inventoryQty": 10
	}]
}
```

```
curl -X POST -d 'skuNo=12OP90&qty=1' http://localhost:8001/add-to-cart

response:
{
	"status": 200,
	"message": "success"
}
```

```
curl http://localhost:8001/list-cart

response:
{
	"status": 200,
	"message": "success",
	"data": {
		"cartProducts": [{
			"name": "Google Home",
			"qty": 1,
			"price": 49.99
		}],
		"total": 49.99
	}
}
```

note: when user try to add to cart, no need session login, this cartId
generate by user-Agent + date

Implement by using net/http and lib mysql
