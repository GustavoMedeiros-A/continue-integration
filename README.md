Apenas pra guardar coisas que eu constumo esquecer

- docker build -t test .
- docker run --rm test

Tirar strategy, não é necessario nesse caso
strategy:
    matrix:
        go: ['1.23', '1.23.2']