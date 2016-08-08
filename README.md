# Procedural generation as a service

## Maze

### GET `/maze`
  * Returns a short intro message

### GET `/maze/generate/backtrack`
  * Parameters:  

  | Name | Description      | Type   | Restrictions  | Default Value | Optional |
  | ---- | ---------------- | ------ | ------------- | ------------- | -------- |
  | w    | Width Dimension  | Number | 1 <= w <= 500 | 10            | Yes      |
  | h    | Height Dimension | Number | 1 <= h <= 500 | 10            | Yes      |
  | seed | Seed             | Number | int64 bounds  | 0             | Yes      |

  * Output (JSON):
  
        {
          "width": Number,
          "height": Number,
          "entrance": {
            "x": Number,
            "y": Number
          },
          "exit": {
            "x": Number,
            "y": Number
          },
          "grid": [
            "point": {
              "x": Number,
              "y": Number
            },
            "next": [{
              "x": Number,
              "y": Number
            }]
          ]
        }

  * Example: `/maze/generate/backtrack?w=10&h=10&seed=12345`



### GET `/maze/generate/backtrack/image`
  * Parameters:  

  | Name | Description      | Type   | Restrictions  | Default Value | Optional |
  | ---- | ---------------- | ------ | ------------- | ------------- | -------- |
  | w    | Width Dimension  | Number | 1 <= w <= 100 | 10            | Yes      |
  | h    | Height Dimension | Number | 1 <= h <= 100 | 10            | Yes      |
  | seed | Seed             | Number | int64 bounds  | 0             | Yes      |

  * Output: PNG
  * Default Output:  
    ![](http://i.imgur.com/l6c7JFz.png)

  * Example: `/maze/generate/backtrack/image?w=10&h=10&seed=12345`
