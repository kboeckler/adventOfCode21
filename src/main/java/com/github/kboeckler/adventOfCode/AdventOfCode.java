package com.github.kboeckler.adventOfCode;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.Set;
import java.util.stream.Collectors;

public class AdventOfCode {

  public static void main(String[] args) {
    Set<Class<?>> allClassesUsingClassLoader = findAllClassesUsingClassLoader(
        AdventOfCode.class.getPackageName());
    allClassesUsingClassLoader.stream().filter(Solution.class::isAssignableFrom)
        .filter(clazz -> !clazz.equals(Solution.class))
        .map(clazz -> {
          try {
            return (Solution) clazz.getDeclaredConstructors()[0].newInstance();
          } catch (Exception e) {
            throw new RuntimeException("Error creating solution " + clazz.getSimpleName(), e);
          }
        }).forEach(solution -> System.out.printf("Solution %s - Part1: %d Part2: %d%n",
            solution.getClass().getSimpleName(),
            solution.solvePart1(FileInput.of(
                solution.getClass().getSimpleName() + ".txt").asString()),
            solution.solvePart2(FileInput.of(
                solution.getClass().getSimpleName() + ".txt").asString())));
  }

  private static Set<Class<?>> findAllClassesUsingClassLoader(String packageName) {
    InputStream stream = ClassLoader.getSystemClassLoader()
        .getResourceAsStream(packageName.replaceAll("[.]", "/"));
    BufferedReader reader = new BufferedReader(new InputStreamReader(stream));
    return reader.lines()
        .filter(line -> line.endsWith(".class"))
        .map(line -> getClass(line, packageName))
        .collect(Collectors.toSet());
  }

  private static Class<?> getClass(String className, String packageName) {
    try {
      return Class.forName(packageName + "."
          + className.substring(0, className.lastIndexOf('.')));
    } catch (ClassNotFoundException e) {
      throw new RuntimeException("Error getting class", e);
    }
  }

}
