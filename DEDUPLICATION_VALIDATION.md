# Проверка дедупликации — 2026-11-14

Все семь этапов выполнены локально. `config.yaml` использует `enforce`.

## Согласование после обновления main

База ветки — `ae88b6b168ee396069bc66c0505f6053038a2c48` от 2026-11-14.
18 конфликтов autostash затронули только производные index.json. Новые исходники
из main включены через offline replay `Archive.Upsert` с сохранённым временем
сбора `2026-11-14T12:36:38.232091035Z`, затем выполнены plan/apply и rebuild.

- 1 207 source records, 1 291 ревизия, 1 206 canonical incidents.
- Все 317 vendor JSON побайтово равны обновлённому main.
- Все прежние 1 195 snapshots и 1 194 canonical IDs сохранены.
- Сохранён SourceStatus из upstream для всех 18 конфликтовавших недель.
- Все 84 недели пересобраны; повтор дал 1 949 побайтово идентичных файлов.
- Membership, anchors и ссылки на evidence проверены; pending/conflicts — 0/0.
- `make test` с race detector, `make lint` и `bin/dedup validate` проходят.
- Каждый кодовый коммит проверен на изолированном staged-снимке. Генератор
  Markdown завершает файл одним переводом строки; staged diff проходит
  стандартную проверку whitespace без исключений.

Отчёт: `reports/dedup/main-refresh/verification.json`; план и diff — в том же
каталоге. Резервная копия состояния до разрешения конфликтов:
`/private/tmp/weekly-incident-main-conflicts-9yh2pqem/worktree`. Autostash сохранён.
Код, миграция, выпуски, CI и документация оформлены отдельными атомарными коммитами
в `feature/incident-deduplication`. Push и уведомления не выполнялись.

## Первоначальная миграция до обновления main

Исходный commit для сравнения: `af7a0dc9f1cef868cd03b3b53715c454ceb25337`. Этот commit обозначает исходный baseline, а не последующую реализацию
или актуальную базу ветки.
Время миграции и всех сравнений: `2026-11-14T00:00:00Z`.

## Результат полного прохода

| Проверка | Результат |
| --- | --- |
| Исходные записи | 1 195, 14 вендоров |
| Vendor JSON | 316 файлов; побайтово совпадают с baseline |
| Сохранённые ревизии / snapshots при bootstrap | 1 195 / 1 195 |
| Канонические инциденты | 1 194 |
| Объединения | Одна группа из двух AWS-записей |
| Pending pairs / неразрешённые конфликты | 0 / 0 по текущим точным правилам |
| Разобранные изменения полей/membership | 24 группы: AWS merge + 23 unknown duration |
| Недели с индексом и полным рендером | 84; прежде digest был у 16, создано ещё 68 |
| Одинаковые названия разных компонентов | 21 запись; кратность компонентов сохранена |
| Повторный rebuild с тем же now | 1 852 файла data/weeks побайтово идентичны |
| Membership и ссылки | Каждый source key в одной группе, каждый canonical ID в одном недельном индексе; anchors и ссылки на исходники проверены |

Группа `ci-845c080d871358f626ba3e26` содержит
`aws/9fb26c8e8bd2a9c4` и `aws/6a8fc17b2bdd3e4d`. Все 16 сообщений второй записи
совпадают с частью 22 сообщений первой. В canonical timeline осталось 22
уникальных обновления, RSS GUID извлечены offline из сохранённого raw.
Основание: `aws-containment-v1`.

На неделе 2026-W09 число записей в представлении уменьшилось с 8 до 7, число
выше threshold — с 6 до 5. AWS занимал первые два места, теперь занимает одно.
Его новый score 741 496 при выбранном now равен score более полной исходной
цепочки, а не сумме двух дублей. Это ongoing score по сохранённому investigating,
не подтверждение реальной длительности аварии до ноября.

23 записи содержат завершение раньше начала, причём противоречие присутствует
в raw. Для них созданы решения автора Codex в `dedup/decisions.yaml`, с hash pins
и ссылками на конкретные исходники. Даты сохранены, длительность неизвестна,
вклад длительности в score равен нулю. У всех остальных самостоятельных записей
score и threshold при одинаковом now сохранены. Равные названия компонентов
в 21 записи относятся к различным IDs в raw; их схлопывание было бы ошибкой.

## Проверки реализации

- `make test`: проходят все пакеты с race detector.
- `make lint`: gofmt, go vet, golangci-lint; 0 issues.
- `make build`: collect, digest, dedup собираются.
- `bin/dedup validate`: согласованы snapshots, registry, decisions, canonical
  builder и хеши всех derived artifacts.
- Миграция сначала выполнена на отдельной полной копии, затем план рабочего
  архива сопоставлен с её proposed registry/canonical hashes и применён локально.
- Полный rebuild повторён с фиксированным временем; проверены все 1 852 файла.
- CLI smoke: stdout `bin/digest -week 2026-W09 -now 2026-11-14T00:00:00Z`
  содержит ровно `2026-W09` и перевод строки.
- Digest проверяется в shadow и enforce: legacy membership в shadow,
  canonical membership в enforce, согласованность score и timeline, отказ при
  повреждённом manifest/artifact, сохранность исходников.
- Unit/integration-проверки покрывают stale/edited plan, merge/split/aliases,
  глобальный different_incident, неоднозначный мост, перенос через ISO-год,
  пустую прежнюю неделю, исправление update с тем же ID, A → B → A,
  усечение AWS timeline и добавление resolution, unknown legacy order,
  восстановление после неполной записи, pins устаревших overrides.
- Workflow source-only commit и retry/resync проверены в изолированных локальных
  Git-репозиториях с mock push/fetch/reset. При ошибке canonical build в commit
  не попадают изменённые индекс/registry/canonical files, job остаётся ошибочным.

Held-out corpus: отдельный captured RSS, 5 цепочек, 3 положительных
синтетических усечения и 13 отрицательных сравнений. Все 16 проходят.
Пара 22/16 — отдельный положительный regression fixture из архива.
Синтетические примеры не доказывают полноту или precision на реальном потоке.

## Воспроизводимые артефакты

- План и полный diff: `reports/dedup/plan.json`, `reports/dedup/plan.md`.
- Аудит исходных кандидатов: `reports/dedup/audit.json`, `audit.md`.
- Машинная сводка проверки: `reports/dedup/validation.json`.
- Постоянные результаты: `data/observations/`, `data/identity/registry.json`,
  `data/canonical/`, недельные индексы и `weeks/`.
- Baseline локально: `/private/tmp/weekly-incident-dedup-baseline-20261114`.
- Проверенная копия: `/private/tmp/weekly-incident-dedup-pilot-20261114`.

`reports/dedup` исключён из Git; это воспроизводимые локальные отчёты. Baseline
в /private/tmp временный и не заменяет долговременного backup. Из-за появления
реестра исходный migration plan после apply становится устаревшим; для новых
решений требуется новый plan, а для повтора производных данных — rebuild.

Fingerprint исходного архива:
`f72dadbd5b05a600371ae61c651c6299eeeaf252d956c36f49d63e4e5594ee14`.

Hash реестра после миграции:
`c19ae25e5d9be4d7c50cd07ca606d00b08cc1abc5f940fc1b4afed07747c2a7b`.

Hash canonical JSON:
`61b4553c6c89384ba1f649c6c17e5450d6c5a40493303781f312b7dd617d3279`.

## Что не проверялось в production

GitHub Actions на удалённом runner, реальные push и Telegram-уведомления не
запускались. Исторический rebuild выполнен без уведомлений. Live feed fetch
для миграции не нужен и не выполнялся; использован полный локальный архив.
Не восстанавливались несобранные версии или вся история Git. Автоматический
fuzzy matching и межвендорские merge отсутствуют по границам первой версии.

Число 1 194 означает группы собранных сообщений после текущих подтверждённых
объединений. Ноль pending-кандидатов не означает отсутствия других дублей.
Команды дальнейшей работы и отката: [DEDUPLICATION.md](DEDUPLICATION.md).
