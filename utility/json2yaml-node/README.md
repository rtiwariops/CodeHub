<!-- @format -->

## JSON to YAML Converter 💻🔧

This is a simple Node.js program that converts a JSON file to a YAML file. It uses the js-yaml library to perform the conversion.

# Usage 📖

To use this program, you'll need to have Node.js installed on your computer. Then, run the following command in your terminal or command prompt:

```
node index.js <input_file_path> <output_file_path>
```

Replace <input_file_path> with the path to your JSON file, and <output_file_path> with the desired path for the output YAML file.

## Example Output 💻🔜💾

Given the following JSON file:

```
{
  "name": "John Doe",
  "age": 30,
  "address": {
    "street": "123 Main St",
    "city": "Anytown",
    "state": "CA"
  }
}
```

Running the program with the following command:

```
node index.js test.json output.yaml
```

Will produce the following YAML file:

```
name: John Doe
age: 30
address:
  street: 123 Main St
  city: Anytown
  state: CA
```

⚡️ And that's it! Your JSON file has been successfully converted to YAML.
