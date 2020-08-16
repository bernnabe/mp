# mp

Run:
---
docker-compose up


Api:
----
Local: http://localhost:8080/

Pública: https://dry-sea-61276.herokuapp.com/


CURL de ejempos:


Pruebas al endpoint TopSecret:
---

Determina la posición y devuelve el mensaje "Este es un mensaje secreto":
---


curl --header "Content-Type: application/json" --request POST --data '{"distance": { "kenobi": 5, "skywalker": 3, "sato":5 }, "message":{ "kenobi":["este", "", "", "mensaje", ""], "skywalker":["", "es", "","", "secreto"], "sato":["este", "", "un","", ""] }}' https://dry-sea-61276.herokuapp.com/topsecret


Determina la posición y devuelve el mensaje "hola1 hola2 hola3 hola4 hola5":  
---


curl --header "Content-Type: application/json" --request POST --data '{"distance": { "kenobi": 5, "skywalker": 3, "sato":5 }, "message":{ "kenobi":["hola1", "", "", "", ""], "skywalker":["", "", "","", "hola5"], "sato":["", "hola2", "hola3","hola4", ""] }}' https://dry-sea-61276.herokuapp.com/topsecret


Pruebas al endpoint TopSecret_Slip:
--


Post del mensaje de kenobi:
---

curl --header "Content-Type: application/json" --request POST --data '{"distance": {"kenobi": 5 },"message":{"kenobi":["este", "", "", "mensaje", ""]}}' https://dry-sea-61276.herokuapp.com/topsecret_slip


Post del mensaje de skywalker:
---


curl --header "Content-Type: application/json" --request POST --data '{"distance": {"skywalker": 3},"message":{"skywalker":["", "es", "","", "secreto"]}}' https://dry-sea-61276.herokuapp.com/topsecret_slip


Post del mensaje de sato:  
---

curl --header "Content-Type: application/json" --request POST --data '{ "distance": { "sato":5},"message":{ "sato":["este", "", "un","", ""] } }' https://dry-sea-61276.herokuapp.com/topsecret_slip


Get del mensaje:  
---
curl --header "Content-Type: application/json" --request GET https://dry-sea-61276.herokuapp.com/topsecret_slip

curl --header "Content-Type: application/json" --request GET http://localhost:8080/topsecret_slip
