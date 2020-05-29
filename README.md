<div align="center">
  <img width="494" height="244" src="https://github.com/goavengers/go-principles/blob/master/img/solid_2x.png">
  <h1>Философия архитектуры ООП, SOLID-принципы, Dry, KISS и YAGNI</h1>
  <h5>Вместе мы разберемся!</h5>
</div>

## Содержание

- [ ] OOP (Object Oriented Programming)
- [x] [SOLID](#solid)
- [ ] DRY - Don’t repeat yourself
- [ ] KISS (Keep it simple, stupid!)
- [ ] Avoid Creating a YAGNI (You aren’t going to need it)
- [ ] LOD (Law of Demeter)

## <a name="solid"></a> SOLID

За этой аббревиатурой скрываются 5 базовых принципов ООП, предложенные __Робертом Мартином__ в его статье [«Принципы проектирования и шаблоны проектирования»](https://web.archive.org/web/20150906155800/http://www.objectmentor.com/resources/articles/Principles_and_Patterns.pdf). Следование их духу сделает код легко тестируемым, расширяемым, читаемым и поддерживаемым.

Вот шпаргалка по этим принципам:

- (__S__) Single Responsibility Principle (принцип единственности ответственности)
- (__O__) Open/Closed Principle (принцип открытости/закрытости)
- (__L__) Liskov Substitution Principle (принцип подстановки Барбары Лисков)
- (__I__) Interface Segregation Principle (принцип разделения интерфейса) 
- (__D__) Dependency Inversion Principle (принцип инверсии зависимостей) 

По словам __Роберта С. Мартина__, симптомы гниющего дизайна или плохого кода:

- __Жесткость__ (трудно менять код, так как простое изменение затрагивает много мест);
- __Неподвижность__ (сложно разделить код на модули, которые можно использовать в других программах);
- __Вязкость__ (разрабатывать и/или тестировать код довольно тяжело);
- __Ненужная Сложность__ (в коде есть неиспользуемый функционал);
- __Ненужная Повторяемость__ (Copy/Paste);
- __Плохая Читабельность__ (трудно понять что код делает, трудно его поддерживать);
- __Хрупкость__ (легко поломать функционал даже небольшими изменениями);

Но это улучшение, теперь мы можем сказать что то вроде "мне не нравится это потому, что слишком сложно модифицировать", или "мне не нравится это потому, что я не могу сказать, что этот код пытается сделать", но что насчет того, чтобы вести обсуждение позитивно?

Разве это не было бы здорово, если бы существовал способ описать свойства хорошего дизайна, а не только плохого и иметь возможность рассуждать объективными терминами? Для этого нам на помощь спешат принципы проектирования архитектуры и написания программного кода.

Сейчас мы рассмотрим эти принципы на схематичных примерах. Обратите внимание на то, что главная цель примеров заключается в том, чтобы помочь читателю понять принципы __SOLID__, узнать, как их применять и как следовать им, проектируя приложения. Автор материала не стремился к тому, чтобы выйти на работающий код, который можно было бы использовать в реальных проектах.

В golang в качестве отдельных частей у нас есть - пакаджи, структуры, методы и интерфейсы. Давайте расссмотрим SOLID в этих терминах.

- [x] [Single Responsibility Principle (принцип единственности ответственности)](./docs/solid/Single%20Responsibility%20Principle.md)
- [x] [Open/Closed Principle (принцип открытости/закрытости)](./docs/solid/Open%20Closed%20Principle.md)
- [ ] [Liskov Substitution Principle (принцип подстановки Барбары Лисков)](./docs/solid/Liskov%20Substitution%20Principle.md)
- [ ] [Interface Segregation Principle (принцип разделения интерфейса)](./docs/solid/Interface%20Segregation%20Principle.md)
- [ ] [Dependency Inversion Principle (принцип инверсии зависимостей)](./docs/solid/Dependency%20Inversion%20Principle.md)

## <a name="solid"></a> Object Oriented Programming
## <a name="solid"></a> DRY - Don’t repeat yourself
## <a name="solid"></a> KISS (Keep it simple, stupid!)
## <a name="solid"></a> Avoid Creating a YAGNI (You aren’t going to need it)
## <a name="solid"></a> LOD (Law of Demeter)
