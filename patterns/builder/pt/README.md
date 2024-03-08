O padrão de design Builder é utilizado para separar a construção de um objeto complexo de sua representação, de maneira que o mesmo processo de construção possa criar diferentes representações. Esse padrão é especialmente útil quando um objeto precisa ser criado com muitas opções possíveis e nem todas elas são necessárias em cada instância.

### Aspectos Fundamentais do Padrão Builder Explicados com o Código

#### Separação da Construção da sua Representação
- **Como é Implementado**: O padrão permite construir diferentes "representações" (ou seja, estados ou configurações) do objeto `Person` sem alterar o processo subjacente. No código, isso é alcançado por meio do uso de `PersonJobBuilder` e `PersonAddressBuilder`, que gerenciam diferentes aspectos do `Person`.
- **Benefício**: Essa separação facilita a construção de variações complexas do objeto `Person` usando o mesmo construtor principal (`PersonBuilder`).

#### Encapsulamento da Construção
- **Como é Implementado**: Os detalhes de como o objeto `Person` é montado estão ocultos dentro dos construtores específicos. Por exemplo, o cliente (`main`) não precisa saber como `PersonJobBuilder` define o campo `CompanyName`.
- **Benefício**: O usuário do construtor só precisa interagir com a interface pública do construtor, sem se preocupar com os detalhes internos da construção do objeto.

#### Flexibilidade na Construção
- **Como é Implementado

**: Permitindo configurar diferentes aspectos do objeto `Person` de forma independente (trabalho e endereço), o código facilita modificações no objeto construído sem a necessidade de alterar o código existente.
- **Benefício**: Mudanças na estrutura ou nos atributos de `Person` podem ser feitas com mínimas alterações no código que utiliza o padrão Builder.

#### Interface Fluida ("Method Chaining")
- **Como é Implementado**: Os métodos de `PersonJobBuilder` e `PersonAddressBuilder` retornam referências a si mesmos, o que permite encadear chamadas de forma legível e expressiva, como mostrado na função `main`.
- **Benefício**: Esta característica melhora a legibilidade e simplifica o código cliente, tornando o processo de construção do objeto mais intuitivo e fluido.

Em resumo, este código ilustra como o padrão Builder não apenas simplifica a criação de objetos complexos, mas também fornece uma estrutura flexível, mantível e fácil de usar para configurar objetos com múltiplos atributos e representações.