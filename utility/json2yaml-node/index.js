/** @format */

const fs = require("fs");
const jsYaml = require("js-yaml");

const inputFilePath = process.argv[2];
const outputFilePath = process.argv[3];

try {
  const jsonData = fs.readFileSync(inputFilePath, "utf8");

  const parsedData = JSON.parse(jsonData);

  const yamlData = jsYaml.dump(parsedData);

  fs.writeFileSync(outputFilePath, yamlData, "utf8");
} catch (error) {
  console.error(`Error reading/writing file: ${error.message}`);
  process.exit(1);
}
