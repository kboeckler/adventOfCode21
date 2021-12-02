package com.github.kboeckler.adventOfCode;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public class FileInput {

  private final Path filepath;

  private FileInput(String filepath) {
    this.filepath = Path.of(filepath);
  }

  static FileInput of(String filepath) {
    String filepathFromResource = FileInput.class.getClassLoader().getResource(filepath).getFile();
    return new FileInput(filepathFromResource);
  }

  String asString() {
    try {
      return Files.readString(filepath);
    } catch (IOException e) {
      throw new RuntimeException("Error reading " + filepath, e);
    }
  }
}
