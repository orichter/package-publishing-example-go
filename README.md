# [DPRO Public Artifacts Publishing Overview](https://confluence.eng.nutanix.com:8443/display/~owen/DPRO+Public+Artifacts+Overview)

Publishing public artifacts such as Software Development Kits (SDKs) which are meant to reflect official releases from Nutanix can be a challenging prospect. It requires software developers to understand not only the requirements of the underlying SDK, but also the publishing process to publish to various external repositories. Furthermore, the credentials to access those repositories must be shared across all those who wish to publish under the official Nutanix accounts. We have developed a reusable way for Nutanix developers to publish public artifacts indirectly using our private Nutanix Github namespaces as a staging area completely under the control of developers, followed by a publishing job which pulls from those private staging repositories and then publishes them first to public staging repositories for review, followed by publication of reviewed and validated release-candidates to public repositories for consumption. The SDK developers maintain complete control over the publication process to internal servers, but the Dev Ops Team maintains control over keys and process for publishing these internal packages publicly.

Our first example for the Nutanix API SDK involves publishing packages for Javascript, Java, Python, and Golang to  https://www.npmjs.com, https://search.maven.org, https://pypi.org, and https://github.com/nutanix respectively.

The developer team first publishes their internal staging publications to the following repositories:
https://github.com/nutanix-core/ntnx-api-javascript-sdk-external
https://github.com/nutanix-core/ntnx-api-java-sdk-external
https://github.com/nutanix-core/ntnx-api-python-sdk-external
https://github.com/nutanix-core/ntnx-api-golang-sdk-external

In the case of javascript and java, the process uses Github Packages to publish packages in a manner nearly identical to npmjs.com, and Maven Central, but simply with a different endpoint. In the case of Python and Golang, the packages are published directly as Github Repositories.
Once a release-candidate has been tagged internally, those tags are used by the Publishing Orchestration Repository to validate the release, run security scans, and publish the release candidate to staging repositories such as:
https://www.npmjs.com/package/@nutanix-scratch/release-canadidate-javascript-sdk
https://test.pypi.org/project/release-candidate-categories-sdk/
Once validate, these release candidates will be promoted to the final public repositories.
We are. currently entering the final stages of our prototype publication for the Nutanix API SDK, but once we have completed this project, we will be reaching out to other teams who have individually published public SDKs to see if we can help facilitate a more streamlined, secure, verifiable and trusted publication system for Nutanix as a whole.

[Detailed Overview of Internal Publishing](https://docs.google.com/document/d/1WLuUmSzDS6cdTuVJKz9uGjKwQ0MIlBT33k0bCPr7bz8) 
