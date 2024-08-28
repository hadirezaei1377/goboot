# postman :

# post request

address : `http://localhost:8000/item`
type : `POST`
Body: 
   json, raw, body
  {
    "name": "Laptop",
    "price": 999.99
  }
  
# get request

address : `http://localhost:8000/item`
type : `GET`
`http://localhost:8000/items`
`http://localhost:8000/item/ID1`

# put request

address : `http://localhost:8000/item`
type : `PUT`  
`http://localhost:8000/item/{id}` 
Body: 
   json, body
  {
    "name": "Updated Laptop",
    "price": 899.99
  }
  
# delete

address : `http://localhost:8000/item/{id}`
type : `DELETE`



