# DDC-SDK-Go

[![Go](https://github.com/bianjieai/ddc-sdk-go/actions/workflows/go.yml/badge.svg)](https://github.com/bianjieai/ddc-sdk-go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go)](https://goreportcard.com/report/github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go)
[![License：Apache](https://camo.githubusercontent.com/13258d937f88709447768f3df4a63170ec889e741d0feaa1d5b2c3f8536dc567/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6963656e73652f697269736e65742f697269736875622e737667)](https://www.apache.org/licenses/LICENSE-2.0)

- [DDC-SDK-Go](#ddc-sdk-go)
    - [平台方可调用的如下方法：](#平台方可调用的如下方法)
        - [1.初始化Client (连接测试网)](#1初始化client-连接测试网)
        - [2.BSN-DDC-权限管理](#2bsn-ddc-权限管理)
        - [3.BSN-DDC-费用管理](#3bsn-ddc-费用管理)
        - [4.BSN-DDC-721](#4bsn-ddc-721)
        - [5.BSN-DDC-1155](#5bsn-ddc-1155)
        - [6.BSN-DDC-交易查询](#6bsn-ddc-交易查询)
        - [7.BSN-DDC-区块查询](#7bsn-ddc-区块查询)
        - [8.BSN-DDC-数据解析](#8bsn-ddc-数据解析)
        - [9.离线账户创建](#9离线账户创建)
    - [终端用户可调用的如下方法：](#终端用户可调用的如下方法)
        - [1.初始化Client (连接测试网)](#1初始化client-连接测试网-1)
        - [2.BSN-DDC-权限管理](#2bsn-ddc-权限管理-1)
        - [3.BSN-DDC-费用管理](#3bsn-ddc-费用管理-1)
        - [4.BSN-DDC-721](#4bsn-ddc-721-1)
        - [5.BSN-DDC-1155](#5bsn-ddc-1155-1)
        - [6.BSN-DDC-交易查询](#6bsn-ddc-交易查询-1)
        - [7.BSN-DDC-区块查询](#7bsn-ddc-区块查询-1)
        - [8.BSN-DDC-数据解析](#8bsn-ddc-数据解析-1)
        - [9.离线账户创建](#9离线账户创建-1)
    - [测试用例](#测试用例)

## 平台方可调用的如下方法：

### 1.初始化Client (连接测试网)