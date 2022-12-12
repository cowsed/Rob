Error if we misalign the types:
I can't reconcile the types for this subscription. `Encoder.Position` is implemented as `int`(c++) (from Builtin/Encoder) yet its source, `$name.read(vex::rotationUnit::rev)`, returns a double (c++)


# Standards

Component.rob

```elm
new: Platform.Port -> Component

type EventType = {}
type OtherEventType = {}

onEvent: Sub EventType
onOtherEvent: Sub OtherEventType


type Command 
    = SetA
    | SetB

setSomething Command

```


Some platforms are defined in the stdlib but if you want to use your own platform create a 'platform/' directory and make files like those seen in components/module defintions

## Example
String library has implementation for c++, python
we want 


# Types

## Robot
Basically a main program

## Component
an abstraction over one piece of hardware. Examples: an Encoder, a Motor

these are instanced in a special way so that you can have subscriptions and commands to specific components.

## Module
more rob code but contained in a different file. These operate completely within `rob` so they need not provide implementations

## External Module

modules that have parts written in other languages. Example, you want to be able to call a searching algorithm written in c++ from elm code. 


TODO: CLarify difference between this and a Component