# Seminario Go
###### product

## Getting started

Console command line
```
go run cmd/product/product.go -config ./config/config.yaml
```

## Rest

| Methods       | Endpoint                          | 
| ------------- |:--------------------------------: |
| GET           | localhost:8080/product     |
| GET           | localhost:8080/product/:id |
| POST          | localhost:8080/product/    |
| PUT           | localhost:8080/product/    |
| DELETE        | localhost:8080/product/:id |

#### Post body parameters

```
{
  "type": "Bitcoin",
  "quantity": 12
}
```

#### Put body parameters

```
{
  "ID": 1
  "type": "Bitcoin",
  "quantity": 18
}
```
