# ecologi-client-go

[![Treeware (Trees)](https://img.shields.io/treeware/trees/benmarsden/ecologi-client-go)](https://plant.treeware.earth/benmarsden/ecologi-client-go)

Go client for communicating with Ecologi's Impact API.

## Supported Endpoints


| HTTP Method | Provided by (method)      | API Endpoint                      | Description                                                                                   |
| ----------- | ------------------------- | --------------------------------- | --------------------------------------------------------------------------------------------- |
| POST        | `Plant`                   | `/impact/trees`                   | Purchase 1 or more trees                                                                      |
| GET         | `GetTrees`                | `/users/<username>/trees`         | Get the number of trees associated with an Ecologi user                                       |
| GET         | `GetCarbonOffset`         | `/users/<username>/carbon-offset` | Get the number (tonnes) of carbon offsets associated with an Ecologi user                     |
| GET         | `GetTreesAndCarbonOffset` | `/users/<username>/impact`        | Get the number of trees and number (tonnes) of carbon offsets associated with an Ecologi user |


## Contributing

See [How to contribute to ecologi-client-go](CONTRIBUTING.md).

## Software Distribution

This package is [Treeware](https://treeware.earth). If you would like to thank me for creating it, I ask that you [**buy the world a tree**](https://plant.treeware.earth/benmarsden/ecologi-client-go).
