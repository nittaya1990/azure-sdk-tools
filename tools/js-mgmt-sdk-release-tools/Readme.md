# Azure Track2 JS SDK Tools
This package contains 3 tools:
1. Use Track2 Codegen to generate codes by inputting swagger, and then run `rush update & rush build`.
2. Use Track2 Codegen to generate codes in swagger pipeline.
3. Generate changelog and bump version for Azure Track2 JS SDK.

## Install
```shell script
npm install -g @azure-tools/js-mgmt-release-tools
```

## Tools

### Use Track2 Codegen to generate codes by inputting swagger
You can use this command to generate codes, and then run rush update, rush build, and pack the generated codes.
```shell script
cd azure-sdk-for-js
codegen-automation --use=@autorest/typescript@6.0.0-beta.5 --readme=../azure-rest-api-specs/specification/iotspaces/resource-manager/readme.md
```
hint: you can change the `readme` to yours.

### Use Track2 Codegen to generate codes in swagger pipeline
This command is only used in swagger pipeline.
```shell script
cd azure-sdk-for-js
codegen-automation-for-pipeline --inputJsonPath=a.json --outputJsonPath=b.json --use=@autorest/typescript@6.0.0-beta.5
```

### Generate changelog and bump version for Azure Track2 JS SDK
This tool is used to generate changelog and bump version automatically.

####Prerequisites
This tools is based on the api.md file generated by [API Extractor](https://api-extractor.com/). So make sure the sdk package contains the api.md file

#### How to Use
```shell script
cd azure-sdk-for-js
changelog-tool ${package-path}
```
example:
```
track2-js-sdk-changelog-tool sdk/compute/arm-compute
```

#### Overview - How Changelog Tool Works
In order to compute this changelog, we do the following steps:
1. Input the package path.
2. Get the package name and download the latest release from NPM.
3. If not found it in NPM, it's first release and generate changelog for first release. If found it, check whether package in npm is track2.
4. If the npm package is track1, generated changelog for migrating track1 to track2. If not, generate change log by compare api.md extracted by API Extractor.
5. Changelog Tool parses the typescript blocks in the api.md to ASTs, and extracts useful information from the ASTs.    
6. Changelog Tool compares extracted information and generates the changelog.

##### Convert the TypeScript Block in api.md to AST And Extract Useful information from the AST
The data structure of the extracted information:
```typescript
export class TSExportedMetaData {
    public typeAlias = {};
    public operationInterface = {};
    public modelInterface = {};
    public enums = {};
    public classes = {};
}
```
This is an example of the extracted information:
```json
{
  "typeAlias": {
    "AttachedDatabaseConfiguration": {
      "type": {
        "inherits": [
          "ProxyResource"
        ],
        "typeLiteralDeclarations": [
          {
            "properties": [
              {
                "name": "location",
                "visibility": 2,
                "type": "string",
                "isOptional": true,
                "isStatic": false,
                "start": 290,
                "end": 308
              }              
            ],
            "methods": [],
            "intersectionDeclarations": [],
            "typeLiteralDeclarations": [],
            "name": "",
            "start": 284,
            "end": 606
          }
        ],
        "name": "",
        "start": 268,
        "end": 606
      },
      "methods": [],
      "name": "AttachedDatabaseConfiguration",
      "isExported": true,
      "start": 224,
      "end": 607
    }
  }
}
```

##### Compare Extracted Information And Generate the Changelog
we have a set of rules that we apply for each entry in the diff. Currently template rules are:
```
## Features

Added operation group {}
Added operation {}.{}
Added Interface {}
Added Class {}
Added Type Alias {}
Interface {} has a new optional parameter {}
Class {} has a new optional parameter {}
Add parameters of {} to TypeAlias {}
Type Alias {} has a new parameter {}
Added Enum {}
Enum {} has a new value {}

## Breaking Changes

Removed operation group {}
Removed operation {}.{}
Operation {}.{} has a new signature
Class {} has a new signature
Interface {} no longer has parameter {}
Interface {} has a new required parameter {}
Parameter {} of interface {} is now required
Class {} no longer has parameter {}
Class {} has a new required parameter {}
Parameter {} of class {} is now required
Delete parameters of {} in TypeAlias {}
Type Alias {} no longer has parameter {}
Type Alias {} has a new parameter {}
Parameter {} of Type Alias {} is now required
Removed Enum {}
Enum {} no longer has value None
```
