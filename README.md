
# Shuffle Sort Visualization

This project is a visualization of the "shuffle sort" algorithm using Go and the `tcell` library. The program generates a random array of integers and then continuously shuffles it until it becomes sorted, displaying each shuffle attempt in real-time. Performance metrics such as elapsed time, iteration count, and operations per second are shown on the screen.

## Features

- **Real-Time Visualization**: Displays the array in a grid, with each element represented by a vertical bar.
- **Performance Metrics**: Shows the current iteration, total elapsed time, and operations per second.
- **Shuffle Sort Algorithm**: Demonstrates the inefficient, but visually interesting, shuffle sort algorithm.

## Installation

1. **Clone this repository**.
2. Install the required `tcell` package:
   ```bash
   go get github.com/gdamore/tcell/v2
   ```
3. Run the program:
   ```bash
   go run main.go
   ```

## How It Works

The program uses the shuffle sort algorithm:
1. Randomly initializes an integer array.
2. Checks if the array is sorted.
3. If not, shuffles the array and counts the attempt.
4. If sorted, displays the final sorted array along with metrics.

To quit the program, press any key.

## Code Overview

### Key Functions

- `sliceRandomInitialize`: Initializes the array with random integers.
- `shuffleSlice`: Shuffles the array.
- `isSorted`: Checks if the array is sorted.
- `visualizeArray`: Updates the screen with the array and performance metrics.

## Example Output

The program will display an array visualization with the following information:

```
Iteration: [current iteration]
Speed: [operations per second] ops/sec
Time Elapsed: [elapsed time]
```

## Dependencies

- `tcell` library for handling terminal graphics.

## Notes

This project serves as an example of inefficient sorting algorithms and is best used for educational or entertainment purposes. Due to the nature of shuffle sort, the number of iterations may be very high.

## License

This project is licensed under the MIT License.
