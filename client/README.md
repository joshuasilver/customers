# Go API client for client

Customers focuses on solving authentic identification of humans who are legally able to hold and transfer currency within the US. Primarily this project solves [Know Your Customer](https://en.wikipedia.org/wiki/Know_your_customer) (KYC), [Customer Identification Program](https://en.wikipedia.org/wiki/Customer_Identification_Program) (CIP), [Office of Foreign Asset Control](https://www.treasury.gov/about/organizational-structure/offices/Pages/Office-of-Foreign-Assets-Control.aspx) (OFAC) checks and verification workflows to comply with US federal law and ensure authentic transfers. Customers has an objective to be a service for detailed due diligence on individuals and companies for Financial Institutions and services in a modernized and extensible way.

Customer phone numbers and addresses are stored and partially used in KYC/OFAC validation. Arbitrary key/value pairs can be stored for a Customer. Documents and Disclaimers, and their acknowledgement are also stored under a Customer as they're accepted. Bank Accounts, which can be validated with micro-deposits currently, are stored under each Customer.

![](https://raw.githubusercontent.com/adamdecaf/customers/create-accounts/docs/images/customer.png)


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/moov-io/customers](https://github.com/moov-io/customers)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8087*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CustomersApi* | [**AcceptDisclaimer**](docs/CustomersApi.md#acceptdisclaimer) | **Post** /customers/{customerID}/disclaimers/{disclaimerID} | Accept customer disclaimer
*CustomersApi* | [**AddCustomerAddress**](docs/CustomersApi.md#addcustomeraddress) | **Post** /customers/{customerID}/address | Add customer address
*CustomersApi* | [**CreateCustomer**](docs/CustomersApi.md#createcustomer) | **Post** /customers | Create customer
*CustomersApi* | [**CreateCustomerAccount**](docs/CustomersApi.md#createcustomeraccount) | **Post** /customers/{customerID}/accounts | Create Customer Account
*CustomersApi* | [**DecryptAccountNumber**](docs/CustomersApi.md#decryptaccountnumber) | **Post** /customers/{customerID}/accounts/{accountID}/decrypt | Decrypt Account Number
*CustomersApi* | [**DeleteCustomerAccount**](docs/CustomersApi.md#deletecustomeraccount) | **Delete** /customers/{customerID}/accounts | Delete Customer Account
*CustomersApi* | [**GetCustomer**](docs/CustomersApi.md#getcustomer) | **Get** /customers/{customerID} | Retrieve customer
*CustomersApi* | [**GetCustomerAccounts**](docs/CustomersApi.md#getcustomeraccounts) | **Get** /customers/{customerID}/accounts | Get Customer Accounts
*CustomersApi* | [**GetCustomerDisclaimers**](docs/CustomersApi.md#getcustomerdisclaimers) | **Get** /customers/{customerID}/disclaimers | Get customer disclaimers
*CustomersApi* | [**GetCustomerDocumentContents**](docs/CustomersApi.md#getcustomerdocumentcontents) | **Get** /customers/{customerID}/documents/{documentID} | Get customer document
*CustomersApi* | [**GetCustomerDocuments**](docs/CustomersApi.md#getcustomerdocuments) | **Get** /customers/{customerID}/documents | Get customer documents
*CustomersApi* | [**GetLatestOFACSearch**](docs/CustomersApi.md#getlatestofacsearch) | **Get** /customers/{customerID}/ofac | Get latest OFAC search
*CustomersApi* | [**Ping**](docs/CustomersApi.md#ping) | **Get** /ping | Ping Customers
*CustomersApi* | [**RefreshOFACSearch**](docs/CustomersApi.md#refreshofacsearch) | **Put** /customers/{customerID}/refresh/ofac | Refresh customer OFAC search
*CustomersApi* | [**ReplaceCustomerMetadata**](docs/CustomersApi.md#replacecustomermetadata) | **Put** /customers/{customerID}/metadata | Update customer metadata
*CustomersApi* | [**UpdateCustomerStatus**](docs/CustomersApi.md#updatecustomerstatus) | **Put** /customers/{customerID}/status | Update customer status
*CustomersApi* | [**UploadCustomerDocument**](docs/CustomersApi.md#uploadcustomerdocument) | **Post** /customers/{customerID}/documents | Upload document
*CustomersApi* | [**ValidateAccount**](docs/CustomersApi.md#validateaccount) | **Post** /customers/{customerID}/accounts/{accountID}/validate | Validate Account


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountStatus](docs/AccountStatus.md)
 - [AccountType](docs/AccountType.md)
 - [CreateAccount](docs/CreateAccount.md)
 - [CreateCustomer](docs/CreateCustomer.md)
 - [CreateCustomerAddress](docs/CreateCustomerAddress.md)
 - [CreatePhone](docs/CreatePhone.md)
 - [Customer](docs/Customer.md)
 - [CustomerAddress](docs/CustomerAddress.md)
 - [CustomerMetadata](docs/CustomerMetadata.md)
 - [CustomerStatus](docs/CustomerStatus.md)
 - [CustomerType](docs/CustomerType.md)
 - [Disclaimer](docs/Disclaimer.md)
 - [Document](docs/Document.md)
 - [Error](docs/Error.md)
 - [InstitutionAddress](docs/InstitutionAddress.md)
 - [InstitutionDetails](docs/InstitutionDetails.md)
 - [OfacSearch](docs/OfacSearch.md)
 - [Phone](docs/Phone.md)
 - [TransitAccountNumber](docs/TransitAccountNumber.md)
 - [UpdateCustomerStatus](docs/UpdateCustomerStatus.md)
 - [UpdateValidation](docs/UpdateValidation.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



