O padrão de design Factory é usado para criar objetos sem necessitar especificar a classe exata do objeto que será criado. Este padrão é particularmente útil em situações onde é necessário criar diferentes tipos de objetos que compartilham uma classe base comum ou interface, mas têm características ou dados iniciais diferentes.

### Aspectos Fundamentais do Padrão Factory

#### Separação da Criação do Uso
- **Como é Implementado**: Ambas as abordagens, funcional e estrutural, permitem criar objetos `Employee` especificando certos detalhes (como `Position` e `AnnualIncome`) sem necessitar conhecer a implementação exata da criação de `Employee`.
- **Benefício**: Isso permite alterar a implementação de como os objetos `Employee` são criados sem afetar o código que os utiliza.

#### Flexibilidade na Criação de Objetos
- **Como é Implementado**: A abordagem estrutural, em particular, permite ajustes pós-criação antes da instância final do objeto (como modificar `AnnualIncome` do `bossFactory`).
- **Benefício**: Fornece uma maneira flexível de ajustar os detalhes dos objetos antes da sua criação final.

#### Encapsulamento da Lógica de Instanciação
- **Como é Implementado**: Tanto a abordagem funcional quanto a estrutural encapsulam o processo de criação de objetos `Employee`, escondendo os detalhes da sua instanciação.
- **Benefício**: Simplifica o processo de criação de objetos e centraliza a lógica de instanciação, facilitando sua manutenção e modificação.

Este exemplo mostra claramente como o padrão Factory pode ser usado para criar objetos de maneira flexível e desacoplada, permitindo a extensão e manutenção fácil do código.