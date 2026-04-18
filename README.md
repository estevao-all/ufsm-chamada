## Como fazer o desenvolvimento:

### Requisitos
- NodeJS + pnpm
- Go

### No Frontend:

Para instalar as dependências:

```bash
pnpm i
```

Para rodar o servidor local e servir o HTML (e automaticamente recompilar o site quando houverem mudanças):
```bash
pnpm run dev
```

### No Backend:

Para rodar o servidor local utilizado como API:
```bash
make dev
# ou
go run main.go
```
É necessário rodar o servidor novamente quando houverem mudanças.