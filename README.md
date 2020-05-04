
   Go by example:
      - The principal goal here is show some aproach how concurrence partners can help you to make your application more faster to fetch data in another rest api 


Before everything make sure run: ```docker-compose up -d users-service && docker-compose up dev```


First in order to compare perfomance from sequencial approach copy script and then do request 
  ```cp examples/sequencial-example.txt main.go```

```
 curl -i -XGET 'http://localhost:8070/proxy/users/'
``` 

Will have two concurrence partner to test each steps bellow  
  -  ```cp examples/generator-concurrence.txt main.go```
  -  ```cp examples/fifo-concurrence.txt main.go```


