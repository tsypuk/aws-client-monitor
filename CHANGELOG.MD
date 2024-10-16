# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

<!-- insertion marker -->
## [v0.0.7](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.7) - 2024-10-16

<small>[Compare with v0.0.6](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.6...v0.0.7)</small>

### Features

- include react-app into gin static endpoint ([bc9a6bf](https://github.com/tsypuk/aws-client-monitor/commit/bc9a6bfd795c42afef8d31a640e047256cd37943) by Roman Tsypuk).
- added ui admin panel react-app ([b902739](https://github.com/tsypuk/aws-client-monitor/commit/b90273914aa83c5637ca6f2d18ef78486fb78f49) by Roman Tsypuk).
- update channel naming for logging ([89802aa](https://github.com/tsypuk/aws-client-monitor/commit/89802aa1fa687e9e46b66e98146184f781810eb5) by Roman Tsypuk).

## [v0.0.6](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.6) - 2024-10-14

<small>[Compare with v0.0.5](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.5...v0.0.6)</small>

### Features

- post lint refactor ([ae7a5f7](https://github.com/tsypuk/aws-client-monitor/commit/ae7a5f78651764ccb4301bfbb8197e3f44ce3fad) by Roman Tsypuk).
- added constructor ([5cf400c](https://github.com/tsypuk/aws-client-monitor/commit/5cf400cb98492670883487b8d0cc14d63cb99531) by Roman Tsypuk).
- added omitempty fields ([67f3b83](https://github.com/tsypuk/aws-client-monitor/commit/67f3b836cd1ba5183012a2bf6a170dd525eadb5e) by Roman Tsypuk).
- added ApiCall test ok ([6e5389a](https://github.com/tsypuk/aws-client-monitor/commit/6e5389ad288471a5ff04654e366a570847b3bc19) by Roman Tsypuk).
- added ApiCall test ([263e3bb](https://github.com/tsypuk/aws-client-monitor/commit/263e3bb787dbc372e7412482b302504c4b9fa4ae) by Roman Tsypuk).
- added ApiCall constructor ([98f3d8c](https://github.com/tsypuk/aws-client-monitor/commit/98f3d8cf9c65d3a5b6f186a48bc1a4cb5b773025) by Roman Tsypuk).
- added ApiCallAttempt validation ([8e77de7](https://github.com/tsypuk/aws-client-monitor/commit/8e77de72cb86f5035524dd9b0eb7bfe4e2e247a8) by Roman Tsypuk).
- added apicall validation ([10ce2d3](https://github.com/tsypuk/aws-client-monitor/commit/10ce2d35fd44dfa5d7ae2f33a60451910260945c) by Roman Tsypuk).
- dockerfile build static content included ([5e98053](https://github.com/tsypuk/aws-client-monitor/commit/5e98053130e75ab73aa4b9ec5fc6735854c5abff) by Roman Tsypuk).
- updated styles ([231da9b](https://github.com/tsypuk/aws-client-monitor/commit/231da9b50f65ae0f26a988f40cd74d83d8d05d60) by Roman Tsypuk).
- integrated realtime dashboard ([74a790f](https://github.com/tsypuk/aws-client-monitor/commit/74a790fb6981537b3ba2656f1a8e7fb447044c6c) by Roman Tsypuk).
- loading menu with table sorting ([5612e7e](https://github.com/tsypuk/aws-client-monitor/commit/5612e7ed1802710e75916b1b934f999bff46143f) by Roman Tsypuk).
- added dashboard with style ([6c97526](https://github.com/tsypuk/aws-client-monitor/commit/6c97526ed33df4ad7e2fae0fe156d596aae47b0c) by Roman Tsypuk).
- serve realtime dashboard ([d72340b](https://github.com/tsypuk/aws-client-monitor/commit/d72340b3f73ef80827ea4f24617dc8dff7eafc68) by Roman Tsypuk).
- added udp to processing side ([1e8e01e](https://github.com/tsypuk/aws-client-monitor/commit/1e8e01e0e3f1be130c88f9e0caeced68d97b39a1) by Roman Tsypuk).
- added udp payload ([9cd93a1](https://github.com/tsypuk/aws-client-monitor/commit/9cd93a1b8cb8fd050acd9c4926d5af4e8a39f140) by Roman Tsypuk).
- updated dockefile build ([232d624](https://github.com/tsypuk/aws-client-monitor/commit/232d62481cefd77c17bc9595a57fe0249e8fb08d) by Roman Tsypuk).
- moved main under cmd ([ba6ad7a](https://github.com/tsypuk/aws-client-monitor/commit/ba6ad7a7ad99e1823e515f34d02b4ce6b606bded) by Roman Tsypuk).
- move to structure ([3906f54](https://github.com/tsypuk/aws-client-monitor/commit/3906f54019fb43b43c48b0d1db4085ccbafd36fc) by Roman Tsypuk).
- moved domain events ([a3e6f34](https://github.com/tsypuk/aws-client-monitor/commit/a3e6f346f98216aad3363f0611df32fc146f64b8) by Roman Tsypuk).

### Tests

- added ApiCallAttempt ok sqs test ([82f458e](https://github.com/tsypuk/aws-client-monitor/commit/82f458e3edad829558ccaab89e045218d1f9370a) by Roman Tsypuk).
- added ApiCallAttempt error test ([6744b0b](https://github.com/tsypuk/aws-client-monitor/commit/6744b0bc8e4e965da321317628d98204aafce862) by Roman Tsypuk).

### Chore

- updated test CI pipeline ([39c60fd](https://github.com/tsypuk/aws-client-monitor/commit/39c60fd515d88c6781dbc5309d3373840cd93afa) by Roman Tsypuk).
- added docker image push latest on main commit ([48096fd](https://github.com/tsypuk/aws-client-monitor/commit/48096fd2c2052464653879a2617ec9c627925e6d) by Roman Tsypuk).
- updated go install for linter ([6b9cbc0](https://github.com/tsypuk/aws-client-monitor/commit/6b9cbc02291c05fee6e14419e3639de91b027e20) by Roman Tsypuk).
- updated go install ([70ed37a](https://github.com/tsypuk/aws-client-monitor/commit/70ed37a51a6805d339093f068be675f5dec36187) by Roman Tsypuk).
- remap go version ([3f3d7c4](https://github.com/tsypuk/aws-client-monitor/commit/3f3d7c456fa4c6470ebdba5b30775b96089de261) by Roman Tsypuk).
- added tests CI ([a6d82e5](https://github.com/tsypuk/aws-client-monitor/commit/a6d82e5432119ccbf30d2e49674c41c64504c2ce) by Roman Tsypuk).
- updated gitignore ([dadd361](https://github.com/tsypuk/aws-client-monitor/commit/dadd3612b60bc8977558ceaec6f452684133775b) by Roman Tsypuk).

## [v0.0.5](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.5) - 2024-10-13

<small>[Compare with v0.0.4](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.4...v0.0.5)</small>

### Docs

- updated changelog ([1e69984](https://github.com/tsypuk/aws-client-monitor/commit/1e69984cc352966b879b24724e3c4fed3b4ddf39) by Roman Tsypuk).

### Features

- switch to thread safe json serialization library ([79db4e2](https://github.com/tsypuk/aws-client-monitor/commit/79db4e2cb1fe49cc84cbb6ea30320b8ac93238b9) by Roman Tsypuk).
- switch to locked impl ([4d8e855](https://github.com/tsypuk/aws-client-monitor/commit/4d8e855e3a734a392ac622a28fa6bebcd0a93d07) by Roman Tsypuk).
- stabilize websocket ([3368a4e](https://github.com/tsypuk/aws-client-monitor/commit/3368a4e527d3eee34160fdc6f5e2d856fb1f2ee2) by Roman Tsypuk).
- moved to router ([c55dcf8](https://github.com/tsypuk/aws-client-monitor/commit/c55dcf8c6fceaff1110421d5fab611d978df6b46) by Roman Tsypuk).
- docker build is working with swagger endpoint ([a1cbcfc](https://github.com/tsypuk/aws-client-monitor/commit/a1cbcfc89ee833017d734cc754ce6e2649da308c) by Roman Tsypuk).
- switch import to local path ([0d66981](https://github.com/tsypuk/aws-client-monitor/commit/0d66981363ba2a3e28b16ba6d1c317b08853b454) by Roman Tsypuk).
- updated swagger online documentation ([be709c5](https://github.com/tsypuk/aws-client-monitor/commit/be709c53ddef24c6945628427608a36498d1732c) by Roman Tsypuk).
- fixed import of module ([862d381](https://github.com/tsypuk/aws-client-monitor/commit/862d3818f6a8de8a9be1c6673469bd982e2dd667) by Roman Tsypuk).

### Chore

- update go mod with swagger module ([7eb39fd](https://github.com/tsypuk/aws-client-monitor/commit/7eb39fda50b37c7d0faf8a194dd5879b2eebe0e9) by Roman Tsypuk).
- updated go.mod github repo refference ([967043f](https://github.com/tsypuk/aws-client-monitor/commit/967043f8832b93494a831f3c42baa7eb6094186b) by Roman Tsypuk).
- updated dockefile ([1b2e857](https://github.com/tsypuk/aws-client-monitor/commit/1b2e857081f99c37604b570fa28334dfb33baec3) by Roman Tsypuk).

## [v0.0.4](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.4) - 2024-10-12

<small>[Compare with v0.0.3](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.3...v0.0.4)</small>

### Docs

- generated swagger documentation ([5fb2aaf](https://github.com/tsypuk/aws-client-monitor/commit/5fb2aafebbb9cd23668ec441751b080f1123e6b9) by Roman Tsypuk).

### Features

- included v1 swagger endpoint ([670d035](https://github.com/tsypuk/aws-client-monitor/commit/670d035bcdfd0719552c3473da96fb96c8779cfe) by Roman Tsypuk).

### Chore

- added swagger to makefile ([85fed8d](https://github.com/tsypuk/aws-client-monitor/commit/85fed8dd7a92ddbf14b45e0d70260057265d3850) by Roman Tsypuk).
- added swagger dependencies ([c384ae1](https://github.com/tsypuk/aws-client-monitor/commit/c384ae1501ffb0398f646f8a65e755506faaea3a) by Roman Tsypuk).

## [v0.0.3](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.3) - 2024-10-12

<small>[Compare with v0.0.2](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.2...v0.0.3)</small>

### Docs

- added changelog generated ([caf21d4](https://github.com/tsypuk/aws-client-monitor/commit/caf21d41d6ca771ce63668296496cd177c9c3d1a) by Roman Tsypuk).

### Features

- added parsing of ApiCallAttempt payload ([da4ec5c](https://github.com/tsypuk/aws-client-monitor/commit/da4ec5c5dc027c359f8741e468a7b6a2d7209202) by Roman Tsypuk).
- change fatal to logging ([6e72f44](https://github.com/tsypuk/aws-client-monitor/commit/6e72f4409acb277b825854ee50b6f399be7098b7) by Roman Tsypuk).

### Chore

- added makefile ([f5f7c4e](https://github.com/tsypuk/aws-client-monitor/commit/f5f7c4e7608a345e7470f2a21620ee508b1f20e0) by Roman Tsypuk).
- updated github actions builds and push ([4f2f03b](https://github.com/tsypuk/aws-client-monitor/commit/4f2f03b30452c52cae02f4f76f83da8c66d7648a) by Roman Tsypuk).

## [v0.0.2](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.2) - 2024-10-08

<small>[Compare with v0.0.1](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.1...v0.0.2)</small>

### Features

- added github auth ([ba341ff](https://github.com/tsypuk/aws-client-monitor/commit/ba341ffdba7345610db8ee1f92a810fe3aaff9bf) by Roman Tsypuk).
- added release github actions ([319167c](https://github.com/tsypuk/aws-client-monitor/commit/319167c3a5990db394da076ec079e36b024711a1) by Roman Tsypuk).
- updated secrets propagation to git actions ([61079f8](https://github.com/tsypuk/aws-client-monitor/commit/61079f892605bc2c5b84d3e2e4a2309040903951) by Roman Tsypuk).
- added css to docker ([35b0367](https://github.com/tsypuk/aws-client-monitor/commit/35b0367c13b832355fdb256fe3f418f564a68dc4) by Roman Tsypuk).
- added loading and rendering of static css content ([4765ed7](https://github.com/tsypuk/aws-client-monitor/commit/4765ed7367f668bed32b024957b45144ce580276) by Roman Tsypuk).
- added generic message parsing ([6329022](https://github.com/tsypuk/aws-client-monitor/commit/63290226b3e28522a7465ef9ecec26b3200da56f) by Roman Tsypuk).
- added dockerfile build ([def8e54](https://github.com/tsypuk/aws-client-monitor/commit/def8e54b01633dac96e0fb5c91300f9b9612f3cf) by Roman Tsypuk).
- added docker multiarch build ([f968eb5](https://github.com/tsypuk/aws-client-monitor/commit/f968eb5a4a6925ad90bd52e4e16fff49610bf1c0) by Roman Tsypuk).
- added dashboard with websocket for latency measurement ([a97eb76](https://github.com/tsypuk/aws-client-monitor/commit/a97eb7647a759d5cda42df81bfbbf1fcffac7912) by Roman Tsypuk).

## [v0.0.1](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.1) - 2024-10-06

<small>[Compare with v0.0.0](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.0...v0.0.1)</small>

### Features

- added UDP listener with output channel ([55dc6c2](https://github.com/tsypuk/aws-client-monitor/commit/55dc6c211acae007ef71500bd51534a7c7cd97a0) by Roman Tsypuk).

## [v0.0.0](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.0) - 2024-10-06

<small>[Compare with first commit](https://github.com/tsypuk/aws-client-monitor/compare/eee74d101b5d4b93ceef4f124d37951b77fd6bfb...v0.0.0)</small>

### Chore

- added gitignore ([eee74d1](https://github.com/tsypuk/aws-client-monitor/commit/eee74d101b5d4b93ceef4f124d37951b77fd6bfb) by Roman Tsypuk).

