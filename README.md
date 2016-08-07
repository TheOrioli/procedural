# Procedural generation-as-a-service

## Maze

### GET `/maze`
  * Returns a short intro message

### GET `/maze/generate/backtrack?w=10&h=10&seed=12345`
  * Returns a JSON encoded generated maze
  * width and height are capped at 500

### GET `/maze/generate/backtrack/image?w=10&h=10&seed=12345`
  * Returns a png image of a generated maze
  * width and height are capped at 100
