curl --location 'localhost:8080' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4OTY3NzIsInVzZXJuYW1lIjoiYWRtaW4ifQ.s4XU_6-uLOyv-AkI6HhRrWYpVhZY-QmOBc1RM7O4KiM'

curl --location 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"admin",
    "password":"admin"
}'

curl --location 'localhost:8080/auth/protected1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4OTY0ODgsInVzZXJuYW1lIjoiYWRtaW4ifQ.YJSj-5Jaqx-CLsRZZ-h2aWH8iTMOeUdKGMB7RvrA6SQ'

curl --location 'localhost:8080/auth/protected2?id=123' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4OTM1NzEsInVzZXJuYW1lIjoiYWRtaW4ifQ.tIZ2HZOoh36cNd-303AIc8nlM4rBAtzNQuwwe0aO02s'

curl --location 'localhost:8080/auth/protected3/santoso' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4OTM1NzEsInVzZXJuYW1lIjoiYWRtaW4ifQ.tIZ2HZOoh36cNd-303AIc8nlM4rBAtzNQuwwe0aO02s'