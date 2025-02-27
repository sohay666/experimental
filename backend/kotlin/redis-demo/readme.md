# Experimental Caching With Redis

Kotlin as a Backend for rest API and implemented caching using redis.

The database i used Redis not a relation databases such as Postgree or Mysql, so here's the sample table in the Redis with the keyname `UserTbl` & `ProductTbl`

#### Table : UserTbl
| ColumnName  | DataType |
| ------------- |:-------------:|
| id      | string     |
| name      | string     |
| gender      | string     |
| age      | int     |

#### Table : ProductTbl
| ColumnName  | DataType |
| ------------- |:-------------:|
| id      | string     |
| name      | string     |
| category      | string     |
| qty      | int     |


also for caching, have an prefix:
`user:cache:${id}` & `product:cache:${id}`

## Clean and Build the project.
./gradlew clean build

## Run app 
./gradlew bootRun


```
set to java 11:
/usr/libexec/java_home -V
export JAVA_HOME=$(/usr/libexec/java_home -v 11)
java -version
```

## Endpoints

## Api User

1. Save a User

```
curl -X POST "http://localhost:8081/users" -H "Content-Type: application/json" -d '{"id":"1","name":"Alice","gender":"female","age":18}'
curl -X POST "http://localhost:8081/users" -H "Content-Type: application/json" -d '{"id":"2","name":"Bob","gender":"male","age":18}'
```


2. Get User by ID

```
curl -X GET "http://localhost:8081/users/1"
```

3. Get All Users

```
curl -X GET "http://localhost:8081/users"
```

4. Delete a User

```
curl -X DELETE "http://localhost:8081/users/1"
```


## Api Product

1. Save a Product

```
curl -X POST "http://localhost:8081/products" -H "Content-Type: application/json" -d '{"id":"1","name":"Indomie Kari Ayam","category":"noodle","qty":30}'
curl -X POST "http://localhost:8081/products" -H "Content-Type: application/json" -d '{"id":"2","name":"Indomie Ayam Bawang","category":"noodle","qty":30}'
```


2. Get User by ID

```
curl -X GET "http://localhost:8081/products/1"
```

3. Get All Users

```
curl -X GET "http://localhost:8081/products"
```

4. Delete a User

```
curl -X DELETE "http://localhost:8081/products/1"
```