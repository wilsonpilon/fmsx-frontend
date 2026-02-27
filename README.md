# fmsx-frontend

Frontend em Go para o emulador **fMSX**, com duas interfaces planejadas no mesmo projeto:

- **TUI com `tview`** (habilitada com `-tags tview`)
- **GUI com `fyne`** (habilitada com `-tags fyne`)

O objetivo é integrar:

1. Documentação do fMSX: <https://fms.komkon.org/fMSX/fMSX.html>
2. Repositório File Hunter MSX: <https://download.file-hunter.com/>
3. Game Database (SQLite zip): <https://romdb.vampier.net/archive/sqllite3db-msxromsdb.zip>
4. Persistência local em SQLite para configurações e catálogos

## Estrutura inicial

```text
.
├── cmd/fmsx-frontend/main.go
├── internal/app/app.go
├── internal/config/sources.go
├── internal/db/store.go
├── internal/ui/tui/tui_basic.go
├── internal/ui/tui/tui_tview.go
├── internal/ui/gui/gui_stub.go
├── internal/ui/gui/gui_fyne.go
├── configs/sources.json
├── Justfile
└── README.md
```

## Configuração de downloads

Arquivo: `configs/sources.json`

- URLs oficiais do fMSX, File Hunter e MSX Game DB
- Caminhos locais de cache e extração
- Caminho do binário local do emulador

## Banco de dados SQLite

`internal/db/store.go` já define o schema inicial:

- `app_config`
- `file_hunter_entries`
- `game_db_roms`

> Nota: o driver SQLite (ex: `modernc.org/sqlite` ou `mattn/go-sqlite3`) será conectado no próximo passo. O código já está preparado para detectar quando o driver está ativo.

## Build e execução (Just)

```bash
just fmt
just test
just build
just run
```

### Build com TUI (`tview`)

```bash
just build-tui
just run-tui
```

### Build com GUI (`fyne`)

```bash
just build-gui
just run-gui
```

## Próximos passos sugeridos

1. Adicionar driver SQLite e migrações versionadas.
2. Implementar downloader de fontes (File Hunter + Game DB).
3. Importar checksums do Game DB para `game_db_roms`.
4. Criar listagem de ROMs e busca por checksum.
5. Integrar execução do fMSX com perfil de máquina e parâmetros.
