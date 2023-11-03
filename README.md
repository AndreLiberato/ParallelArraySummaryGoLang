# ParallelArraySummary
Implementação do projeto com Go Lang

# Como buildar

```
    go build
```

Esse comando criará um arquivo nomeado `ParallelArraySummary`

# Como executar 

```
    ./ParallelArraySummary [N] [T]
```

Sendo `N` aplicado a 10^N elementos e `T` o número de threads criada

# Arquivo de ids

Ao final da execução, será gerado dois arquivos.

> `id_less_five.csv`: ids menores que cinco 

> `id_greater_five.csv`: ids maiores ou iguais a cinco

Dependendo do número de elementos, esses arquivos podem ficar pesado.