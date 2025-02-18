<p align="center">
    <a href="https://openim.io">
        <img src="./assets/logo-gif/openim-logo.gif" width="60%" height="30%"/>
    </a>
</p>

<div align="center">

[![Stars](https://img.shields.io/github/stars/openimsdk/open-im-server?style=for-the-badge&logo=github&colorB=ff69b4)](https://github.com/openimsdk/open-im-server/stargazers)
[![Forks](https://img.shields.io/github/forks/openimsdk/open-im-server?style=for-the-badge&logo=github&colorB=blue)](https://github.com/openimsdk/open-im-server/network/members)
[![Codecov](https://img.shields.io/codecov/c/github/openimsdk/open-im-server?style=for-the-badge&logo=codecov&colorB=orange)](https://app.codecov.io/gh/openimsdk/open-im-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/openimsdk/open-im-server?style=for-the-badge)](https://goreportcard.com/report/github.com/openimsdk/open-im-server)
[![Go Reference](https://img.shields.io/badge/Go%20Reference-blue.svg?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/openimsdk/open-im-server/v3)
[![License](https://img.shields.io/badge/license-Apache--2.0-green?style=for-the-badge)](https://github.com/openimsdk/open-im-server/blob/main/LICENSE)
[![Slack](https://img.shields.io/badge/Slack-500%2B-blueviolet?style=for-the-badge&logo=slack&logoColor=white)](https://join.slack.com/t/openimsdk/shared_invite/zt-2ijy1ys1f-O0aEDCr7ExRZ7mwsHAVg9A)
[![Best Practices](https://img.shields.io/badge/Best%20Practices-purple?style=for-the-badge)](https://www.bestpractices.dev/projects/8045)
[![Good First Issues](https://img.shields.io/github/issues/openimsdk/open-im-server/good%20first%20issue?style=for-the-badge&logo=github)](https://github.com/openimsdk/open-im-server/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22good+first+issue%22)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Gurubase](https://img.shields.io/badge/Gurubase-Ask%20OpenIM%20Guru-006BFF?style=for-the-badge)](https://gurubase.io/g/openim)

<p align="center">
  <a href="./README.md">English</a> · 
  <a href="./README_zh_CN.md">中文</a> · 
  <a href="./docs/readme/README_uk.md">Українська</a> · 
  <a href="./docs/readme/README_cs.md">Česky</a> · 
  <a href="./docs/readme/README_hu.md">Magyar</a> · 
  <a href="./docs/readme/README_es.md">Español</a> · 
  <a href="./docs/readme/README_fa.md">فارسی</a> · 
  <a href="./docs/readme/README_fr.md">Français</a> · 
  <a href="./docs/readme/README_de.md">Deutsch</a> · 
  <a href="./docs/readme/README_pl.md">Polski</a> · 
  <a href="./docs/readme/README_id.md">Indonesian</a> · 
  <a href="./docs/readme/README_fi.md">Suomi</a> · 
  <a href="./docs/readme/README_ml.md">മലയാളം</a> · 
  <a href="./docs/readme/README_ja.md">日本語</a> · 
  <a href="./docs/readme/README_nl.md">Nederlands</a> · 
  <a href="./docs/readme/README_it.md">Italiano</a> · 
  <a href="./docs/readme/README_ru.md">Русский</a> · 
  <a href="./docs/readme/README_pt_BR.md">Português (Brasil)</a> · 
  <a href="./docs/readme/README_eo.md">Esperanto</a> · 
  <a href="./docs/readme/README_ko.md">한국어</a> · 
  <a href="./docs/readme/README_ar.md">العربي</a> · 
  <a href="./docs/readme/README_vi.md">Tiếng Việt</a> · 
  <a href="./docs/readme/README_da.md">Dansk</a> · 
  <a href="./docs/readme/README_el.md">Ελληνικά</a> · 
  <a href="./docs/readme/README_tr.md">Türkçe</a>
</p>

</div>

</p>

## :busts_in_silhouette: Join Our Community

- 💬 [Follow us on Twitter](https://twitter.com/founder_im63606)
- 🚀 [Join our Slack](https://join.slack.com/t/openimsdk/shared_invite/zt-2ijy1ys1f-O0aEDCr7ExRZ7mwsHAVg9A)
- :eyes: [Join our WeChat Group](https://openim-1253691595.cos.ap-nanjing.myqcloud.com/WechatIMG20.jpeg)

## Ⓜ️ About OpenIM

Unlike standalone chat applications such as Telegram, Signal, and Rocket.Chat, OpenIM offers an open-source instant messaging solution designed specifically for developers rather than as a directly installable standalone chat app. Comprising OpenIM SDK and OpenIM Server, it provides developers with a complete set of tools and services to integrate instant messaging functions into their applications, including message sending and receiving, user management, and group management. Overall, OpenIM aims to provide developers with the necessary tools and framework to implement efficient instant messaging solutions in their applications.

![App-OpenIM Relationship](./docs/images/oepnim-design.png)

## 🚀 Introduction to OpenIMSDK

**OpenIMSDK**, designed for **OpenIMServer**, is an IM SDK created specifically for integration into client applications. It supports various functionalities and modules:

- 🌟 Main Features:

  - 📦 Local Storage
  - 🔔 Listener Callbacks
  - 🛡️ API Wrapping
  - 🌐 Connection Management

- 📚 Main Modules:
  1. 🚀 Initialization and Login
  2. 👤 User Management
  3. 👫 Friends Management
  4. 🤖 Group Functions
  5. 💬 Session Handling

Built with Golang and supports cross-platform deployment to ensure a consistent integration experience across all platforms.

👉 **[Explore the GO SDK](https://github.com/openimsdk/openim-sdk-core)**

## 🌐 Introduction to OpenIMServer

- **OpenIMServer** features include:
  - 🌐 Microservices Architecture: Supports cluster mode, including a gateway and multiple rpc services.
  - 🚀 Diverse Deployment Options: Supports source code, Kubernetes, or Docker deployment.
  - Massive User Support: Supports large-scale groups with hundreds of thousands, millions of users, and billions of messages.

### Enhanced Business Functions:

- **REST API**: Provides a REST API for business systems to enhance functionality, such as group creation and message pushing through backend interfaces.

- **Webhooks**: Expands business forms through callbacks, sending requests to business servers before or after certain events.

  ![Overall Architecture](./docs/images/architecture-layers.png)

## :rocket: Quick Start

Experience online for iOS/Android/H5/PC/Web:

👉 **[OpenIM Online Demo](https://www.openim.io/en/commercial)**

To facilitate user experience, we offer various deployment solutions. You can choose your preferred deployment method from the list below:

- **[Source Code Deployment Guide](https://docs.openim.io/guides/gettingStarted/imSourceCodeDeployment)**
- **[Docker Deployment Guide](https://docs.openim.io/guides/gettingStarted/dockerCompose)**

## System Support

Supports Linux, Windows, Mac systems, and ARM and AMD CPU architectures.

## :link: Links

- **[Developer Manual](https://docs.openim.io/)**
- **[Changelog](https://github.com/openimsdk/open-im-server/blob/main/CHANGELOG.md)**

## :writing_hand: How to Contribute

We welcome contributions of any kind! Please make sure to read our [Contributor Documentation](https://github.com/openimsdk/open-im-server/blob/main/CONTRIBUTING.md) before submitting a Pull Request.

- **[Report a Bug](https://github.com/openimsdk/open-im-server/issues/new?assignees=&labels=bug&template=bug_report.md&title=)**
- **[Suggest a Feature](https://github.com/openimsdk/open-im-server/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=)**
- **[Submit a Pull Request](https://github.com/openimsdk/open-im-server/pulls)**

Thank you for contributing to building a powerful instant messaging solution!

## :closed_book: License

For more details, please refer to [here](./LICENSE).

## 🔮 Thanks to our contributors!

<a href="https://github.com/openimsdk/open-im-server/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=openimsdk/open-im-server" />
</a>
