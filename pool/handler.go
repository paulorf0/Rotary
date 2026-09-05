// Package pool é responsável pela lógica do gerenciamento e rotacionamento de ips da pool
package pool

// Ele vai receber uma chamada, pedindo um ip, para uma companhia especifica.
// Portanto, ele precisa ter informações sobre as proxys que a companhia cadastrou e que pode ser usada para gerar os ips.
// Ele deve gerenciar isso para diversas requisições simultaneas que vai chegar.

// O ip vai ser retirado ou menos rotacionado no momento em que o ip for reportado como ruim (conceito de fila de prioridade com envelhecimento?)
// O report de ip ruim é retirada definitiva?
// Não reportar nada significa que o handler deve entregar o mesmo ip ou vai ter outra lógica de entrega?
//
// Talvez o sistema de report pode ser algo como: reporta até chegar em 0, baseado em intervalo de tempo.
// Quanto menor o intervalo de tempo de report, mais rápido a prioridade do ip decaí. Quando chega em zero, é retirado do pool e dado lugar a outro ip novo.

// TODO: Validar se existe ip no pool para buscar valores.

// TODO: Report deve carregar motivo da falha (403/429/timeout/5xx), não só OK/NOK.
// Reação difere: 429 é cooldown temporário, 403 repetido decai score rápido, timeout != ban.

// TODO: transição para Burned deve ser por threshold de Score (ex: score < X),
// não decisão binária. Definir o valor do threshold.
// Deve ser dinâmico o threashold?
