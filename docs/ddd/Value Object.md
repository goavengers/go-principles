### DDD Value Object

Давайте начнем путешествие по практическому предметно-ориентированному дизайну на 
Голанге с самого важного паттерна - объекта-ценности.

#### Просто, но красиво

Объект Value - это простой шаблон на первый взгляд. 
Он группирует несколько атрибутов в единый блок, который обеспечивает определенное поведение. 
Эта единица представляет собой определенное качество или количество, которое мы можем найти в реальном мире и привязать к более сложному объекту. 
Он обладает определенной ценностью или характеристиками. 
Это может быть цвет или деньги ( подтип объекта-значения), номер телефона или любой другой небольшой объект, который предоставляет некоторую ценность, как в блоке кода ниже.

```go
type Money struct {
  Value    float64
  Currency Currency
}

func (m Money) ToHTML() string {
  returs fmt.Sprintf(`%.2f%s`, m.Value, m.Currency.HTML)
}

type Salutation string

func (s Salutation) IsPerson() bool {
  returs s != "company" 
}  

type Color struct {
  Red   byte
  Green byte
  Blue  byte
}

func (c Color) ToCSS() string {
  return fmt.Sprintf(`rgb(%d, %d, %d)`, c.Red, c.Green, c.Blue)
}

type Address struct {
  Street   string
  Number   int
  Suffix   string
  Postcode int
}

type Phone struct {
  CountryPrefix string
  AreaCode      string
  Number        string
}
```

В Golang объекты-значения могут быть представлены как новые структуры или путем расширения какого-либо примитивного типа. 
В обоих случаях идея состоит в том, чтобы обеспечить дополнительное поведение, уникальное для этого отдельного значения или группы значений. 
Во многих случаях объект значения может предоставлять определенные методы для форматирования строк, чтобы определить, 
как значения должны вести себя при кодировании или декодировании JSON. 
Тем не менее, основная цель этих методов должна заключаться в поддержке бизнес-инвариантов, связанных с этой характеристикой или качеством в реальной жизни.

#### Идентичность и равенство

Объект Value не имеет идентичности, и это его критическое отличие от шаблона Entity . Шаблон сущности имеет идентификатор как описание его уникальности. Если две Сущности имеют некоторую идентичность, это означает, что мы говорим об одних и тех же объектах. Объект-значение не имеет этого идентификатора. Объект значения имеет только несколько полей, которые лучше описывают его значение. Чтобы проверить равенство между двумя объектами-значениями, нам нужно проверить равенство всех полей, как в блоке кода ниже.

```go
// checking equality for value objects
func (c Color) EqualTo(other Color) bool {
  return c.Red == other.Red && c.Green == other.Green && c.Blue == other.Blue
}

// checking equality for value objects
func (m Money) EqualTo(other Money) bool {
  return m.Value == other.Value && m.Currency.EqualTo(other.Currency)
}

// checking equality for entities
func (c Currency) EqualTo(other Currency) bool {
  return c.ID.String() == other.ID.String()
}
```

В приведенном выше примере структуры Money и Color определили методы EqualTo, которые проверяют все свои поля. С другой стороны, Currency проверяет равенство Identities, которым в этом примере является UUID.
Как вы могли заметить, объект Value также может ссылаться на некоторый объект Entity, например Money и Currency в этом случае. Он также может содержать некоторые другие объекты-значения меньшего размера, например структуру Coin, которая содержит Color и Money. Или определить срез как набор цветов.

```go
type Coin struct {
  Value Money
  Color Color
}

type Colors []Color
```

В одном ограниченном контексте у нас могут быть десятки объектов-значений. Тем не менее, некоторые из них могут быть сущностями внутри других ограниченных контекстов. Так будет с валютой. В простом веб-сервисе, где мы хотим отобразить деньги, мы можем рассматривать валюту как объект-значение, привязанный к нашим деньгам, которые мы не планируем менять. С другой стороны, в Платежной службе мы хотим получать обновления в реальном времени с помощью некоторого API службы Exchange, где нам нужно использовать удостоверения внутри модели домена. В этом случае у нас будут разные реализации валюты на разных сервисах.

```go
// value object on web service
type Currency struct {
  Code string
  HTML int
}

// entity on payment service
type Currency struct {
  ID   uuid.UUID
  Code string
  HTML int
}
```

Шаблон, который мы хотим использовать, объект-значение или сущность, зависит только от того, что этот объект представляет в ограниченном контексте. Если это многократно используемый объект, независимо хранящийся в базе данных, может изменяться и применяться ко многим другим объектам или связан с некоторой внешней сущностью, которая требуется для изменения при изменении внешнего, мы говорим о сущности. Но если объект описывает какое-то значение, принадлежит определенной сущности, является простой копией из внешней службы или не должен существовать независимо в базе данных, тогда это объект-значение.

#### Явность

Самая полезная особенность Value Object - это его ясность. Он обеспечивает ясность для внешнего мира в тех случаях, когда исходные типы из Golang (или любого другого языка программирования) не поддерживают конкретное поведение или поддерживаемое поведение не является интуитивно понятным. Мы можем работать с заказчиком по многим проектам, которые должны соответствовать некоторым бизнес-инвариантам, например, быть взрослым или представлять какое-либо юридическое лицо. В таких случаях допустимы более явные типы, такие как Birthday и LegalForm.

```go
type Birthday time.Time

func (b Birthday) IsYoungerThen(other time.Time) bool {
  return time.Time(b).After(other)
}

func (b Birthday) IsAdult() bool {
  return time.Time(b).AddDate(18, 0, 0).Before(time.Now())
}

const (
  Freelancer = iota
  Partnership
  LLC
  Corporation
)

type LegalForm int

func (s LegalForm) IsIndividual() bool {
  return s == Freelancer
}

func (s LegalForm) HasLimitedResponsability() bool {
  return s == LLC || s == Corporation
}
```

Иногда объект-значение не нужно явно определять как часть какой-либо другой сущности или объекта-значения. Тем не менее, мы можем определить объект значения как вспомогательный объект, который обеспечивает ясность для последующего использования в коде. Так обстоит дело с Клиентом, которым может быть Лицо или Компания. В зависимости от типа клиента у нас есть разные потоки в приложении. Одним из лучших подходов может быть преобразование клиентов, чтобы с ними было легче справляться.

```go
type Customer struct {
  ID        uuid.UUID
  Name      string
  LegalForm LegalForm
  Date      time.Time
}

func (c Customer) ToPerson() Person {
  return Person{
    FullName: c.Name,
    Birthday: c.Date,
 }
}

func (c Customer) ToCompany() Company {
  return Company{
    Name: c.Name,
    CreationDate: c.Date,
  }
}

type Person struct {
  FullName string
  Birthday Birthday
}

type Company struct {
  Name         string
  CreationDate time.Time
}
```

Хотя случаи с преобразованием могут происходить в некоторых проектах, в большинстве случаев они говорят нам, что мы должны добавить эти объекты-значения как реальную часть нашей модели предметной области. Фактически, всякий раз, когда мы замечаем, что какая-то конкретная меньшая группа полей постоянно взаимодействует друг с другом, но они находятся внутри какой-то более крупной группы, это уже знак того, что мы должны сгруппировать их в объект-значение и использовать его таким же образом внутри нашей большей группы. (который теперь становится меньше).

#### Неизменность

Объекты-значения неизменны. Не существует единой причины, причины или другого аргумента для изменения состояния объекта-значения в течение его жизненного цикла. Иногда несколько объектов могут содержать один и тот же объект-значение (хотя это не идеальное решение). В таких ситуациях мы определенно не хотим изменять объекты-значения в неожиданных местах. Итак, всякий раз, когда мы хотим изменить внутреннее состояние объекта-значения или объединить несколько из них, нам всегда нужно вернуть новый экземпляр с новым состоянием, как в блоке кода ниже.

```go
// wrong way to change the state inside value object
func (m *Money) AddAmount(amount float64) {
  m.Amount += amount
}

// right way to return new value objects with new state
func (m Money) WithAmount(amount float64) Money {
  return Money {
    Amount:   m.Amount + amount,
    Currency: m.Currency,
  }
}

// wrong way to change the state inside value object
func (m *Money) Deduct(other Money) {
  m.Amount -= other.Amount
}

// right way to return new value objects with new state
func (m Money) DeductedWith(other Money) Money {
  return Money {
    Amount:   m.Amount + other.Amount,
    Currency: m.Currency,
  }
}

// wrong way to change the state inside value object
func (c *Color) KeppOnlyGreen() {
  c.Red = 0
  c.Bed = 0
}

// right way to return new value objects with new state
func (c Color) WithOnlyGreen() Color {
  return Color {
    Red:   0,
    Green: c.Green,
    Blue:  0,
  }
}
```

Во всех примерах единственный правильный способ - всегда возвращать свежие экземпляры и оставлять старые нетронутыми. Хорошая практика в Golang - всегда привязывать функции к значениям вместо ссылок на объекты значений, чтобы быть уверенным, что мы никогда не изменим внутреннее состояние.

```go
func (m Money) Deduct(other Money) (Money, error) {
  if !m.Currency.EqualTo(other.Currency) {
    return Money{}, errors.New("currencies must be identical")
  }
  
  if other.Amount > m.Amount {
    return Money{}, errors.New("there is not enough amount to deduct")
  }
  
  return Money {
    Amount:   m.Amount - other.Amount,
    Currency: m.Currency,
  }
}
```

Эта неизменяемость означает, что мы не должны проверять объект-значение в течение всего его жизненного цикла, а только при создании, как это показано в приведенном выше примере. Когда мы хотим создать новый объект-значение, мы всегда должны выполнять проверку и возвращать ошибки, если бизнес-инварианты не выполняются, и создавать объект-значение только в том случае, если он действителен. С этого момента больше не нужно проверять объект значения.

#### Богатое поведение

Value Object обеспечивает множество различных вариантов поведения. Его основная цель - предоставить доступный интерфейс. Если это анемия, нам, вероятно, следует подумать о причине ее существования без каких-либо методов. Если Value Object действительно имеет смысл в каком-то конкретном месте кода, то он предоставляет огромное количество дополнительных бизнес-инвариантов, которые намного лучше описывают проблему, которую мы хотим решить.

```go
func (c Color) ToBrighter() Color {
  return Color {
    Red:   math.Min(255, c.Red + 10),
    Green: math.Min(255, c.Green + 10),
    Blue:  math.Min(255, c.Blue + 10),
  }
}

func (c Color) ToDarker() Color {
  return Color {
    Red:   math.Max(0, c.Red - 10),
    Green: math.Max(0, c.Green - 10),
    Blue:  math.Max(0, c.Blue - 10),
  }
}

func (c Color) Combine(other Color) Color {
  return Color {
    Red:   math.Min(255, c.Red + other.Red),
    Green: math.Min(255, c.Green + other.Green),
    Blue:  math.Min(255, c.Blue + other.Blue),
  }
}

func (c Color) IsRed() bool {
  return c.Red == 255 && c.Green == 0 && c.Blue == 0
}

func (c Color) IsYellow() bool {
  return c.Red == 255 && c.Green == 255 && c.Blue == 0
}

func (c Color) IsMagenta() bool {
  return c.Red == 255 && c.Green == 0 && c.Blue == 255
}

func (c Color) ToCSS() string {
  return fmt.Sprintf(`rgb(%d, %d, %d)`, c.Red, c.Green, c.Blue)
}
```

Декомпозиция всей модели предметной области на небольшие части, такие как объекты-значения (и сущности), делает код понятным и приближенным к бизнес-логике в реальном мире. Каждый объект-значение может описывать некоторые небольшие компоненты и поддерживать многие модели поведения, аналогичные обычным бизнес-процессам. В конце концов, это значительно упрощает весь процесс модульного тестирования и помогает охватить все крайние случаи.