package com.github.kboeckler.adventOfCode;

public class Day2 implements Solution {

  @Override
  public int solvePart1(String input) {
    String[] rows = input.split("\n");
    int horizontalPos = 0;
    int depth = 0;
    for (String row : rows) {
      String[] move = row.split(" ");
      String moveCommand = move[0];
      int moveValue = Integer.parseInt(move[1]);
      switch (moveCommand) {
        case "forward": {
          horizontalPos += moveValue;
          break;
        }
        case "down": {
          depth += moveValue;
          break;
        }
        case "up": {
          depth -= moveValue;
          break;
        }
        default:
          throw new IllegalArgumentException("Unknown move command: " + moveCommand);
      }
    }
    return horizontalPos * depth;
  }

  @Override
  public int solvePart2(String input) {
    String[] rows = input.split("\n");
    int horizontalPos = 0;
    int depth = 0;
    int currentAim = 0;
    for (String row : rows) {
      String[] move = row.split(" ");
      String moveCommand = move[0];
      int moveValue = Integer.parseInt(move[1]);
      switch (moveCommand) {
        case "forward": {
          horizontalPos += moveValue;
          depth += currentAim * moveValue;
          break;
        }
        case "down": {
          currentAim += moveValue;
          break;
        }
        case "up": {
          currentAim -= moveValue;
          break;
        }
        default:
          throw new IllegalArgumentException("Unknown move command: " + moveCommand);
      }
    }
    return horizontalPos * depth;
  }

}
