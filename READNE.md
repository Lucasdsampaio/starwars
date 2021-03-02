# StarWars

StarWars API - Create, List and Find Planets

[![N|Solid](https://s2.glbimg.com/QqEDyJyWVPTI9tU-5izKpJls6UE=/620x520/smart/e.glbimg.com/og/ed/f/original/2020/11/30/baby-yoda.jpg)]()

### Prerequisites

What things you need to install the software and how to install them

```
GOlang 1.15 >=
Docker
```
## How to Build QA and Deploy
make build-qa && make push-qa

## How to build PROD
make build-prod && make push-prod

## How use
### Endpoints

Create Planet
- /planet 
Method: POST
Payload = {
	"name": "Tatooine",
	"climate": "arid",
	"terrain": "desert"
}

List all Planet
- /planets
Method: GET

Find planet by name
- /planet-name/:name
Method: GET

Find planet by ID
- /planet-id/:id
Method: GET

Delete planet by ID
- /planet/:id
Method: DELETE