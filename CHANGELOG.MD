# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

<!-- insertion marker -->
## [v0.0.11](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.11) - 2024-12-21

<small>[Compare with v0.0.10](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.10...v0.0.11)</small>

### Features

- api call & attempt console output updated ([28aa9d4](https://github.com/tsypuk/aws-client-monitor/commit/28aa9d4cd70a15a31581887bc4fb7be1edd25ec3) by Roman Tsypuk).
- updated cli output ([081f17c](https://github.com/tsypuk/aws-client-monitor/commit/081f17cf99c900edabcf671b3d832cf588d5d38e) by Roman Tsypuk).
- logging to console subscription ([b542a1e](https://github.com/tsypuk/aws-client-monitor/commit/b542a1e7d9845c632b1683f44667f4d3676d6e1a) by Roman Tsypuk).
- ui updated ([dece98e](https://github.com/tsypuk/aws-client-monitor/commit/dece98ee513ba431b5ce21c6fccf382b6d66c79d) by Roman Tsypuk).
- websocket reconnect on backend available ([9d03c94](https://github.com/tsypuk/aws-client-monitor/commit/9d03c949273f0dca380f02259cfef28d156670b8) by Roman Tsypuk).
- default color for clients dashboard ([4efcc99](https://github.com/tsypuk/aws-client-monitor/commit/4efcc99f868916edbfb029eb6fc3f0caa565c072) by Roman Tsypuk).
- default color for empty clients ([e7cd9c9](https://github.com/tsypuk/aws-client-monitor/commit/e7cd9c96c6aa84a4c6d18787aaf06e6d84f31d86) by Roman Tsypuk).

### Chore

- move to color pallete ([fef7f3e](https://github.com/tsypuk/aws-client-monitor/commit/fef7f3ecae78fdb5dd43d23a46a15b93cfc3e3ed) by Roman Tsypuk).
- updated eslints ([b08d540](https://github.com/tsypuk/aws-client-monitor/commit/b08d5404b99c964df688531920c1db63601e9d49) by Roman Tsypuk).

## [v0.0.10](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.10) - 2024-10-20

<small>[Compare with v0.0.9](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.9...v0.0.10)</small>

### Features

- clients counter has progress sectors per each client with individual colors ([46cd0b7](https://github.com/tsypuk/aws-client-monitor/commit/46cd0b7b17ab4e702321e5a9966868efaea482bb) by Roman Tsypuk).
- added multi box component ([993507e](https://github.com/tsypuk/aws-client-monitor/commit/993507e1e0f416bd0736a6f945192fa18949a68d) by Roman Tsypuk).
- websocket move client counter to statistics ([41c7c43](https://github.com/tsypuk/aws-client-monitor/commit/41c7c43432946ed0321e0cd0c97411ba21a1be12) by Roman Tsypuk).
- added muti circle progress bar component ([b51aa6e](https://github.com/tsypuk/aws-client-monitor/commit/b51aa6e538f0e2915394fd01519ab262cbde9df4) by Roman Tsypuk).
- added base api type; changed validation flow ([f15e11c](https://github.com/tsypuk/aws-client-monitor/commit/f15e11c04df8b3bfa70c1ca6eaa627de6433ba2f) by Roman Tsypuk).
- added aws region traffic realtime monitoring ([0646044](https://github.com/tsypuk/aws-client-monitor/commit/0646044d54cfdd31d5d30e9ce447f4ee70c14750) by Roman Tsypuk).
- added aws clients realtime counter ([539d5c5](https://github.com/tsypuk/aws-client-monitor/commit/539d5c5be0867d9973e79c00d55a861c881d082f) by Roman Tsypuk).
- added client and region tracking ([afd422a](https://github.com/tsypuk/aws-client-monitor/commit/afd422ab41b6b27036464bf9fae7b9564a4b67ac) by Roman Tsypuk).
- updated APiCallAttempts component ([023190e](https://github.com/tsypuk/aws-client-monitor/commit/023190e1d5730a007be42434fdd8cf00a43160f7) by Roman Tsypuk).
- updated websocket listeners ([5116ac1](https://github.com/tsypuk/aws-client-monitor/commit/5116ac18b729d8df73f0b89aaa63e72c3769f316) by Roman Tsypuk).
- added parsing of type in socket service ([bf7744d](https://github.com/tsypuk/aws-client-monitor/commit/bf7744dc5f517fc9459629e2d1aa643fcbf6b733) by Roman Tsypuk).
- injected ApiCallsAttempts.tsx component to main page ([c1eda7e](https://github.com/tsypuk/aws-client-monitor/commit/c1eda7e26b68ece2fe571f82253137283f13ed3c) by Roman Tsypuk).
- created ApiCall component ([5d7a5f0](https://github.com/tsypuk/aws-client-monitor/commit/5d7a5f02928bba6b733f025101b813c77cdd9eae) by Roman Tsypuk).
- updated socket structure ([355f924](https://github.com/tsypuk/aws-client-monitor/commit/355f924181fc6f2ae1484f19f39a8523cda6100c) by Roman Tsypuk).
- added websocket for apicall and apicallattempt ([b60ca9b](https://github.com/tsypuk/aws-client-monitor/commit/b60ca9bdb75b2edb15f4c95079973b331f2f900b) by Roman Tsypuk).

### Chore

- updated changelog ([23aa89e](https://github.com/tsypuk/aws-client-monitor/commit/23aa89ec7818b2789895294dea306130dec0e8f8) by Roman Tsypuk).

## [v0.0.9](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.9) - 2024-10-18

<small>[Compare with v0.0.8](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.8...v0.0.9)</small>

### Features

- resolve websocket origin check ([2a2eb08](https://github.com/tsypuk/aws-client-monitor/commit/2a2eb0817aedc4c43dafeaedc4750ad71f919336) by Roman Tsypuk).
- updated all pages with react directive ([1ac384e](https://github.com/tsypuk/aws-client-monitor/commit/1ac384e31173c26c660349cd69c33a1e0430309d) by Roman Tsypuk).
- updated router ([c3d67ab](https://github.com/tsypuk/aws-client-monitor/commit/c3d67ab7f34126d2bcf7370050a80a07c1886b4e) by Roman Tsypuk).
- added backend staus indicator component ([ab75e6f](https://github.com/tsypuk/aws-client-monitor/commit/ab75e6fbbf9e947e60b87cfc09301014b0a1f9b7) by Roman Tsypuk).
- added websocket indicator component ([319873a](https://github.com/tsypuk/aws-client-monitor/commit/319873ad26ffffac726c973442227119261c3855) by Roman Tsypuk).
- dockerfile updated with react admin builder step ([a4b273b](https://github.com/tsypuk/aws-client-monitor/commit/a4b273b22c81adc1214e1cf8cf4da722633fe34f) by Roman Tsypuk).
- added websocket service ([9161bd5](https://github.com/tsypuk/aws-client-monitor/commit/9161bd53b8b91c12efd504bab8f72a5802f4ed14) by Roman Tsypuk).

### Chore

- updated github test actions ([59fd8c5](https://github.com/tsypuk/aws-client-monitor/commit/59fd8c5878d906e06c786c07c16a4cda4b96f086) by Roman Tsypuk).
- updated es lint config ([6f97696](https://github.com/tsypuk/aws-client-monitor/commit/6f976965c6309bd26a8e8e5a4e45e3d83d80a62c) by Roman Tsypuk).
- updated npm command with force ([b7023e5](https://github.com/tsypuk/aws-client-monitor/commit/b7023e5195bed74abff13992959eee9b15ad20c1) by Roman Tsypuk).
- added build react to tests github actions ([9190289](https://github.com/tsypuk/aws-client-monitor/commit/919028990b58b38f56ed7f1d96d317cb5e455ded) by Roman Tsypuk).
- updated eslintrc with exclusion ([547030b](https://github.com/tsypuk/aws-client-monitor/commit/547030be5a1991db3b2eb4256520dc12ba0e5a2f) by Roman Tsypuk).

## [v0.0.8](https://github.com/tsypuk/aws-client-monitor/releases/tag/v0.0.8) - 2024-10-17

<small>[Compare with v0.0.7](https://github.com/tsypuk/aws-client-monitor/compare/v0.0.7...v0.0.8)</small>

### Features

- enabled cors for react:3000 app ([7527ac9](https://github.com/tsypuk/aws-client-monitor/commit/7527ac9db28931bcb6f929d98a56cb76728624dc) by Roman Tsypuk).
- status indicator added to index ([c473ae0](https://github.com/tsypuk/aws-client-monitor/commit/c473ae00d29c7c9d97f5d579b915340b8d16fe20) by Roman Tsypuk).
- added status indicator component ([28fc901](https://github.com/tsypuk/aws-client-monitor/commit/28fc9018b688740bf5e33f1a457431cb2dc07087) by Roman Tsypuk).
- added ES ts support ([43711d5](https://github.com/tsypuk/aws-client-monitor/commit/43711d547047afbdda4ccb4b902cc8d00d2bcb77) by Roman Tsypuk).
- app adoption ([2178529](https://github.com/tsypuk/aws-client-monitor/commit/21785290e4a838068d0e25a9eb2c77976d8cceb7) by Roman Tsypuk).

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

