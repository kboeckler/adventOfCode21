package com.github.kboeckler.adventOfCode;

public class Day1 implements Solution {

  @Override
  public int solvePart1(String input) {
    String[] rows = input.split("\n");
    Integer lastValue = null;
    int countOfIncrements = 0;
    for (String row : rows) {
      int value = Integer.parseInt(row);
      if (lastValue != null && value > lastValue) {
        countOfIncrements++;
      }
      lastValue = value;
    }
    return countOfIncrements;
  }

  @Override
  public int solvePart2(String input) {
    String[] rows = input.split("\n");
    Integer lastValue = null;
    int countOfIncrements = 0;
    for (int i = 0; i <= rows.length - 3; i++) {
      int value =
          Integer.parseInt(rows[i]) + Integer.parseInt(rows[i + 1]) + Integer.parseInt(rows[i + 2]);
      if (lastValue != null && value > lastValue) {
        countOfIncrements++;
      }
      lastValue = value;
    }
    return countOfIncrements;
  }
}
