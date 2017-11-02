package local

// Comment ...
type Comment struct {
  SelfName string
  ButtonNew string
}

// Concept ...
type Concept struct {
  SelfName string
  ButtonNew string
  LabelTitle string
}

// Context ...
type Context struct {
  SelfName string
  ButtonNew string
}

// Discussion ...
type Discussion struct {
  SelfName string
  ButtonNew string
  LabelTitle string
  LabelText string
}

// Perspective ...
type Perspective struct {
  SelfName string
  ButtonNew string
  LabelText string
}

// Prerequisite ...
type Prerequisite struct {
  SelfName string
  ButtonNew string
}

// Prompts ...
type Prompts struct {
  WrongUsernameOrPassword string
  NoResultsFound string
  NotWhatYouWereLookingFor string
}

// Reasoning ...
type Reasoning struct {
  SelfName string
  LabelText string
}

// User ...
type User struct {
  SelfName string
  LabelUsername string
  LabelPassword string
  ButtonRegister string
  ButtonLogin string
}

// Strings ...
type Strings struct {
  Comment Comment
  Concept Concept
  Context Context
  Discussion Discussion
  Perspective Perspective
  Prerequisite Prerequisite
  Prompts Prompts
  Reasoning Reasoning
  User User
}

// NewStrings ...
func NewStrings() *Strings {
  comment := Comment{}
  concept := Concept{}
  context := Context{}
  discussion := Discussion{}
  perspective := Perspective{}
  prerequisite := Prerequisite{}
  prompts := Prompts{}
  reasoning := Reasoning{}
  user := User{}
  strings := Strings{
    Comment: comment,
    Concept: concept,
    Context: context,
    Discussion: discussion,
    Perspective: perspective,
    Prerequisite: prerequisite,
    Prompts: prompts,
    Reasoning: reasoning,
    User: user,
  }
  return &strings
}
