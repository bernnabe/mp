# mp

Run:
---
docker-compose up


Api:
----
Local: http://localhost:8080/
Pública: https://dry-sea-61276.herokuapp.com/


CURL de ejempos:

Pruebas al endpoint TopSecret
-----------------------------

curl --header "Content-Type: application/json" --request POST  --data '{"distance": { "kenobi": 5, "skywalker": 3, "sato":5 }, "message":{ "kenobi":["este", "", "", "mensaje", ""], "skywalker":["", "es", "","", "secreto"], "sato":["este", "", "un","", ""] }}'  http://localhost:8080/topsecret
  
 curl --header "Content-Type: application/json" --request POST --data '{"distance": { "kenobi": 5, "skywalker": 3, "sato":5 }, "message":{ "kenobi":["hola1", "", "", "", ""], "skywalker":["", "", "","", "hola5"], "sato":["", "hola2", "hola3","hola4", ""] }}' http://localhost:8080/topsecret

Pruebas al endpoint TopSecret_Slip
----------------------------------

curl --header "Content-Type: application/json" --request POST --data '{"distance": {"kenobi": 5	},"message":{"kenobi":["este", "", "", "mensaje", ""]}}' http://localhost:8080/topsecret_slip

curl --header "Content-Type: application/json" --request POST --data '{"distance": {"skywalker": 3},"message":{"skywalker":["", "es", "","", "secreto"]}}'  http://localhost:8080/topsecret_slip

curl --header "Content-Type: application/json" --request POST --data '{ "distance": {	"sato":5},"message":{	"sato":["este", "", "un","", ""] } }' http://localhost:8080/topsecret_slip


curl --header "Content-Type: application/json" --request GET http://localhost:8080/topsecret_slip


https://dry-sea-61276.herokuapp.com/