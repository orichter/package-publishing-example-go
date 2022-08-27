# Golang SDKs for Nutanix APIs
Staging Area for Golang SDKs of Nutanix APIs to be published externally

# Versions
| Namespace    | Description                                  | Version          |
|--------------|----------------------------------------------|------------------|
| vmm          | Virtual Machine Management                   | 4.0.1-alpha.1    |
| prism        | Tasks and Monitoring                         | 4.0.1-alpha.1    |
| clustermgmt  | Infra Management                             | 4.0.1-alpha.1    |
| aiops        | Alerts and Resources                         | 4.0.1-alpha.1    |
| storage      | Storage Containers                           | 4.0.1-alpha.2    |
| iam          | Identity and Access Management               | 4.0.1-alpha.1    |

# Project Structure
Project will contain a top level directory corresponding to each namespace as listed above.The directory name would be
{namespace}-go-client

Each namespace directory will contain the following sub-directory/files

|Name                  | Description                                       |  
|----------------------|---------------------------------------------------|
|go.mod                | Root dependency management                        |
|go.sum                | hash of dependencies                              |
|License.txt           | License for the client                            |
|README.md             | README  for the client                            |
|api                   | Client containing Go source files for api         |
|client                | Client containing Go source files for client      |
|model                 | Client containing Go source files for model       |

# Publishing process

- Checkout the main branch
- At the root of the directory, run the following command
    ```shell
    git grep -l 'orichter/package-publishing-example-go' | xargs sed -i '' -e 's#orichter/package-publishing-example-go#nutanix/ntnx-api-go-clients#g'
    ```
- Step above will change all the go imports to the new location (https://github.com/nutanix/ntnx-api-go-clients)
- For each namespace directory
  - Push directory to https://github.com/nutanix/ntnx-api-go-clients creating a commit 
  - Tag commit with {namesapce}/{version} from table above
    ```
       vmm-go-client/v4.0.0-alpha.1
    ```

## Architecture

![image (1)](https://user-images.githubusercontent.com/50940815/174028904-fb7c5869-c033-4d19-a101-25388f9d09aa.png)

## Important Note:
```diff
- SDK packages are uploaded through an automated pipeline. No manual changes must be made.
```
Click [here](https://docs.github.com/en/packages/learn-github-packages/introduction-to-github-packages) for more info about github packages.

#### Contact List
1. [Sujeet Mishra(Staging Publishing)](mailto:sujeet.mishra@nutanix.com)
2. [Athar Ejaz(Staging Publishing)](mailto:athar.ejaz@nutanix.com)
3. [Gagan Singh(Staging Publishing)](mailto:gagan.singh@nutanix.com)
4. [Owen Richter(External Publishing)](mailto:owen.richter@nutanix.com)

