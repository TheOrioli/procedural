FORMAT: 1A

# Procedural

Procedural generation as a service API documentation.

## Group Maze

A [Maze](https://en.wikipedia.org/wiki/Maze) is represented as a graph of `Point` objects that are connected to other `Point` objects. An entrance an is also provided for easier graph traversal, but all `Points` in the Maze can reach all other `Points`

## Message [GET /maze]

Return a simple sub-service description, with a message and the current version.

- Response 200 (application/json)

  - Attributes (Intro Object)

      - message: `Maze-as-a-service`

## Generate [/maze/generate/{algorithm}{?w,h,seed}]

Generates a new maze using a specified `algorithm`.

If you want to share a maze, it's just a copy-paste away.

## JSON [GET]

- Parameters
  - algorithm: `backtrack` (required, enum[string])
      - Members:
        - `backtrack`
  - w: `3` (optional, number)
      - Default: 10
  - h: `3` (optional, number)
      - Default: 10
  - seed: `5` (optional, number)
      - Default: 0

- Request (application/json)

- Response 200 (application/json)
  - Attributes (Maze Object)
      - entrance
          - x: `0` (number)
          - y: `2` (number)
      - exit
          - x: `0` (number)
          - y: `2` (number)
      - width: `3` (number)
      - height: `3` (number)
      - grid (array)
          - (Grid Point Object)
              - point
                  - x: `0` (number)
                  - y: `2` (number)
              - next (array)
                  - (Point Object)
                      - x: `0` (number)
                      - y: `1` (number)
          - (Grid Point Object)
              - point
                  - x: `1` (number)
                  - y: `0` (number)
              - next (array)
                  - (Point Object)
                      - x: `0` (number)
                      - y: `0` (number)
                  - (Point Object)
                      - x: `2` (number)
                      - y: `0` (number)
          - (Grid Point Object)
              - point
                  - x: `2` (number)
                  - y: `0` (number)
              - next (array)
                  - (Point Object)
                      - x: `1` (number)
                      - y: `0` (number)
                  - (Point Object)
                      - x: `2` (number)
                      - y: `1` (number)
          - (Grid Point Object)
              - point
                  - x: `1` (number)
                  - y: `2` (number)
              - next (array)
                  - (Point Object)
                      - x: `2` (number)
                      - y: `2` (number)
                  - (Point Object)
                      - x: `1` (number)
                      - y: `1` (number)
          - (Grid Point Object)
              - point
                  - x: `0` (number)
                  - y: `1` (number)
              - next (array)
                  - (Point Object)
                      - x: `0` (number)
                      - y: `2` (number)
                  - (Point Object)
                      - x: `0` (number)
                      - y: `0` (number)
          - (Grid Point Object)
              - point
                  - x: `0` (number)
                  - y: `0` (number)
              - next (array)
                  - (Point Object)
                      - x: `0` (number)
                      - y: `1` (number)
                  - (Point Object)
                      - x: `1` (number)
                      - y: `0` (number)
          - (Grid Point Object)
              - point
                  - x: `2` (number)
                  - y: `1` (number)
              - next (array)
                  - (Point Object)
                      - x: `2` (number)
                      - y: `0` (number)
                  - (Point Object)
                      - x: `2` (number)
                      - y: `2` (number)
          - (Grid Point Object)
              - point
                  - x: `2` (number)
                  - y: `2` (number)
              - next (array)
                  - (Point Object)
                      - x: `2` (number)
                      - y: `1` (number)
                  - (Point Object)
                      - x: `1` (number)
                      - y: `2` (number)
          - (Grid Point Object)
              - point
                  - x: `1` (number)
                  - y: `1` (number)
              - next (array)
                  - (Point Object)
                      - x: `1` (number)
                      - y: `2` (number)

## Image [GET /maze/generate/{algorithm}/image{?w,h,seed}]

- Parameters
  - algorithm: `backtrack` (required, enum[string])
      - Members:
        - `backtrack`
  - w: `3` (optional, number)
      - Default: 10
  - h: `3` (optional, number)
      - Default: 10
  - seed: `5` (optional, number)
      - Default: 0

- Request (application/json)

- Response 200 (image/png)

## Group Dungeon

A Dungeon is a collection of rooms connected by hallways.

## Message [GET /dungeon]

Return a simple sub-service description, with a message and the current version.

- Response 200 (application/json)

  - Attributes (Intro Object)

      - message: `Dungeon-as-a-service`

## Generate [/dungeon/generate{?size,rooms,seed}]

Generates a new dungeon.

## JSON [GET]

- Parameters
  - size: `20` (optional, number)
      - Default: 5
  - rooms: `3` (optional, number)
      - Default: 10
  - seed: `1` (optional, number)
      - Default: 0

- Request (application/json)

- Response 200 (application/json)
  - Attributes (Dungeon Object)

## Image [GET /dungeon/generate/image{?size,rooms,seed}]

- Parameters
  - size: `20` (optional, number)
      - Default: 5
  - rooms: `3` (optional, number)
      - Default: 10
  - seed: `1` (optional, number)
      - Default: 0

- Request (application/json)

- Response 200 (image/png)

# Data Structures

## Intro Object (object)

- message (required, string) - A short intro message describing the sub-service
- version (required, Version Object) - Version following the SEMVER approach

## Version Object (object)

- major (required, number)
- minor (required, number)
- patch (required, number)

## Point Object (object)

- x (required, number) - X coordinate
- y (required, number) - Y coordinate

## Grid Point Object (object)

- point (required, Point Object) - the current point coordinates
- next (required, array) - coordinates accessible from these coordinates.

  - (Point Object)

## Maze Object (object)

- width (required, number) - Maze width
- height (required, number) - Maze height
- entrance (required, Point Object) - Maze entrance coordinates
- exit (required, Point Object) - Maze exit coordinates
- grid (required, array) - Grid of all maze points, arranged into a graph

  - (Grid Point Object)

## Dungeon Object (object)
- width (required, number) - Dungeon width
- height (required, number) - Dungeon height
- grid (required, array) - Grid of all dungeon points, arranged into a graph

  - (Grid Point Object)
