# Exemplos de Requests da API RecipeHub

## Buscar todas as receitas
```bash
curl -X GET http://localhost:8080/recipes
```

## Buscar receita por ID
```bash
curl -X GET http://localhost:8080/recipes/1
```

## Criar nova receita
```bash
curl -X POST http://localhost:8080/recipes \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "name": "Pão de Açúcar",
    "description": "Receita tradicional brasileira",
    "instructions": "Misture todos os ingredientes...",
    "ingredients": "Farinha, açúcar, ovos, leite",
    "category": "Sobremesa",
    "subcategory": "Pão doce",
    "tags": "tradicional, brasileiro"
  }'
```

## Atualizar receita
```bash
curl -X PUT http://localhost:8080/recipes/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "name": "Pão de Açúcar Atualizado",
    "description": "Receita tradicional brasileira melhorada",
    "instructions": "Misture todos os ingredientes...",
    "ingredients": "Farinha, açúcar, ovos, leite, canela",
    "category": "Sobremesa",
    "subcategory": "Pão doce",
    "tags": "tradicional, brasileiro, canela"
  }'
```

## Deletar receita
```bash
curl -X DELETE http://localhost:8080/recipes/1
```
