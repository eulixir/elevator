# Rules

Um elevador de 10 andares (exemplo)

  - Um elevador pode subir até o andar X
  - Um elevador pode descer até o andar x
  - Se eu estiver no 2 andar, e eu quiser subir. Tem que aguardar.
  - Se eu quiser descer do 2 andar para para o primeiro. Se alugém quiser subir, tem que esperar outro elevador.
  - A primeira pessoa que entrar no elevador, tem que selecionar o andar se não ele fica parado (Até alguém chamar). 
  - Se entrar em um elevador que já tenha pessoas e não selecionar o andar, ao fechar as portas, vai seguindo o fluxo. 
  - Se o elevador está subindo e alguém coloca um andar a baixo, ele primeiro termina de subir para descer depois.
  <!-- - Se passar 3 minutos e não tiver movimentação, ele volta para o térreo. -->
  
- Requisitos técnicos

  - Ter um tempo de 5 segundos entre um andar e outro
  - Ter um observador para observar o estado dos elevadores
  
  Um elevador pode ser uma fila?
  
  [{
    direction: up,
    floor: 2
  }, [direction: up,
    floor: 4], []]
    
- Agrupar pela direção e ordenar pelo andar
