# mp

Run:

docker-compose up


Api:

http://localhost:8080/

CURL de ejempos:

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"distance": { "kenobi": 5, "skywalker": 3, "sato":5 }, "message":{ "kenobi":["este", "", "", "mensaje", ""], "skywalker":["", "es", "","", "secreto"], "sato":["este", "", "un","", ""] }}' \
  http://localhost:8080/topsecret
  
 curl --header "Content-Type: application/json" \\
  --request POST \
  --data '{"distance": { "kenobi": 5, "skywalker": 3, "sato":5 }, "message":{ "kenobi":["hola1", "", "", "", ""], "skywalker":["", "", "","", "hola5"], "sato":["", "hola2", "hola3","hola4", ""] }}' \
  http://localhost:8080/topsecret
