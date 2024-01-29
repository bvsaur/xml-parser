# Invoice XML Parser

This is a utility that parses invoice XML files and export certain data to a CSV file.

## How to use it?

1. Clone this repository
2. Build the application using the `go build` command
3. Create this folder structure:
    ```bash
    ├── xml-parser (executable)
    ├── files (invoices folder)
    │   ├── Invoice-1.xml
    │   └── ...
    └──
    ```
4. Run the executable:
    ```bash
    ./xml-parser --folder [folder-name] --file [file-name]
    ```

Note:
- The `folder` option points the name of the folder that contains the invoices
- The `file` option indicates the name of the CSV file that will be exported.