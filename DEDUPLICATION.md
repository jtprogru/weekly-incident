# Работа с дубликатами

В `config.yaml` включён `dedup.mode: enforce`. Архив уже обработан целиком:
1 207 исходных записей → 1 206 канонических инцидентов после обновления main
до `ae88b6b`. Сохранена 1 291 ревизия, включая данные до обновления. Оригинальные vendor JSON
сохранены. Повторная запись исключается из счётчика и выпуска через membership
реестра, а её свидетельства остаются доступны.

Результат относится к локальному архиву и правилам `identity-v1` на 2026-11-14.
Он не доказывает, что найдены все семантические дубли. Проверки миграции описаны
в [DEDUPLICATION_VALIDATION.md](DEDUPLICATION_VALIDATION.md).

## Повторный проход по всему архиву

Из корня репозитория:

```bash
make build
bin/dedup audit -data data -config config.yaml \
  -out reports/dedup -now 2026-11-14T00:00:00Z
bin/dedup plan -data data -config config.yaml \
  -out reports/dedup -now 2026-11-14T00:00:00Z
```

Для нового разбора выбери актуальное фиксированное время вместо даты примера.
Просмотри `reports/dedup/audit.md`, затем `plan.md` и `plan.json`. В JSON есть
состав групп, изменения дат, длительности, score и threshold, top-списки каждой
недели, конфликты, `source_heads` для привязки решений и `revision_options` для
неупорядоченных legacy-версий. Audit показывает исходные кандидаты, поэтому уже
объединённая AWS-пара останется в audit; pending в плане исключает принятые связи
и явно разделённые пары. `plan.changes` сравнивает source-представление с canonical,
а не обещает число новых операций очередного запуска.

После проверки результата:

```bash
bin/dedup apply -data data -config config.yaml -plan reports/dedup/plan.json
bin/dedup rebuild -data data -config config.yaml -weeks weeks \
  -now 2026-11-14T00:00:00Z
bin/dedup validate -data data -config config.yaml
```

`apply` использует время из плана. Любое изменение входов или proposed hashes
делает план устаревшим: сформируй его заново. Одинаковый план второй раз после
первого apply тоже может стать устаревшим, поскольку реестр уже изменился.
`rebuild` пересчитывает данные по сохранённому реестру и не принимает новых
identity-решений. Параметр `-weeks` дополнительно создаёт `digest.md`, `brief.md`
и `summary.txt` для каждой недели архива, включая пустые после переноса/merge.
Без `-weeks` обновляются только canonical JSON и индексы. Команды не отправляют
уведомлений и не выполняют Git commit/push.

Сокращения Makefile: `make dedup-audit`, `make dedup-plan`, `make dedup-apply`,
`make dedup-rebuild`, `make dedup-validate`. Для audit/plan доступен `REPORT_DIR`,
для расчётов — `NOW`. Для нестандартных путей используй бинарник.

## Проверка на копии

Создай отдельный каталог и скопируй туда `data/`, `weeks/`, `dedup/` и
`config.yaml`, сохраняя их взаимное расположение. Затем укажи абсолютные пути:

```bash
bin/dedup plan -data /tmp/incident-review/data \
  -config /tmp/incident-review/config.yaml \
  -out /tmp/incident-review/reports -now 2026-11-14T00:00:00Z
bin/dedup apply -data /tmp/incident-review/data \
  -config /tmp/incident-review/config.yaml \
  -plan /tmp/incident-review/reports/plan.json
bin/dedup rebuild -data /tmp/incident-review/data \
  -config /tmp/incident-review/config.yaml \
  -weeks /tmp/incident-review/weeks -now 2026-11-14T00:00:00Z
bin/dedup validate -data /tmp/incident-review/data \
  -config /tmp/incident-review/config.yaml
```

Путь `dedup.decisions` считается относительно файла конфигурации, а не data root.
Общий флаг `-decisions` у maintenance-команд позволяет явно заменить этот путь.
Report directory не может пересекать исходный архив или каталог выпусков;
входные config/decisions также нельзя помещать внутрь каталога отчёта.

## Что сохраняется

| Путь | Назначение |
| --- | --- |
| `data/YYYY/Wnn/<vendor>.json` | Последние исходные записи; существующий key сохраняет физический путь после исправления даты |
| `data/observations/<vendor>/<hash>.json` | Неизменяемое содержимое нормализованной записи с raw |
| `data/observations/manifest.json` | Порядок ревизий, время наблюдения, импорт legacy-данных |
| `data/identity/registry.json` | Устойчивые IDs, anchors, membership, aliases, автоматические связи, журнал merge/split |
| `dedup/decisions.yaml` | Редактируемые решения с автором, основанием, evidence и supersession |
| `data/canonical/incidents.json` | Поля, timeline, история исправлений сообщений, provenance, конфликты и warnings |
| `data/canonical/manifest.json` | Fingerprints входов, хеши derived artifacts, pending-кандидаты |
| `data/YYYY/Wnn/index.json`, `weeks/` | Скоринг и представления канонических инцидентов |

Исчезновение записи из фида ничего не удаляет. Повтор одинакового fetch не
создаёт ревизию. A → B → A даёт три ревизии и два snapshots. При bootstrap
импортируется только доступное содержимое; дата импорта не выдаётся за дату
редактирования у вендора. У разных legacy-версий одного key порядок неизвестен
до нового fetch или решения `source_revision`. Если старый key физически
повторялся в нескольких файлах, новый fetch обновляет все эти совместимые копии;
canonical membership всё равно учитывает source key один раз.

## Ручные решения и отмена merge

Редактируй `dedup/decisions.yaml`, затем запускай `plan`/`apply`. Не меняй вручную
реестр, snapshots и canonical JSON. `same_incident` объединяет участников одного
вендора; `different_incident` запрещает каждой паре участников оказаться в одной
группе, в том числе через третью запись. При конфликте решений apply блокируется.

Пример разделения текущей автоматической AWS-группы — шаблон для случая, если
появятся доказательства ошибочного merge. Это не рекомендация разделять её сейчас.
Добавить элемент в существующий `decisions`, заполнив автора, дату и evidence:

```yaml
  - id: separate-aws-w09-after-review
    type: different_incident
    members: [aws/9fb26c8e8bd2a9c4, aws/6a8fc17b2bdd3e4d]
    reason: "ОПИСАТЬ НОВОЕ ДОКАЗАТЕЛЬСТВО РАЗНЫХ СОБЫТИЙ"
    evidence: ["ПУТЬ ИЛИ URL С КОНКРЕТНЫМ СВИДЕТЕЛЬСТВОМ"]
    author: "ИМЯ АВТОРА РЕШЕНИЯ"
    at: 2026-11-14T00:00:00Z
```

Если заменяется ручное решение, новый элемент содержит
`supersedes: ID_ПРЕЖНЕГО_РЕШЕНИЯ`; прежний элемент остаётся в истории. Для split
автоматической связи supersedes не требуется. Неизвестные ссылки, циклы и две
замены одного решения отклоняются.

После merge сохраняется один существующий ID, прочие становятся aliases. После
split ID сохраняет группа с исходным anchor. Если старый alias охватывает несколько
новых групп, он становится явно неоднозначным; `Registry.Resolve` возвращает
ошибку, а не выбирает одну сторону. Ранее выданный ID не переиспользуется.
`auto_aws: false` отключает новые автоматические связи; уже подтверждённые связи
сохраняются до явного решения о разделении.

`field_override` поддерживает `started_at`, `resolved_at`, `status`, `impact`,
`components`, `title`, `source_revision` и `duration_unknown`. Укажи `field`,
`value`, `members`, обычные атрибуты решения и `expected_hashes` для **каждого**
source key текущей группы. Хеши доступны в `plan.json → source_heads`.
`source_revision` требует одного member и hash принадлежащего ему snapshot в value.
Привязка проверяется по текущему head, даже если выбрана более ранняя ревизия.

Любое изменение snapshot делает override устаревшим. Перепроверь evidence,
добавь решение с новым ID и `supersedes`, обнови pins и повтори plan/apply.
Простое удаление старого решения теряет историю и не используется в этом процессе.

Для 23 legacy-записей уже оформлены решения `duration_unknown: true` авторства
Codex: сохранённый raw содержит завершение раньше начала. Этот override допустим
только при таком противоречии; он не переписывает даты. Длительность остаётся
неизвестной, вклад длительности в score равен нулю, событие не считается ongoing.
Threshold по impact продолжает работать. Одинаковые имена компонентов могут
обозначать разные vendor IDs: их количество сохраняется.

## Режимы и восстановление

- `off`: исходные записи в скоринге; сохранение новых observations остаётся включено.
- `shadow`: исходное представление плюс план сравнения; collect пишет
  `reports/dedup/shadow.json` рядом с config, digest — `shadow-dedup.json` в out root.
  Реестр и canonical artifacts не изменяются.
- `enforce`: collect сохраняет sources, затем reconcile/build; digest читает только
  согласованное canonical-представление без принятия новых identity-решений.

При изменении режима, правил или решений сформируй новый plan/apply перед
следующим enforce digest. После изменения только весов/threshold достаточно
rebuild. Изменение конфигурации инвалидирует прежний manifest; его нельзя
молча игнорировать.

Если `validate` сообщает о повреждённом индексе или canonical JSON, выполни
`dedup rebuild` и повтори validate. Это восстановление derived artifacts по
сохранённому реестру. Если прерван initial apply и manifest наблюдений уже записан,
но реестр не готов, нужен новый plan/apply. Повреждённый immutable snapshot или
сам реестр восстанавливается из согласованной резервной копии/Git, а не выводится
заново из выпусков. При `source projection differs from history` следующий
успешный collect восстанавливает совместимую запись из текущего ответа вендора;
до этого digest блокируется. При недоступном источнике используй сохранённую
согласованную копию для расследования.

Mutating-команды берут локальный `data/.archive.lock`; workflow дополнительно
использует `concurrency: archive`. После аварийного завершения lock может остаться.
Сначала проверь PID внутри файла и отсутствие работающего collector/digest,
затем перемести stale lock за пределы data root. Автоматического снятия чужой
блокировки нет. Читатели проверяют manifest, записываемый последним: отдельные
atomic rename не означают транзакцию всего каталога.

Collect выдаёт JSON `-result` с `source_saved` и `canonical_saved`. Если источники
сохранены, а canonical build упал, workflow коммитит только observations,
vendor JSON и state; частичные derived files не входят в этот commit. Job
остаётся ошибочным. После resync выполняется новый collect/reconcile. Локальная
команда dedup сама ничего не коммитит и не отправляет.

Для отката ошибочного объединения предпочтителен split с новым решением.
Для сравнения с прежней системой используй отдельную копию baseline и `mode: off`,
сохранив текущие observations и registry. Не смешивай старые индексы с новым
manifest. Резервная копия этой миграции находится локально в
`/private/tmp/weekly-incident-dedup-baseline-20261114`; временный каталог не заменяет
долговечную резервную копию. Публикация исправления и повторное уведомление —
отдельные действия, которые исторический rebuild не выполняет.

## Границы достоверности

Автоматическое правило AWS требует containment точных сообщений одной
service-region, минимум двух разных нетривиальных текстов и отсутствия
противоречий закрытия/нового события. Неоднозначные мосты A–B–C не склеиваются
транзитивно. Остальные адаптеры участвуют в полном аудите и ручных решениях;
автоматический fuzzy matching и межвендорские merge не включены.

Held-out проверка использует синтетические усечения отдельно захваченного AWS
фида: 3 положительных и 13 отрицательных сравнений. Это regression-проверка,
не независимая разметка всех реальных аварий и не оценка precision/recall.
AWS onset остаётся временем первого наблюдаемого сообщения. Недостоверный
старый investigating может завышать ongoing score; дедупликация сама по себе
не подтверждает, что авария всё ещё продолжается.

Пересборка старых недель с сегодняшними snapshots и фиксированным now является
пересмотром архива. Для точного воспроизведения прежнего выпуска нужны прежние
source snapshots, registry, decisions, config, код правил и прежний now.
