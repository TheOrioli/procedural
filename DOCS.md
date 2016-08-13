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

## Group Music

It's a bit hard to define `Music` so let's just call it a semi-random collection of air pressure changes that sound like something,

## Message [GET /music]

Return a simple sub-service description, with a message and the current version.

- Response 200 (application/json)

  - Attributes (Intro Object)

      - message: `Music-as-a-service`

## Generate [/music/generate{?size,seed,smoke_on_the_water}]

## JSON [GET]
The `duration` of a note is the musical notation duration of a note, e.g. full, half, quarter, eight.

The `start_at` attribute of a note describes at which point in the song should the note start playing, and it is equal to the sum of the duration of notes that came before it.

By using the `tempo`, the `duration` and the `start_at` attribute, the correct length and position of a note in seconds can be determined.

- Parameters
  - size: `64` (optional, number) - a modifier parameter, influences the length of the finished song by a little bit
      - Default: 10
  - seed: `5` (optional, number)
      - Default: 0
  - smoke_on_the_water: true (optional, boolean) - just a little easter egg, useful for checking out the sound generation
      - Default: false

- Response 200 (application/json)
  - Attributes (Song Object)
      - scale (array)
          - 2 (number)
          - 1 (number)
          - 2 (number)
          - 2 (number)
          - 1 (number)
          - 2 (number)
          - 2 (number)
      - key (Note Object)
          - note: G (string)
          - octave: 2 (number)
          - frequency: 97.9988589954373 (number)
      - tempo: 0.75 (number)
      - song (array)
          - (Song Note Object)
              - note: D (string)
              - octave: 2 (number)
              - frequency: 73.41619197935186 (number)
              - duration: 0.75 (number)
              - star_at: 0 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 0.75 (number)
              - star_at: 0 (number)
          - (Song Note Object)
              - note: F (string)
              - octave: 2 (number)
              - frequency: 87.30705785825096 (number)
              - duration: 0.75 (number)
              - start_at: 0.75 (number)
          - (Song Note Object)
              - note: A (string),
              - octave: 2 (number)
              - frequency: 116.54094037952244 (number)
              - duration: 0.75 (number)
              - start_at: 0.75 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 1 (number)
              - start_at: 1.5 (number)
          - (Song Note Object)
              - note: C (string)
              - octave: 2 (number)
              - frequency: 65.40639132514963 (number)
              - duration: 1 (number)
              - start_at: 1.5 (number)
          - (Song Note Object)
              - note: D (string)
              - octave: 2 (number)
              - frequency: 73.41619197935186 (number)
              - duration: 0.75 (number)
              - start_at: 2.5 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 0.75 (number)
              - start_at: 2.5 (number)
          - (Song Note Object)
              - note: F (string)
              - octave: 2 (number)
              - frequency: 87.30705785825096 (number)
              - duration: 0.75 (number)
              - start_at: 3.25 (number)
          - (Song Note Object)
              - note: A (string),
              - octave: 2 (number)
              - frequency: 116.54094037952244 (number)
              - duration: 0.75 (number)
              - start_at: 3.25 (number)
          - (Song Note Object)
              - note: G (string),
              - octave: 2 (number)
              - frequency: 103.82617439498627 (number)
              - duration: 0.5 (number)
              - star_at: 4 (number)
          - (Song Note Object)
              - note: C (string),
              - octave: 2 (number)
              - frequency: 69.295657744218 (number)
              - duration: 0.5 (number)
              - star_at: 4 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 1 (number)
              - start_at: 4.5 (number)
          - (Song Note Object)
              - note: C (string)
              - octave: 2 (number)
              - frequency: 65.40639132514963 (number)
              - duration: 1 (number)
              - start_at: 4.5 (number)
          - (Song Note Object)
              - note: D (string)
              - octave: 2 (number)
              - frequency: 73.41619197935186 (number)
              - duration: 0.75 (number)
              - start_at: 5.5 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 0.75 (number)
              - start_at: 5.5 (number)
          - (Song Note Object)
              - note: F (string)
              - octave: 2 (number)
              - frequency: 87.30705785825096 (number)
              - duration: 0.75 (number)
              - start_at: 6.25 (number)
          - (Song Note Object)
              - note: A (string),
              - octave: 2 (number)
              - frequency: 116.54094037952244 (number)
              - duration: 0.75 (number)
              - start_at: 6.25 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 1 (number)
              - star_at: 7 (number)
          - (Song Note Object)
              - note: C (string)
              - octave: 2 (number)
              - frequency: 65.40639132514963 (number)
              - duration: 1 (number)
              - star_at: 7 (number)
          - (Song Note Object)
              - note: F (string)
              - octave: 2 (number)
              - frequency: 87.30705785825096 (number)
              - duration: 0.75 (number)
              - star_at: 8 (number)
          - (Song Note Object)
              - note: A (string),
              - octave: 2 (number)
              - frequency: 116.54094037952244 (number)
              - duration: 0.75 (number)
              - star_at: 8 (number)
          - (Song Note Object)
              - note: D (string)
              - octave: 2 (number)
              - frequency: 73.41619197935186 (number)
              - duration: 2 (number)
              - start_at: 8.75 (number)
          - (Song Note Object)
              - note: G (string)
              - octave: 2 (number)
              - frequency: 97.9988589954373 (number)
              - duration: 2 (number)
              - start_at: 8.75 (number)

## WAVE [GET /music/generate/wave{?size,seed,smoke_on_the_water}]

- Parameters
  - size: `64` (optional, number) - a modifier parameter, influences the length of the finished song by a little bit
      - Default: 10
  - seed: `5` (optional, number)
      - Default: 0
  - smoke_on_the_water: true (optional, boolean) - just a little easter egg, useful for checking out the sound generation
      - Default: false

- Response 200 (audio/wav)

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

## Note Object (object)
- note (required, string) - Musical note name
- octave (required, number) - octave of the note
- frequency (required, number) - physical frequency of the note

## Song Note Object (Note Object)
- duration (required, number) - length of the note in musical notation (full, quarter, half etc.)
- start_at (required, number) - start at the specified point in the song, specified in musical notation as a sum of duration passed

## Song Object (object)
- scale (required, array) - the scale pattern in halfstep jumps from the key of the scale
  - (number)
- key (required, Note Object) - the note that sets key of the scale and the song
- Tempo (required, number) - the relative length of a full(1.0) duration note
- song (required, array) - an array of notes that define the song
  - (Song Note Object)
